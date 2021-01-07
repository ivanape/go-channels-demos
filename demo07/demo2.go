package main

import (
	"log"
	"time"
)

func task1(name string) {
	log.Printf("Hello %s from task1", name)
	time.Sleep(time.Duration(defaultDuration) * time.Second)
}
