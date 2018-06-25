package main

import (
	"fmt"
	"math"
)

// SymbolToNumber - Map used for calculation
var SymbolToNumber = map[string]int{
	"A": 0,
	"C": 1,
	"G": 2,
	"T": 3,
}

// NumberToSymbol - Map used for calculation
var NumberToSymbol = map[int]string{
	0: "A",
	1: "C",
	2: "G",
	3: "T",
}

// SymbolComplement - Reverse Complement for nucleotides
var SymbolComplement = map[string]string{
	"A": "T",
	"T": "A",
	"G": "C",
	"C": "G",
}

func main() {
	fmt.Println(len("GAGCCACCGCGATA"))
	fmt.Println(CalculateSkew("GAGCCACCGCGATA"))
}

func CalculateSkew(genome string) []int {
	var skew []int
	for i := 0; i < len(genome)+1; i++ {
		skew = append(skew, 0)
	}
	for i := 0; i < len(genome); i++ {
		if genome[i:i+1] == "G" {
			skew[i+1] = skew[i] + 1
		} else if genome[i:i+1] == "C" {
			skew[i+1] = skew[i] - 1
		} else {
			skew[i+1] = skew[i]
		}
	}
	return skew
}

func ComputingFrequencies(text string, k int) []int {
	combinations := math.Pow(4, float64(k))
	var frequencyArray []int
	for i := 0; i < int(combinations); i++ {
		frequencyArray = append(frequencyArray, 0)
	}
	for x := 0; x <= len(text)-k; x++ {
		pattern := text[x:(x + k)]
		number := PatternToNumber(pattern)
		frequencyArray[number] = frequencyArray[number] + 1
	}
	return frequencyArray
}

func PatternToNumber(pattern string) int {
	if len(pattern) == 0 {
		return 0
	}
	lastIndex := (len(pattern) - 1)
	lastCharacter := pattern[lastIndex:]
	return 4*PatternToNumber(pattern[:lastIndex]) + SymbolToNumber[lastCharacter]
}

func NumberToPattern(index int, k int) string {
	if k == 1 {
		return NumberToSymbol[index]
	}
	prefixIndex := index / 4
	r := index % 4
	prefixPattern := NumberToPattern(prefixIndex, k-1)
	return prefixPattern + NumberToSymbol[r]
}

// PatternCount - calculates how many times a pattern appears in a text
func PatternCount(text, pattern string) int {
	count := 0
	for i := 0; i <= len(text)-len(pattern); i++ {
		window := text[i : len(pattern)+i]
		if window == pattern {
			count++
		}
	}
	return count
}

func FasterFrequentWords(text string, k int) []string {
	var frequentPatterns []string
	frequencyArray := ComputingFrequencies(text, k)
	maxCount := GetMaximum(frequencyArray)
	for i := 0; i < int(math.Pow(4, float64(k))); i++ {
		if frequencyArray[i] == maxCount {
			pattern := NumberToPattern(i, k)
			frequentPatterns = append(frequentPatterns, pattern)
		}
	}
	return frequentPatterns
}

func ClumpFinding(genome string, k, L, t int) []string {
	var frequentPatterns []string
	var clump []int
	combinations := int(math.Pow(4, float64(k))) - 1
	//initiate clumps
	for i := 0; i < combinations; i++ {
		clump = append(clump, 0)
	}
	text := genome[0:L]
	frequencyArray := ComputingFrequencies(text, k)
	for i := 0; i < combinations; i++ {
		if frequencyArray[i] >= t {
			clump[i] = 1
		}
	}
	for i := 1; i < len(genome)-L; i++ {
		firstPattern := genome[i-1 : (i-1)+k]
		index := PatternToNumber(firstPattern)
		frequencyArray[index] = frequencyArray[index] - 1
		lastPattern := genome[i+L-k : (i+L-k)+k]
		index = PatternToNumber(lastPattern)
		frequencyArray[index] = frequencyArray[index] + 1
		if frequencyArray[index] >= t {
			clump[index] = 1
		}
	}
	for i := 0; i < combinations; i++ {
		if clump[i] == 1 {
			pattern := NumberToPattern(i, k)
			frequentPatterns = append(frequentPatterns, pattern)
		}
	}
	return frequentPatterns
}

// GetMaximum - Gets the maximum occurrence of an integer in an array
func GetMaximum(array []int) int {
	max := 0
	for _, v := range array {
		if v > max {
			max = v
		}
	}
	return max
}

func FrequentWords(text string, k int) []string {
	var maxOccurrences int = 0
	sequences := []string{}
	for i := 0; i <= len(text)-k; i++ {
		var pattern = text[i : k+i]
		if Index(sequences, pattern) == -1 {
			occurrences := PatternCount(text, pattern)
			if occurrences == maxOccurrences {
				sequences = append(sequences, pattern)
			} else if occurrences > maxOccurrences {
				maxOccurrences = occurrences
				sequences = []string{}
				sequences = append(sequences, pattern)
			}
		}
	}
	return sequences
}

func ReverseComplement(pattern string) string {
	var reversecomp string = ""
	for i := len(pattern) - 1; i >= 0; i-- {
		reversecomp += GetComplement(pattern[i : i+1])
	}
	return reversecomp
}

func GetComplement(input string) string {
	return SymbolComplement[input]
}

func FindStartingPositions(input string, pattern string) []int {
	var positions []int
	for i := 0; i <= (len(input) - len(pattern)); i++ {
		if input[i:len(pattern)+i] == pattern {
			positions = append(positions, i)
		}
	}
	return positions
}

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}
