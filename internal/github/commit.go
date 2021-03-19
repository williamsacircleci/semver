package github

import (
	"context"
	"fmt"

	"github.com/cenkalti/backoff/v4"
	"github.com/google/go-github/v33/github"
	"log"
)

func GetReleaseByCommit(ctx context.Context, client *github.Client, owner, repo, sha string) string {

	var version string
	var tagList []*github.RepositoryTag

	operation := func() error {

		opts := &github.ListOptions{
			Page: 0,
		}

		for {
			tags, resp, err := client.Repositories.ListTags(ctx, owner, repo, opts)
			if err != nil {
				log.Fatalln(err)
			}

			tagList = append(tagList, tags...)
			if resp.NextPage == 0 {
				break
			}
			opts.Page = resp.NextPage
		}

		for i := 0; i < len(tagList); i++ {
			if tagList[i].GetCommit().GetSHA() == sha {
				version = tagList[i].GetName()
				break
			}
		}

		if version == "" {
			return fmt.Errorf("%s %s", "No Tags found with commit", sha)
		}
		return nil
	}

	err := backoff.Retry(operation, backoff.NewExponentialBackOff())
	if err != nil {
		log.Fatalln("No Tags found with commit ", sha)
	}

	return version
}
