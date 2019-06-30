package main

import "regexp"

var (
	RawPatternsAsOne  = `(?m)^.*?/\..*?$|^.*?\.(lock)?$|^[^/]+$|^.*?\.\..*?$|^.*?[\000-\037\177 ~^:?*[]+.*?$|^\..*?$|^.*?/$|^.*?//.*?$|^.*?@\{.*?$|^@$|^.*?\\.*?$`
	RawPatternsAsList = []string{
		`(?m)^.*?/\..*?$`,
		`^.*?\.(lock)?$`,
		`^[^/]+$`,
		`^.*?\.\..*?$`,
		`^.*?[\000-\037\177 ~^:?*[]+.*?$`,
		`^\..*?$`,
		`^.*?/$`,
		`^.*?//.*?$`,
		`^.*?@\{.*?$`,
		`^@$`,
		`^.*?\\.*?$`,
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
	}
)

func BuildSingleRegexPattern() *regexp.Regexp {
	return regexp.MustCompile(RawPatternsAsOne)
}

func BuildListOfPatterns() []*regexp.Regexp {
	patterns := make([]*regexp.Regexp, len(RawPatternsAsList))
	for index, pattern := range RawPatternsAsList {
		patterns[index] = regexp.MustCompile(pattern)
	}
	return patterns
}

func MatchAgainstSinglePattern(single_pattern *regexp.Regexp, string_to_test string) bool {
	return single_pattern.MatchString(string_to_test)
}

func MatchAgainstListOfPatterns(list_of_patterns []*regexp.Regexp, string_to_test string) bool {
	for _, single_pattern := range list_of_patterns {
		if !MatchAgainstSinglePattern(single_pattern, string_to_test) {
			return False
		}
	}
	return true
}
