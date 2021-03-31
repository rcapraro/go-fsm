package complex

import (
	"fmt"
)

const desiredReplicas = 2

const (
	Creating StateType = iota +1
	Requeuing
	Retrying
	Running
	Stopped
	Failed
)

func (s StateType) String() string {
	return [...]string{"Default", "Creating", "Requeuing", "retrying", "Running", "Stopped", "Failed"}[s]
}

const (
	CreateDeployment EventType = iota +1
	RequeueDeployment
	RetryDeployment
	RunDeployment
	StopDeployment
	FailDeployment
)

func (s EventType) String() string {
	return [...]string{"No Operation", "Create", "Requeue", "Retry", "Run", "Stop", "Fail"}[s]
}

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

type RunDeploymentAction struct {}

func (r RunDeploymentAction) Execute(ctx EventContext) EventType {
	fmt.Println("Deployment is running, updating its status...")
	return NoOp
}

func NewDeploymentFSM() *StateMachine {
	return &StateMachine{
		Current:  0,
		States:   States{
			Default: State{
				Events: Events {
					CreateDeployment: Creating,
				},
			},
			Creating: State{
				Action: &CreationDeploymentAction{},
				Events: Events{
					RequeueDeployment: Requeuing,
					RunDeployment:     Running,
					RetryDeployment:   Retrying,
				},
			},
			Requeuing: State{
				Action: &RequeueDeploymentAction{},
				Events: Events {
					CreateDeployment: Creating,
					RetryDeployment:  Retrying,
					RunDeployment:    Running,
				},
			},
			Retrying: State{
				Events: Events{
					FailDeployment: Failed,
					RunDeployment:  Running,
				},
			},
			Running: State {
				Action: &RunDeploymentAction{},
				Events: Events{

				},
			},
			Stopped: State{
				Events: Events{
					StopDeployment: Stopped,
				},
			},
		},
	}
}

