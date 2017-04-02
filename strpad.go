package iso8583

import "strings"

func leftPad(s string, length int, pad string) string {
	if len(s) >= length {
		return s
	}
	padding := strings.Repeat(pad, length-len(s))
	return padding + s
}
