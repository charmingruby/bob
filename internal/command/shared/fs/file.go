package fs

import (
	"fmt"
	"os"
	"text/template"

	"github.com/charmingruby/bob/tpl"
)

type File struct {
	Identifier string
	HasTest    bool
	ActionType string
	Directory  string
	Name       string
	Suffix     string
	Data       any
}

func GenerateFile(file File) error {
	if file.HasTest {
		testFile := fmt.Sprintf("%s_test", file.Identifier)

		testTmpl, err := createTemplate(testFile, file.ActionType)
		if err != nil {
			fmt.Println("Error creating test template:", err)
			return err
		}

		if err := createFile(file.Name, "_test", file.Directory, file.Data, testTmpl); err != nil {
			return err
		}
	}

	tmpl, err := createTemplate(file.Identifier, file.ActionType)
	if err != nil {
		fmt.Println("Error creating template:", err)
		return err
	}

	return createFile(file.Name, file.Suffix, file.Directory, file.Data, tmpl)
}

func createTemplate(file, actionType string) (*template.Template, error) {
	templatePath := fmt.Sprintf("%s/%s", actionType, formatTplFile(file))

	tplContent, err := tpl.GenerateTemplateFS.ReadFile(templatePath)
	if err != nil {
		fmt.Println("Error reading template:", err)
		return nil, err
	}

	tmpl, err := template.New(file).Parse(string(tplContent))
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return nil, err
	}

	return tmpl, err
}

func createFile(name, suffix, directory string, data any, template *template.Template) error {
	destinyDir, err := getDestinationDirectory(directory)
	if err != nil {
		fmt.Println("Error getting destination directory:", err)
		return err
	}

	var finalFileName string = name
	if suffix != "" {
		finalFileName += suffix
	}

	fileName := fmt.Sprintf("%s/%s", destinyDir, formatGoFile(finalFileName))

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
