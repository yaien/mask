package crypto

import (
	"testing"
)

func TestCaesarEncrypt(t *testing.T) {
	testcases := []struct {
		message  string
		expected string
		offset   int
	}{
		{message: "hello", expected: "ifmmp", offset: 1},
		{message: "google", expected: "iqqing", offset: 2},
		{message: "0b9 ", expected: " c0a", offset: 1},
	}

	for _, testcase := range testcases {
		result, _ := Encrypt(testcase.message, testcase.offset)
		if result != testcase.expected {
			t.Errorf("Message '%s' should be '%s', but received '%s'", testcase.message, testcase.expected, result)
		}
	}
}

func TestCaesarDecrypt(t *testing.T) {
	testcases := []struct {
		message  string
		expected string
		offset   int
	}{
		{message: "ifmmp", expected: "hello", offset: 1},
		{message: "iqqing", expected: "google", offset: 2},
		{message: " c0a", expected: "0b9 ", offset: 1},
	}

	for _, testcase := range testcases {
		result, _ := Decrypt(testcase.message, testcase.offset)
		if result != testcase.expected {
			t.Errorf("Message '%s' should be '%s', but received '%s'", testcase.message, testcase.expected, result)
		}
	}
}

func TestCaesarFailDecrypt(t *testing.T) {
	_, err := Decrypt(".-`", 1)
	if err == nil {
		t.Error("error should't be null")
	}
}

func TestCaesarFailEncrypt(t *testing.T) {
	_, err := Encrypt(".-`", 1)
	if err == nil {
		t.Error("error should't be null")
	}
}
