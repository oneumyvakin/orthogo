package orthogo

import "fmt"

func ExampleOrthogonate() {
	input := []Row{}
	input = append(input, Row{"browser": "Firefox", "os": "Windows"})
	input = append(input, Row{"browser": "Firefox", "os": "Windows"})
	input = append(input, Row{"browser": "Chrome", "os": "Windows"})
	input = append(input, Row{"browser": "Chrome", "os": "Windows"})
	output := Orthogonate(input)
	for i, item := range output {
		fmt.Printf("Case #%d for %s on %s\n", i, item["browser"], item["os"])
	}

	// Output:
	// Case #0 for Firefox on Windows
	// Case #1 for Chrome on Windows
}
