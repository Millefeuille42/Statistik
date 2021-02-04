package main

import (
	"fmt"
	"os"
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
		return err
	}

	newScore := sortScore(best, total)
	oldScore, err := readTemp(tempPath)
	if err != nil {
		return err
	}

	err = os.Remove(tempPath)
	if logError(err) {
		return err
	}

	newScore.RunDate = time.Now()
	for k := range oldScore.Total {
		if oldScore.Total[k] != 0 && oldScore.Total[k] != newScore.Total[k]{
			newScore.Total[fmt.Sprintf("%s_percent", k)] = (newScore.Total[k] / oldScore.Total[k]) * 100
		}
		newScore.Total[k] -= oldScore.Total[k]
	}
	newScore.Total["run_approx_duration_min"] = time.Since(oldScore.RunDate).Minutes()

	err = writeJsonFile(fmt.Sprintf("%s/%s", conf.StatsOutputDirectory, newScore.RunDate.Format(time.Stamp)), newScore)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	conf, err := readConfFile("./synthetikScanner_conf.json")
	if err != nil {
		return
	}

	printTitle()
	fmt.Println("\n---------------------")
	fmt.Print("Commands : start stop exit\n\n")
	microCli(conf)
}