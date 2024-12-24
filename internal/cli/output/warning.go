package output

func Warning(message string) {
	title := "⚠️ Warning:"

	orangeBoldPrinter(title)
	print(" ")
	orangePrinter(message)
	println()
}
