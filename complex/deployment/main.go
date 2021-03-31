package main

import (
	"fmt"
	. "github.com/rcapraro/go-fsm/complex"
)

func main() {

	deploymentFSM := NewDeploymentFSM()

	err := deploymentFSM.SendEvent(CreateDeployment,  &DeploymentCreationContext{NbReplicas: 1})
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}

	//Should be still REQUEUING because the desired number of replicas is not reached
	fmt.Println(deploymentFSM.Current)

	err = deploymentFSM.SendEvent(CreateDeployment, &DeploymentCreationContext{NbReplicas: 2})
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}

	//Should be RUNNING
	fmt.Println(deploymentFSM.Current)
}