package main

import (
	"fmt"
	"service"
)

func main2() {
	messageService := service.GetMessageService()
	// messageService.AddMessage(pojo.NewMessage(2,1,1,"你也好啊","",false))

	// messages := messageService.GetMessageByReceiveID(2)
	messages := messageService.GetNoReadCountMessageByReceiveID(2)
	fmt.Println(messages)
}
