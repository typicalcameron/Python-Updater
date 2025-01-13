package logToFile

import (
	"io"
	"log"
	"os"
	"time"
)

var (
	Log *log.Logger
)

func init() {
	logPath := "C:\\temp\\logToFile-" + time.Now().Format("2006-01-02") + ".log"

	logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}
	multiWriter := io.MultiWriter(logFile, os.Stdout)

	Log = log.New(multiWriter, "", 0)
	Log.SetFlags(0)
	Log.SetPrefix(time.Now().Format("2006-01-02 15:04:05") + ": ")
}
