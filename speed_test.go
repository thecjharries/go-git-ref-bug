package main

import (
	"regexp"
	"testing"
)

func singlePatternSuccesses(single_pattern *regexp.Regexp, strings_to_test []string, t *testing.T) {
	result := MatchAgainstSinglePattern(single_pattern, string_to_test)

}
