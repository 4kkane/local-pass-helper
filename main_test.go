package main

import (
	"regexp"
	"strings"
	"testing"
)

// Test password length
func TestPasswordLength(t *testing.T) {
	tests := []struct {
		name     string
		length   int
		expected int
	}{
		{"Default Length", 12, 12},
		{"Custom Length", 16, 16},
		{"Minimum Length", 6, 6},
		{"Below Minimum Length", 3, 4}, // generatePassword函数中仍保持最小长度为4
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			password := generatePassword(tc.length)
			if len(password) != tc.expected {
				t.Errorf("Expected length %d, got %d", tc.expected, len(password))
			}
		})
	}
}

// Test if password contains all required character types
func TestPasswordContainsAllTypes(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{"Default Length", 12},
		{"Custom Length", 16},
		{"Minimum Length", 6},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			password := generatePassword(tc.length)

			// Check for lowercase letters
			hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
			if !hasLower {
				t.Error("Password does not contain lowercase letters")
			}

			// Check for uppercase letters
			hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
			if !hasUpper {
				t.Error("Password does not contain uppercase letters")
			}

			// Check for numbers
			hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
			if !hasNumber {
				t.Error("Password does not contain numbers")
			}

			// Check for special characters
			hasSpecial := false
			for _, c := range password {
				if strings.ContainsRune(specialCharSet, c) {
					hasSpecial = true
					break
				}
			}
			if !hasSpecial {
				t.Error("Password does not contain special characters")
			}
		})
	}
}

// Test if special character is at the last position
func TestSpecialCharAtLastPosition(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{"Default Length", 12},
		{"Custom Length", 16},
		{"Minimum Length", 6},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			password := generatePassword(tc.length)
			lastChar := password[len(password)-1]

			isSpecial := false
			for _, c := range specialCharSet {
				if lastChar == uint8(c) {
					isSpecial = true
					break
				}
			}

			if !isSpecial {
				t.Errorf("Last character '%c' is not a special character", lastChar)
			}
		})
	}
}

// Test if generated passwords are unique
func TestPasswordUniqueness(t *testing.T) {
	passwords := make(map[string]bool)
	count := 100

	for i := 0; i < count; i++ {
		password := generatePassword(12)
		if passwords[password] {
			t.Errorf("Duplicate password: %s", password)
		}
		passwords[password] = true
	}

	if len(passwords) != count {
		t.Errorf("Expected %d different passwords, got %d", count, len(passwords))
	}
}

// Test character type placement function
func TestEnsureCharType(t *testing.T) {
	password := []byte("ABCDEFGH")

	// Test placing lowercase letter at position 0
	result := ensureCharType(password, lowerCharSet, 0)

	if result[0] < 'a' || result[0] > 'z' {
		t.Errorf("Position 0 should be a lowercase letter, got: %c", result[0])
	}
}

// Test empty password handling
func TestEmptyPassword(t *testing.T) {
	result := ensureCharType([]byte{}, lowerCharSet, 0)
	if len(result) != 0 {
		t.Error("Empty password handling failed")
	}
}

// Test out-of-bounds position
func TestPositionOutOfBounds(t *testing.T) {
	password := []byte("ABCD")
	result := ensureCharType(password, lowerCharSet, 10)

	// Should return the original password, unchanged
	if string(result) != "ABCD" {
		t.Errorf("Out-of-bounds position handling failed, expected original password, got: %s", string(result))
	}
}

// 测试字符分布 - 确保第一位为小写字母，第二位为大写字母，第三位为数字
func TestCharacterDistribution(t *testing.T) {
	for i := 0; i < 10; i++ {
		password := generatePassword(12)

		// 检查第一位是否为小写字母
		if password[0] < 'a' || password[0] > 'z' {
			t.Errorf("First character should be a lowercase letter, got: %c", password[0])
		}

		// 检查第二位是否为大写字母
		if password[1] < 'A' || password[1] > 'Z' {
			t.Errorf("Second character should be an uppercase letter, got: %c", password[1])
		}

		// 检查第三位是否为数字
		if password[2] < '0' || password[2] > '9' {
			t.Errorf("Third character should be a number, got: %c", password[2])
		}
	}
}

