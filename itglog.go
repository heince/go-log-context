package itglog

import (
	log "github.com/sirupsen/logrus"
	"os"
)

// Context - return default log context
func Context() *log.Entry {
	/*
		Entry fields:
		- Hostname
		- User ID
		- Process ID
		- Working Directory
		- Program Name
	*/

	hostname, err := os.Hostname()
	if err != nil {
		panic("can't get hostname")
	}

	userID := os.Getuid()
	processID := os.Getpid()

	workingDir, err := os.Getwd()
	if err != nil {
		panic("can't get current working directory")
	}

	// Program Name can be set on Environment variable
	// set "ITG_PROGRAM_NAME"
	programName := os.Getenv("ITG_PROGRAM_NAME")
	if programName == "" {
		programName = "empty"
	}

	// reuse fields
	contextLog := log.WithFields(log.Fields{
		"hostname":     hostname,
		"pid":          processID,
		"uid":          userID,
		"working dir":  workingDir,
		"program name": programName,
	})

	return contextLog
}
