package complex

import (
	"errors"
	"fmt"
)

var UnhandledEvent = errors.New("unhandled event")

type StateType int

const (
	Default StateType = iota
	Off
	On
)

type EventType int

const (
	NoOp EventType = iota
	SwitchOn
	SwitchOff
)

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
	return Default, UnhandledEvent
}

func (s *StateMachine) SendEvent(event EventType, ctx EventContext) error {
	for {
		nextState, err := s.getNextState(event)
		if err != nil {
			return UnhandledEvent
		}

		state, ok := s.States[nextState]
		if !ok || state.Action == nil {
			// configuration error
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

type OnAction struct{}

func (o OnAction) Execute(ctx EventContext) EventType {
	fmt.Println("Light turned on")
	return NoOp
}

type OffAction struct{}

func (o OffAction) Execute(ctx EventContext) EventType {
	fmt.Println("Light turned off")
	return NoOp
}

func NewLightSwitchFSM() *StateMachine {
	return &StateMachine{
		States: States{
			Default: State{
				Events: Events{
					SwitchOff: Off,
				},
			},
			Off: State{
				Action: &OffAction{},
				Events: Events{
					SwitchOn: On,
				},
			},
			On: State{
				Action: &OnAction{},
				Events: Events{
					SwitchOff: Off,
				},
			},
		},
	}
}



