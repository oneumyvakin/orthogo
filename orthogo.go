// Package generate orthogonal array for pair-wise / all-pairs software testing technique
package orthogo

// Row keeps table field name as key and field value
type Row map[string]string

// Orthogonate receive slice of Rows and returns orthogonal array
func Orthogonate(input []Row) (output []Row) {
	if len(input) < 2 {
		return input
	}

	keys := []string{}
	for k, _ := range input[0] {
		keys = append(keys, k)
	}
	keysLen := len(keys)
	if keysLen < 2 {
		return input
	}

	allPairs := map[string]bool{}
	for _, iItem := range input {
		pairs := []string{}
		for iKey, ivKey := range keys {
			yKey := iKey + 1
			if yKey >= keysLen {
				if keysLen > 2 {
					pairs = append(pairs, iItem[keys[0]]+iItem[keys[keysLen-1]])
				}
				break
			}
			yvKey := keys[yKey]
			pairs = append(pairs, iItem[ivKey]+iItem[yvKey])
		}

		uniquePair := true
		for _, pair := range pairs {
			if found := allPairs[pair]; found {
				uniquePair = false
			}
		}
		if uniquePair {
			output = append(output, iItem)
			for _, pair := range pairs {
				allPairs[pair] = true
			}
		}
	}

	return
}
