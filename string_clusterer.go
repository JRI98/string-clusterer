package string_clusterer

import (
	"github.com/adrg/strutil/metrics"
)

type btreeNode struct {
	left, right *btreeNode
	values      []string
}

func traverse(node *btreeNode, groups [][]string) [][]string {
	if node == nil {
		return groups
	}

	groups = append(groups, node.values)

	groups = traverse(node.left, groups)
	groups = traverse(node.right, groups)

	return groups
}

// SimilarityMetric represents a metric for measuring the similarity between strings.
type SimilarityMetric interface {
	Compare(a, b string) float64
}

// Cluster groups a slice of strings according to a similarity metric and a threshold.
func Cluster(inputStrings []string, similarityMetric SimilarityMetric, threshold float64, iterations uint64) [][]string {
	if len(inputStrings) == 0 {
		return [][]string{}
	}

	result := make([][]string, len(inputStrings))
	for i, v := range inputStrings {
		result[i] = []string{v}
	}

	for range iterations {
		bTree := &btreeNode{nil, nil, result[0]}

		for _, cluster := range result[1:] {
			inputHead := cluster[0]
			node := bTree
			for {
				nodeHead := node.values[0]
				similarity := similarityMetric.Compare(inputHead, nodeHead)
				if similarity >= threshold {
					node.values = append(node.values, cluster...)
					break
				}

				if node.left == nil {
					node.left = &btreeNode{nil, nil, cluster}
					break
				}

				leftHead := node.left.values[0]
				leftSimilarity := similarityMetric.Compare(inputHead, leftHead)

				if node.right == nil {
					if leftSimilarity >= threshold {
						node.left.values = append(node.left.values, cluster...)
						break
					} else {
						node.right = &btreeNode{nil, nil, cluster}
						break
					}
				}

				rightHead := node.right.values[0]
				rightSimilarity := similarityMetric.Compare(inputHead, rightHead)

				if leftSimilarity >= rightSimilarity {
					if leftSimilarity >= threshold {
						node.left.values = append(node.left.values, cluster...)
						break
					}
					node = node.left
				} else {
					if rightSimilarity >= threshold {
						node.right.values = append(node.right.values, cluster...)
						break
					}
					node = node.right
				}
			}
		}

		result = traverse(bTree, make([][]string, 0))
	}

	return result
}

// NewHamming returns a new Hamming similarity metric.
func NewHamming(caseSensitive bool) SimilarityMetric {
	metric := metrics.NewHamming()
	metric.CaseSensitive = caseSensitive
	return metric
}

// NewJaccard returns a new Jaccard similarity metric.
func NewJaccard(caseSensitive bool) SimilarityMetric {
	metric := metrics.NewJaccard()
	metric.CaseSensitive = caseSensitive
	return metric
}

// NewJaro returns a new Jaro similarity metric.
func NewJaro(caseSensitive bool) SimilarityMetric {
	metric := metrics.NewJaro()
	metric.CaseSensitive = caseSensitive
	return metric
}

// NewJaroWinkler returns a new JaroWinkler similarity metric.
func NewJaroWinkler(caseSensitive bool) SimilarityMetric {
	metric := metrics.NewJaroWinkler()
	metric.CaseSensitive = caseSensitive
	return metric
}

// NewLevenshtein returns a new Levenshtein similarity metric.
func NewLevenshtein(caseSensitive bool) SimilarityMetric {
	metric := metrics.NewLevenshtein()
	metric.CaseSensitive = caseSensitive
	return metric
}

// NewOverlapCoefficient returns a new OverlapCoefficient similarity metric.
func NewOverlapCoefficient(caseSensitive bool) SimilarityMetric {
	metric := metrics.NewOverlapCoefficient()
	metric.CaseSensitive = caseSensitive
	return metric
}

// NewSmithWatermanGotoh returns a new SmithWatermanGotoh similarity metric.
func NewSmithWatermanGotoh(caseSensitive bool) SimilarityMetric {
	metric := metrics.NewSmithWatermanGotoh()
	metric.CaseSensitive = caseSensitive
	return metric
}

// NewSorensenDice returns a new SorensenDice similarity metric.
func NewSorensenDice(caseSensitive bool) SimilarityMetric {
	metric := metrics.NewSorensenDice()
	metric.CaseSensitive = caseSensitive
	return metric
}
