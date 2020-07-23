/*
 * 结构体
 * @Author: cxx<cxx2320@foxmail.com>
 * @Date: 2020-07-22 19:05:23
 * @LastEditors: cxx
 * @LastEditTime: 2020-07-23 10:02:47
 */

package main

// 为fmt定义别名
import mm "fmt"

// 定义一个对象 （有a和b成员）
type Mutatable struct {
	// 后面的字符串是对变量的tag
	a int `一个数字a`
	b int `一个数字b`
}

// 为对象Mutatable添加一个StayTheSame方法
func (m Mutatable) StayTheSame() {
	m.a = 5
	m.b = 7
}

// 为对象Mutatable添加一个Mutate方法，且引用原来的对象（使用*号引用）在此函数内修改函数成员会作用到调用者身上
func (m *Mutatable) Mutate() {
	m.a = 5
	m.b = 7
}

// 程序运行入口（在init函数运行后执行）
func main() {
	// 实例化一个Mutatable对象（&用处未知）？
	// m := &Mutatable{10, 20}

	// 另一种实例化对象方式（和上一种的区别）？
	// m := new(Mutatable)

	// 另一种实例化对象方式（和上一种的区别）？
	var m Mutatable

	// 打印m对象
	mm.Println(m)

	// 调用StayTheSame方法
	m.StayTheSame()
	mm.Println(m)
	m.Mutate()
	mm.Println(m)
}

// 初始化函数（在此使用Println打印时数字正常.字符串报错[定义变量，直接打印都会]）？
func init() {
	// mm.Println(1235464)
}
