package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func microCli(conf conf) {
	reader := bufio.NewReader(os.Stdin)
	runState := false

	for {
		fmt.Print("-> ")

		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		text = strings.TrimSpace(text)

		switch {
		case text == "start":
			if runState {
				fmt.Println("A run is already started")
				continue
			}
			runState = true
			fmt.Println("\tRun Started")
			_ = runStart(conf.TempFilePath, conf.StatsPath)

		case text == "stop":
			if err := runEnd(conf.TempFilePath, conf.StatsPath, conf); err != nil {
				fmt.Println(err)
				continue
			}
			runState = false
			fmt.Println("\tRun Ended")
		case text == "exit":
			if runState {
				fmt.Println("Exit current run before exiting")
				continue
			}
			os.Exit(0)
		default:
			fmt.Println("\tInvalid Command")
		}
	}
}