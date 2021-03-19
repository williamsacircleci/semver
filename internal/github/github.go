package github

import (
	"context"
	"log"
	"strings"

	"github.com/google/go-github/v33/github"
	"github.com/heycar-uk/semver/internal/utils"
)

func GetRelease(ctx context.Context, client *github.Client, owner, repo, sha string) utils.PR {
	currentVersion := getLatestRelease(ctx, client, owner, repo)
	pr := getPRbyCommit(ctx, client, owner, repo, sha)
	prLabel := getPRLabel(pr)
	newVersion := utils.IncrementCurrentRelease(currentVersion, prLabel)
	return utils.PR{
		Repo:          repo,
		Name:          pr[0].GetTitle(),
		Author:        pr[0].GetUser().GetLogin(),
		ReleaseNumber: newVersion,
		ReleaseNotes:  filterPRBody(pr[0].GetBody()),
	}
}

func filterPRBody(body string) string {

	start := "[//]: RELEASE_NOTES_START"
	end := "[//]: RELEASE_NOTES_END"

	i := strings.Index(body, start)
	if i == -1 {
		return ""
	}
	notesEnd := strings.Index(body, end)
	if notesEnd == -1 {
		return ""
	}
	notesStart := i + len(start)
	if notesStart >= notesEnd {
		return ""
	}
	return body[notesStart:notesEnd]
}

// GetLatestRelease returns a string with the latest SemVer release
func getLatestRelease(ctx context.Context, client *github.Client, owner, repo string) string {
	var latestRelease = "0.0.0"

	release, _, err := client.Repositories.GetLatestRelease(ctx, owner, repo)
	if err != nil {
		return latestRelease
	}

	latestRelease = release.GetTagName()
	return latestRelease
}

func getPRbyCommit(ctx context.Context, client *github.Client, owner, repo, sha string) []*github.PullRequest {
	pr, _, err := client.PullRequests.ListPullRequestsWithCommit(ctx, owner, repo, sha, nil)
	if err != nil {
		log.Fatalf("Could not find PR assigned to commit %s", sha)
	}

	return pr
}

// getPRLabel returns the label assigned to the PR
func getPRLabel(pr []*github.PullRequest) utils.SemVer {
	var cMajor int = 0
	var cMinor int = 0
	var cPatch int = 0
	var currentLabel utils.SemVer

	for _, label := range pr[0].Labels {
		switch label.GetName() {
		case "major":
			cMajor++
			currentLabel = utils.Major
		case "minor":
			cMinor++
			currentLabel = utils.Minor
		case "patch":
			cPatch++
			currentLabel = utils.Patch
		}
	}

	if cMajor+cMinor+cPatch == 0 {
		log.Fatalf("Could not find major, minor nor patch label")
	} else if cMajor+cMinor+cPatch > 1 {
		log.Fatalf("Found multiple labels, please use major, minor or patch")
	}

	return currentLabel
}
