package data

type ContainerData struct {
	GoVersion string
}

func NewContainerData(
	goVersion string,
) ContainerData {
	return ContainerData{
		GoVersion: goVersion,
	}
}
