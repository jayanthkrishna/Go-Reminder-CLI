package main

import (
	"github.com/whyrusleeping/go-notifier"
)

func main() {

	var n notifier.Notifier

	n.Notify(map[string]interface{}{
		"title":      "Unknown title",
		"message":    "Unknown message",
		"sound":      true,
		"wait":       true,
		"reply":      true,
		"closeLabel": "Completed?",
		"timeout":    15,
	})

}
