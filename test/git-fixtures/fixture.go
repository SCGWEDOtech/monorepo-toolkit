package gitfixture

import "fmt"

var (
	basic *gitHubRepository = nil
)

func BasicRepository() GitRepository {
	if basic == nil {
		basic = &gitHubRepository{name: "WhatTheFar/monorepo-toolkit-git-fixture-basic"}
	}
	return basic
}

type GitRepository interface {
	CompareURL(from, to string) string
}

type gitHubRepository struct {
	name string
}

func (r *gitHubRepository) repositoryURL() string {
	return fmt.Sprintf("https://github.com/%s", r.name)
}

func (r *gitHubRepository) CompareURL(from, to string) string {
	return fmt.Sprintf("%s/compare/%s..%s", r.repositoryURL(), from, to)
}
