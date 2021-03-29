package main

import (
	"fmt"
	. "github.com/rcapraro/go-fsm/complex"
)

const (
	Creating StateType = iota +1
	Retrying
	Running
	Stopped
	Failed
)

const (
	CreateDeployment EventType = iota +1
	RetryDeployment
	RunDeployment
	StopDeployment
	FailDeployment
)

type DeploymentCreationContext struct {
	nbReplicas int
}

type CreationDeploymentAction struct {

}

func (c *CreationDeploymentAction) Execute(ctx EventContext) EventType {
	context := ctx.(*DeploymentCreationContext)
	fmt.Printf("Creating deployment with nbReplicas=%d\n", context.nbReplicas)
	// if
	return RetryDeployment
}

func main() {


}