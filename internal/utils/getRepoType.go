package utils

func GetRepoType(repoType string) PRType {
	if repoType == "frontend" {
		return Frontend
	} else if repoType == "backend" {
		return Backend
	} else if repoType == "frontend-services" {
		return Graphql
	} else {
		return DevOps
	}
}
