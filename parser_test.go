package main

import "testing"

func TestExtractPathFromGitUrl(t *testing.T) {
	tables := []struct {
		url      string
		hostname string
		path     string
	}{
		// Github
		{"git@github.com:MattKetmo/git-get", "github.com", "MattKetmo/git-get"},
		{"git@github.com:MattKetmo/git-get.git", "github.com", "MattKetmo/git-get"},
		{"https://github.com/MattKetmo/git-get", "github.com", "MattKetmo/git-get"},
		{"https://github.com/MattKetmo/git-get.git", "github.com", "MattKetmo/git-get"},

		// Custom
		{"git@git.example.org:MattKetmo/git-get", "git.example.org", "MattKetmo/git-get"},
		{"git@git.example.org:MattKetmo/git-get.git", "git.example.org", "MattKetmo/git-get"},
		{"ssh://git@git.example.org/MattKetmo/git-get", "git.example.org", "MattKetmo/git-get"},
		{"ssh://git@git.example.org/MattKetmo/git-get.git", "git.example.org", "MattKetmo/git-get"},
		{"https://mattketmo@git.example.org/MattKetmo/git-get.git", "git.example.org", "MattKetmo/git-get"},

		// Examples from https://stackoverflow.com/questions/31801271/what-are-the-supported-git-url-formats
		// SSH
		{"ssh://user@example.org:1337/path/to/repo.git/", "example.org", "path/to/repo"},
		{"ssh://user@example.org/path/to/repo.git/", "example.org", "path/to/repo"},
		{"ssh://example.org:1337/path/to/repo.git/", "example.org", "path/to/repo"},
		{"ssh://example.org/path/to/repo.git/", "example.org", "path/to/repo"},
		{"ssh://user@example.org/path/to/repo.git/", "example.org", "path/to/repo"},
		{"ssh://example.org/path/to/repo.git/", "example.org", "path/to/repo"},
		{"ssh://user@example.org/~user/path/to/repo.git/", "example.org", "~user/path/to/repo"},
		{"ssh://example.org/~user/path/to/repo.git/", "example.org", "~user/path/to/repo"},
		{"ssh://user@example.org/~/path/to/repo.git", "example.org", "~/path/to/repo"},
		{"ssh://example.org/~/path/to/repo.git", "example.org", "~/path/to/repo"},
		{"user@example.org:/path/to/repo.git/", "example.org", "path/to/repo"},
		{"example.org:/path/to/repo.git/", "example.org", "path/to/repo"},
		{"user@example.org:~user/path/to/repo.git/", "example.org", "~user/path/to/repo"},
		{"example.org:~user/path/to/repo.git/", "example.org", "~user/path/to/repo"},
		{"user@example.org:path/to/repo.git", "example.org", "path/to/repo"},
		{"example.org:path/to/repo.git", "example.org", "path/to/repo"},
		{"rsync://example.org/path/to/repo.git/", "example.org", "path/to/repo"},
		// Git
		{"git://example.org/path/to/repo.git/", "example.org", "path/to/repo"},
		{"git://example.org/~user/path/to/repo.git/", "example.org", "~user/path/to/repo"},
		// HTTP(S)
		{"http://example.org/path/to/repo.git/", "example.org", "path/to/repo"},
		{"https://example.org/path/to/repo.git/", "example.org", "path/to/repo"},
	}

	for _, table := range tables {
		h, p, err := parseUrl(table.url)
		if err != nil {
			t.Error(err)
		}
		if h != table.hostname || p != table.path {
			t.Errorf(
				"Parser for %s was incorrect, got: (%s, %s), want: (%s, %s).",
				table.url,
				h,
				p,
				table.hostname,
				table.path,
			)
		}
	}
}
