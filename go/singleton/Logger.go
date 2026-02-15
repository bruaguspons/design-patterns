package main

import (
	"fmt"
	"sync"
)

type Logger struct{}

var instance *Logger
var once sync.Once

func GetInstance() *Logger {
	once.Do(func() {
		instance = &Logger{}
	})
	return instance
}

func (l *Logger) Log(message string) {
	fmt.Printf("[LOG]: %s\n", message)
}
