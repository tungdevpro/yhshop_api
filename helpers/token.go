package helpers

import "strings"

func ExtraTokenFromHeader(s string) string {
	items := strings.Split(s, " ")
	bearer := strings.TrimSpace(items[0])

	if len(bearer) < 2 || bearer == " " || bearer != "Bearer" {
		return ""
	}
	return items[1]
}
