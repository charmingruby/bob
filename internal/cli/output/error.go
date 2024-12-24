package output

import "os"

func ShutdownWithError(message string) {
	symbol := "✖️"
	title := "Error"

	redBoldPrinter(symbol)
	print(blankSpace)

	redBoldUnderlinePrinter(title)
	redBoldPrinter(separator)

	println(message + ".")

	os.Exit(1)
}
