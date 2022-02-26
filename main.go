package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"time"

	"github.com/TheCreeper/go-notify"
)

type alert struct {
	Output       string
	Priority     string
	Rule         string
	Tags         []string
	Time         time.Time
	OutputFields map[string]interface{} `json:"output_fields"`
}

func notificationFromAlert(a alert) notify.Notification {
	return notify.Notification{
		AppName: "falco",
		AppIcon: "emblem-important",
		Summary: a.Rule,
		Body:    a.Output,
		Hints: map[string]interface{}{
			notify.HintUrgency: notify.UrgencyCritical,
		},
	}
}

func main() {
	dec := json.NewDecoder(os.Stdin)
	for {
		var a alert
		if err := dec.Decode(&a); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		if _, err := notificationFromAlert(a).Show(); err != nil {
			log.Fatal(err)
		}
	}
}
