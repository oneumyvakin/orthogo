package orthogo

import "testing"

func TestOneColumnOneRow(t *testing.T) {
	input := []Row{}

	input = append(input, Row{"browser": "Firefox"})

	output := Orthogonate(input)
	t.Logf("Output: %#v", output)
	if len(output) != 1 {
		t.Errorf("output size != 1:\ninput:\n%#v\noutput:\n%#v", input, output)
	}
}

func TestOneColumnTwoRows(t *testing.T) {
	input := []Row{}

	input = append(input, Row{"browser": "Firefox"})
	input = append(input, Row{"browser": "Firefox"})

	output := Orthogonate(input)
	t.Logf("Output: %#v", output)
	if len(output) != 2 {
		t.Errorf("output size != 2:\ninput:\n%#v\noutput:\n%#v", input, output)
	}
}

func TestOneColumnThreeRows(t *testing.T) {
	input := []Row{}

	input = append(input, Row{"browser": "Firefox"})
	input = append(input, Row{"browser": "Firefox"})
	input = append(input, Row{"browser": "Firefox"})

	output := Orthogonate(input)
	t.Logf("Output: %#v", output)
	if len(output) != 3 {
		t.Errorf("output size != 2:\ninput:\n%#v\noutput:\n%#v", input, output)
	}
}

func TestOneColumn(t *testing.T) {
	input := []Row{}

	input = append(input, Row{"browser": "Firefox"})
	input = append(input, Row{"browser": "Firefox"})
	input = append(input, Row{"browser": "Firefox"})
	input = append(input, Row{"browser": "Firefox"})

	output := Orthogonate(input)
	t.Logf("Output: %#v", output)
	if len(output) != 4 {
		t.Errorf("output size != 4:\ninput:\n%#v\noutput:\n%#v", input, output)
	}
}

func TestTwoColumns(t *testing.T) {
	input := []Row{}

	input = append(input, Row{"browser": "Firefox", "os": "Windows"})
	input = append(input, Row{"browser": "Chrome", "os": "Windows"})
	input = append(input, Row{"browser": "Firefox", "os": "Windows"})
	input = append(input, Row{"browser": "Chrome", "os": "Windows"})

	output := Orthogonate(input)
	t.Logf("Output: %#v", output)
	if len(output) != 2 {
		t.Errorf("output size != 1:\ninput:\n%#v\noutput:\n%#v", input, output)
	}
}

func TestThreeColumns(t *testing.T) {
	input := []Row{}

	input = append(input, Row{"browser": "Firefox", "os": "Windows", "connection": "Fast"})
	input = append(input, Row{"browser": "Firefox", "os": "OSX", "connection": "Bad"})
	input = append(input, Row{"browser": "Firefox", "os": "Linux", "connection": "Slow"})
	input = append(input, Row{"browser": "Chrome", "os": "Windows", "connection": "Good"})

	output := Orthogonate(input)
	t.Logf("Output: %#v", output)
	if len(output) != 4 {
		t.Errorf("output size != 4:\ninput:\n%#v\noutput:\n%#v", input, output)
	}
}

func TestFourColumns(t *testing.T) {
	input := []Row{}

	input = append(input, Row{"browser": "Firefox", "os": "Windows", "arch": "x86", "connection": "Fast"})
	input = append(input, Row{"browser": "Firefox", "os": "Windows", "arch": "x86", "connection": "Slow"})
	input = append(input, Row{"browser": "Firefox", "os": "Windows", "arch": "x86", "connection": "Good"})
	input = append(input, Row{"browser": "Firefox", "os": "Windows", "arch": "x86", "connection": "Bad"})

	input = append(input, Row{"browser": "Firefox", "os": "Linux", "arch": "x86", "connection": "Fast"})
	input = append(input, Row{"browser": "Firefox", "os": "Linux", "arch": "x86", "connection": "Slow"})
	input = append(input, Row{"browser": "Firefox", "os": "Linux", "arch": "x86", "connection": "Good"})
	input = append(input, Row{"browser": "Firefox", "os": "Linux", "arch": "x86", "connection": "Bad"})

	input = append(input, Row{"browser": "Firefox", "os": "Windows", "arch": "x64", "connection": "Fast"})
	input = append(input, Row{"browser": "Firefox", "os": "Windows", "arch": "x64", "connection": "Slow"})
	input = append(input, Row{"browser": "Firefox", "os": "Windows", "arch": "x64", "connection": "Good"})
	input = append(input, Row{"browser": "Firefox", "os": "Windows", "arch": "x64", "connection": "Bad"})

	input = append(input, Row{"browser": "Firefox", "os": "Linux", "arch": "x64", "connection": "Fast"})
	input = append(input, Row{"browser": "Firefox", "os": "Linux", "arch": "x64", "connection": "Slow"})
	input = append(input, Row{"browser": "Firefox", "os": "Linux", "arch": "x64", "connection": "Good"})
	input = append(input, Row{"browser": "Firefox", "os": "Linux", "arch": "x64", "connection": "Bad"})

	input = append(input, Row{"browser": "Chrome", "os": "Windows", "arch": "x86", "connection": "Fast"})
	input = append(input, Row{"browser": "Chrome", "os": "Windows", "arch": "x86", "connection": "Slow"})
	input = append(input, Row{"browser": "Chrome", "os": "Windows", "arch": "x86", "connection": "Good"})
	input = append(input, Row{"browser": "Chrome", "os": "Windows", "arch": "x86", "connection": "Bad"})

	input = append(input, Row{"browser": "Chrome", "os": "Linux", "arch": "x86", "connection": "Fast"})
	input = append(input, Row{"browser": "Chrome", "os": "Linux", "arch": "x86", "connection": "Slow"})
	input = append(input, Row{"browser": "Chrome", "os": "Linux", "arch": "x86", "connection": "Good"})
	input = append(input, Row{"browser": "Chrome", "os": "Linux", "arch": "x86", "connection": "Bad"})

	input = append(input, Row{"browser": "Chrome", "os": "Windows", "arch": "x64", "connection": "Fast"})
	input = append(input, Row{"browser": "Chrome", "os": "Windows", "arch": "x64", "connection": "Slow"})
	input = append(input, Row{"browser": "Chrome", "os": "Windows", "arch": "x64", "connection": "Good"})
	input = append(input, Row{"browser": "Chrome", "os": "Windows", "arch": "x64", "connection": "Bad"})

	input = append(input, Row{"browser": "Chrome", "os": "Linux", "arch": "x64", "connection": "Fast"})
	input = append(input, Row{"browser": "Chrome", "os": "Linux", "arch": "x64", "connection": "Slow"})
	input = append(input, Row{"browser": "Chrome", "os": "Linux", "arch": "x64", "connection": "Good"})
	input = append(input, Row{"browser": "Chrome", "os": "Linux", "arch": "x64", "connection": "Bad"})

	output := Orthogonate(input)
	t.Logf("Output: %#v", output)
	if len(output) != 4 {
		t.Errorf("output size != 4:\ninput:\n%#v\noutput:\n%#v", input, output)
	}
}
