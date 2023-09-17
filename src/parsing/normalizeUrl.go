package parsing

import (
	"net/url"
	"strings"
)

// normalizeURL takes a string representing a URL and returns a normalized version of the URL.
//
// It accepts a single parameter: link (string) - the URL to be normalized.
// It returns two values: string - the normalized URL, and error - any error that occurred during normalization.
func normalizeURL(link string) (string, error) {
	u, err := url.Parse(link)
	if err != nil {
		return "", err
	}

	u.Scheme = strings.ToLower(u.Scheme)
	u.Host = strings.ToLower(u.Host)
	u.Fragment = ""                 // Ignore fragment
	u.RawQuery = u.Query().Encode() // Alphabetize the query parameters

	return u.String(), nil
}
