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
	for _, row := range input {
		pairs := []string{}
		for iKey, ivKey := range keys {
			yKey := iKey + 1
			if yKey >= keysLen {
				if keysLen > 2 {
					pair := keys[0]+"|||"+row[keys[0]]+"|||"+keys[keysLen-1]+"|||"+row[keys[keysLen-1]]
					pairs = append(pairs, pair)
				}
				break
			}
			yvKey := keys[yKey]
			pair := ivKey+"|||"+row[ivKey]+"|||"+yvKey+"|||"+row[yvKey]
			pairs = append(pairs, pair)
		}

		hasUniquePair := false
		for _, pair := range pairs {
			if found := allPairs[pair]; !found {
				hasUniquePair = true
			}
		}
		if hasUniquePair {
			output = append(output, row)
			for _, pair := range pairs {
				allPairs[pair] = true
			}
		}
	}

	return
}

// Combines all possible fields values to Rows
func AllCombinations(input map[string][]string) (output []Row) {
	deleteDups := func(result []Row) []Row {
		hashes := map[string]bool{}
		fieldNames := []string{}
		for fieldName := range input {
			fieldNames = append(fieldNames, fieldName)
		}
		cleanedResult := []Row{}
		for _, row := range result {
			hash := ""
			for _, fieldName := range fieldNames {
				hash += fieldName + "|||" + row[fieldName] + "|||"
			}
			if _, found := hashes[hash]; !found {
				cleanedResult = append(cleanedResult, row)
				hashes[hash] = true
			}
		}
		return cleanedResult
	}

	addRow := func(resRow Row, fieldName string, val string) Row {
		if _, ok := resRow[fieldName]; !ok {
			resRow[fieldName] = val
			return resRow
		}
		if rVal, ok := resRow[fieldName]; ok && rVal == val {
			return nil
		}
		if rVal, ok := resRow[fieldName]; ok && rVal != val {
			newRow := Row{}
			for resRowFieldName, resRowFieldVal := range resRow {
				newRow[resRowFieldName] = resRowFieldVal
			}
			newRow[fieldName] = val
			return newRow
		}
		return nil
	}

	result := []Row{}
	for fieldName, fieldVals := range input {
		for _, val := range fieldVals {
			if len(result) == 0 {
				newRow := Row{}
				newRow[fieldName] = val
				result = append(result, newRow)
			}
			for _, resRow := range result {
				if newRow := addRow(resRow, fieldName, val); newRow != nil {
					result = append(result, newRow)
				}
			}

			result = deleteDups(result)
		}
	}

	return deleteDups(result)
}
