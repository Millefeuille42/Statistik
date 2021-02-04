package main

import "time"

type conf struct {
	TempFilePath			string	`json:"tempFilePath"`
	StatsPath				string	`json:"statsPath"`
	StatsOutputDirectory	string	`json:"statsOutputDirectory"`
}

type ScoreSort struct {
	RunDate	time.Time
	Total	map[string]float64
	Best	map[string]float64
}
