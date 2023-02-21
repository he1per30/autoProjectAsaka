package internal

import (
	"fmt"
	"os"
	"strings"
)

// RenameFolder переименовывает скопированную папку в нужный appName и возвращает путь до системны файлов
func RenameFolder(appName string) string {
	fmt.Println(appName)
	newPath := "./NewService/" + appName
	serviceSlice := strings.Split(appName, ".")
	fmt.Println(serviceSlice)
	//Переименовываем корневаю папку
	err := os.Rename("./NewService/test", newPath)
	if err != nil {
		fmt.Println(err)
	}

	//Переименовываем папку с названием системы адаптера
	helpPath := newPath + "/ad/"
	newPath = helpPath + strings.ToLower(serviceSlice[1])
	err = os.Rename(helpPath+"/testsystem", newPath)
	if err != nil {
		fmt.Println(err)
	}

	//Переименовываем папку с названием сущности адаптера
	helpPath = newPath + "/"
	newPath = helpPath + strings.ToLower(serviceSlice[2])
	err = os.Rename(helpPath+"testapp", newPath)
	if err != nil {
		fmt.Println(err)
	}

	return newPath + "/v1/"
}
