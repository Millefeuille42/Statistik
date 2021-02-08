package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func runStart(tempPath, statPath string) error {
	best, total, err := readFile(statPath)
	if err != nil {
		return err
	}
	score := sortScore(best, total)
	err = writeJsonFile(tempPath, score)
	if err != nil {
		return err
	}
	return nil
}

func runEnd(tempPath, statPath string, conf conf) error {
	best, total, err := readFile(statPath)
	if err != nil {
		fmt.Println("No run started")
		return err
	}

	newScore := sortScore(best, total)
	oldScore, err := readTemp(tempPath)
	if err != nil {
		return err
	}


	newScore.RunDate = time.Now()
	for k := range oldScore.Total {
		if oldScore.Total[k] != 0 && oldScore.Total[k] != newScore.Total[k]{
			newScore.Total[fmt.Sprintf("%s_percent", k)] = ((newScore.Total[k] - oldScore.Total[k]) / oldScore.Total[k]) * 100
		}
		newScore.Total[k] -= oldScore.Total[k]
	}
	fName := strings.Replace(newScore.RunDate.Format(time.Stamp), ":", "_", -1)
	fName = strings.Replace(fName, " ", "_", -1)
	err = writeJsonFile(fmt.Sprintf("%s/%s.stats", conf.StatsOutputDirectory, fName), newScore)
	if err != nil {
		return err
	}
	err = os.Remove(tempPath)
	if logError(err) {
		return err
	}
	return nil
}

func main() {
	conf, err := readConfFile("./statistik_conf.json")
	if err != nil {
		return
	}

	printTitle()
	fmt.Println("\n---------------------")
	fmt.Print("Commands : start stop exit\n\n")
	microCli(conf)
}