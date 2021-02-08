package main

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"log"
	"strconv"
	"strings"
)

func printTitle() {
	figure.NewFigure("Statistik", "", true).Print()
	fmt.Println("A Synthetik stats tool. (Beta V1.0)")
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