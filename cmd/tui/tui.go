package main

import (
	"fmt"
	"github.com/fetaro/gcal_forcerun_go/lib/installer"
)

func main() {
	config, err := installer.NewInstaller().MakeConfigFromUserInput()
	if err != nil {
		panic(err)
	}
	fmt.Printf("config: %v\n", config)
}
