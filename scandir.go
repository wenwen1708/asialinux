package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"
)

func main() {
	u, _ := user.Current()
	fmt.Println("当前用户：", u.Name)
	dirname := u.HomeDir
	postfix := readinput()
	//dirname := "c:\\MyGo\\src\\tmp"
	fmt.Println("开始查找，请稍等 ...")
	listfile(dirname, postfix, 0)
	fmt.Println("查找结束，结果已写入当前目录中的test.txt文件 ...")

}

func listfile(dirname string, postfix string, level int) {
	n := 0
	s := "|--"
	for i := 0; i < level; i++ {
		s = "| " + s
	}
	fileinfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		fmt.Println("MMP 出错了。。。")
		log.Fatal(err)
		return
	}
	for _, fi := range fileinfos {
		filename := dirname + "\\" + fi.Name()
		s1 := s + filename + "\n"
		ps := []byte(postfix)
		var newbyte []byte = ps[:len(ps)-2]
		x := string(newbyte)
		bl := strings.Contains(fi.Name(), x)
		if bl == true {
			writeresult(s1)
			n++
		}
		if fi.IsDir() && err == nil {
			listfile(filename, postfix, level+1)
		}
	}
	fmt.Printf("找到%d个符合的文件。\n", n)
}

func writeresult(str string) {
	filename := "./test.txt"
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	w1 := bufio.NewWriterSize(file, 1024)
	w1.WriteString(str)
	w1.Flush()
}
func printerr(info error) {
	if info != nil {
		fmt.Println("err :", info)
	}
}

func readinput() (input string) {
	inputreader := bufio.NewReader(os.Stdin)
	fmt.Printf("这是一个在当前用户查找包含指定关键字文件的程序\n请输入要查找的关键字如: 苍井空 回车结束输入 :")
	input, err := inputreader.ReadString('\n')
	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
	}
	fmt.Println(input)
	return input
}
