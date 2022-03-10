package gostars_test

import (
	"io"
	"testing"
	"time"

	"github.com/KEINOS/gostars/gostars"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
//  Function Test
// ----------------------------------------------------------------------------
//  For Golden cases see: examples_test.go
// ----------------------------------------------------------------------------

func TestCoolDown(t *testing.T) {
	start := time.Now()

	// Cool down
	gostars.CoolDown()

	e1 := time.Since(start).Seconds()
	e2 := float64(1)
	assert.GreaterOrEqual(t, e1, e2, "it should sleep more than equal to 1 second")
}

func TestGetContentURL(t *testing.T) {
	// Backup and defer restore
	oldIOCopy := gostars.IOCopy
	defer func() {
		gostars.IOCopy = oldIOCopy
	}()

	// Mock IOCopy to force error
	gostars.IOCopy = func(dst io.Writer, src io.Reader) (written int64, err error) {
		return 0, errors.New("forced error")
	}

	for _, test := range []struct {
		url     string
		contain string
	}{
		{"https://github.com/KEINOS/" + string([]byte{0x7f}), "failed to parse URL before request"},
		{"https://foo.bar.com/", "dial tcp: lookup foo.bar.com"},
		{"https://github.com/KEINOS/unknownrepo/", "returned status: 404"},
		{"https://github.com/KEINOS/dev-go/", "failed to copy response data"},
	} {
		_, err := gostars.GetContentURL(test.url)

		require.Error(t, err)
		assert.Contains(t, err.Error(), test.contain)
	}
}

func TestGetDistance(t *testing.T) {
	for _, test := range []struct {
		name   string
		args   []int
		expect int
	}{
		{name: "no dim", args: []int{}, expect: 0},
		{name: "one dim", args: []int{10}, expect: 10},
		{name: "two dim", args: []int{10, 10}, expect: 14},
		{name: "three dim", args: []int{10, 10, 10}, expect: 17},
		{name: "four dim", args: []int{10, 10, 10, 10}, expect: 20},
	} {
		expect := test.expect
		actual := gostars.GetAttractionGravity(test.args...)
		assert.Equal(t, expect, actual, "failed test: %v", test.name)
	}

	p1 := gostars.GetAttractionGravity(10, 10, 10, 10)
	p2 := gostars.GetAttractionGravity(100, 10, 10, 10)
	assert.Greater(t, p2, p1, "greater point should have more distance")

	p1 = gostars.GetAttractionGravity(10, 10, 10, 10)
	p2 = gostars.GetAttractionGravity(-10, -10, -10, -10)
	assert.Equal(t, p1, p2, "negative position but same distance should be equal")
}

func TestGetURLGitHub(t *testing.T) {
	// URL not in the aliases list
	urlUnknown := "https://foo.bar/"

	expect := urlUnknown
	actual := gostars.GetURLGitHub(urlUnknown)

	assert.Equal(t, expect, actual, "unlisted URL in the aliases map should return the input")
}

func TestPrettyFormatJSON(t *testing.T) {
	foo := func() {}

	output, err := gostars.PrettyFormatJSON(foo)

	require.Error(t, err)
	assert.Empty(t, output, "it should be empty on error")
}

// ----------------------------------------------------------------------------
//  PkgInfo
// ----------------------------------------------------------------------------

func TestPkgInfo_bad_package_name(t *testing.T) {
	namePkg := "github.com/KEINOS/undefined"

	pkgInfo, err := gostars.NewPkgInfo(namePkg)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get 'imported by' information")
	assert.Nil(t, pkgInfo)
}

func TestUpdateURLRepository_fail(t *testing.T) {
	pkgInfo := &gostars.PkgInfo{
		Name: "github.com/KEINOS/undefined",
	}

	err := pkgInfo.UpdateURLRepository()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get package information")
}

// ----------------------------------------------------------------------------
//  RepoInfo
// ----------------------------------------------------------------------------

func TestRepoInfo_bad_url(t *testing.T) {
	for _, test := range []struct {
		url     string
		contain string
	}{
		{"https://foo.bar.com/", "the URL must be under GitHub host"},
		{"https://github.com/", "missing repo owner and/or repo name"},
		{"https://github.com/KEINOS", "missing repo owner and/or repo name"},
		{"https://github.com/KEINOS/" + string([]byte{0x7f}), "invalid control character in URL"},
		{"https://github.com/KEINOS/undefined", "failed to update repository info"},
	} {
		repoInfo, err := gostars.NewRepoInfo(test.url)

		require.Error(t, err, "'%v' should be an error", test.url)
		assert.Nil(t, repoInfo, "it shuld be nil on error")
		assert.Contains(t, err.Error(), test.contain)
	}
}

func TestRepoInfo_Update_bad_credential(t *testing.T) {
	oldGithubToken := gostars.GithubToken
	defer func() {
		gostars.GithubToken = oldGithubToken
	}()

	gostars.GithubToken = "<undefined>"

	repoInfo := gostars.RepoInfo{
		Name:  "dev-go",
		Owner: "KEINOS",
		URL:   new(gostars.URLInfo),
	}

	err := repoInfo.Update()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "faild to get repository info")
}
