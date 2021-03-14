package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readall(){
	fileobj,err:=os.Open("./xx.txt")//打开文件
	if err!=nil{
		fmt.Println("出错啦")
		return
	}
	defer fileobj.Close()//关闭文件
	//读取文件
	for{
		var tmp=make([]byte,128)//设置了读取的字节数大小
		n,err:=fileobj.Read(tmp)//读取的是从文件获得的字节数，也就是读取了多少字节
		if err==io.EOF {//打印文件末尾
			fmt.Println(string(tmp[:n]))
			return
		}
		if err!=nil{
			fmt.Println("打印错误")
			return
		}
		fmt.Printf("read %d bytes from files.\n",n)
		fmt.Println(string(tmp[:n]))
	}
}
func readbybufio(){
	//打开文件+关闭文件
	fileobj,err:=os.Open("./xx.txt")//打开文件
	if err!=nil{
		fmt.Println("出错啦")
		return
	}
	defer fileobj.Close()//关闭文件
	reader:=bufio.NewReader(fileobj)
	line,err:=reader.ReadString('\n')
if err!=nil{
	fmt.Println("错误")
	return
}
fmt.Println(line)
}
func write(){
	fileobj,err:=os.OpenFile("./xx.txt",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0644)

}
func main() {
//readall()
	readbybufio()
}
