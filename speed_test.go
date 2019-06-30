package main

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) { check.TestingT(t) }

type SpeedSuite struct{}

var _ = check.Suite(&SpeedSuite{})

func (s *SpeedSuite) simplifyAllSinglePatternRuns(single_pattern *regexp.Regexp, c *check.C) {
	success := false
	fmt.Println(single_pattern.String())
	for _, string_to_test := range StringsToTest {
		fmt.Println(strings.Repeat(" ", 8), string_to_test)
		result := MatchAgainstSinglePattern(single_pattern, string_to_test)
		if result {
			success = true
			fmt.Println("success", single_pattern.String(), string_to_test)
		}
	}
	fmt.Println(single_pattern.String())
	c.Assert(success, check.Equals, true)

}

func (s *SpeedSuite) TestAllSinglePatterns(c *check.C) {

	for _, single_pattern := range CompiledSlowListPattern {
		s.simplifyAllSinglePatternRuns(single_pattern, c)
	}
	for _, single_pattern := range CompiledTunedListPattern {
		s.simplifyAllSinglePatternRuns(single_pattern, c)
	}

	s.simplifyAllSinglePatternRuns(CompiledSlowSinglePattern, c)

	s.simplifyAllSinglePatternRuns(CompiledTunedSinglePattern, c)

}

func (s *SpeedSuite) TestAllListPatterns(c *check.C) {
	for _, string_to_test := range StringsToTest {
		slow_result := MatchAgainstListOfPatterns(CompiledSlowListPattern, string_to_test)
		c.Assert(slow_result, check.Equals, false) // fmt.Sprintf("The slower set failed on %s", string_to_test)
		tuned_result := MatchAgainstListOfPatterns(CompiledSlowListPattern, string_to_test)
		c.Assert(tuned_result, check.Equals, false) // fmt.Sprintf("The tuned set failed on %s", string_to_test)
	}
}
