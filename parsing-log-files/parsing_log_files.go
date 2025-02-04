package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[^a-zA-Z0-9]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)["']([^"']*\bpassword\b[^"']*)["']`)

	count := 0
	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
		count += len(matches)
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\S+)`)

	for i, line := range lines {
		// Check if there is a match
		match := re.FindStringSubmatch(line)
		if match != nil {
			// Extract the username from the first capturing group
			username := match[1]
			// Modify the log line by prefixing it with "[USR] <username>"
			lines[i] = fmt.Sprintf("[USR] %s %s", username, line)
		}
	}

	return lines
}
