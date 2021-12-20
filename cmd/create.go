package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/google/go-github/v41/github"
	"github.com/spf13/cobra"
)

var createGistCmd = &cobra.Command{
	Args:    cobra.MinimumNArgs(1),
	Use:     "create filePaths...",
	Aliases: []string{"c"},
	Short:   "Creates a gist given a list of file paths.",
	Long: `Creates a gist given a list of file paths.
	
If a directory is specified, will include all files within that directory recursively, with depth specified by depth flag.`,
}

func init() {
	var (
		isPublic bool
		maxDepth int
	)

	createGistCmd.Flags().BoolVarP(&isPublic, "public", "p", false, "If set, creates a public gist.")
	createGistCmd.Flags().IntVarP(&maxDepth, "depth", "d", 1, "Maximum depth to recurse for directories.")

	createGistCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("expected at least 1 argument")
		}

		client, err := getGithubClient()
		if err != nil {
			return fmt.Errorf("error initializing github client: %w", err)
		}

		gist, err := buildGist(args, isPublic, maxDepth)
		if err != nil {
			return fmt.Errorf("error constructing gist: %w", err)
		}

		created, err := client.CreateGist(cmd.Context(), gist)
		if err != nil {
			return err
		}

		fmt.Printf("Created gist %s\nAvailable at %s\n", *created.ID, *created.HTMLURL)

		return nil
	}
}

func buildGist(filePaths []string, isPublic bool, depth int) (*github.Gist, error) {
	gist := &github.Gist{
		Files:  map[github.GistFilename]github.GistFile{},
		Public: &isPublic,
	}

	if err := addGistFiles(gist, filePaths, depth, 0); err != nil {
		return nil, fmt.Errorf("error adding gist files: %w", err)
	}

	for key := range gist.Files {
		fmt.Println(key)
	}

	if len(gist.Files) == 0 {
		return nil, fmt.Errorf("could not find any files at the specified paths")
	}

	return gist, nil
}

func addGistFiles(gist *github.Gist, filePaths []string, maxDepth int, depth int) error {
	if depth > maxDepth {
		return nil
	}
	for _, filePath := range filePaths {
		fs, err := os.Stat(filePath)
		if err != nil {
			return fmt.Errorf("unable to read file %s: %w", filePath, err)
		}
		if !fs.IsDir() {
			content, err := ioutil.ReadFile(filePath)
			if err != nil {
				return fmt.Errorf("unable to read content at path %s: %w", filePath, err)
			}
			contentStr := string(content)
			if len(contentStr) == 0 {
				continue
			}
			gist.Files[github.GistFilename(path.Base(filePath))] = github.GistFile{
				Content: &contentStr,
			}
			continue
		}

		dirents, err := os.ReadDir(filePath)
		if err != nil {
			return fmt.Errorf("error reading dirents at path %s: %w", filePath, err)
		}

		var direntFilePaths []string
		for _, dirent := range dirents {
			direntFilePaths = append(direntFilePaths, path.Join(filePath, dirent.Name()))
		}

		addGistFiles(gist, direntFilePaths, maxDepth, depth+1)
	}
	return nil
}
