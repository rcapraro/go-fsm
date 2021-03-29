package main

import (
	"fmt"
	. "github.com/rcapraro/go-fsm/complex"
)

const (
	Off StateType = iota + 1
	On
)

const (
	SwitchOn EventType = iota + 1
	SwitchOff
)

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

func main() {
	lightSwitchFSM := NewLightSwitchFSM()

	// Set the initial "off" state in the state machine.
	err := lightSwitchFSM.SendEvent(SwitchOff, nil)
	if err != nil {
		fmt.Printf("Couldn't set the initial state of the state machine, err: %v", err)
	}

	// Send the switch-off event again
	err = lightSwitchFSM.SendEvent(SwitchOff, nil)
	if err != nil {
		fmt.Println("Could not switch off again")
	}

	// Send the switch-on event and transition to the "on" state.
	err = lightSwitchFSM.SendEvent(SwitchOn, nil)
	if err != nil {
		fmt.Println(err)
	}

	// Send the switch-on event again
	err = lightSwitchFSM.SendEvent(SwitchOn, nil)
	if err != nil {
		fmt.Println("Could not switch on again")
	}

	// Send the switch-off event and transition back to the "off" state.
	err = lightSwitchFSM.SendEvent(SwitchOff, nil)
	if err != nil {
		fmt.Println(err)
	}
}
