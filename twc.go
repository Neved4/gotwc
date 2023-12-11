package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	showHumanReadable := flag.Bool("h", false, "Print human-readable format")
	formatSpecifier := flag.String("s", time.RFC3339, "Specify time format")
	filePath := flag.String("f", "", "Specify timezone file")
	timezoneFlag := flag.String("t", "", "Specify timezone directly")

	flag.Parse()

	format := *formatSpecifier
	if *showHumanReadable {
		format = "2006-01-02 15:04:05"
	}

	timezones := []string{}
	timezones = []string{"UTC"}
	if *timezoneFlag != "" {
		timezones = []string{*timezoneFlag}
	} else if tzFromEnv := os.Getenv("TZ"); tzFromEnv != "" {
		timezones = []string{tzFromEnv}
	} else if *filePath != "" {
		fileContent, err := os.ReadFile(*filePath)
		if err != nil {
			fmt.Printf("error reading file: %s\n", err)
			return
		}

		lines := strings.Split(string(fileContent), "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line != "" && !strings.HasPrefix(line, "#") {
				timezones = append(timezones, line)
			}
		}
	}

	maxWidth := 0
	for _, tz := range timezones {
		tz = strings.TrimSpace(tz)
		if tz == "" || strings.HasPrefix(tz, "#") {
			continue
		}
		if len(tz) > maxWidth {
			maxWidth = len(tz)
		}
	}
	maxWidth += 1

	for _, tz := range timezones {
		tz = strings.TrimSpace(tz)
		if loc, err := time.LoadLocation(tz); err == nil {
			timeInTZ := time.Now().UTC().In(loc)
			fmt.Printf("%-*s %s\n", maxWidth, tz, timeInTZ.Format(format))
		} else {
			fmt.Printf("error loading timezone %s: %s\n", tz, err)
		}
	}
}
