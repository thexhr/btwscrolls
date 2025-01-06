package clog

import "log"

var DebugLog = false

func Debug(msg string, args ...interface{}) {
	if DebugLog {
		log.Printf("DEBUG: %s\n", msg)
	}
}
