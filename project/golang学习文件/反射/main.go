package main
//https://www.liwenzhou.com/posts/Go/13_reflect/网址
import (
	"fmt"
	"reflect"
)

//具体反射是什么，还不知道，但是知道反射涉及到类型和纸两个

//判断空接口传进来的是什么类型的数据
func reflecttype(x interface{}){
	a:=reflect.TypeOf(x)
	fmt.Println(a,a.Name(),a.Kind())//也就是说在反射出来的类型中，可再分出来名字和种类
	fmt.Printf("%v\n",a)
}

type cat struct {}
type dog struct {}

func reflectvalue(x interface{}){//定义一个值类型
	b:=reflect.ValueOf(x)
	fmt.Printf("%v     %T\n",b,b)
	//如何获取原始值
	k:=b.Kind()
	switch k {//主要看哪个作为对比参数
	case reflect.Float32:
	//把反射值转化成一个int32类型的变量
	ret:=float32(b.Float())
	fmt.Printf("%v %T\n",ret,ret)
	case reflect.Int32:
		res:=int32(b.Int())
		fmt.Printf("%v %T",res,res)
	}
}

func reflectsetvalue(x interface{}){
	a:=reflect.ValueOf(x)
	k:=a.Elem().Kind()//elem()  把指针转化成相对应的值,所以说这个案例是，b其实是一个指针
	switch k {
	case reflect.Int32:
		a.Elem().SetInt(26)//把值设置成一个int类型，（）内为设置的值
	case reflect.Float32:
		a.Elem().SetFloat(3.21)
	}
}
//通过反射修改某个值，注意的是地址操作，而不是值操作
func main() {
//var a float64=1.222
////reflecttype(a)
//reflectvalue(a)
//var b bool=true
////reflecttype(b)
//	reflectvalue(b)
//var d cat//这个猫结构体的种类里面的名字就是CAT,种类就是struct。具体有哪些种类见网页
////reflecttype(d)
//	reflectvalue(d)
var aa int32=100
reflectsetvalue(&aa)
fmt.Println(aa)
}
