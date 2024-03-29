package main

import (
	"regexp"
)

var (
	RawPatternsAsOne  = `^.*?/\..*?$|^.*?\.(lock)?$|^[^/]+$|^.*?\.\..*?$|^.*?[\000-\037\177 ~^:?*[]+.*?$|^\..*?$|^.*?/$|^.*?//.*?$|^.*?@\{.*?$|^@$|^.*?\\.*?$`
	RawPatternsAsList = []string{
		`^@$`,
		`^\..*?$`,
		`^.*?\.(lock)?$`,
		`^.*?/$`,
		`^.*?/\..*?$`,
		`^.*?\.\..*?$`,
		`^.*?[\000-\037\177 ~^:?*[]+.*?$`,
		`^.*?//.*?$`,
		`^.*?@\{.*?$`,
		`^.*?\\.*?$`,
		`^[^/]+$`,
	}
	RawTunedPatternsAsOne  = `^(@|\.)|(.(lock)?|/)$|/(\.|/)|^[^/]+$|([\000-\037\177 ~^:?*\\[]+|\.\.|@\{)`
	RawTunedPatternsAsList = []string{
		`@{`,
		`^(@|\.)`,
		`(.(lock)?|/)$`,
		`/(\.|/)`,
		`^[^/]+$`,
		`[\000-\037\177 ~^:?*[]+`,
		`\.\.`,
	}
	StringsToTest = []string{
		`not/.allowed`,
		`this/ends/in.lock`,
		`hasnopath`,
		`has/to..do/dots`,
		`.starts/with/a/dot`,
		`ends/with/`,
		`has//slash`,
		`ends/with/dot.`,
		`has/con@{trol/chars`,
		`has/a\.bad/setup`,
		`@`,
		"\a",
	}
)

func BuildSingleRegexPattern(raw_single_pattern string) *regexp.Regexp {
	return regexp.MustCompilePOSIX(raw_single_pattern)
}

func BuildListOfPatterns(raw_list_of_patterns []string) []*regexp.Regexp {
	patterns := make([]*regexp.Regexp, len(RawPatternsAsList))
	for index, pattern := range raw_list_of_patterns {
		patterns[index] = regexp.MustCompilePOSIX(pattern)
	}
	return patterns
}

func MatchAgainstSinglePattern(single_pattern *regexp.Regexp, string_to_test string) bool {
	result := single_pattern.Match([]byte(string_to_test))
	return result
}

func MatchAgainstListOfPatterns(list_of_patterns []*regexp.Regexp, string_to_test string) bool {
	for _, single_pattern := range list_of_patterns {
		if !MatchAgainstSinglePattern(single_pattern, string_to_test) {
			return false
		}
	}
	return true
}

func CompileAndMatchAllStringsSinglePattern(raw_single_pattern string, strings_to_test []string) bool {
	compiled_pattern := BuildSingleRegexPattern(raw_single_pattern)
	success := true
	for _, untested_string := range strings_to_test {
		if !MatchAgainstSinglePattern(compiled_pattern, untested_string) {
			success = false
		}
	}
	return success
}

func CompileAndMatchAllStringsListPattern(raw_list_of_patterns []string, strings_to_test []string) bool {
	compiled_patterns := BuildListOfPatterns(raw_list_of_patterns)
	success := true
	for _, untested_string := range strings_to_test {
		if !MatchAgainstListOfPatterns(compiled_patterns, untested_string) {
			success = false
		}
	}
	return success
}
