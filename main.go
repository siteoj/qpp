package main

import (
	"Mrs4s/MiraiGo/client"
	"Mrs4s/MiraiGo/message"
	//"encoding/base64"
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	//"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"time"
	//"image"
)

type sendmsgpp struct {
	Receiver  string `json:"rcv"`
	Msgg  string `json:"msg"`
	//Name  string `json:"temp"`
}
var qpp = client.NewClient(账号,"密码")
func main() {


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
		if(groupMessage.Sender.Uin==169829974){
			var reply = message.NewSendingMessage()
			reply.Append(message.NewReply(groupMessage))
			reply.Append(message.NewText(groupMessage.ToString()))
			reply.Append(message.NewFace(178))
			qpp.SendGroupMessage(groupMessage.GroupCode,reply)

		}


	})
	//回复
	var tempimagemsg = message.NewSendingMessage()
	//var img Image =NewImage(d, 100, 40)

	ff, _ := ioutil.ReadFile("C:\\Users\\liwen\\Desktop\\logo4.png")
	i := message.NewImage(ff)
	//var fm = (message.FriendImageElement{})
	fm, _ := qpp.UploadGroupImage(755956002, i.Data)

	//encodeString := base64.StdEncoding.EncodeToString(ff)
	//fmt.Println(encodeString)
	tempimagemsg.Append(fm)
	tempimagemsg.Append(message.NewText("Hahaha"))
	qpp.SendGroupMessage(755956002,tempimagemsg)
	http.HandleFunc("/prisend", sendingHandle)
	http.HandleFunc("/grosend", groupsend)
	http.HandleFunc("/getnick",getnick)
	http.ListenAndServe(":3000", nil)
	for{
		time.Sleep(time.Millisecond*100)
	}
	//堵塞进程
}
func sendingHandle(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	con, _ := ioutil.ReadAll(r.Body) //获取post的数据
	//b := []byte(`{"tset" : "hhhuasidhad" , "temp" : "temp"}`)
	//println(b)
	fmt.Println(string(con))
	var tt sendmsgpp
	err := json.Unmarshal(con, &tt)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(tt)

	var msg = message.NewSendingMessage()
	msg.Append(message.NewText(tt.Msgg))
	rcv, err := strconv.Atoi(tt.Receiver)
	//qpp.SendPrivateMessage(int64(rcv), msg)
	m:=sp(tt.Msgg,int64(rcv))
	qpp.SendPrivateMessage(169829974,m)
}
func groupsend(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	con, _ := ioutil.ReadAll(r.Body) //获取post的数据
	//b := []byte(`{"tset" : "hhhuasidhad" , "temp" : "temp"}`)
	//println(b)
	fmt.Println(string(con))
	var tt sendmsgpp
	err := json.Unmarshal(con, &tt)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(tt)
	var msg = message.NewSendingMessage()
	msg.Append(message.NewText(tt.Msgg))
	rcv, err := strconv.Atoi(tt.Receiver)
	qpp.SendGroupMessage(int64(rcv), msg)
}

type response struct {
	Status  string `json:"status"`
	Resdata string `json:"response"`
}



func getnick(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	con, _ := ioutil.ReadAll(r.Body) //获取post的数据
	//b := []byte(`{"tset" : "hhhuasidhad" , "temp" : "temp"}`)
	//println(b)
	fmt.Println(string(con))
	var tt sendmsgpp
	err := json.Unmarshal(con, &tt)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(tt)
	var res =response{
		Status: "200",
		Resdata: qpp.Nickname,
	}
	var byteres []byte
	byteres,_=json.Marshal(res)
	w.Write(byteres)
}
//验证接口
//func check(w http.ResponseWriter, r *http.Request) {
//
//	msg, _ := json.Marshal({Code: 400, Msg: "验证失败"})
//
//	w.Header().Set("content-type","text/json")
//		if res {
//			msg, _ = json.Marshal(tools.JsonResult{Code: 200, Msg: "验证成功"})
//			w.Write(msg)
//		} else {
//			w.WriteHeader(400)
//			w.Write(msg)
//		}
//	}
//
//}
//func getnick(w http.ResponseWriter, r *http.Request) {
//	defer r.Body.Close()
//	con, _ := ioutil.ReadAll(r.Body) //获取post的数据
//	//println(b)
//	fmt.Println(string(con))
//	var tt sendmsgpp
//	err := json.Unmarshal(con, &tt)
//	if err != nil {
//		fmt.Println("json err:", err)
//	}
//	fmt.Println(tt)
//
//}
func sp(s string,target int64) *message.SendingMessage{
	var rgx = regexp.MustCompile(`\[(.*?)\]`)
	//s := `[tag]SomeText`
	m:=message.NewSendingMessage()
	for strings.Contains(s,"["){
		rs := rgx.FindStringSubmatch(s)
		println(len(rs))
		fmt.Println(rs[1])
		s=strings.Replace(s,rs[0],"",-1)
		//fmt.Println(s)
		m.Append(spp(rs[1],target))
	}
	m.Append(message.NewText(s))
	fmt.Println(m)
	return m
}
func spp(s string, target int64) *message.FriendImageElement {
	var rgx = regexp.MustCompile(`\((.*?)\)`)
	rs := rgx.FindStringSubmatch(s)
	switch rs[1] {
	case "image":
		rs[1]=strings.Replace(s,rs[0],"",-1)
		//ddd, _ := base64.StdEncoding.DecodeString(rs[1])
		fmt.Println(rs[1])
		ff, _ := ioutil.ReadFile(rs[1])
		fm, _ :=qpp.UploadPrivateImage(target,ff)
		return fm
	}
	return nil
}
