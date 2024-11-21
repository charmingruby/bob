package generator

import (
	"fmt"
	"os"
	"text/template"

	"github.com/charmingruby/gentoo/internal/command/shared/constant"
)

type GenerateFileInput struct {
	Directory    string
	Module       string
	Resource     string
	ResourceName string
	Suffix       string
	Data         any
}

func GenerateFile(input GenerateFileInput) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	templatePath := constant.TEMPLATE_DIR + formatTplFile(input.Resource)

	tplContent, err := os.ReadFile(templatePath)
	if err != nil {
		fmt.Println("Error reading template:", err)
		return err
	}

	tmpl, err := template.New(input.Resource).Parse(string(tplContent))
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return err
	}

	destinyDir := fmt.Sprintf("%s/%s", currentDir, input.Directory)

	var finalResourceName string = input.ResourceName
	if input.Suffix != "" {
		finalResourceName += input.Suffix
	}

	fileName := fmt.Sprintf("%s/%s", destinyDir, formatGoFile(finalResourceName))

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	err = tmpl.Execute(file, input.Data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return err
	}

	return nil
}
