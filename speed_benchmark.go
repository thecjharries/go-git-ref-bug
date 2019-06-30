package main

import (
	"math/rand"
	"regexp"
	"testing"
	"time"
)

func SeedEnv() bool {
	rand.Seed(time.Now().UnixNano())
	return true
}

var (
	IsEnvSeeded                = SeedEnv()
	CompiledSlowSinglePattern  = BuildSingleRegexPattern(RawPatternsAsOne)
	CompiledSlowListPattern    = BuildListOfPatterns(RawPatternsAsList)
	CompiledTunedSinglePattern = BuildSingleRegexPattern(RawTunedPatternsAsOne)
	CompiledTunedListPattern   = BuildListOfPatterns(RawTunedPatternsAsList)
	NumberOfStringsToTest      = len(StringsToTest)
	MatchSuccess               bool
	SinglePatternSuccess       *regexp.Regexp
	ListPatternSuccess         []*regexp.Regexp
)

func benchmarkBuildList(raw_pattern_list []string, b *testing.B) {
	var pattern []*regexp.Regexp
	for n := 0; n < b.N; n++ {
		pattern = BuildListOfPatterns(raw_pattern_list)
	}
	ListPatternSuccess = pattern

}

func benchmarkBuildSingle(raw_pattern_single string, b *testing.B) {
	var pattern *regexp.Regexp
	for n := 0; n < b.N; n++ {
		pattern = BuildSingleRegexPattern(raw_pattern_single)
	}
	SinglePatternSuccess = pattern
}

func benchmarkMatchingList(compiled_pattern_list []*regexp.Regexp, b *testing.B) {
	var success bool
	for n := 0; n < b.N; n++ {
		success = MatchAgainstListOfPatterns(
			compiled_pattern_list,
			StringsToTest[rand.Intn(NumberOfStringsToTest)],
		)
	}
	MatchSuccess = success
}

func benchmarkMatchingSingle(compiled_pattern_single *regexp.Regexp, b *testing.B) {
	var success bool
	for n := 0; n < b.N; n++ {
		success = MatchAgainstSinglePattern(
			compiled_pattern_single,
			StringsToTest[rand.Intn(NumberOfStringsToTest)],
		)
	}
	MatchSuccess = success
}

func benchmarkEverythingList(raw_pattern_list []string, b *testing.B) {
	var success bool
	for n := 0; n < b.N; n++ {
		success = CompileAndMatchAllStringsListPattern(
			raw_pattern_list,
			StringsToTest,
		)
	}
	MatchSuccess = success
}

func benchmarkEverythingSingle(raw_pattern_single string, b *testing.B) {
	var success bool
	for n := 0; n < b.N; n++ {
		success = CompileAndMatchAllStringsSinglePattern(
			raw_pattern_single,
			StringsToTest,
		)
	}
	MatchSuccess = success
}

func BenchmarkBuildSingleSlowRegexPattern(b *testing.B) { benchmarkBuildSingle(RawPatternsAsOne, b) }

func BenchmarkBuildSlowListOfPatterns(b *testing.B) { benchmarkBuildList(RawPatternsAsList, b) }
func BenchmarkBuildSingleTunedRegexPattern(b *testing.B) {
	benchmarkBuildSingle(RawTunedPatternsAsOne, b)
}

func BenchmarkBuildTunedListOfPatterns(b *testing.B) { benchmarkBuildList(RawTunedPatternsAsList, b) }

func BenchmarkMatchAgainstSlowSinglePattern(b *testing.B) {
	benchmarkMatchingSingle(CompiledSlowSinglePattern, b)
}

func BenchmarkMatchAgainstSlowListOfPatterns(b *testing.B) {
	benchmarkMatchingList(CompiledSlowListPattern, b)
}

func BenchmarkMatchAgainstTunedSinglePattern(b *testing.B) {
	benchmarkMatchingSingle(CompiledTunedSinglePattern, b)
}

func BenchmarkMatchAgainstTunedListOfPatterns(b *testing.B) {
	benchmarkMatchingList(CompiledTunedListPattern, b)
}

func BenchmarkTheWholeProcessMatchAgainstSlowSinglePattern(b *testing.B) {
	benchmarkEverythingSingle(RawPatternsAsOne, b)
}

func BenchmarkTheWholeProcessMatchAgainstSlowListOfPatterns(b *testing.B) {
	benchmarkEverythingList(RawPatternsAsList, b)
}

func BenchmarkTheWholeProcessMatchAgainstTunedSinglePattern(b *testing.B) {
	benchmarkEverythingSingle(RawTunedPatternsAsOne, b)
}

func BenchmarkTheWholeProcessMatchAgainstTunedListOfPatterns(b *testing.B) {
	benchmarkEverythingList(RawTunedPatternsAsList, b)
}