// 测试密码最小长度验证（main函数中的逻辑）
// 注意：由于此测试涉及到main函数中的代码逻辑，无法直接测试
// 这里提供测试validatePasswordLength函数的示例，该函数需要从main函数中提取
func TestMinimumPasswordLengthValidation(t *testing.T) {
	// 提取validatePasswordLength函数的示例代码:
	/*
		func validatePasswordLength(length int) error {
			if length < 6 {
				return fmt.Errorf("Error: Password length must be greater than 6")
			}
			return nil
		}
	*/

	// 可以添加以下测试，如果你决定提取validatePasswordLength函数
	/*
		tests := []struct {
			name        string
			length      int
			expectError bool
		}{
			{"Valid Length", 6, false},
			{"Valid Longer Length", 12, false},
			{"Invalid Length", 5, true},
			{"Invalid Zero Length", 0, true},
			{"Invalid Negative Length", -1, true},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				err := validatePasswordLength(tc.length)
				hasError := err != nil
				if hasError != tc.expectError {
					t.Errorf("Expected error: %v, got: %v", tc.expectError, hasError)
				}
			})
		}
	*/
}

// 测试密码随机性和字符分布
func TestPasswordRandomnessAndDistribution(t *testing.T) {
	const (
		passwordCount  = 1000
		passwordLength = 12
	)

	// 字符集分布计数
	lowerCount := 0
	upperCount := 0
	numberCount := 0
	specialCount := 0

	// 生成多个密码并统计
	for i := 0; i < passwordCount; i++ {
		password := generatePassword(passwordLength)

		for _, char := range password {
			c := string(char)
			switch {
			case regexp.MustCompile(`[a-z]`).MatchString(c):
				lowerCount++
			case regexp.MustCompile(`[A-Z]`).MatchString(c):
				upperCount++
			case regexp.MustCompile(`[0-9]`).MatchString(c):
				numberCount++
			default:
				specialCount++
			}
		}
	}

	// 计算总字符数
	totalChars := passwordCount * passwordLength

	// 计算比例
	lowerRatio := float64(lowerCount) / float64(totalChars)
	upperRatio := float64(upperCount) / float64(totalChars)
	numberRatio := float64(numberCount) / float64(totalChars)
	specialRatio := float64(specialCount) / float64(totalChars)

	// 打印分布情况以供参考
	t.Logf("Character distribution in %d passwords:", passwordCount)
	t.Logf("Lowercase: %.2f%% (%d/%d)", lowerRatio*100, lowerCount, totalChars)
	t.Logf("Uppercase: %.2f%% (%d/%d)", upperRatio*100, upperCount, totalChars)
	t.Logf("Numbers: %.2f%% (%d/%d)", numberRatio*100, numberCount, totalChars)
	t.Logf("Special: %.2f%% (%d/%d)", specialRatio*100, specialCount, totalChars)

	// 验证特殊字符的数量是否符合预期（每个密码最后一位是特殊字符）
	expectedSpecialCount := passwordCount
	if specialCount != expectedSpecialCount {
		t.Errorf("Expected %d special characters, got %d", expectedSpecialCount, specialCount)
	}

	// 验证固定位置字符类型
	// 由于我们固定第一位是小写字母，第二位是大写字母，第三位是数字
	minLowerCount := passwordCount  // 至少每个密码有一个小写字母（第一位）
	minUpperCount := passwordCount  // 至少每个密码有一个大写字母（第二位）
	minNumberCount := passwordCount // 至少每个密码有一个数字（第三位）

	if lowerCount < minLowerCount {
		t.Errorf("Expected at least %d lowercase characters, got %d", minLowerCount, lowerCount)
	}

	if upperCount < minUpperCount {
		t.Errorf("Expected at least %d uppercase characters, got %d", minUpperCount, upperCount)
	}

	if numberCount < minNumberCount {
		t.Errorf("Expected at least %d numbers, got %d", minNumberCount, numberCount)
	}

	// 验证是否有合理分布（这些阈值可能需要调整）
	if lowerRatio < 0.2 || lowerRatio > 0.5 {
		t.Errorf("Lowercase ratio (%.2f) outside expected range", lowerRatio)
	}

	if upperRatio < 0.2 || upperRatio > 0.5 {
		t.Errorf("Uppercase ratio (%.2f) outside expected range", upperRatio)
	}

	if numberRatio < 0.1 || numberRatio > 0.4 {
		t.Errorf("Number ratio (%.2f) outside expected range", numberRatio)
	}
}
