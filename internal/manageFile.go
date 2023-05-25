package internal

import (
	"autoProjectAsaka/internal/domain"
	"autoProjectAsaka/pkg"
	"fmt"
	"strings"
)

const (
	AdMqPath   = ".\\testModel\\Ad\\Ad.TestSystem.testAppMq.V1"
	AdHttpPath = ".\\testModel\\Ad\\Ad.TestSystem.testApp.V1"
	AcMqPath   = ".\\testModel\\Ac\\Ac.TestSystem.testAppMq.V1"
	AcHttpPath = ".\\testModel\\Ac\\Ac.TestSystem.testApp.V1"
)

func ManageFile(appName domain.AppAsaka) {
	service := strings.ToLower(strings.Split(appName.ServiceName, ".")[0])
	//inputType := strings.ToLower(appName.InputType)
	if service == "ad" {
		changeAdapterFiles(appName)
	} else if service == "ac" {
		changeAсFiles(appName)
	} else {
		fmt.Println("another service")
	}
	//pathToCopyDir := manageAdapterPathToCopy(appName)
	//if pathToCopyDir == "" {
	//	return
	//}

	//err := pkg.CopyDir(pathToCopyDir, ".\\NewService\\test")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//ChangeProjectFile(appName.ServiceName)
	//newPath := RenameFolder(appName.ServiceName)
	//ChangeMsgFlow(newPath+"rq.msgflow", appName.ServiceName)
	//
	//if service != "ac" && inputType != "http" {
	//	ChangeMsgFlow(newPath+"rs.msgflow", appName.ServiceName)
	//}
	//
	//changeAdapterBrokerSchema(newPath, appName.ServiceName)
}

func manageAdapterPathToCopy(appName domain.AppAsaka) string {
	service := strings.ToLower(strings.Split(appName.ServiceName, ".")[0])
	inputType := strings.ToLower(appName.InputType)
	var pathToCopyDir string
	if service == "ad" {
		if inputType == "http" {
			pathToCopyDir = AdHttpPath
		} else if inputType == "mq" {
			pathToCopyDir = AdMqPath
		} else {
			fmt.Println("another inputType")
			return ""
		}
		//} else if service == "ac" {
		//	if inputType == "http" {
		//		pathToCopyDir = AcHttpPath
		//	} else if inputType == "mq" {
		//		pathToCopyDir = AdMqPath
		//	} else {
		//		fmt.Println("another inputType")
		//		return ""
		//	}
	}
	return pathToCopyDir
}

func manageAcPathToCopy(appName domain.AppAsaka) string {
	//service := strings.ToLower(strings.Split(appName.ServiceName, ".")[0])
	inputType := strings.ToLower(appName.InputType)
	var pathToCopyDir string
	if inputType == "http" {
		pathToCopyDir = AcHttpPath
	} else if inputType == "mq" {
		pathToCopyDir = AcMqPath
	} else {
		fmt.Println("another inputType")
		return ""
	}
	return pathToCopyDir
}

func changeAсFiles(appName domain.AppAsaka) {
	pathToCopyDir := manageAcPathToCopy(appName)
	if pathToCopyDir == "" {
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
	//ChangeMsgFlow(newPath+"rs.msgflow", appName.ServiceName)
	changeAcBrokerSchema(newPath, appName.ServiceName)
}

func changeAdapterFiles(appName domain.AppAsaka) {
	pathToCopyDir := manageAdapterPathToCopy(appName)
	if pathToCopyDir == "" {
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
	changeAdapterBrokerSchema(newPath, appName.ServiceName)
}

func changeAcBrokerSchema(dir, appName string) {
	changeEsqlBrokerSchema(dir+"rq_ReqLogic.esql", appName)
	changeEsqlBrokerSchema(dir+"rq_RespLogic.esql", appName)
	//changeEsqlBrokerSchema(dir+"rs_RespLogic.esql", appName)
	//changeEsqlBrokerSchema(dir+"rs_ErrRespLogic.esql", appName)
}

func changeAdapterBrokerSchema(dir, appName string) {
	changeEsqlBrokerSchema(dir+"rq_ReqLogic.esql", appName)
	changeEsqlBrokerSchema(dir+"rq_ErrReqLogic.esql", appName)
	changeEsqlBrokerSchema(dir+"rs_RespLogic.esql", appName)
	changeEsqlBrokerSchema(dir+"rs_ErrRespLogic.esql", appName)
}
