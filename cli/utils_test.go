package cli

import "testing"

func TestSafeInt32ConvertIsNil(t *testing.T) {
	expected := int32(0)
	resp := SafeInt32Convert(nil)
	if resp != expected {
		t.Errorf("expected %d != %d", expected, resp)
	}
}

func TestSafeInt32ConvertIsNotNil(t *testing.T) {
	expected := int32(0)
	resp := SafeInt32Convert(&expected)
	if resp != expected {
		t.Errorf("expected %d != %d", expected, resp)
	}
}

func TestSafeStringConvertIsNil(t *testing.T) {
	resp := SafeStringConvert(nil)
	if resp != "" {
		t.Errorf("expected empty != %s", resp)
	}
}

func TestSafeStringConvertIsNotNil(t *testing.T) {
	expected := "content"
	resp := SafeStringConvert(&expected)
	if resp != expected {
		t.Errorf("expected %s != %s", expected, resp)
	}
}

func TestEscapeStringToJson(t *testing.T) {
	password := EscapeStringToJson("§±!@#$%^&*()_+-=[]{};'\\:\"/.,?><`~")
	expectedPassword := "§±!@#$%^&*()_+-=[]{};'\\\\:\\\"/.,?><`~"
	if password != expectedPassword {
		t.Errorf("expected %s != %s", expectedPassword, password)
	}
}
