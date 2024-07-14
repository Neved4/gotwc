package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var (
	progName   string
)

func errMsg(err error, message string) {
	if err != nil {
		fmt.Printf("%s: %s\n", message, err)
		os.Exit(1)
	}
}

func getTz(filePath, tzFlag string) ([]string, error) {
	if tzFlag != "" {
		tzFlag = strings.Replace(tzFlag, "UTC-0", "UTC", 1)
		return []string{tzFlag}, nil
	}

	tzFromEnv := os.Getenv("TZ")
	if tzFromEnv != "" {
		return []string{tzFromEnv}, nil
	}

	if filePath == "" {
		filePath = getConfigPath()
	}

	return readTzFile(filePath)
}

func getConfigPath() string {
	xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")
	homeDir, err := os.UserHomeDir()
	errMsg(err, "failed to get user home directory")

	configPath := filepath.Join(homeDir, ".config", "twc", "tz.conf")

	if xdgConfigHome != "" {
		configPath = filepath.Join(xdgConfigHome, "twc", "tz.conf")
	}

	return configPath
}

func readTzFile(filePath string) ([]string, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return []string{"UTC"}, nil
	}

	re := regexp.MustCompile(`(?m)^[ \t]*(#.*|\n)`)
	fileContent = re.ReplaceAll(fileContent, []byte{})
	fileContent = []byte(strings.ReplaceAll(string(fileContent), "UTC-0", "UTC"))

	var timezones []string
	for _, line := range strings.Split(string(fileContent), "\n") {
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "#") {
			timezones = append(timezones, line)
		}
	}

	return timezones, nil
}

func usage() {
	usage := `usage: %s [-h] [-f path] [-s format] [-t timezone] ...

Options:
  -f path
        Read config from path (default "$HOME/.config/twc/tz.conf")
  -h    Print in human-readable format
  -s format
        Set desired time format (e.g. "%%Y-%%m-%%d")
  -t timezone
        Set a specific timezone (e.g. "Asia/Tokyo")
`
	usageStr := fmt.Sprintf(usage, progName)
	fmt.Fprint(flag.CommandLine.Output(), usageStr)
}

func main() {
	progName = filepath.Base(os.Args[0])
	flag.Usage = usage

	fmtHuman := flag.Bool("h", false, "Print human-readable format")
	fmtSpec := flag.String("s", time.RFC3339, "Specify time format")
	filePath := flag.String("f", "", "Specify timezone file")
	tzFlag := flag.String("t", "", "Specify timezone directly")

	flag.Parse()

	format := *fmtSpec
	if *fmtHuman {
		format = "2006-01-02 15:04:05"
	}

	timezones, err := getTz(*filePath, *tzFlag)
	errMsg(err, "failed to get timezones")

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

		loc, err := time.LoadLocation(tz)
		if err != nil {
			fmt.Printf("error loading timezone %s: %s\n", tz, err)
			continue
		}

		tzTime := time.Now().UTC().In(loc)
		fmt.Printf("%-*s %s\n", maxWidth, tz, tzTime.Format(format))
	}
}
