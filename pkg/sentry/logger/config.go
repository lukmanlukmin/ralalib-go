package logger

import "strings"

// Config :nodoc:
type Config struct {
	URL         string `json:"url"`
	Debug       bool   `json:"debug"`
	Environment string `json:"environment"`
	Level       string `json:"level"`
}

const (
	// LogTypePrint this log type for print
	LogTypePrint = "print"

	// LogTypeFile this log type for file
	LogTypeFile = "file"

	// LogFormatJSON this log format json
	LogFormatJSON = "json"
)

var envs = map[string]string{
	"production":  "production",
	"staging":     "staging",
	"development": "development",
	"prod":        "production",
	"stg":         "staging",
	"dev":         "development",
	"prd":         "production",
	"green":       "green",
	"blue":        "blue",
}

// Environment :nodoc:
func Environment(env string) string {
	v, ok := envs[strings.ToLower(strings.Trim(env, " "))]

	if !ok {
		return ""
	}

	return v
}
