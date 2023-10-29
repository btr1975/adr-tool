package main

import (
	"fmt"
	"github.com/btr1975/adr-tool/pkg/records"
)

func main() {
	/*
		thing := adr_templates.NewLongTemplate("My Title", "My Deciders", "My Statement", []string{"Option 1", "Option 2"})

		err := records.WriteNewADR("./temp", thing)

		if err != nil {
			panic(err)
		}
	*/

	adrs, err := records.GetADRs("./temp")

	if err != nil {
		panic(err)
	}

	fmt.Printf(fmt.Sprintf("%v", adrs))
}
