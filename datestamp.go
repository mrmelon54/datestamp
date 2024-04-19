package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	dryFlag := flag.Bool("d", false, "dry run: run without making changes")
	interactiveFlag := flag.Bool("i", false, "interactive: manually accept each change")
	flag.Parse()

	entries, err := os.ReadDir(".")
	if err != nil {
		log.Fatal("Failed to read files in directory:", err)
		return
	}
	for _, i := range entries {
		i.Name()
		info, err := i.Info()
		if err != nil {
			log.Printf("Failed to get file info for '%s': %v\n", i.Name(), err)
			continue
		}
		t, timeType := getEarliestFileTime(info)

		// only log changes on dry run
		if *dryFlag {
			log.Printf("[+] Would rename '%s' using '%s' of '%s'\n", info.Name(), timeType, t.Format(time.DateOnly))
			continue
		}

		// ask the user in interactive mode
		if *interactiveFlag {
			check := promptCheckbox(fmt.Sprintf("[?] Rename '%s' using '%s' of '%s'? [y/N] ", info.Name(), timeType, t.Format(time.DateOnly)))
			if !check {
				continue
			}
		}

		// try to rename
		newName := fmt.Sprintf("%s %s", t.Format(time.DateOnly), i.Name())
		err = os.Rename(i.Name(), newName)
		if err != nil {
			log.Printf("Failed to rename '%s' to '%s': %v\n", i.Name(), newName, err)
		}
	}
}
