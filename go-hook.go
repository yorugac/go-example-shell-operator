package main

import (
	"fmt"
	"os"

	"github.com/savaki/jq"
)

func main() {
	args := os.Args
	if len(args) > 1 && args[1] == "--config" {
		fmt.Println(`configVersion: v1
kubernetes:
- apiVersion: v1
  kind: Pod
  executeHookOnEvent: ["Added"]
`)
	} else {
		ctxVar, ok := os.LookupEnv("BINDING_CONTEXT_PATH")
		if !ok {
			fmt.Println("LookupEnv absent BINDING_CONTEXT_PATH")
			os.Exit(1)
		}

		data, err := os.ReadFile(ctxVar)
		if err != nil {
			fmt.Println("os.ReadFile err:", err)
			os.Exit(1)
		}

		op, err := jq.Parse(".[0].object.metadata.name")
		if err != nil {
			fmt.Println("jq.Parse err:", err)
			os.Exit(1)
		}
		value, err := op.Apply(data)
		if err != nil {
			fmt.Println("jq Apply err:", err)
			os.Exit(1)
		}
		fmt.Println(string(value))
	}
}
