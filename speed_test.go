package main

import "testing"

func BenchmarkBuildSingleRegexPattern(b *testing.B) {
	for n := 0; n < b.N; n++ {
		BuildSingleRegexPattern()
	}
}

func BenchmarkBuildListOfPatterns(b *testing.B) {
	for n := 0; n < b.N; n++ {
		BuildListOfPatterns()
	}
}

var (
	CompiledSinglePattern = BuildSingleRegexPattern()
	CompiledListPattern   = BuildListOfPatterns()
	NumberOfStringsToTest = len(StringsToTest)
)

func BenchmarkMatchAgainstSinglePattern(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MatchAgainstSinglePattern(
			CompiledSinglePattern,
			StringsToTest[n%NumberOfStringsToTest],
		)
	}
}

func BenchmarkMatchAgainstListOfPatterns(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MatchAgainstListOfPatterns(
			CompiledSinglePattern,
			StringsToTest[n%NumberOfStringsToTest],
		)
	}
}
