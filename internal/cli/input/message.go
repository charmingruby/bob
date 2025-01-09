package input

import "fmt"

func ChooseSectionMessage(section string) string {
	return fmt.Sprintf("Choose %s:", section)
}

func EnterValueMessage(key string) string {
	return fmt.Sprintf("Enter %s value:", key)
}
