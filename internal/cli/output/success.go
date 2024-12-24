package output

import "fmt"

func CommandSuccess(command string) {
	title := "Command success"
	message := fmt.Sprintf("%s ran successfully", command)
	baseSuccessPrinter(title, message)
}

func ComponentCreated(identifier string) {
	title := "Component created"
	baseSuccessPrinter(title, identifier)
}

func baseSuccessPrinter(title, message string) {
	symbol := "✔️"

	greenBoldPrinter(symbol)
	print(blankSpace)

	greenBoldUnderlinePrinter(title)
	greenBoldPrinter(separator)

	println(message + ".")
}
