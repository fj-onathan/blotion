package config

import (
	"strings"
)

func Contains(str string, subs ...string) bool {
	isCompleteMatch := false
	for _, sub := range subs {
		if strings.Contains(str, sub) {
			isCompleteMatch = true
		}
	}
	return isCompleteMatch
}

// Convert ID page to Notion
func ToDashID(id string) string {
	res := id[:8] + "-" + id[8:12] + "-" + id[12:16] + "-" + id[16:20] + "-" + id[20:]
	return res
}
