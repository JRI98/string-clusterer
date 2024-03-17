# string-clusterer

Go package for clustering strings. Given a slice of strings, a similarity metric and a threshold, the input strings are clustered according to their similarity.

Similarity metrics are provided by <https://github.com/adrg/strutil>.

## Installation

```shell
go get github.com/JRI98/string-clusterer
```

## Example

```golang
input := []string{"apple", "aple", "banana", "bananna", "orange", "ornge"}
result := Cluster(input, NewJaroWinkler(false), 0.9, 1)
fmt.Println(result) // [[apple aple] [banana bananna] [orange ornge]]
```

## Available Similarity Metrics

```golang
NewHamming(caseSensitive bool)
NewJaccard(caseSensitive bool)
NewJaro(caseSensitive bool)
NewJaroWinkler(caseSensitive bool)
NewLevenshtein(caseSensitive bool)
NewOverlapCoefficient(caseSensitive bool)
NewSmithWatermanGotoh(caseSensitive bool)
NewSorensenDice(caseSensitive bool)
```

## Repository Maintenance

### Run Tests

```go test```

### Run Benchmarks

```go test -bench=. -run=^#```

### Run Fuzzing

```go test -fuzz=FuzzCluster -run=^#```
