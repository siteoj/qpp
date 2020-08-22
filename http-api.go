package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type tust struct {
	Tst  string `json:"tst"`
	//Name  string `json:"temp"`
}

func main() {
	http.HandleFunc("/json", myHandle)
	http.ListenAndServe(":3000", nil)

}

func myHandle(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	con, _ := ioutil.ReadAll(r.Body) //获取post的数据
	//b := []byte(`{"tset" : "hhhuasidhad" , "temp" : "temp"}`)
	//println(b)
	fmt.Println(string(con))
	var tt tust
	err := json.Unmarshal(con,&tt)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(tt)
	//ddd, _ := base64.StdEncoding.DecodeString(tt.Tst)
	//ioutil.WriteFile("./logo4.png", ddd, 0667)
	var rgx = regexp.MustCompile(`\[(.*?)\]`)
	rs1 := rgx.FindStringSubmatch(tt.Tst)
	rgx = regexp.MustCompile(`\((.*?)\)`)
	//fmt.Println(rs1)
	rs2 := rgx.FindStringSubmatch(rs1[1])
	//fmt.Println(rs2)
	tt.Tst=strings.Replace(rs1[1],rs2[0],"",-1)
	//tt.Tst=strings.Replace(,rs2[0],"",-1)
	fmt.Println(tt)
	//ddd, _ := base64.StdEncoding.DecodeString(tt.Tst)
	//ioutil.WriteFile("./logo4.png", ddd, 0667)
}