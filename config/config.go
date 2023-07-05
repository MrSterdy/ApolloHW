package config

import (
	"encoding/json"
)

var File = new(Config)

type Config struct {
	Messages struct {
		Templates struct {
			User string `json:"user"`
			Date struct {
				Layout   string `json:"layout"`
				Template string `json:"template"`
			} `json:"date"`
		} `json:"templates"`
		Commands struct {
			Start string `json:"start"`
		} `json:"commands"`
		Notifications struct {
			Timetable string `json:"timetable"`
			Logs      struct {
				Timetable   string `json:"timetable"`
				Class       string `json:"class"`
				Holidays    string `json:"holidays"`
				Subjects    string `json:"subjects"`
				Application string `json:"application"`
			} `json:"logs"`
		} `json:"notifications"`
		Weekdays []string `json:"weekdays"`
	} `json:"messages"`
}

func Init(file string) {
	err := json.Unmarshal([]byte(file), File)
	if err != nil {
		panic(err)
	}
}
