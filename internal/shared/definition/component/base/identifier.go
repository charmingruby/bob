package base

import "fmt"

func BuildIdentifier(module, content, destinationDirectory string) string {
	return fmt.Sprintf("[%s] %s, at: %s", module, content, directoryRootSafety(destinationDirectory))
}

func BuildNonModuleIdentifier(resource string, content, destinationDirectory string) string {
	return fmt.Sprintf("[bob | %s] %s, at: %s", resource, content, directoryRootSafety(destinationDirectory))
}

func BuildBobIdentifier(content, destinationDirectory string) string {
	return fmt.Sprintf("[bob] %s, at: %s", content, directoryRootSafety(destinationDirectory))
}

func directoryRootSafety(dir string) string {
	if dir == "." {
		return "root"
	}

	return dir
}
