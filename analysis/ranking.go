package analysis

import "sort"

func RankResults(results []string) []string {

	sort.Slice(results, func(i, j int) bool {
		return len(results[i]) > len(results[j])
	})

	return results
}
