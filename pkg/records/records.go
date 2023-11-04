package records

import (
	"errors"
	"fmt"
	"github.com/btr1975/adr-tool/pkg/adr_templates"
	"os"
	"regexp"
	"strings"
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
//		fileName, err := records.WriteNewADR("./", thing)
func WriteNewADR(path string, template adr_templates.RenderTemplate) (fileName string, err error) {
	if !DirectoryExists(path) {
		return "", fmt.Errorf("directory %s does not exist", path)
	}

	found, err := GetADRs(path)

	if err != nil {
		return "", err
	}

	nextADR := GetADRNextNumber(found)

	fullFileName := fmt.Sprintf("%s-%s", GetADRNumberFromInteger(nextADR), template.GetFileName())

	fullPath := fmt.Sprintf("%s/%s", path, fullFileName)

	if FileExists(fullPath) {
		return "", fmt.Errorf("file %s already exists", fullPath)
	}

	render, err := template.Render()

	if err != nil {
		return "", err
	}

	err = os.WriteFile(fullPath, []byte(render), 0644)

	return fullFileName, err
}

// SupersedeADR supersede the given ADR with a new ADR.
//
// Example:
//
//	    thing := adr_templates.NewLongTemplate("My Title", "My Deciders", "My Statement", []string{"Option 1", "Option 2"})
//		err := records.SupersedeADR("./", thing, "0001-my-title.md")
func SupersedeADR(path string, template adr_templates.RenderTemplate, adr string) (fileName string, err error) {
	if !DirectoryExists(path) {
		return "", fmt.Errorf("directory %s does not exist", path)
	}

	err = ChangeADRStatus(path, adr, Superseded, true)

	if err != nil {
		return "", err
	}

	fileName, err = WriteNewADR(path, template)

	if err != nil {
		return "", err
	}

	err = AppendSupersededBy(path, adr, fileName)

	if err != nil {
		return fileName, err
	}

	return fileName, nil
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
//	records.ChangeADRStatus("./", "0001-my-title.md", records.Accepted, false)
func ChangeADRStatus(path string, adr string, status Status, supersede bool) (err error) {
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

	if !supersede {
		if strings.Contains(regexStatus.FindString(string(adrFile)), string(Accepted)) {
			return fmt.Errorf("ADR %s is accepted can not change status", adr)
		}
	}

	adrFile = regexStatus.ReplaceAll(adrFile, []byte(fmt.Sprintf("* Status: %s", status)))

	err = os.WriteFile(fullPath, adrFile, 0644)

	if err != nil {
		return err
	}

	return nil
}

// AppendToFile appends the given string to the given file.
//
// Example:
//
//	records.AppendToFile("./0001-my-title.md", "My Append")
func AppendToFile(fullPath string, append string) (err error) {
	if !FileExists(fmt.Sprintf("%s", fullPath)) {
		return fmt.Errorf("file %s does not exist", fullPath)
	}

	file, err := os.OpenFile(fullPath, os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()

	if err != nil {
		return err
	}

	_, err = file.WriteString(fmt.Sprintf("%s", append))

	if err != nil {
		return err
	}

	return nil
}

// AppendSupersededBy appends the superseded by and supersedes to the given ADRs.
//
// Example:
//
//	records.AppendSupersededBy("./", "0001-my-title.md", "0002-my-title.md")
func AppendSupersededBy(path string, adr string, supersededBy string) (err error) {
	supersededFullPath := fmt.Sprintf("%s/%s", path, adr)
	supersededAppend := fmt.Sprintf("* [Superseded By: %s](./%s)\n", supersededBy, supersededBy)
	newFullPath := fmt.Sprintf("%s/%s", path, supersededBy)
	newAppend := fmt.Sprintf("* [Supersedes: %s](./%s)\n", adr, adr)

	err = AppendToFile(supersededFullPath, supersededAppend)

	if err != nil {
		return err
	}

	err = AppendToFile(newFullPath, newAppend)

	return nil
}
