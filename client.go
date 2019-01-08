package main

import (
	log "github.com/sirupsen/logrus"
	"itg/log"
)

func main() {
	itglog.Context().Info("from itg")
	itglog.Context().WithFields(log.Fields{
		"owner": "heince",
	}).Info("aloha")
}
