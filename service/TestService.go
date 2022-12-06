package service

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

var loopCount = 100000000
func TestA(){
	//testStruct()
	testString1()

}

func TestB()  {
	//testMap()
	testString2()
}

func TestC()  {
	testString3()
}

func TestD()  {
	testString4()
}

func testStruct(){
	t1 := time.Now().UnixNano()/1e6
	fmt.Println("t1:",t1)
	for i:=0;i<loopCount;i++{
		var test struct{
			name string
			age int8
		}
		test.name = "dan"
		test.age = 28
	}
	t2 := time.Now().UnixNano()/1e6
	fmt.Println("t2:",t2)
	fmt.Println("t2-t1:",t2-t1)
}

func testMap(){
	t1 := time.Now().UnixNano()/1e6
	fmt.Println("t1:",t1)
	for i:=0;i<loopCount;i++{
		var test = map[string]interface{}{}
		test["name"] = "dan"
		test["age"] = 28
	}
	t2 := time.Now().UnixNano()/1e6
	fmt.Println("t2:",t2)
	fmt.Println("t2-t1:",t2-t1)
}

func testString1()  {
	t1 := time.Now().UnixNano() / 1e6
	fmt.Println("t1 == ", t1)

	s := "xiao"
	for i := 0; i < 500000; i++ {
		s += "motong"
	}

	t2 := time.Now().UnixNano() / 1e6
	fmt.Println("t2 == ", t2)
	fmt.Println("t2 - t1 == ", t2-t1)
}

func testString2()  {
	t1 := time.Now().UnixNano() / 1e6
	fmt.Println("t1 == ", t1)

	s := "xiao"
	for i := 0; i < 500000; i++ {
		s = fmt.Sprintf("%s%s",s,"motong")
	}

	t2 := time.Now().UnixNano() / 1e6
	fmt.Println("t2 == ", t2)
	fmt.Println("t2 - t1 == ", t2-t1)
}

func testString3()  {
	t1 := time.Now().UnixNano() / 1e6
	fmt.Println("t1 == ", t1)

	s := []string{}
	s = append(s,"xiao")
	for i := 0; i < 500000; i++ {
		s = append(s ,"motong")
	}
	strings.Join(s,"")

	t2 := time.Now().UnixNano() / 1e6
	fmt.Println("t2 == ", t2)
	fmt.Println("t2 - t1 == ", t2-t1)
}

func testString4()  {
	t1 := time.Now().UnixNano() / 1e6
	fmt.Println("t1 == ", t1)

	s := bytes.NewBufferString("xiao")
	for i := 0; i < 500000; i++ {
		s.WriteString("motong")
	}

	t2 := time.Now().UnixNano() / 1e6
	fmt.Println("t2 == ", t2)
	fmt.Println("t2 - t1 == ", t2-t1)
}