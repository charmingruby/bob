package filesystem

import (
	"fmt"
	"os"
	"text/template"

	"github.com/charmingruby/bob/internal/command/shared/formatter"
	"github.com/charmingruby/bob/tpl"
)

const (
	TEMPLATE_DIR = "tpl/gen/"
)

type File struct {
	FileName             string // file name
	FileSuffix           string // file suffix
	CommandType          string // ex: gen, new...
	TemplateName         string // handler, model...
	TemplateData         any    // data to be used in the template
	DestinationDirectory string // directory where the file will be created
	HasTest              bool
}

func (f *File) format() {
	f.FileName = formatter.ToSnakeCase(f.FileName)
	f.FileSuffix = formatter.ToSnakeCase(f.FileSuffix)
}

func (f *Manager) GenerateFile(file File) error {
	file.format()

	if file.HasTest {
		testFile := fmt.Sprintf("%s_test", file.TemplateName)

		testTmpl, err := createTemplate(testFile, file.CommandType)
		if err != nil {
			fmt.Println("Error creating test template:", err)
			return err
		}

		if err := createFile(file.FileName, "_test", file.DestinationDirectory, file.TemplateData, testTmpl); err != nil {
			return err
		}
	}

	tmpl, err := createTemplate(file.TemplateName, file.CommandType)
	if err != nil {
		fmt.Println("Error creating template:", err)
		return err
	}

	return createFile(file.FileName, file.FileSuffix, file.DestinationDirectory, file.TemplateData, tmpl)
}

func createTemplate(fileName, command string) (*template.Template, error) {
	templatePath := fmt.Sprintf("%s/%s", command, formatTplFile(fileName))

	tplContent, err := tpl.GenerateTemplateFS.ReadFile(templatePath)
	if err != nil {
		fmt.Println("Error reading template:", err)
		return nil, err
	}

	tmpl, err := template.New(fileName).Parse(string(tplContent))
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
		finalFileName += "_" + suffix
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
