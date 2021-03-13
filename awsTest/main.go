package main

import "fmt"

func main() {

	instances := []string{
		"ec2-00-000-00-000.compute-1.amazonaws.com",
	}

	names := make(map[string]string)
	names["PEPITO JUANITO"] = "PEPITO.LOCO@EMPRESA.com"

	index := 0
	for k, i := range names {
		fmt.Printf("|Correo: %s Nombre: %s | -> Instancia: ssh -i mxintech-key.pem ec2-user@%s \n\n", i, k, instances[index])
		index++
	}

}
