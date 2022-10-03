package clr

type Format string

const (
	black        Format = "30"
	darkRed      Format = "31"
	darkGreen    Format = "32"
	darkYellow   Format = "33"
	darkBlue     Format = "34"
	darkMagenta  Format = "35"
	darkCyan     Format = "36"
	lightGray    Format = "37"
	darkGray     Format = "90"
	lightRed     Format = "91"
	lightGreen   Format = "92"
	lightYellow  Format = "93"
	lightBlue    Format = "94"
	lightMagenta Format = "95"
	lightCyan    Format = "96"
	white        Format = "97"
	// Background
	bgBlack        Format = "40"
	bgDarkRed      Format = "41"
	bgDarkGreen    Format = "42"
	bgDarkYellow   Format = "43"
	bgDarkBlue     Format = "44"
	bgDarkCyan     Format = "46"
	bgLightGray    Format = "100"
	bgLightRed     Format = "101"
	bgLightGreen   Format = "102"
	bgLightYellow  Format = "103"
	bgLightBlue    Format = "104"
	bgLightMagenta Format = "105"
	bgLightCyan    Format = "106"
	bgWhite        Format = "107"
	// Format
	bold        Format = "1"
	underline   Format = "4"
	noUnderline Format = "24"
	// ReverseText  Format = "7"
	positiveText Format = "27"
	reset        Format = "0"
)
