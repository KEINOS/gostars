package gostars

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"github.com/russross/blackfriday"
)

// ============================================================================
//  Package Functions
// ============================================================================

// CoolDown is a sleep function to avoid a large number of requests to each API.
//
// It is currently forced to 1 second.
func CoolDown() {
	// Ref:
	//   Unauthenticated request: 60 req/hour ≅ 1 req/min
	//   Authenticated request: 5,000 req/hour ≅ 1.4 req/sec
	sleepSec := 1

	time.Sleep(time.Duration(sleepSec*1000) * time.Millisecond)
}

// GetAttractionGravity returns the distance from the point 0 to the point of
// "points" dimensions.
//
// Each point should be the value of comparison. Such as number of stars, number
// of forks, etc.
func GetAttractionGravity(points ...int) int {
	if len(points) == 0 {
		return 0
	}

	sumPoints := 0

	for _, point := range points {
		sumPoints += point * point
	}

	return int(math.Sqrt(float64(sumPoints)))
}

// GetContentURL returns the content of a given URL.
//
// To avoid a large number of requests to the target server, it sleeps for about
// one second.
func GetContentURL(urlTarget string) ([]byte, error) {
	CoolDown()

	urlParsed, err := url.Parse(urlTarget)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse URL before request")
	}

	response, err := http.Get(urlParsed.String())
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch contents from the URL")
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.Errorf("failed to featch. returned status: %v",
			response.StatusCode)
	}

	buf := new(bytes.Buffer)

	// Copy data from the response to the buffer
	if _, err = IOCopy(buf, response.Body); err != nil {
		return nil, errors.Wrap(err, "failed to copy response data")
	}

	return buf.Bytes(), nil
}

// GetURLGitHub will return the URL of the GitHub repository if urlOrigin matches
// the alias list.
func GetURLGitHub(urlOrigin string) string {
	urlGitHub, ok := URLAliases[urlOrigin]
	if !ok {
		return urlOrigin
	}

	return urlGitHub
}

// Hash256 returns the SHA2-256 (FIPS 180-4) hashed hex string from input.
//
// This function is not suitable for large files, as it copies all bytes into
// memory.
func Hash256(input []byte) string {
	return fmt.Sprintf("%x", sha256.Sum256(input))
}

// NewQuery returns a query object that processes HTML documents in a simple,
// jQuery-like manner. Powered by GoQuery.
func NewQuery(sourceHTML []byte) (*goquery.Document, error) {
	buf := bytes.NewBuffer(sourceHTML)

	return goquery.NewDocumentFromReader(buf)
}

// ParseMarkdownToHTML parses the markdown content to HTML.
func ParseMarkdownToHTML(markdown []byte) (html string) {
	return fmt.Sprintf("<body>%s</body>", blackfriday.MarkdownCommon(markdown))
}

// PrettyFormatJSON is a formatter for printing objects in a pretty way.
func PrettyFormatJSON(v interface{}) (string, error) {
	byteJSON, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", errors.Wrap(err, "failed to marshal")
	}

	return string(byteJSON), nil
}
