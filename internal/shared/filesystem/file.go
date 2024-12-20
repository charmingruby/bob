package filesystem

import (
	"fmt"
	"os"
	"text/template"

	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/pkg/formatter"
	"github.com/charmingruby/bob/tooling/tpl"
)

type File struct {
	CommandType          string // ex: gen, new...
	Extension            string
	DestinationDirectory string // directory where the file will be created
	FileName             string // file name
	FileSuffix           string // file suffix
	TemplateName         string // handler, model...
	TemplateData         any    // data to be used in the template
	HasTest              bool
}

func (f *File) format() {
	if f.Extension == definition.GO_EXTENSION {
		f.FileName = formatter.ToSnakeCase(f.FileName)
		f.FileSuffix = formatter.ToSnakeCase(f.FileSuffix)
	}
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

		if err := generateFileIfNotExists(file.FileName, "_test", file.DestinationDirectory, file.TemplateData, testTmpl, file.Extension); err != nil {
			return err
		}
	}

	tmpl, err := createTemplate(file.TemplateName, file.CommandType)
	if err != nil {
		fmt.Println("Error creating template:", err)
		return err
	}

	return generateFileIfNotExists(file.FileName, file.FileSuffix, file.DestinationDirectory, file.TemplateData, tmpl, file.Extension)
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

func generateFileIfNotExists(name, suffix, directory string, data any, tmpl *template.Template, extension string) error {
	destinyDir, err := getDestinationDirectory(directory)
	if err != nil {
		fmt.Println("Error getting destination directory:", err)
		return err
	}

	var finalFileName string = name
	if suffix != "" {
		finalFileName += "_" + suffix
	}

	var filePath string
	switch extension {
	case definition.GO_EXTENSION:
		filePath = fmt.Sprintf("%s/%s", destinyDir, formatGoFile(finalFileName))
	case definition.NO_EXTENSION:
		filePath = fmt.Sprintf("%s/%s", destinyDir, finalFileName)
	default:
		filePath = fmt.Sprintf("%s/%s.%s", destinyDir, finalFileName, extension)
	}

	if _, err := os.Stat(filePath); err == nil {
		fmt.Printf("File already exists: %s\n", filePath)
		return nil
	} else if !os.IsNotExist(err) {
		fmt.Println("Error checking file existence:", err)
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return err
	}

	return nil
}
