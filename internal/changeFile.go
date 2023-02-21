package internal

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func ChangeProjectFile(appName string) {
	dir := "./NewService/test/.project"
	input, err := os.ReadFile(dir)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")
	for i, line := range lines {
		subString := "Ad.TestSystem.testApp.V1"
		if strings.Contains(line, subString) {
			lines[i] = strings.Replace(lines[i], subString, appName, -1)
		}
	}

	output := strings.Join(lines, "\n")
	err = os.WriteFile(dir, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}

}

// ChangeRqMsg Меняет все вхождения testsystem и testapp на части из appName
func ChangeMsgFlow(dir, appName string) {
	//rqDir := dir + "rq.msgflow"
	input, err := os.ReadFile(dir)
	splitLowerAppName := strings.Split(strings.ToLower(appName), ".")
	splitAppName := strings.Split(appName, ".")

	fmt.Println("this is app name: ", splitLowerAppName)
	if err != nil {
		log.Fatalln(err, " in ChangeMsgFlow")
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		subString := "/testsystem/testapp"
		if strings.Contains(line, subString) {
			newSubString := "/" + splitLowerAppName[1] + "/" + splitLowerAppName[2]
			lines[i] = strings.Replace(lines[i], subString, newSubString, -1)
		}

		if strings.Contains(line, "testsystem") {
			lines[i] = strings.Replace(lines[i], "testsystem", splitLowerAppName[1], -1)
		}

		if strings.Contains(line, "Ad.TestSystem.testApp.V1") {
			lines[i] = strings.Replace(lines[i], "Ad.TestSystem.testApp.V1", appName, -1)
		}

		if strings.Contains(line, "testapp") {
			lines[i] = strings.Replace(lines[i], "testapp", splitLowerAppName[2], -1)
		}

		if strings.Contains(line, "Sr.test.V1") {
			lines[i] = strings.Replace(lines[i], ".test.", "."+splitAppName[2]+".", -1)
		}

		if strings.Contains(line, "TEST.RS") || strings.Contains(line, "TEST.SN") {
			lines[i] = strings.Replace(lines[i], "TEST", strings.ToUpper(appName), -1)
		}
	}
	output := strings.Join(lines, "\n")
	err = os.WriteFile(dir, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func changeEsqlBrokerSchema(dir, appName string) {
	input, err := os.ReadFile(dir)
	if err != nil {
		log.Fatalln(err)
	}
	splitLowerAppName := strings.Split(strings.ToLower(appName), ".")

	lines := strings.Split(string(input), "\n")
	for i, line := range lines {
		subString := "testsystem.testapp"
		if strings.Contains(line, subString) {
			newSubString := splitLowerAppName[1] + "." + splitLowerAppName[2]
			lines[i] = strings.Replace(lines[i], subString, newSubString, -1)
		}
		output := strings.Join(lines, "\n")
		err = os.WriteFile(dir, []byte(output), 0644)
		if err != nil {
			log.Fatalln(err)
		}
		return
	}
}
