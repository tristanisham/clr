package clr

import "fmt"

type formatter struct {
	formatString string
	body         string
}

func This(inputs ...any) *formatter {
	f := &formatter{formatString: "\x1b["}

	for _, i := range inputs {
		f.body += fmt.Sprintf("%v", i)
	}
	return f
}

func (f *formatter) formats(ft ...Format) *formatter {
	if len(f.formatString) == 0 {
		f.formatString = "\x1b["
	}

	for _, v := range ft {
		f.formatString += fmt.Sprintf("%v", v)
	}
	f.formatString += "m"

	return f
}

func (f formatter) String() string {
	return fmt.Sprintf("%s%v\x1b[%sm", f.formatString, f.body, reset)
}
