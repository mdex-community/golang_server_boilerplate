package port

import (
	"fmt"
)

func GetPort(cArgs []string) string {
	var port string

	for i, v := range cArgs {
		if v == "--port" || v == "-p" {
			if len(cArgs) <= i+1 {
				fmt.Println("Defaulting to port 3004")
				port = "3004"
			} else {
				port = cArgs[i+1]
			}
		}
	}

	if port == "" {
		fmt.Println("Defaulting to port 3004")
		port = "3004"
	}

	return port
}
