package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	RED       = "\033[31m"
	COLOR_END = "\033[0m"
)

func ReadChars(file string) string {
	return fmt.Sprint(len(file))
}

func ReadLines(file string) string {
	return fmt.Sprint(strings.Count(file, "\n"))
}

func ReadWords(file string) string {
	words := strings.Fields(file)
	return fmt.Sprint(len(words))
}

func ReadRunes(file string) string {
	runes := []rune(file)
	return fmt.Sprint(len(runes))
}

func formatLine(chars []string) string {
	content := strings.Join(chars, " --- ")
	bar := fmt.Sprint(strings.Repeat("-", len(content)))
	return fmt.Sprint(bar, "\n", content, "\n", bar)
}

func main() {
	arguments := make(map[string]bool)
	files := make(map[*os.File]string)
	for i, arg := range os.Args {
		if i == 0 {
			continue
		}
		if strings.HasPrefix(arg, "-") {
			arguments[arg] = true
		} else {
			file, err := os.Open(arg)
			// TODO: handle error properly
			if err == nil {
				files[file] = arg
			}
		}
	}
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		files[os.Stdin] = "stdin"
	}
	if len(files) < 1 {
		fmt.Println("Please provide file names")
		return
	}
	var output string
	for file, fileName := range files {
		var fileOutput string
		buf, err := io.ReadAll(file)
		if err != nil {
			continue
		}
		data := string(buf)
		for action := range arguments {
			switch action {
			case "-c":
				fileOutput = strings.Join([]string{fileOutput, ReadChars(data)}, " chars ")
			case "-l":
				fileOutput = strings.Join([]string{fileOutput, ReadLines(data)}, " lines ")
			case "-w":
				fileOutput = strings.Join([]string{fileOutput, ReadWords(data)}, " words ")
			case "-m":
				fileOutput = strings.Join([]string{fileOutput, ReadRunes(data)}, " runes ")
			default:
				fileOutput = strings.Join([]string{fileOutput, fmt.Sprintf("Invalid argument: %v%s%v", RED, action, COLOR_END)}, " ")
			}
		}

		if len(arguments) == 0 {
			fileOutput = fmt.Sprintf(" chars %s, lines %s, words %s ", ReadChars(data), ReadLines(data), ReadWords(data))
		}

		output = fmt.Sprint(output, "\n", formatLine([]string{fileOutput, fileName}))
	}
	fmt.Println(output)
}
