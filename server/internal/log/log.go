package log

import "fmt"

type Level int

const (
	Fatal = iota
	Moderate
	Common
)

var LevelString = map[Level]string{
	Fatal:    "Fatal",
	Moderate: "Moderate",
	Common:   "Common",
}

func Log(lvl Level, msg string) {
	fmt.Printf("[%s] %s\n", LevelString[lvl], msg)
}
