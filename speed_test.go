package main

import "testing"

var (
	CompiledSlowSinglePattern  = BuildSingleRegexPattern(RawPatternsAsOne)
	CompiledSlowListPattern    = BuildListOfPatterns(RawPatternsAsList)
	CompiledTunedSinglePattern = BuildSingleRegexPattern(RawTunedPatternsAsOne)
	CompiledTunedListPattern   = BuildListOfPatterns(RawTunedPatternsAsList)
	NumberOfStringsToTest      = len(StringsToTest)
)

func BenchmarkBuildSingleSlowRegexPattern(b *testing.B) {
	for n := 0; n < b.N; n++ {
		BuildSingleRegexPattern(RawPatternsAsOne)
	}
}

func BenchmarkBuildSlowListOfPatterns(b *testing.B) {
	for n := 0; n < b.N; n++ {
		BuildListOfPatterns(RawPatternsAsList)
	}
}
func BenchmarkBuildSingleTunedRegexPattern(b *testing.B) {
	for n := 0; n < b.N; n++ {
		BuildSingleRegexPattern(RawTunedPatternsAsOne)
	}
}

func BenchmarkBuildTunedListOfPatterns(b *testing.B) {
	for n := 0; n < b.N; n++ {
		BuildListOfPatterns(RawTunedPatternsAsList)
	}
}

func BenchmarkMatchAgainstSlowSinglePattern(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MatchAgainstSinglePattern(
			CompiledSlowSinglePattern,
			StringsToTest[n%NumberOfStringsToTest],
		)
	}
}

func BenchmarkMatchAgainstSlowListOfPatterns(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MatchAgainstListOfPatterns(
			CompiledTunedListPattern,
			StringsToTest[n%NumberOfStringsToTest],
		)
	}
}

func BenchmarkMatchAgainstTunedSinglePattern(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MatchAgainstSinglePattern(
			CompiledSlowSinglePattern,
			StringsToTest[n%NumberOfStringsToTest],
		)
	}
}

func BenchmarkMatchAgainstTunedListOfPatterns(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MatchAgainstListOfPatterns(
			CompiledTunedListPattern,
			StringsToTest[n%NumberOfStringsToTest],
		)
	}
}

func BenchmarkTheWholeProcessMatchAgainstSlowSinglePattern(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CompileAndMatchAllStringsSinglePattern(RawPatternsAsOne, StringsToTest)
	}
}

func BenchmarkTheWholeProcessMatchAgainstSlowListOfPatterns(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CompileAndMatchAllStringsListPattern(RawPatternsAsList, StringsToTest)
	}
}

func BenchmarkTheWholeProcessMatchAgainstTunedSinglePattern(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CompileAndMatchAllStringsSinglePattern(RawTunedPatternsAsOne, StringsToTest)
	}
}

func BenchmarkTheWholeProcessMatchAgainstTunedListOfPatterns(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CompileAndMatchAllStringsListPattern(RawTunedPatternsAsList, StringsToTest)
	}
}
