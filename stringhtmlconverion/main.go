package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Hello text"

	var isHtml bool = containHtml(str)

	fmt.Println("is this html:", isHtml)

}


func containHtml(str string) bool{
	if strings.Contains(str, "<") && strings.Contains(str, ">") {
		return true
	}

	// Check for escaped HTML entities
	if strings.Contains(str, "&lt;") || strings.Contains(str, "&gt;") {
		return true
	}

	return false
}


// func extractTextFromHTML(html string) string{
//      unescaped := html
// 	 strings.ReplaceAll(unescaped, "&lt", "<")
// 	 strings.ReplaceAll(unescaped, "&gt", ">")
// 	 strings.ReplaceAll(unescaped, "&lt", "<")
// }

