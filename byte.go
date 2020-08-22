package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	a:=[]byte(`AAA`)
	s:="abc"
	test()
	println(string(s[1]))
	fmt.Println(a)
	fmt.Println(string(a))
	http.HandleFunc("/",upload)
	http.ListenAndServe(":8080",nil)
}
func upload(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	con, _ := ioutil.ReadAll(r.Body)
	fmt.Println(con)
	fmt.Printf(string(con))
	ddd, _ := base64.StdEncoding.DecodeString(string(con))
	ioutil.WriteFile("./output.png", ddd, 0667)
}


func test() {
	var rgx = regexp.MustCompile(`\[(.*?)\]`)
	s := `[[tag55]][22]SomeText`
	for strings.Contains(s,"["){
		rs := rgx.FindStringSubmatch(s)
		println(len(rs))
		fmt.Println(rs[1])
		s=strings.Replace(s,rs[0],"hhh",-1)
		fmt.Println(s)
	}

}