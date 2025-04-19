package main

import (
	"log"

	"github.com/mujhtech/agent-cli/pkg/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
