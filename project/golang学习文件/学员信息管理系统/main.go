package main

import (
	"fmt"
	"os"
)

//需求：
//1.添加学员信息
//2编辑学员信息
//3展示所有学员信息
func showmenu() {
	fmt.Println("欢迎来到学员管理系统")
	fmt.Println("1.添加学员信息")
	fmt.Println("2.编辑学员信息")
	fmt.Println("3.展示所有学员信息")
	fmt.Println("4.退出")
}


//获取用户输入的学员信息
func getinput() *student {
	var(//声明变量也可以批量声明
		id int
		name string
		class string
	)
	fmt.Println("请安要求输入学员信息")
	fmt.Print("请输入学员序号")
	fmt.Scanf("%d\n",&id)
	fmt.Print("请输入学员姓名")
	fmt.Scanf("%s\n",&name)
	fmt.Print("请输入学员班级")
	fmt.Scanf("%s\n",&class)
	//调用造学生构造函数，造一个学生
	stu:= Newstundent(id,name,class)
	return stu
}


func main() {
	sm := newstudentmgr() //有一个管理类去调用ADD
	for {
		//1.打印系统菜单
		showmenu()
		//2等待用户选择
		var input int
		fmt.Println("请入你要操作的序号：")
		fmt.Scanf("%d\n", &input) //扫描用户输入，并且赋值给input，要改变其值，用指针
		fmt.Println("用户输入的是", input)
		//3执行选项
		switch input {
		//添加学生
		case 1:
			stu:=getinput()
			sm.addstudent(stu)
		case 2:
			//编辑学员信息
			stu:=getinput()
			sm.modifystudent(stu)
		case 3:
			sm.showstudent()
		//打印学员信息
		case 4:
			os.Exit(0)
		}

	}
}
