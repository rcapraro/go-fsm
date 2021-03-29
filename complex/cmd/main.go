package main


import (
	"fmt"
	"github.com/rcapraro/go-fsm/complex"
)

func main() {
	lightSwitchFSM := complex.NewLightSwitchFSM()

	// Set the initial "off" state in the state machine.
	err := lightSwitchFSM.SendEvent(complex.SwitchOff, nil)
	if err != nil {
		fmt.Printf("Couldn't set the initial state of the state machine, err: %v", err)
	}

	// Send the switch-off event again
	err = lightSwitchFSM.SendEvent(complex.SwitchOff, nil)
	if err != nil {
		fmt.Println("Could not switch off again")
	}

	// Send the switch-on event and transition to the "on" state.
	err = lightSwitchFSM.SendEvent(complex.SwitchOn, nil)
	if err != nil {
		fmt.Println(err)
	}

	// Send the switch-on event again
	err = lightSwitchFSM.SendEvent(complex.SwitchOn, nil)
	if err != nil {
		fmt.Println("Could not switch on again")
	}

	// Send the switch-off event and transition back to the "off" state.
	err = lightSwitchFSM.SendEvent(complex.SwitchOff, nil)
	if err != nil {
		fmt.Println(err)
	}
}
