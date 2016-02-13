package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	var input string;
	fmt.Scanln(&input)
	var data []map[string]string
	res,_ := http.Get("http://www.google.com/finance/info?q="+input)
	reqJson,_ := ioutil.ReadAll(res.Body);
	json.Unmarshal(reqJson[3:], &data)
	fmt.Println("Index of ",input,":",data[0]["l_fix"])
	res.Body.Close();

}
