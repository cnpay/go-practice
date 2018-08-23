package test

import "testing"

type Human interface {
	getName() string
}

type Student struct {
	name  string
	class string
}

func (s Student) getName() string { //struct只要实现了接口定义的方法，就算实现了接口
	return s.name
}

type Teacher struct {
	name   string
	course string
}

func (t Teacher) getName() string {
	return t.name
}

func printName(x Human) { //接口使用典型示例，面向接口类型编程
	println(x.getName())
	// println(x.(type)) //x.(type) 只能在type switch语句中使用
}

func Test_interface1(t *testing.T) {
	var s1 = Student{"jack", "class 3"}
	var t1 = Teacher{"liunix", "blockchain"}
	printName(s1)
	printName(t1)
}

func Test_interface2(t *testing.T) { //获取接口内部struct的方法
	var p Human = new(Student)
	p.(*Student).name = "jack" //找出实现
	println(p.getName())

	p = new(Teacher)
	p.(*Teacher).name = "liunix"
	println(p.getName())

}
