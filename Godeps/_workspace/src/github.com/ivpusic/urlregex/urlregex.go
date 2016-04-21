package urlregex

import (
	"errors"
	"regexp"
	"strings"
)

type UrlRegex struct {
	Regex *regexp.Regexp
}

// Method will collect named parameters values.
// If there are not matches, method will return error.
func (u *UrlRegex) Match(url string) (map[string]string, error) {
	params := u.Regex.SubexpNames()
	matches := u.Regex.FindAllStringSubmatch(url, -1)

	if matches == nil {
		return nil, errors.New("cannot match " + url)
	}

	result := map[string]string{}
	for i, n := range matches[0] {
		if len(params[i]) > 0 {
			result[params[i]] = n
		}
	}

	return result, nil
}

// Converting regex into UrlRegex object
// Function will build regex with named groups to be able to extract parameters later
func Pattern(pattern string) UrlRegex {
	parts := strings.Split(pattern, "/")
	regex := "^"

	for i, part := range parts {
		if len(part) > 0 {
			if i > 0 {
				regex += "\\/"
			}

			// do we have special character?
			switch part[0] {
			case ':':
				groupName := "(?P<" + part[1:] + ">"
				regex += groupName + ".[^\\/]*)"
			case '*':
				//support named wildcards
				if len(part) > 1 {
					groupName := "(?P<" + part[1:] + ">"
					regex += groupName + ".*)"
				} else {
					regex += ".*"
				}
				if part[len(part)-1] == '/' {
					regex += "\\/"
				}
			default:
				regex += regexp.QuoteMeta(part)
			}
		}
	}

	if pattern[len(pattern)-1] == '/' {
		regex += "\\/$"
	} else {
		regex += "$"
	}

	return UrlRegex{
		Regex: regexp.MustCompile(regex),
	}
}
