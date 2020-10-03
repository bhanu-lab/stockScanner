package utils

import (
	"golang.org/x/net/html"
	"strings"
)

func ReplaceUnnecessaryHtmlData(eg html.Token) string {
	actualString := strings.ReplaceAll(eg.String(), "&lt;\\/td&gt;", "")
	actualString = strings.ReplaceAll(actualString, `<a href="\&#34;javascript:" toajaxtableeditor('order_by_changed',="" new="" array('`, "")
	actualString = strings.ReplaceAll(actualString, `','desc'));\"="">`, "")
	actualString = strings.ReplaceAll(actualString, `','asc'));\"="">`, "")
	return actualString
}
