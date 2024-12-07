package scaffold

func TransportPath(moduleDir, protocol string, complementaryPath []string) string {
	var path = moduleDir + "/" + TRANSPORT_PACKAGE + "/" + REST_PACKAGE

	for _, p := range complementaryPath {
		path += "/" + p
	}

	return path
}

func CorePath(moduleDir string, complementaryPath []string) string {
	var path = moduleDir + "/" + CORE_PACKAGE

	for _, p := range complementaryPath {
		path += "/" + p
	}

	return path
}

func PersistencePath(moduleDir string, complementaryPath []string) string {
	var path = moduleDir + "/" + PERSISTENCE_PACKAGE

	for _, p := range complementaryPath {
		path += "/" + p
	}

	return path
}

func CustomErrPath(moduleDir string, complementaryPath []string) string {
	var path = moduleDir + "/" + CUSTOM_ERR_PACKAGE

	for _, p := range complementaryPath {
		path += "/" + p
	}

	return path

}
