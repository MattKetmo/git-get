package main

import (
	"net/url"
	"regexp"
	"strings"
)

func parseUrl(gitUrl string) (string, string, error) {
	h := "" // hostname
	p := "" // path

	// Look if URL as a scheme
	matched, _ := regexp.MatchString("^[a-z]+://", gitUrl)
	if matched {
		u, err := url.Parse(gitUrl)
		if err != nil {
			return "", "", err
		}

		h = u.Hostname()
		p = u.EscapedPath()
	} else {
		r := regexp.MustCompile(`^(?:[\w\.\-]+@)?([\w\.]+):([\w\.~/-]+)$`)
		finds := r.FindStringSubmatch(gitUrl)
		if len(finds) == 3 {
			h = finds[1]
			p = finds[2]
		}
	}

	// Clean path
	p = strings.Trim(p, "/")
	p = strings.TrimSuffix(p, ".git")

	return h, p, nil
}
