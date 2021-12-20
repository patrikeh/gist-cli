package githubclient

import (
	"context"
	"fmt"

	"github.com/google/go-github/v41/github"
	"golang.org/x/oauth2"
)

type IClient interface {
	CreateGist(*github.Gist) (*github.Gist, error)
}

type Client struct {
	client *github.Client
}

func New(accessToken string) *Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	return &Client{
		client: github.NewClient(tc),
	}
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
