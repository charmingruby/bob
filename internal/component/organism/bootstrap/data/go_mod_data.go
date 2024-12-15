package data

type GoModData struct {
	ProjectName string
	BaseURL     string
	GoVersion   string
}

func NewGoModData(
	projectName, baseURL, goVersion string,
) GoModData {
	return GoModData{
		ProjectName: projectName,
		BaseURL:     baseURL,
		GoVersion:   goVersion,
	}
}
