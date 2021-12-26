package githubclient

import (
	"context"
	"fmt"
	"net/url"

	"github.com/google/go-github/v41/github"
	"golang.org/x/oauth2"
)

type IClient interface {
	CreateGist(*github.Gist) (*github.Gist, error)
}

type Client struct {
	client *github.Client
}

func New(accessToken, githubUrl string) (*Client, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	url, err := url.Parse(githubUrl)
	if err != nil {
		return nil, fmt.Errorf("error parsing github url: %w", err)
	}
	client.BaseURL = url

	return &Client{
		client,
	}, nil
}

func (c Client) CreateGist(ctx context.Context, gist *github.Gist) (*github.Gist, error) {
	created, _, err := c.client.Gists.Create(ctx, gist)
	if err != nil {
		return nil, fmt.Errorf("error creating gist: %w", err)
	}

	return created, nil
}

func (c Client) DeleteGist(ctx context.Context, gistId string) error {
	_, err := c.client.Gists.Delete(ctx, gistId)
	if err != nil {
		return fmt.Errorf("error deleting gist: %w", err)
	}
	return nil
}

func (c Client) GetGist(ctx context.Context, gistId string) (*github.Gist, error) {
	gist, _, err := c.client.Gists.Get(ctx, gistId)
	if err != nil {
		return nil, fmt.Errorf("error getting gist: %w", err)
	}
	return gist, nil
}

// In case userId is empty string, lists gists for authenticated user.
func (c Client) ListGists(ctx context.Context, userId string) ([]*github.Gist, error) {
	gists, _, err := c.client.Gists.List(ctx, userId, &github.GistListOptions{})
	if err != nil {
		return nil, fmt.Errorf("error listing gists: %w", err)
	}
	return gists, nil
}
