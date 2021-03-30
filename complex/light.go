package complex

import (
	"fmt"
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

