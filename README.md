# QPP

**本项目遵守AGPL->v3协议**，请各位开发者遵守该协议的内容。

目前正在完善**http-api**，所以请大家在**issue**中提交未完成待实现的**http-api**

**提交格式：**

```go
http.HandleFunc("/网页api地址", 函数名)

type 结构体类型 struct{
    大写字母.... 数据类型 `json:"json数据昵称"`
    ...
}

func 函数名(w http.ResponseWriter, r *http.Request) {
   defer r.Body.Close()
   con, _ := ioutil.ReadAll(r.Body) //获取post的数据
   fmt.Println(string(con))
   var tt 结构体类型
   err := json.Unmarshal(con, &tt)//解析json文件
   if err != nil {
      fmt.Println("json err:", err)
   }
   fmt.Println(tt)
   //处理数据...
}
```

#### 例子：

```go
http.HandleFunc("/prisend", sendingHandle)//私聊信息发送

type sendmsgpp struct {
	Receiver  string `json:"rcv"`//收信者
	Msgg  string `json:"msg"`//信息内容
}

func sendingHandle(w http.ResponseWriter, r *http.Request) {
   defer r.Body.Close()
   con, _ := ioutil.ReadAll(r.Body) //获取post的数据
   fmt.Println(string(con))
   var tt sendmsgpp
   err := json.Unmarshal(con, &tt)//解析
   if err != nil {
      fmt.Println("json err:", err)
   }
   fmt.Println(tt)
   //处理数据，发送信息
   var msg = message.NewSendingMessage()
   msg.Append(message.NewText(tt.Msgg))
   rcv, err := strconv.Atoi(tt.Receiver)
   //qpp.SendPrivateMessage(int64(rcv), msg)
   m:=sp(tt.Msgg,int64(rcv))
   qpp.SendPrivateMessage(169829974,m)
}
```



