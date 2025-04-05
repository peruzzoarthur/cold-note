package file

import (
	"encoding/json"
	"os"

	"github.com/charmbracelet/huh"
)

// Tag represents a tag in the tag selection system
type Tag struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Selected bool   `json:"selected"`
}

// LoadTagsFromJSON loads tags from a JSON file
func LoadTagsFromJSON(path string) ([]huh.Option[string], error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var tags []Tag
	if err := json.Unmarshal(data, &tags); err != nil {
		return nil, err
	}

	options := make([]huh.Option[string], len(tags))
	for i, tag := range tags {
		options[i] = huh.NewOption(tag.Name, tag.Value)
	}

	return options, nil
}

// SaveTagToJSON adds a new tag to the JSON file
func SaveTagToJSON(path string, name, value string) error {
	// Read existing tags
	var tags []Tag
	data, err := os.ReadFile(path)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		// If file doesn't exist, start with an empty slice
		tags = []Tag{}
	} else {
		if err := json.Unmarshal(data, &tags); err != nil {
			return err
		}
	}

	// Check if tag with same value already exists
	for _, tag := range tags {
		if tag.Value == value || tag.Name == name {
			return nil // Tag already exists, no need to add
		}
	}

	// Add new tag
	tags = append(tags, Tag{
		Name:     name,
		Value:    value,
		Selected: false,
	})

	// Write back to file
	updatedData, err := json.MarshalIndent(tags, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, updatedData, 0644)
}
