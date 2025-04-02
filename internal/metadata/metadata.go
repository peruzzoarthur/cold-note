package metadata

import (
	"strings"
	"time"
)

type Metadata struct {
	Title   string
	Tags    []string
	Aliases []string
}

func FormatMetadata(content string, metadata Metadata) string {
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	tomorrow := now.AddDate(0, 0, 1)

	replacements := map[string]string{
		"{{date:YYYYMMDD}}":        now.Format("20060102"),
		"{{time:HHmm}}":            now.Format("1504"),
		"{{date:YYYY-MM-DD}}":      now.Format("2006-01-02"),
		"{{yesterday:YYYY-MM-DD}}": yesterday.Format("2006-01-02"),
		"{{tomorrow:YYYY-MM-DD}}":  tomorrow.Format("2006-01-02"),
		"{{title}}":                metadata.Title,
		"{{tags}}":                 strings.Join(metadata.Tags, ", "),
		"{{alias}}":                strings.Join(metadata.Aliases, ", "),
	}

	for placeholder, value := range replacements {
		content = strings.ReplaceAll(content, placeholder, value)
	}

	return content
}
