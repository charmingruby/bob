package output

import "fmt"

func ComingSoon(feature, section string) {
	title := "Coming soon"
	message := fmt.Sprintf("%s %s", feature, section)
	baseWarningPrinter(title, message)
}

func baseWarningPrinter(title, message string) {
	symbol := "⚠️"

	orangeBoldPrinter(symbol)
	print(blankSpace)

	orangeBoldUnderlinePrinter(title)
	orangeBoldPrinter(separator)

	println(message + ".")
}
