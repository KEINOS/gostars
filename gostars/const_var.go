package gostars

import "io"

// ============================================================================
//  Constants and Package Variables (Both local and exposed)
// ============================================================================

// ----------------------------------------------------------------------------
//  Constants
// ----------------------------------------------------------------------------

const urlAwesomeGoDefault = "https://raw.githubusercontent.com/avelino/awesome-go/main/README.md"

// ----------------------------------------------------------------------------
//  Variables
// ----------------------------------------------------------------------------

// URLAwesomeGo is the URL of Awesome-Go's README.md. Which is the markdown
// file of the awesome Go packages.
var URLAwesomeGo = urlAwesomeGoDefault

// GithubToken is a personal access token for the GitHub API that must be assigned
// by the caller and must not be hard-coded in the source code.
var GithubToken string

// IOCopy is a copy of io.Copy to ease test.
var IOCopy = io.Copy

// URLAliases are a mapping between the site URL and the actual URL of the GitHub repository.
var URLAliases = map[string]string{
	"https://joe-bot.net/": "https://github.com/go-joe/joe",
}
