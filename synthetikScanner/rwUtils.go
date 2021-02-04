package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

func readConfFile(path string) (conf, error) {
	fileReadData := conf{}
	fileData, err := ioutil.ReadFile(path)

	err = json.Unmarshal(fileData, &fileReadData)
	if logError(err) {
		return conf{}, err
	}

	return fileReadData, nil
}

func writeJsonFile(path string, sorted ScoreSort) error {
	jsonWriteData, err := json.MarshalIndent(sorted, "", "\t")
	if logError(err) {
		return err
	}

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if logError(err) {
			return err
		}
		defer file.Close()
	}
	err = ioutil.WriteFile(path, jsonWriteData, 0644)
	if logError(err) {
		return err
	}
	return nil
}

func readTemp(path string) (ScoreSort, error) {
	fileReadData := ScoreSort{}
	fileData, err := ioutil.ReadFile(path)

	err = json.Unmarshal(fileData, &fileReadData)
	if logError(err) {
		return ScoreSort{}, err
	}

	return fileReadData, nil
}

func readFile(path string) (bestMap, totalMap map[string]float64, err error) {
	bestMap = make(map[string]float64)
	totalMap = make(map[string]float64)

	file, err := os.Open(path)
	if logError(err) {
		return bestMap, totalMap, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		switch {
		case strings.HasPrefix(line, "total_"):
			line, score := beautifyString(line, "total_")
			totalMap[line] = score
		case strings.HasPrefix(line, "best_"):
			line, score := beautifyString(line, "best_")
			bestMap[line] = score
		}
	}
	return bestMap, totalMap, nil
}