package generator

import (
	"fmt"
	"os"
	"text/template"

	"github.com/charmingruby/gentoo/tpl"
)

type GenerateFileInput struct {
	Directory    string
	Module       string
	Resource     string
	ResourceName string
	Suffix       string
	Data         any
	ActionType   string
	HasTest      bool
}

func GenerateFile(input GenerateFileInput) error {
	if input.HasTest {
		testResource := fmt.Sprintf("%s_test", input.Resource)

		testTmpl, err := createTemplate(testResource, input.ActionType)
		if err != nil {
			fmt.Println("Error creating test template:", err)
			return err
		}

		if err := createFile(input.ResourceName, "_test", input.Directory, input.Data, testTmpl); err != nil {
			return err
		}
	}

	tmpl, err := createTemplate(input.Resource, input.ActionType)
	if err != nil {
		fmt.Println("Error creating template:", err)
		return err
	}

	return createFile(input.ResourceName, input.Suffix, input.Directory, input.Data, tmpl)
}

func createTemplate(resource, actionType string) (*template.Template, error) {
	templatePath := fmt.Sprintf("%s/%s", actionType, formatTplFile(resource))

	tplContent, err := tpl.GenerateTemplateFS.ReadFile(templatePath)
	if err != nil {
		fmt.Println("Error reading template:", err)
		return nil, err
	}

	tmpl, err := template.New(resource).Parse(string(tplContent))
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return nil, err
	}

	return tmpl, err
}

func createFile(resourceName, suffix, directory string, data any, template *template.Template) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	destinyDir := fmt.Sprintf("%s/%s", currentDir, directory)

	var finalResourceName string = resourceName
	if suffix != "" {
		finalResourceName += suffix
	}

	fileName := fmt.Sprintf("%s/%s", destinyDir, formatGoFile(finalResourceName))

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	err = template.Execute(file, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return err
	}

	return nil
}
