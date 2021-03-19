package github

import (
	"context"
	"fmt"
	"github.com/google/go-github/v33/github"
	"github.com/heycar-uk/semver/internal/utils"
	"log"
)

func GetReleaseNotesBacklog(ctx context.Context, client *github.Client, owner, repo, currentDeployed, repoType string) string {
	releases, _, err := client.Repositories.ListReleases(ctx, owner, repo, nil)
	if err != nil {
		log.Fatalln(err)
	}

	code := utils.GetRepoType(repoType)

	var PRReleases []utils.PR

	for _, release := range releases {
		name := release.GetName()
		tag := release.GetTagName()
		notes := utils.CleanReleaseNotes(release.GetBody())

		if currentDeployed == tag || len(currentDeployed) == 0 {
			break
		} else if currentDeployed != tag {
			PRReleases = append(PRReleases, utils.PR{
				Repo:          repo,
				Type:          code,
				Name:          name,
				ReleaseNumber: tag,
				ReleaseNotes:  notes,
			})
		}
	}

	if len(releases) == 0 {
		log.Fatalln("Did not find any releases")
	}

	if len(PRReleases) == 0 {
		for _, release := range releases {
			name := release.GetName()
			tag := release.GetTagName()
			notes := utils.CleanReleaseNotes(release.GetBody())
			fmt.Print("Re-Release | ")
			PRReleases = append(PRReleases, utils.PR{
				Repo:          repo,
				Type:          code,
				Name:          name,
				ReleaseNumber: tag,
				ReleaseNotes:  notes,
			})
			break
		}
	}

	releaseNotes := ""
	repoTypeLogo := ":bookmark:"
	for _, release := range PRReleases {
		if release.Type == utils.Frontend {
			repoTypeLogo = ":react:"
		} else if release.Type == utils.Backend {
			repoTypeLogo = ":python:"
		} else if release.Type == utils.DevOps {
			repoTypeLogo = ":terraform:"
		} else if release.Type == utils.Graphql {
			repoTypeLogo = ":graphql:"
		}

		releaseNotes += fmt.Sprintf("%s | *%s | %s*\n", repoTypeLogo, release.Repo, release.ReleaseNumber)
		releaseNotes += fmt.Sprintf("_%s_\n\n", release.Name)
		releaseNotes += fmt.Sprintf("%s\n\n", release.ReleaseNotes)
	}

	return releaseNotes
}
