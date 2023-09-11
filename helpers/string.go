package helpers

import (
	"regexp"
	"strings"
)

func SnakeCase(text string) string {
	// Chuyển đổi thành chữ thường
	text = strings.ToLower(text)

	// Sử dụng biểu thức chính quy để tìm và thay thế các ký tự không hợp lệ
	re := regexp.MustCompile(`[-_\s.]+(.)?`)
	text = re.ReplaceAllStringFunc(text, func(match string) string {
		if len(match) > 1 {
			return strings.ToUpper(string(match[1]))
		}
		return ""
	})

	// Biến đổi ký tự đầu tiên thành chữ thường
	if len(text) > 0 {
		text = strings.ToLower(text[0:1]) + text[1:]
	}

	return text
}
