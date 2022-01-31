package gostars

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/google/go-github/v42/github"
	"golang.org/x/oauth2"
)

// ============================================================================
//  Type: RepoInfo
// ============================================================================

// RepoInfo contains information about the repository on GitHub.
// It is mainly used to retrieve the number of stars, forks, followers, etc.
// from the repository.
type RepoInfo struct {
	URL         *URLInfo `json:"url"`         // Parsed URL info of the repo
	Description string   `json:"description"` // Desctiption of the repo
	Name        string   `json:"name"`        // Name of the repo
	Owner       string   `json:"owner"`       // Name of the repo owner
	Stars       int      `json:"stars"`       // Number of stars of the repo
	Forks       int      `json:"forks"`       // Number of forked repo of the repo
	Followers   int      `json:"followers"`   // Number of watching people
}

// ============================================================================
//  Constructor
// ============================================================================

// NewRepoInfo returns the initialized object of RepoInfo from the given GitHub's URL.
func NewRepoInfo(urlRepo string) (*RepoInfo, error) {
	// Get the actual URL of the GitHub repository from the mapping. (URLAliases)
	urlRepo = GetURLGitHub(urlRepo)

	urlInfo, err := NewURLInfo(urlRepo)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get URL information")
	}

	repoInfo := &RepoInfo{
		URL: urlInfo,
	}

	if !strings.Contains(urlRepo, "https://github.com/") {
		return nil, errors.New("the URL must be under GitHub host")
	}

	nameOwner, err := repoInfo.getNameOwner()
	if err != nil {
		return nil, err
	}

	repoInfo.Owner = nameOwner

	nameRepo, err := repoInfo.getNameRepo()
	if err != nil {
		return nil, err
	}

	repoInfo.Name = nameRepo

	// Update other field
	if err := repoInfo.Update(); err != nil {
		return nil, errors.Wrap(err, "failed to update repository info")
	}

	return repoInfo, nil
}

// ============================================================================
//  Methods
// ============================================================================

// Update retrieves the repository information from GitHub and sets it in the
// corresponding field.
func (r *RepoInfo) Update() error {
	CoolDown()

	ctx := context.Background()
	client := github.NewClient(nil)

	if GithubToken != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: GithubToken},
		)
		tc := oauth2.NewClient(ctx, ts)

		client = github.NewClient(tc)
	}

	repo, resp, err := client.Repositories.Get(ctx, r.Owner, r.Name)
	if err != nil || resp.StatusCode != 200 {
		msgErr := fmt.Sprintf("faild to get repository info. Returned status: %v", resp.StatusCode)

		return errors.Wrap(err, msgErr)
	}

	r.Description = repo.GetDescription()
	r.Stars = repo.GetStargazersCount()
	r.Forks = repo.GetForksCount()
	r.Followers = repo.GetSubscribersCount()

	return nil
}

// Owner returns the owner name from the repository URL.
func (r *RepoInfo) getNameOwner() (string, error) {
	path := r.URL.Path

	if len(path) < 1 {
		return "", errors.New("invalid URL format. missing repo owner and/or repo name")
	}

	return path[0], nil
}

// Name returns the name from the repository from the URL.
func (r *RepoInfo) getNameRepo() (string, error) {
	path := r.URL.Path

	if len(path) < 2 {
		return "", errors.New("invalid URL format. missing repo owner and/or repo name")
	}

	return path[1], nil
}
