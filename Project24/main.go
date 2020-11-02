package main

type tester interface {
	test()
	string() string
}

type data struct{}

func (*data) test()         {}
func (data) string() string { return "" }

func main() {
	var d data
	/*	t := &d
		d.test()
		fmt.Println(d.string())
		t.test()
		fmt.Println(t.string())*/ //并未涉及接口data，方法可以混用
//	var e tester = d //不行接口内的方法中的test没有对应方法
//	e.string()
//	e.test()
	var f tester = &d//可以虽然接口内的方法中的string没有对应方法，但可以隐式转换
	f.string()
	f.test()
}
