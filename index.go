package main

import (
		"log"
		"github.com/dimonchik0036/vk-api"
		"strings"
		"strconv"
	//	"time"
)
func checkCoolGuys(dialog_id int64, client *vkapi.Client) {
	updates, _, err := client.GetLPUpdatesChan(100, vkapi.LPConfig{25, vkapi.LPModeAttachments})
	if err != nil {
		log.Panic(err)
	}
	counter := 0
	for update := range updates {
		if update.Message == nil || !update.IsNewMessage() || update.Message.Outbox(){
			continue
		}

		log.Printf("%s", update.Message.String())
		if strings.Contains(strings.ToLower(update.Message.Text), "да") && update.Message.FromID == dialog_id {
			counter++
			client.SendMessage(vkapi.NewMessage(vkapi.NewDstFromUserID(update.Message.FromID), "пизда"))
		}
		if strings.ToLower(update.Message.Text) == "#бум" && update.Message.FromID == dialog_id {
			client.SendMessage(vkapi.NewMessage(vkapi.NewDstFromUserID(update.Message.FromID), "Кол-во ПИЗДЫ отправлено этим ботом : " + strconv.Itoa(counter)))	
		}

	}
}

func main() {
	client, err := vkapi.NewClientFromLogin("hatimka@mail.ru", "sorryme", vkapi.ScopeMessages)
	if err != nil {
	    log.Panic(err)
	}
	
	client.Log(true)

	if err := client.InitLongPoll(0, 2); err != nil {
		log.Panic(err)
	}
	go checkCoolGuys(211526358, client)
	go checkCoolGuys(157266591, client)
	go checkCoolGuys(210692107, client)
	checkCoolGuys(366864374, client)

	
}
