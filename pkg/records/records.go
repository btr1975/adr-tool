package records

import (
	"errors"
	"fmt"
	"github.com/btr1975/adr-tool/pkg/adr_templates"
	"os"
	"regexp"
)

type Status string

const (
	Proposed   Status = "proposed"
	Accepted   Status = "accepted"
	Rejected   Status = "rejected"
	Deprecated Status = "deprecated"
	Superseded Status = "superseded"
)

// FileExists checks if a file exists and is not a directory before we try using it to prevent further errors.
//
// Example:
//
//	FileExists("./0001-0001-my-title.md")
func FileExists(path string) (exists bool) {
	fileInfo, err := os.Stat(path)
	if err == nil {
		return !fileInfo.IsDir()
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}

// DirectoryExists checks if a directory exists and is not a file before we try using it to prevent further errors.
//
// Example:
//
//	DirectoryExists("./")
func DirectoryExists(path string) (exists bool) {
	directoryInfo, err := os.Stat(path)
	if err == nil {
		return directoryInfo.IsDir()
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}

// WriteNewADR writes a new ADR file to the filesystem.
//
// Example:
//
//	    thing := adr_templates.NewLongTemplate("My Title", "My Deciders", "My Statement", []string{"Option 1", "Option 2"})
//		err := records.WriteNewADR("./", thing)
func WriteNewADR(path string, template adr_templates.RenderTemplate) (err error) {
	if !DirectoryExists(path) {
		return fmt.Errorf("directory %s does not exist", path)
	}

	fullPath := fmt.Sprintf("%s/%s", path, template.GetFileName())

	if FileExists(fullPath) {
		return fmt.Errorf("file %s already exists", fullPath)
	}

	render, err := template.Render()

	if err != nil {
		return err
	}

	err = os.WriteFile(fullPath, []byte(render), 0644)

	return err
}

// GetADRs returns a list of ADRs found in the given directory.
//
// Example:
//
//	records.GetADRs("./")
func GetADRs(path string) (found []string, err error) {
	var foundADRS []string

	if !DirectoryExists(path) {
		return foundADRS, fmt.Errorf("directory %s does not exist", path)
	}

	directoryEntries, err := os.ReadDir(path)

	if err != nil {
		return foundADRS, err
	}

	regexADR, err := regexp.Compile(`\d{4}-\S+\.md$`)

	if err != nil {
		return foundADRS, err
	}

	for _, entry := range directoryEntries {
		if !entry.IsDir() {
			if regexADR.MatchString(entry.Name()) {
				foundADRS = append(foundADRS, entry.Name())
			}
		}
	}

	return foundADRS, nil
}
