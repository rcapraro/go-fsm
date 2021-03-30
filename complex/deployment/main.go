package main

import (
	"fmt"
	. "github.com/rcapraro/go-fsm/complex"
)

func main() {

	deploymentFSM := NewDeploymentFSM()

	deploymentCreationCtx := &DeploymentCreationContext{NbReplicas: 1}

	err := deploymentFSM.SendEvent(CreateDeployment, deploymentCreationCtx)
	if err != nil {
		fmt.Printf("Error %v", err)
	}

	//Should be still REQUEUING because the desired number of replicas is not reached
	fmt.Println(deploymentFSM.Current)
}