package main

import (
	"Mrs4s/MiraiGo/client"
	"Mrs4s/MiraiGo/message"
	"fmt"
	"log"
	"time"
)

func main() {
	var qpp = client.NewClient(123456,"123456")
	//上面输入QQ号和密码
	var res,err = qpp.Login()
	//登录
	if(res.Success){
		log.Println("success")
		//登陆成功
	}else{
		log.Println(err.Error())
		return
		//登录失败
	}
	log.Println(qpp.Nickname)
	//输出用户名
	time.Sleep(time.Second)
	_ = qpp.ReloadFriendList()
	//重载好友和群聊（必须）否则不能接收到信息
	_=qpp.ReloadGroupList()
	//同上
	log.Println(len(qpp.FriendList))
	log.Println(len(qpp.GroupList))
	qpp.OnPrivateMessage(func(qqClient *client.QQClient, message *message.PrivateMessage) {
		fmt.Printf("%s   %d   %s  \n",message.Sender.Nickname,message.Time,message.ToString())
	})
	//添加消息处理事件
	qpp.OnGroupMessage(func(qqClient *client.QQClient, groupMessage *message.GroupMessage) {
		fmt.Printf("%s   %d   %s  %s\n",groupMessage.Sender.Nickname,groupMessage.Time,groupMessage.ToString(),groupMessage.GroupName)
	})
	qpp.OnGroupMessage(func(qqClient *client.QQClient, groupMessage *message.GroupMessage) {
		var reply = message.NewSendingMessage()
		reply.Append(message.NewReply(groupMessage))
		reply.Append(message.NewText(groupMessage.ToString()))
		reply.Append(message.NewFace(178))
		qpp.SendGroupMessage(groupMessage.GroupCode,reply)

	})
	//回复
	for{
		time.Sleep(time.Millisecond*100)
	}
	//堵塞进程
}
