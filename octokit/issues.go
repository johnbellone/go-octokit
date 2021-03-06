package octokit

import (
	"github.com/lostisland/go-sawyer/hypermedia"
	"net/url"
	"time"
)

var (
	RepoIssuesURL = Hyperlink("/repos/{owner}/{repo}/issues{/number}")
)

// Create a IssuesService with the base Hyperlink and the params M to expand the Hyperlink
// If no Hyperlink is passed in, it will use RepoIssuesURL.
func (c *Client) Issues(link *Hyperlink, m M) (issues *IssuesService, err error) {
	if link == nil {
		link = &RepoIssuesURL
	}

	url, err := link.Expand(m)
	if err != nil {
		return
	}

	issues = &IssuesService{client: c, URL: url}
	return
}

type IssuesService struct {
	client *Client
	URL    *url.URL
}

func (i *IssuesService) Get() (issue *Issue, result *Result) {
	result = i.client.Get(i.URL, &issue)
	return
}

func (i *IssuesService) GetAll() (issues []Issue, result *Result) {
	result = i.client.Get(i.URL, &issues)
	return
}

type Issue struct {
	*hypermedia.HALResource

	URL     string `json:"url,omitempty,omitempty"`
	HTMLURL string `json:"html_url,omitempty,omitempty"`
	Number  int    `json:"number,omitempty"`
	State   string `json:"state,omitempty"`
	Title   string `json:"title,omitempty"`
	Body    string `json:"body,omitempty"`
	User    User   `json:"user,omitempty"`
	Labels  []struct {
		URL   string `json:"url,omitempty"`
		Name  string `json:"name,omitempty"`
		Color string `json:"color,omitempty"`
	}
	Assignee  User `json:"assignee,omitempty"`
	Milestone struct {
		URL          string     `json:"url,omitempty"`
		Number       int        `json:"number,omitempty"`
		State        string     `json:"state,omitempty"`
		Title        string     `json:"title,omitempty"`
		Description  string     `json:"description,omitempty"`
		Creator      User       `json:"creator,omitempty"`
		OpenIssues   int        `json:"open_issues,omitempty"`
		ClosedIssues int        `json:"closed_issues,omitempty"`
		CreatedAt    time.Time  `json:"created_at,omitempty"`
		DueOn        *time.Time `json:"due_on,omitempty"`
	}
	Comments    int `json:"comments,omitempty"`
	PullRequest struct {
		HTMLURL  string `json:"html_url,omitempty"`
		DiffURL  string `json:"diff_url,omitempty"`
		PatchURL string `json:"patch_url,omitempty"`
	} `json:"pull_request,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	ClosedAt  *time.Time `json:"closed_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
}
