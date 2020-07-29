package main

import (
	"fmt"
)

type Log struct {
	msg string
}

type Customer struct {
	Name string
	Log
}

func main() {
	// 加上 & 符号后把c变量赋值给b后，修改b的属性Name时会影响到c的Name值（引用）
	c := &Customer{"Barak Obama", Log{"1 - Yes we can!"}}
	// c.Name = "Cxx"
	// b := c
	// b.Name = "XXc"
	// fmt.Println(c.Name)

	c.Add("2 - After me the world will be a better place!")
	// fmt.Println(c)
	fmt.Println(c.Log.String())
}

func (l *Log) Add(s string) {
	l.msg += "---" + s
}

func (l *Log) String() string {
	return l.msg
}

func (c *Customer) String() string {
	return c.Name + "\nLog:" + fmt.Sprintln(c.Log)
}
