package internal

import (
	"autoProjectAsaka/internal/domain"
	"autoProjectAsaka/pkg"
	"fmt"
	"strings"
)

const (
	MqPath   = ".\\testModel\\Ad\\Ad.TestSystem.testAppMq.V1"
	HttpPath = ".\\testModel\\Ad\\Ad.TestSystem.testApp.V1"
)

func ManageFile(appName domain.AppAsaka) {
	service := strings.ToLower(strings.Split(appName.ServiceName, ".")[0])
	inputType := strings.ToLower(appName.InputType)
	var pathToCopyDir string
	if service == "ad" {
		if inputType == "http" {
			pathToCopyDir = HttpPath
		} else if inputType == "mq" {
			pathToCopyDir = MqPath
		} else {
			fmt.Println("another inputType")
			return
		}

		err := pkg.CopyDir(pathToCopyDir, ".\\NewService\\test")
		if err != nil {
			fmt.Println(err)
			return
		}

		ChangeProjectFile(appName.ServiceName)
		newPath := RenameFolder(appName.ServiceName)
		ChangeMsgFlow(newPath+"rq.msgflow", appName.ServiceName)
		ChangeMsgFlow(newPath+"rs.msgflow", appName.ServiceName)
		changeBrokerSchema(newPath, appName.ServiceName)
	} else {
		fmt.Println("another service")
	}

}

func changeBrokerSchema(dir, appName string) {
	changeEsqlBrokerSchema(dir+"rq_ReqLogic.esql", appName)
	changeEsqlBrokerSchema(dir+"rq_ErrReqLogic.esql", appName)
	changeEsqlBrokerSchema(dir+"rs_RespLogic.esql", appName)
	changeEsqlBrokerSchema(dir+"rs_ErrRespLogic.esql", appName)
}
