package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func formatName(line string) string {
	str := strings.TrimSpace(line)
	str = strings.ReplaceAll(str, " ", "-")
	str = regexp.MustCompile(`[^a-zA-Z0-9 _-]+`).ReplaceAllString(str, "")
	str = strings.ToLower(str)
	return str
}

func GetTableOfContents(filename string) (strs []string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "#") {
			continue
		}
		for i := 6; i >= 0; i-- {
			if strings.HasPrefix(line, strings.Repeat("#", i)) {
				line = strings.TrimPrefix(line, strings.Repeat("#", i)+" ")
				str := strings.Repeat("  ", i-1)
				str += "- [" + line + "](#" + formatName(line) + ")"
				strs = append(strs, str)
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Usage: md-table-of-contents <markdownfile.md>")
		os.Exit(1)
	}
	tableOfContents, err := GetTableOfContents(os.Args[1])
	if err != nil {
		panic(err)
	}
	for _, line := range tableOfContents {
		fmt.Println(line)
	}
}
