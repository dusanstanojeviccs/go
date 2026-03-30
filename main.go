package main

import (
	"fmt"
)

func main() {
	examples := []string{
		"0101000710009",
		"0710003730015",
		"1705978730032",
		"1505995800002",
		"invalid",
		"1234567890123",
	}

	for _, input := range examples {
		j, err := Parse(input)
		if err != nil {
			fmt.Printf("%-13s  ❌ %s\n", input, err)
			continue
		}

		fmt.Printf(
			"%-13s  ✅  Born: %s  Age: %d  Region: %s (%s)  Gender: %s  Adult: %v\n",
			input,
			j.GetDate().Format("2006-01-02"),
			j.GetAge(),
			j.RegionText,
			j.Country,
			j.Gender(),
			j.IsAdult(),
		)
	}
}
