// Package clr is a small package for coloring terminal text output with a pleasent to use API.
// Many Go terminal coloring packages are highly functional, which is nice for legibility but a bit of a pain
// when you're just trying to color a sentence or two. 
package clr

// Black takes your input and wraps it in ANSI color code 30
func Black(input ...any) string {
	return This(input...).formats(black).String()
}

// Red takes your input and wraps it in ANSI color code 31
func Red(input ...any) string {
	return This(input...).formats(darkRed).String()
}

// Green takes your input and wraps it in ANSI color code 32
func Green(input ...any) string {
	return This(input...).formats(darkGreen).String()
}

// Yellow takes your input and wraps it in ANSI color code 33
func Yellow(input ...any) string {
	return This(input...).formats(darkYellow).String()
}

// Blue takes your input and wraps it in ANSI color code 34
func Blue(input ...any) string {
	return This(input...).formats(darkBlue).String()
}

// Magenta takes your input and wraps it in ANSI color code 35
func Magenta(input ...any) string {
	return This(input...).formats(darkMagenta).String()
}

// Cyan takes your input and wraps it in ANSI color code 36
func Cyan(input ...any) string {
	return This(input...).formats(darkCyan).String()
}

// White takes your input and wraps it in ANSI color code 97.
func White(input ...any) string {
	return This(input...).formats(white).String()
}
