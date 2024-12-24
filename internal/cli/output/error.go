package output

import "os"

func ShutdownWithError(message string) {
	title := "✖️ Error:"

	redBoldPrinter(title)
	print(" ")
	redPrinter(message + ".")
	println()

	os.Exit(1)
}
