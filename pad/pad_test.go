package pad

import "testing"

func TestPad_Encrypt_Success(t *testing.T) {
	testCases := []struct {
		plaintext string
		key       string
		expected  string
	}{
		{
			"Hello World!",
			string([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}),
			"Hello World!",
		},
		{
			"Hello World!",
			string([]byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}),
			"Ifmmp!Xpsme\"",
		},
	}

	for _, testCase := range testCases {
		pad := New(testCase.key)
		result, err := pad.Encrypt(testCase.plaintext)

		if err != nil {
			t.Errorf("TestPad_Encrypt failed: %v", err)
		}
		if result != testCase.expected {
			t.Errorf("TestPad_Encrypt failed: expected %s, got %s", testCase.expected, result)
		}
	}
}

func TestPad_Encrypt_Fail(t *testing.T) {
	testCases := []struct {
		plaintext string
		key       string
	}{
		{
			"Hello World!",
			"",
		},
		{
			"Hello World!",
			"hsdf",
		},
	}

	for _, testCase := range testCases {
		pad := New(testCase.key)
		_, err := pad.Encrypt(testCase.plaintext)

		if err == nil {
			t.Errorf("TestPad_Encrypt did not return error: %v", err)
		}
	}
}

func TestPad_Decrypt_Success(t *testing.T) {
	testCases := []struct {
		encrypted string
		key       string
		expected  string
	}{
		{
			"Hello World!",
			string([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}),
			"Hello World!",
		},
		{
			"Ifmmp!Xpsme\"",
			string([]byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}),
			"Hello World!",
		},
	}

	for _, testCase := range testCases {
		pad := New(testCase.key)
		result, err := pad.Decrypt(testCase.encrypted)

		if err != nil {
			t.Errorf("TestPad_Decrypt failed: %v", err)
		}
		if result != testCase.expected {
			t.Errorf("TestPad_Decrypt failed: expected %s, got %s", testCase.expected, result)
		}
	}
}

func TestPad_Decrypt_Fail(t *testing.T) {
	testCases := []struct {
		encrypted string
		key       string
	}{
		{
			"Hello World!",
			"",
		},
		{
			"Hello World!",
			"hsdf",
		},
	}

	for _, testCase := range testCases {
		pad := New(testCase.key)
		_, err := pad.Encrypt(testCase.encrypted)

		if err == nil {
			t.Errorf("TestPad_Decrypt did not return error: %v", err)
		}
	}
}
