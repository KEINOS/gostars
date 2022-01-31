package gostars

import (
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

// ============================================================================
//  Type: URLInfo
// ============================================================================

// URLInfo contains information about the parsed URL.
type URLInfo struct {
	RawURL string   // the original url
	Scheme string   // protocol
	Host   string   // host or host:port
	Path   []string // slice of directory path
}

// ============================================================================
//  Constructor
// ============================================================================

// NewURLInfo returns the initialized object of URLInfo from urlTarget.
func NewURLInfo(urlTarget string) (*URLInfo, error) {
	urlInfo := URLInfo{RawURL: urlTarget}

	if err := urlInfo.parse(); err != nil {
		return nil, err
	}

	return &urlInfo, nil
}

// ============================================================================
//  Methods
// ============================================================================

// IsRepoGitHub returns true if the host is GitHub.
func (u *URLInfo) IsRepoGitHub() bool {
	return u.Host == "github.com"
}

func (u *URLInfo) parse() error {
	parsed, err := url.Parse(u.RawURL)
	if err != nil {
		return errors.Wrap(err, "failed to parse given URL")
	}

	u.Scheme = parsed.Scheme
	u.Host = parsed.Host

	pathChunk := strings.Split(parsed.Path, "/")

	// Cleanup
	pathClean := []string{}

	for _, path := range pathChunk {
		if path != "" {
			pathClean = append(pathClean, path)
		}
	}

	u.Path = pathClean

	return nil
}

// String is an implementation of Stringer.
func (u *URLInfo) String() string {
	return u.RawURL
}
