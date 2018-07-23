package orthogo

import (
	"testing"
	"encoding/json"
)

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
	mustLen := 15
	if len(output) != mustLen {
		t.Errorf("output size %d != %d:\n%#v", len(output), mustLen, output)
	}
}

func TestAllCombinations(t *testing.T) {
	input := map[string][]string{
		"RoomCount":  {"1", "2", "3", "4"},
		"AdultCount": {"1", "2", "3", "4"},
		"KidAge1":    {"None", "-", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17"},
		"KidAge2":    {"None", "-", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17"},
		"KidAge3":    {"None", "-", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17"},
		"KidAge4":    {"None", "-", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17"},
	}
	output := AllCombinations(input)

	mustLen := 2560000
	if len(output) != mustLen {
		t.Errorf("output size %d != %d", len(output), mustLen)
	}
}

func TestOrthogonateAllCombinations(t *testing.T) {
	input := map[string][]string{
		"RoomCount":  {"1", "2", "3", "4"},
		"AdultCount": {"1", "2", "3", "4"},
		"KidAge1":    {"None", "-", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17"},
		"KidAge2":    {"None", "-", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17"},
		"KidAge3":    {"None", "-", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17"},
		"KidAge4":    {"None", "-", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17"},
	}
	allCombs := AllCombinations(input)

	output := Orthogonate(allCombs)
	mustLen := 1289
	if len(output) != mustLen {
		j, _ := json.Marshal(output)
		t.Log(string(j))
		t.Errorf("output size %d != %d", len(output), mustLen)
	}
}

func TestOrthogonate2AllCombinations(t *testing.T) {
	input := map[string][]string{
		"browser":  {"Firefox", "Chrome"},
		"os": {"Windows", "Linux"},
		"arch":    {"x86", "x64"},
		"connection":    {"Fast", "Slow", "Good", "Bad"},
	}
	allCombs := AllCombinations(input)

	output := Orthogonate(allCombs)
	mustLen := 15
	if len(output) != mustLen {
		j, _ := json.Marshal(output)
		jAllCombs, _ := json.Marshal(allCombs)
		t.Logf("AllCombinations: %d \n %s \n", len(allCombs), string(jAllCombs))
		t.Errorf("output size %d != %d:\n%#v", len(output), mustLen, string(j))
	}
}
