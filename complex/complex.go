package complex

import (
	"fmt"
)

const (
	Default StateType = 0
)

const (
	NoOp EventType = 0
)

type StateType int

type EventType int

// Action to be executed on a given State, can return another EventType
type Action interface {
	Execute(ctx EventContext) EventType
}

// Context to be passed to the Action
type EventContext interface {
}

// Events and States mapping represents
type Events map[EventType]StateType

// State Holds the Action to be executed and the possible Events it can handle
type State struct {
	Action Action
	Events Events
}

// StateType and State mapping
type States map[StateType]State

type StateMachine struct {
	Previous StateType
	Current StateType
	States States
}

func (s *StateMachine) getNextState(event EventType) (StateType, error) {
	if state, ok := s.States[s.Current]; ok {
		if state.Events != nil {
			if next, ok := state.Events[event]; ok {
				return next, nil
			}
		}
	}
	return Default, fmt.Errorf("unhandled event %v for state %v\n", event, s.Current)
}

func (s *StateMachine) SendEvent(event EventType, ctx EventContext) error {
	for {
		nextState, err := s.getNextState(event)
		if err != nil {
			return err
		}

		state, ok := s.States[nextState]
		if !ok {
			return fmt.Errorf("no State for %v\n",nextState)
		} else if  state.Action == nil {
			return fmt.Errorf("no Action for %v\n",nextState)
		} else if state.Events == nil {
			return fmt.Errorf("no Events for %v\n",nextState)
		}

		// Transition over to the next state.
		s.Previous = s.Current
		s.Current = nextState

		// Execute the next state's action and stop if the event returned is a NoOp.
		nextEvent := state.Action.Execute(ctx)
		if nextEvent == NoOp {
			return nil
		}
		event = nextEvent
	}
}



