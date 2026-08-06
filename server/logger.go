package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

var (
	logMu     sync.Mutex
	logWriter io.Writer
)

func initLogger() error {
	if err := os.MkdirAll("logs", 0o755); err != nil {
		return err
	}
	file, err := os.OpenFile("logs/server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	logWriter = io.MultiWriter(os.Stdout, file)
	return nil
}

type logEntry struct {
	Time    string                 `json:"time"`
	Level   string                 `json:"level"`
	Event   string                 `json:"event"`
	Details map[string]interface{} `json:"details,omitempty"`
}

func logEvent(level, event string, details map[string]interface{}) {
	entry := logEntry{
		Time:    time.Now().UTC().Format(time.RFC3339Nano),
		Level:   level,
		Event:   event,
		Details: details,
	}
	data, err := json.Marshal(entry)
	if err != nil {
		return
	}

	logMu.Lock()
	defer logMu.Unlock()
	if logWriter == nil {
		fmt.Println(string(data))
		return
	}
	fmt.Fprintln(logWriter, string(data))
}

func logInfo(event string, details map[string]interface{})  { logEvent("INFO", event, details) }
func logWarn(event string, details map[string]interface{})  { logEvent("WARN", event, details) }
func logError(event string, details map[string]interface{}) { logEvent("ERROR", event, details) }
