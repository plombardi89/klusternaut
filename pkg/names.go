package names

import (
	"math/rand"
	"strings"
	"fmt"
	"errors"
	"regexp"
)

var prefixRegex = regexp.MustCompile("(?i)[a-z][a-z0-9]{1,8}")
var suffixChars = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

func GenerateClusterName(rand rand.Rand, namePrefix string, region string) (name string, err error) {
	if prefixRegex.MatchString(namePrefix) == false {
		return "", errors.New(fmt.Sprintf("name prefix does not match expected pattern: %s", prefixRegex))
	}

	return fmt.Sprintf("%s-%s-%s", namePrefix, region, "abc"), nil
}

func AbbreviateZoneName(zone string) string {
	return strings.ToLower(zone)
}

