/**
 * ASE/lab1-2/menu.go
 *
 * Copyright (c) 2000, 2023 shuaikai SA22***200
 *
 * Implemented a single linkelist to store commands and their descriptions,
 * And use a callback function to handle the business logic of different commands,
 * to maintain stability of the core.
 *
 * TODO: 把链表定义和 main 函数分离开，就像 C/C++ 中那样
 */

package main

import (
	"fmt"
)

// 单链表
type LinkTable struct {
	head  *node
	tail  *node
	count int
}

// 节点
type node struct {
	cmd     string
	desc    string
	handler func(*LinkTable)
	next    *node
}

// 创建一个空链表
func New() *LinkTable {
	return &LinkTable{nil, nil, 0}
}

// 添加元素 -> 头插法
func (t *LinkTable) AddNode(cmd string, desc string, handler func(*LinkTable)) {
	if t.head == nil {
		t.head = &node{cmd, desc, handler, nil}
		t.tail = t.head
	} else {
		newnode := new(node)
		newnode.cmd = cmd
		newnode.desc = desc
		newnode.handler = handler
		newnode.next = t.head
		t.head = newnode
	}
	t.count++
}

// 追加元素 -> 尾插法
func (t *LinkTable) AppendNode(cmd string, desc string, handler func(*LinkTable)) {
	if t.head == nil {
		t.head = &node{cmd, desc, handler, nil}
		t.tail = t.head
	} else {
		t.tail.next = &node{cmd, desc, handler, nil}
		t.tail = t.tail.next
	}
	t.count++
}

// 链表长度
func (t *LinkTable) Len() int {
	return t.count
}

// 删除命令
func (t *LinkTable) Remove(cmd string) {
	if t.head.next == nil {
		panic("Empty List!")
	}
	cur := t.head
	for cur != nil {
		if cur.cmd == cmd {
			removed := cur
			if removed.next == nil {
				t.head = nil
				t.tail = nil
			} else {
				t.head = removed.next
			}
			t.count--
		}
		cur = cur.next
	}
}

// 找到命令返回指针
func (t *LinkTable) FindCmd(cmd string) *node {
	cur := t.head
	for cur != nil {
		if cur.cmd == cmd {
			return cur
		}
		cur = cur.next
	}
	return nil
}

// 输出元素
func (t *LinkTable) PrintList() {
	i := 1
	cur := t.head
	for cur != nil {
		fmt.Printf("cmd %d: %s -> %s\n", i, cur.cmd, cur.desc)
		i++
		cur = cur.next
	}
}

func main() {
	list := New()
	list.AppendNode("help", "show this help text", help)
	list.AppendNode("version", "2.0.1", version)
	list.AppendNode("exit", "quit this menu", nil)
	list.AppendNode("hello", "give you a greeting from the author", hello)

	for {
		fmt.Printf("==> Please give your command $ ")

		var in_cmd string
		fmt.Scanln(&in_cmd)

		ptr := list.FindCmd(in_cmd)

		if ptr == nil {
			fmt.Println("ERROR: Illeagl Command!\n\tPlease cheak out your input, or get some help by type \"help\"")
			continue
		} else if ptr.cmd == "exit" {
			break
		} else {
			ptr.handler(list)
		}

	}
}

func help(list *LinkTable) {
	list.PrintList()
}

func version(list *LinkTable) {
	cur := list.FindCmd("version")
	fmt.Println(cur.desc)
}

func hello(list *LinkTable) {
	cur := list.FindCmd("hello")
	fmt.Print(cur.desc)
	fmt.Println(", and that means:\n>>> Hello you beautiful | handsome GUY! <<< ")
}

// func todo(*LinkTable) {
// 	fmt.Println("Haven't implement...")
// }
