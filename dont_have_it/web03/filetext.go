package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//打开文件
//file,bool:=os.OpenFile("./1.txt",os.)
//if bool!=nil{
//	fmt.Fprintln("66")
//}
file,bool:=os.Open("./1.txt")
	if bool!=nil{
		fmt.Fprintln("66")
	}
	//关闭文件
defer file.Close()
//读取文件
//file.Read()
reader,bool:=bufio.NewReader(file)
reader.ReadString()
b,err:=ioutil.ReadFile("./1.txt"
	if err!=nil {
		fmt.Println("666")
	}
ioutil.ReadAll()
//写入文件
file.WriteString()
w:=bufio.NewWriter(file)
w.Write()
w.WriteString("456465")
w.Flush()
ioutil.WriteFile()
}
