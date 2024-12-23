package data

type GoModData struct {
	RepoURL   string
	GoVersion string
}

func NewGoModData(
	repoURL, goVersion string,
) GoModData {
	return GoModData{
		RepoURL:   repoURL,
		GoVersion: goVersion,
	}
}
