package util

import "regexp"

var suffixRegex = regexp.MustCompile(`^[\w\-.\\:]+$`)

func IsSuffixValid(s string) bool {
	return suffixRegex.MatchString(s)
}
