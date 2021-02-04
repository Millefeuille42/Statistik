package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func printTitle() {
	fmt.Println("  _____             _   _          _   _      _       _____")
	fmt.Println(" / ____|           | | | |        | | (_)    | |     / ____|")
	fmt.Println("| (___  _   _ _ __ | |_| |__   ___| |_ _  ___| | __ | (___   ___ __ _ _ __  _ __   ___ _ __")
	fmt.Println(" \\___ \\| | | |  _ \\| __|  _ \\ / _ | __| |/ __| |/ /  \\___ \\ / __/ _  |  _ \\|  _ \\ / _ |  __|")
	fmt.Println(" ____) | |_| | | | | |_| | | |  __| |_| | (__|   <   ____) | (_| (_| | | | | | | |  __| |")
	fmt.Println("|_____/ \\__, |_| |_|\\__|_| |_|\\___|\\__|_|\\___|_|\\_\\ |_____/ \\___\\__,_|_| |_|_| |_|\\___|_|  Beta V1.0")
	fmt.Println("          __/ |                                                                              ")
	fmt.Println("         |___/                                             ")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func logError(err error) bool {
	if err != nil {
		log.Print(err)
		return true
	}
	return false
}

func beautifyString(line, prefix string) (string, float64) {
	line = strings.Replace(line, prefix, "", 1)
	lineSplit := strings.Split(line, "=")
	lineScore, err := strconv.ParseFloat(strings.Trim(lineSplit[1], "\""), 64)
	logError(err)

	return lineSplit[0], lineScore
}