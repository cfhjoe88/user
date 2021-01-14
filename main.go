package main

import (
	"crypto/md5"
	"fmt"
	"github.com/cfhjoe88/utils"
)

func auth()bool  {
	passWord:=md5.Sum([]byte("123abc!@#"))
	p:=fmt.Sprintf("%x",passWord)
	fmt.Println(p)
	for i:=0;i<3;i++{
		if utils.Md5text( utils.Input("请输入密码"))==p {
			fmt.Println("密码输入错误")
		}
	}
	return false
	
}

func main()  {
	
}