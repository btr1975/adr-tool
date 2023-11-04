package main

import "github.com/btr1975/adr-tool/pkg/records"

func main() {
	err := records.ChangeADRStatus("./temp", "0001-my-title.md", records.Proposed)

	if err != nil {
		panic(err)
	}
}
