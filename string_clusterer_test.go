package string_clusterer

import (
	"maps"
	"slices"
	"strings"
	"testing"
)

func TestCluster(t *testing.T) {
	testcases := []struct {
		inputs       []string
		metric       SimilarityMetric
		threshold    float64
		iterations   uint64
		wantClusters int
	}{
		{[]string{"apple", "aple", "banana", "bananna", "orange", "ornge"}, NewJaroWinkler(false), 0.9, 1, 3},
		{[]string{"apple", "aple", "banana", "bananna", "orange", "ornge"}, NewJaroWinkler(false), 0, 1, 1},
		{[]string{"apple", "aple", "banana", "bananna", "orange", "ornge"}, NewJaroWinkler(false), 1, 1, 6},
	}

	for _, tc := range testcases {
		result := Cluster(tc.inputs, tc.metric, tc.threshold, tc.iterations)

		if len(result) != tc.wantClusters {
			t.Fatalf("%q clusters, want %q clusters", result, tc.wantClusters)
		}

		inputsMap := make(map[string]struct{}, len(tc.inputs))
		for _, v := range tc.inputs {
			inputsMap[v] = struct{}{}
		}

		resultMap := make(map[string]struct{}, len(result))
		for _, r := range result {
			for _, v := range r {
				resultMap[v] = struct{}{}
			}
		}

		if !maps.Equal(inputsMap, resultMap) {
			t.Fatalf("%q, want %q", inputsMap, resultMap)
		}
	}
}

func FuzzCluster(f *testing.F) {
	metric := NewJaroWinkler(false)

	f.Fuzz(func(t *testing.T, inputString string, threshold float64, iterations uint64) {
		inputs := strings.Fields(inputString)
		if len(inputs) == 0 {
			return
		}

		result := Cluster(inputs, metric, threshold, iterations)

		if len(result) == 0 {
			t.Fatal("result len is never 0")
		}

		if len(result) > len(inputs) {
			t.Fatal("result len is always greater than len inputs")
		}

	outer:
		for _, input := range inputs {
			for _, r := range result {
				if slices.Contains(r, input) {
					continue outer
				}
			}

			t.Fatalf("input %q missing in result %q", input, result)
		}
	})
}

func BenchmarkCluster(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Cluster([]string{"apple", "aple", "banana", "bananna", "orange", "ornge"}, NewJaroWinkler(false), 0.9, 1)
	}
}
