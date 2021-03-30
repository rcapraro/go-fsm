package complex

import (
	"fmt"
)

const desiredReplicas = 2

const (
	CreatingDeployment StateType = iota +1
	RequeuingDeployment
	RetryingDeployment
	RunningDeployment
	StoppedDeployment
	FailedDeployment
)

func (s StateType) String() string {
	return [...]string{"DEFAULT", "CREATING", "REQUEUING", "RETRYING", "RUNNING", "STOPPED", "FAILED"}[s]
}

const (
	CreateDeployment EventType = iota +1
	RequeueDeployment
	RetryDeployment
	RunDeployment
	StopDeployment
	FailDeployment
)

type DeploymentCreationContext struct {
	NbReplicas int
}

type CreationDeploymentAction struct {}

func (c *CreationDeploymentAction) Execute(ctx EventContext) EventType {
	context, ok := ctx.(*DeploymentCreationContext)
	if !ok {
		fmt.Print("Bad context")
		return NoOp
	}
	fmt.Printf("Creating deployment with nbReplicas=%d\n", context.NbReplicas)
	if context.NbReplicas == desiredReplicas {
		return RunDeployment
	} else {
		return RequeueDeployment
	}
}

type RequeueDeploymentAction struct {}

func (r RequeueDeploymentAction) Execute(ctx EventContext) EventType {
	fmt.Println("Requeuing deployment...")
	return NoOp
}

func NewDeploymentFSM() *StateMachine {
	return &StateMachine{
		Current:  0,
		States:   States{
			Default: State{
				Events: Events {
					CreateDeployment: CreatingDeployment,
				},
			},
			CreatingDeployment: State{
				Action: &CreationDeploymentAction{},
				Events: Events{
					RequeueDeployment: RequeuingDeployment,
					RunDeployment:   RunningDeployment,
					RetryDeployment: RetryingDeployment,
				},
			},
			RequeuingDeployment: State{
				Action: &RequeueDeploymentAction{},
				Events: Events {
					RetryDeployment: RetryingDeployment,
					RunDeployment: RunningDeployment,
				},
			},
			RetryingDeployment: State{
				Events: Events{
					FailDeployment: FailedDeployment,
					RunDeployment:  RunningDeployment,
				},
			},
			StoppedDeployment: State{
				Events: Events{
					StopDeployment: StoppedDeployment,
				},
			},
		},
	}
}

