package data

type RequestHelperData struct {
	SourcePath string
}

func NewRequestHelperData(sourcePath string) RequestHelperData {
	return RequestHelperData{
		SourcePath: sourcePath,
	}
}
