package internal

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadFromFile() string {
	dir := "./settings.txt"
	input, err := os.ReadFile(dir)

	if err != nil {
		log.Fatalln(err, " in ReadFromFile")
	}

	lines := strings.Split(string(input), "\n")
	appName := ""
	for i, line := range lines {
		if strings.Contains(line, "appName") {
			appName = strings.Replace(lines[i], "appName: ", "", -1)
		}

	}
	fmt.Println("appName === ", appName)
	return appName
}
