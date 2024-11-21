package create

import (
	"fmt"
	"os"
	"text/template"

	"github.com/ettle/strcase"
)

const (
	TEMPLATE_DIR = "./template/create/"
)

type TemplateParams struct {
	HandlerName string
}

func CreateHandler() {
	handlerName := "CreateWallet"

	params := TemplateParams{
		HandlerName: handlerName,
	}

	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("Current Directory:", currentDir)

	templatePath := TEMPLATE_DIR + "handler.tpl"
	fmt.Println("Template Path:", templatePath)

	tplContent, err := os.ReadFile(templatePath)
	if err != nil {
		fmt.Println("Error reading template:", err)
		panic(err)
	}

	tmpl, err := template.New("handler").Parse(string(tplContent))
	if err != nil {
		fmt.Println("Error parsing template:", err)
		panic(err)
	}

	fileName := strcase.ToSnake(handlerName) + ".go"

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		panic(err)
	}
	defer file.Close()

	err = tmpl.Execute(file, params)
	if err != nil {
		fmt.Println("Error executing template:", err)
		panic(err)
	}
}
