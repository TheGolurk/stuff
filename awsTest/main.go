package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	//	aws ec2 describe-instances --filters "Name=instance-type,Values=t2.micro" --query "Reservations[].Instances[].InstanceId"
	cmd := exec.Command("aws", "ec2", "describe-instances", "--filters", "Name=instance-type,Values=t2.micro", "--query", "Reservations[].Instances[].InstanceId")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	var a interface{}

	if err = json.NewDecoder(stdout).Decode(&a); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", a)

	// Expecting output with or wihin text (json) -> Then save as json to terminate the instances
}
