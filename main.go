package main

import (
	"fmt"
	"net/http"

	"os"

	"github.com/Nickyg56/first_go_server/api"
	port "github.com/Nickyg56/first_go_server/utils"
)

func main() {
	argsWithoutProg := os.Args[1:]
	port := port.GetPort(argsWithoutProg)
	fmt.Println("Server Starting on port:", port)
	srv := api.NewServer()
	http.ListenAndServe(":"+port, srv)
}
