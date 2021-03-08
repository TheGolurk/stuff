package main

import (
	"fmt"
	"os/exec"
)

func main() {
	//	aws ec2 describe-instances --filters "Name=instance-type,Values=t2.micro" --query "Reservations[].Instances[].InstanceId"
	cmd, _ := exec.Command("aws", `ec2 describe-instances --filters "Name=instance-type,Values=t2.micro" --query "Reservations[].Instances[].InstanceId"`).Output()
	fmt.Println("s", string(cmd))
	fmt.Printf("%v", cmd)

	// Expecting output with or wihin text (json) -> Then save as json to terminate the instances
}
