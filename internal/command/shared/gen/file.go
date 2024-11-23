package gen

import (
	"fmt"
	"os"
	"text/template"

	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/tpl"
)

func GenerateFile(component component.Component) error {
	if component.HasTest {
		testComponent := fmt.Sprintf("%s_test", component.Identifier)

		testTmpl, err := createTemplate(testComponent, component.ActionType)
		if err != nil {
			fmt.Println("Error creating test template:", err)
			return err
		}

		if err := createFile(component.Name, "_test", component.Directory, component.Data, testTmpl); err != nil {
			return err
		}
	}

	tmpl, err := createTemplate(component.Identifier, component.ActionType)
	if err != nil {
		fmt.Println("Error creating template:", err)
		return err
	}

	return createFile(component.Name, component.Suffix, component.Directory, component.Data, tmpl)
}

func createTemplate(component, actionType string) (*template.Template, error) {
	templatePath := fmt.Sprintf("%s/%s", actionType, formatTplFile(component))

	tplContent, err := tpl.GenerateTemplateFS.ReadFile(templatePath)
	if err != nil {
		fmt.Println("Error reading template:", err)
		return nil, err
	}

	tmpl, err := template.New(component).Parse(string(tplContent))
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return nil, err
	}

	return tmpl, err
}

func createFile(componentName, suffix, directory string, data any, template *template.Template) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	destinyDir := fmt.Sprintf("%s/%s", currentDir, directory)

	var finalComponentName string = componentName
	if suffix != "" {
		finalComponentName += suffix
	}

	fileName := fmt.Sprintf("%s/%s", destinyDir, formatGoFile(finalComponentName))

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
