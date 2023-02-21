package internal

import (
	"autoProjectAsaka/internal/domain"
	"autoProjectAsaka/pkg"
	"fmt"
)

func ManageFile(appName domain.AppAsaka) {
	err := pkg.CopyDir(".\\testModel\\Ad\\Ad.TestSystem.testApp.V1", ".\\NewService\\test")
	if err != nil {
		fmt.Println(err)
		return
	}
	ChangeProjectFile(appName.ServiceName)
	newPath := RenameFolder(appName.ServiceName)
	ChangeMsgFlow(newPath+"rq.msgflow", appName.ServiceName)
	ChangeMsgFlow(newPath+"rs.msgflow", appName.ServiceName)
	changeBrokerSchema(newPath, appName.ServiceName)

}

func changeBrokerSchema(dir, appName string) {
	changeEsqlBrokerSchema(dir+"rq_ReqLogic.esql", appName)
	changeEsqlBrokerSchema(dir+"rq_ErrReqLogic.esql", appName)
	changeEsqlBrokerSchema(dir+"rs_RespLogic.esql", appName)
	changeEsqlBrokerSchema(dir+"rs_ErrRespLogic.esql", appName)
}
