package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func handleError(err error, message string) {
	if err != nil {
		fmt.Printf("%s: %s\n", message, err)
		os.Exit(1)
	}
}

func getTimezones(filePath, timezoneFlag string) ([]string, error) {
	if timezoneFlag != "" {
		timezoneFlag = strings.Replace(timezoneFlag, "UTC-0", "UTC", 1)
		return []string{timezoneFlag}, nil
	}

	tzFromEnv := os.Getenv("TZ")
	if tzFromEnv != "" {
		return []string{tzFromEnv}, nil
	}

	if filePath == "" {
		filePath = getConfigPath()
	}

	return readTZFromFile(filePath)
}


func getConfigPath() string {
	xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")
	homeDir, err := os.UserHomeDir()
	handleError(err, "failed to get user home directory")

	configPath := filepath.Join(homeDir, ".config", "twc", "tz.conf")

	if xdgConfigHome != "" {
		configPath = filepath.Join(xdgConfigHome, "twc", "tz.conf")
	}

	return configPath
}

func readTZFromFile(filePath string) ([]string, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return []string{"UTC"}, nil
	}

	fileContent = []byte(strings.ReplaceAll(string(fileContent),
		"UTC-0", "UTC"))

	var timezones []string
	for _, line := range strings.Split(string(fileContent), "\n") {
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "#") {
			timezones = append(timezones, line)
		}
	}

	return timezones, nil
}

func main() {
	showHumanRead := flag.Bool("h", false, "Print human-readable format")
	formatSpecifier := flag.String("s", time.RFC3339, "Specify time format")
	filePath := flag.String("f", "", "Specify timezone file")
	timezoneFlag := flag.String("t", "", "Specify timezone directly")
	flag.Parse()

	format := *formatSpecifier
	if *showHumanRead {
		format = "2006-01-02 15:04:05"
	}

	timezones, err := getTimezones(*filePath, *timezoneFlag)
	handleError(err, "failed to get timezones")

	maxWidth := 0
	for _, tz := range timezones {
		tz = strings.TrimSpace(tz)
		if len(tz) > maxWidth {
			maxWidth = len(tz)
		}
	}
	maxWidth++

	for _, tz := range timezones {
		tz = strings.TrimSpace(tz)
		if tz == "" || strings.HasPrefix(tz, "#") {
			continue
		}

		if loc, err := time.LoadLocation(tz); err == nil {
			timeInTZ := time.Now().UTC().In(loc)
			fmt.Printf("%-*s %s\n", maxWidth, tz, timeInTZ.Format(format))
		} else {
			fmt.Printf("error loading timezone %s: %s\n", tz, err)
		}
	}
}
