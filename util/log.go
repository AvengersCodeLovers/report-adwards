package util

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

// GetDurationInMiliseconds takes a start time and returns a duration in milliseconds
func GetDurationInMiliseconds(start time.Time) float64 {
	end := time.Now()
	duration := end.Sub(start)
	milliseconds := float64(duration) / float64(time.Millisecond)
	rounded := float64(int(milliseconds*100+.5)) / 100
	return rounded
}

// UseJSONLogFormat sets up the JSON log formatter
func UseJSONLogFormat() {
	env := GetEnv("APP_ENV", "development")
	appName := GetEnv("APP_NAME", "application-default-name")

	log.SetFormatter(&JSONFormatter{
		AppName: appName,
		Env:     env,
	})

	// so our debug entries appear!
	log.SetLevel(log.DebugLevel)

	f, err := os.OpenFile("storage/logs/logs.json", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	logrus.SetOutput(f)
}

// Timestamps in microsecond resolution (like time.RFC3339Nano but microseconds)
var timeStampFormat = "2006-01-02 15:04:05"

// JSONFormatter is a logger for use with Logrus
type JSONFormatter struct {
	AppName string
	Env     string
}

// Format includes the program, environment, and a custom time format: microsecond resolution
func (f *JSONFormatter) Format(entry *log.Entry) ([]byte, error) {
	data := make(log.Fields, len(entry.Data)+3)
	for k, v := range entry.Data {
		data[k] = v
	}

	data["time"] = entry.Time.Format(timeStampFormat)
	data["msg"] = entry.Message
	data["level"] = strings.ToUpper(entry.Level.String())
	data["app"] = f.AppName
	data["env"] = f.Env

	serialized, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}
	return append(serialized, '\n'), nil
}
