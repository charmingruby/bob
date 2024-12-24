package output

import "github.com/fatih/color"

const (
	blankSpace = " "
	separator  = ": "
)

var (
	redBoldPrinter = color.New(color.Bold, color.Underline, color.FgRed).PrintFunc()
	redPrinter     = color.New(color.FgRed).PrintFunc()

	greenBoldPrinter          = color.New(color.Bold, color.FgGreen).PrintFunc()
	greenBoldUnderlinePrinter = color.New(color.Bold, color.Underline, color.FgGreen).PrintFunc()
	greenPrinter              = color.New(color.FgGreen).PrintFunc()

	orangeBoldPrinter = color.New(color.Bold, color.FgYellow).PrintFunc()
	orangePrinter     = color.New(color.FgYellow).PrintFunc()
)
