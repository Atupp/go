package main

import "fmt"

//学生接结构体
type student struct {
	id    int
	name  string
	class string
}
//对象为学生，的构造函数
func Newstundent(id int, name, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}



//学生管理结构体
type studentmgr struct {
	allstudents []*student //是一个student结构体的一个切片的地址？？？
}

//对象为学生管理，的构造函数
func newstudentmgr() *studentmgr {
	return &studentmgr{
		make([]*student,0,100),
	}
}


//对象为学生管理，的添加学生的方法
func(s *studentmgr)addstudent(newstu *student){
	s.allstudents=append(s.allstudents,newstu)
}
//编辑学生,也就是说你输入的是一个学生的结构体
func(s *studentmgr)modifystudent(newstu *student){
	for i,v:= range s.allstudents{
		if newstu.id==v.id{//当学号相同时，就替换内容
			s.allstudents[i]=newstu
			return
		}
	}
	//如果走到这里说明输入的学生没有找到
	fmt.Printf("输入学生有误，没有这个学生,其学号为：%d\n",newstu.id)
}
//展示学生
func(s *studentmgr)showstudent(newstu *student){
	for _,v:=range s.allstudents{
		fmt.Printf("学号：%d 姓名：%s 班级：%s\n",v.id,v.name,v.class)//格式化输出
	}
}