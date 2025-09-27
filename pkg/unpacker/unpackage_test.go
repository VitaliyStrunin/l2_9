package unpacker

import (
	"testing"
)

func Test001(t *testing.T) {
	packed := "45"
	expected := ""
	actual, err := UnpackString(packed)
	if err == nil {
		t.Errorf("unpackString(%s) should return an error", packed)
	}

	if expected != actual {
		t.Errorf("unpackString(%s) should return %s, got %s", packed, expected, actual)
	}
}

func Test002(t *testing.T) {
	packed := "a4bc2d5e"
	expected := "aaaabccddddde"
	actual, err := UnpackString(packed)
	if err != nil {
		t.Errorf("unpackString(%s) shoundn't return an error", packed)
	}

	if expected != actual {
		t.Errorf("unpackString(%s) should return %s, got %s", packed, expected, actual)
	}
}

func Test003(t *testing.T) {
	packed := ""
	expected := ""
	actual, err := UnpackString(packed)
	if err != nil {
		t.Errorf("unpackString(%s) shoundn't return an error", packed)
	}

	if expected != actual {
		t.Errorf("unpackString(%s) should return %s, got %s", packed, expected, actual)
	}
}

func Test004(t *testing.T) {
	packed := `\55`
	expected := "55555"
	actual, err := UnpackString(packed)
	if err != nil {
		t.Errorf("unpackString(%s) shoundn't return an error", packed)
	}

	if expected != actual {
		t.Errorf("unpackString(%s) should return %s, got %s", packed, expected, actual)
	}
}

func Test005(t *testing.T) {
	packed := `\\5`
	expected := `\\\\\`
	actual, err := UnpackString(packed)
	if err != nil {
		t.Errorf("unpackString(%s) shoundn't return an error", packed)
	}

	if expected != actual {
		t.Errorf("unpackString(%s) should return %s, got %s", packed, expected, actual)
	}
}

func Test006(t *testing.T) {
	packed := `qwe\4\5`
	expected := "qwe45"
	actual, err := UnpackString(packed)
	if err != nil {
		t.Errorf("unpackString(%s) shoundn't return an error", packed)
	}

	if expected != actual {
		t.Errorf("unpackString(%s) should return %s, got %s", packed, expected, actual)
	}
}

func Test007(t *testing.T) {
	packed := "qwe\\45"
	expected := "qwe44444"
	actual, err := UnpackString(packed)
	if err != nil {
		t.Errorf("unpackString(%s) shoundn't return an error", packed)
	}

	if expected != actual {
		t.Errorf("unpackString(%s) should return %s, got %s", packed, expected, actual)
	}
}

func Test008(t *testing.T) {
	packed := "a10b"
	expected := "aaaaaaaaaab"
	actual, err := UnpackString(packed)
	if err != nil {
		t.Errorf("unpackString(%s) shoundn't return an error", packed)
	}

	if expected != actual {
		t.Errorf("unpackString(%s) should return %s, got %s", packed, expected, actual)
	}
}

func Test009(t *testing.T) {
	packed := `qwe\45`
	expected := `qwe44444`
	actual, err := UnpackString(packed)
	if err != nil {
		t.Errorf("unpackString(%s) shoundn't return an error", packed)
	}

	if expected != actual {
		t.Errorf("unpackString(%s) should return %s, got %s", packed, expected, actual)
	}
}

func Test010(t *testing.T) {
	packed := `\`
	expected := ``
	actual, err := UnpackString(packed)
	if err == nil {
		t.Errorf("unpackString(%s) shound return an error", packed)
	}

	if expected != actual {
		t.Errorf("unpackString(%s) should return %s, got %s", packed, expected, actual)
	}
}

func Test011(t *testing.T) {
	packed := `a0`
	expected := ``
	actual, err := UnpackString(packed)
	if err == nil {
		t.Errorf("unpackString(%s) shound return an error", packed)
	}

	if expected != actual {
		t.Errorf("unpackString(%s) should return %s, got %s", packed, expected, actual)
	}
}
