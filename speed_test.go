package main

import (
	"fmt"
	"regexp"

	"gopkg.in/check.v1"
)

type SpeedSuite struct{}

var _ = Suite(&SpeedSuite)

func singlePatternSuccesses(single_pattern *regexp.Regexp, strings_to_test []string, c check.C) {
	for _, string_to_test := range strings_to_test {
	}
}

func (s *SpeedSuite) TestAllSinglePatterns(c *check.C) {
	patterns_to_test = append(CompiledSlowListPattern, CompiledSlowSinglePattern, CompiledTunedSinglePattern, CompiledTunedListPattern...)
	for _, pattern := range patterns_to_test {
		for _, string_to_test := range StringsToTest {
			result := MatchAgainstSinglePattern(single_pattern, string_to_test)
			c.Assert(result, c.IsTrue, fmt.Sprintf("%s failed on %s", single_pattern.String(), string_to_test))
		}
	}
}

func (s *SpeedSuite) TestAllListPatterns(c *check.C) {
	for _, string_to_test := range StringsToTest {
		slow_result := MatchAgainstListOfPatterns(CompiledSlowListPattern, string_to_test)
		c.Assert(slow_result, c.IsTrue, fmt.Sprintf("The slower set failed on %s", string_to_test))
		tuned_result := MatchAgainstListOfPatterns(CompiledSlowListPattern, string_to_test)
		c.Assert(tuned_result, c.IsTrue, fmt.Sprintf("The tuned set failed on %s", string_to_test))
	}
}
