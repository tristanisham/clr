package clr

import "testing"

func TestBlackSingleInput(t *testing.T) {
	got := Black("hi")
	want := "\x1b[30mhi\x1b[0m"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestBlackMultiInput(t *testing.T) {
	got := Black("hi", 1, "😂")
	want := "\x1b[30mhi1😂\x1b[0m"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestRedSingleInput(t *testing.T) {
	got := Red("hi")
	want := "\x1b[31mhi\x1b[0m"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestRedMultiInput(t *testing.T) {
	got := Red("hi", 1, "😂")
	want := "\x1b[31mhi1😂\x1b[0m"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestGreenSingleInput(t *testing.T) {
	got := Green("hi")
	want := "\x1b[32mhi\x1b[0m"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestGreenMultiInput(t *testing.T) {
	got := Green("hi", 1, "😂")
	want := "\x1b[32mhi1😂\x1b[0m"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestYellowSingleInput(t *testing.T) {
	got := Yellow("hi")
	want := "\x1b[33mhi\x1b[0m"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestYellowMultiInput(t *testing.T) {
	got := Yellow("hi", 1, "😂")
	want := "\x1b[33mhi1😂\x1b[0m"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestBlueSingleInput(t *testing.T) {
	got := Blue("hi")
	want := "\x1b[34mhi\x1b[0m"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestBlueMultiInput(t *testing.T) {
	got := Blue("hi", 1, "😂")
	want := "\x1b[34mhi1😂\x1b[0m"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestMagentaSingleInput(t *testing.T) {
	got := Magenta("hi")
	want := "\x1b[35mhi\x1b[0m"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestMagentaMultiInput(t *testing.T) {
	got := Magenta("hi", 1, "😂")
	want := "\x1b[35mhi1😂\x1b[0m"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}


func TestCyanSingleInput(t *testing.T) {
	got := Cyan("hi")
	want := "\x1b[36mhi\x1b[0m"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestCyanMultiInput(t *testing.T) {
	got := Cyan("hi", 1, "😂")
	want := "\x1b[36mhi1😂\x1b[0m"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestWhiteSingleInput(t *testing.T) {
	got := White("hi")
	want := "\x1b[97mhi\x1b[0m"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestWhiteMultiInput(t *testing.T) {
	got := White("hi", 1, "😂")
	want := "\x1b[97mhi1😂\x1b[0m"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

