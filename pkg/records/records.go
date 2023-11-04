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

	found, err := GetADRs(path)

	if err != nil {
		return err
	}

	nextADR := GetADRNextNumber(found)

	fullFileName := fmt.Sprintf("%s-%s", GetADRNumberFromInteger(nextADR), template.GetFileName())

	fullPath := fmt.Sprintf("%s/%s", path, fullFileName)

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

// GetADRNumberFromString returns the ADR number from the given string.
//
// Example:
//
//	records.GetADRNumberFromString("0001-0001-my-title.md")
func GetADRNumberFromString(adr string) (number int, err error) {
	regexADRNumber, err := regexp.Compile(`^\d{4}`)

	if err != nil {
		return number, err
	}

	adrNumberString := regexADRNumber.FindString(adr)

	if adrNumberString == "" {
		return number, fmt.Errorf("could not find ADR number in %s", adr)
	}

	adrNumber, err := fmt.Sscanf(adrNumberString, "%d", &number)

	if err != nil {
		return number, err
	}

	if adrNumber != 1 {
		return number, fmt.Errorf("could not find ADR number in %s", adr)
	}

	return number, nil
}

// GetADRNumberFromInteger returns the ADR number from the given integer.
//
// Example:
//
//	records.GetADRNumberFromInteger(1)
func GetADRNumberFromInteger(number int) (adr string) {
	return fmt.Sprintf("%04d", number)
}

// GetADRNextNumber returns the next ADR number from the given list of ADRs.
//
// Example:
//
//	records.GetADRNextNumber([]string{"0001-0001-my-title.md", "0001-0002-my-title.md"})
func GetADRNextNumber(found []string) (number int) {
	var highest int

	for _, adr := range found {
		adrNumber, err := GetADRNumberFromString(adr)

		if err != nil {
			continue
		}

		if adrNumber > highest {
			highest = adrNumber
		}
	}

	return highest + 1
}

// ChangeADRStatus changes the status of the given ADR.
//
// Example:
//
//	records.ChangeADRStatus("./", "0001-0001-my-title.md", records.Accepted)
func ChangeADRStatus(path string, adr string, status Status) (err error) {
	if !DirectoryExists(path) {
		return fmt.Errorf("directory %s does not exist", path)
	}

	if !FileExists(fmt.Sprintf("%s/%s", path, adr)) {
		return fmt.Errorf("ADR %s does not exist", adr)
	}

	fullPath := fmt.Sprintf("%s/%s", path, adr)

	adrFile, err := os.ReadFile(fullPath)

	if err != nil {
		return err
	}

	regexStatus, err := regexp.Compile(`. Status: \S+`)

	if err != nil {
		return err
	}

	adrFile = regexStatus.ReplaceAll(adrFile, []byte(fmt.Sprintf("* Status: %s", status)))

	err = os.WriteFile(fullPath, adrFile, 0644)

	if err != nil {
		return err
	}

	return nil
}
