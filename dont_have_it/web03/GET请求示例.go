package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp,err:=http.Get("https://www.liwenzhou.com/")
	if err!=nil{
		fmt.Println("555")
		return
	}
	defer resp.Body.Close()

	body,err:=ioutil.ReadAll(resp.Body)
	if err!=nil {
		fmt.Println("56")
	}
	return
	fmt.Print(string(body))
}

//怎么打印啊