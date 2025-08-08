package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"pdfcpu/pdf"
	"pdfcpu/util"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
	// 设置log输出到控制台和文件
	logFile, err := os.OpenFile("pdfcpu.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("无法创建日志文件: %v", err)
	}
	// 将日志同时输出到控制台和文件
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)
}

/**
最终的二进制可执行文件实现
命令行中输入
程序名 密码 路径
之后
执行
*/

func main() {
	args := os.Args
	if len(args) < 3 {
		log.Fatalf("需要输入: 程序名 密码 文件名")
	}
	passwd := args[1]
	log.Printf("密码设置为%v\n", passwd)
	fp := args[2]
	log.Printf("文件名为%v\n", fp)
	dir := filepath.Dir(fp)
	pdfs, err := util.GetPDFFiles(dir)
	if err != nil {
		log.Fatalf("搜索同路径下文件发生错误:%v\n", err)
	}
	for i, p := range pdfs {
		log.Printf("正在处理第%d/%d个文件", i+1, len(pdfs))
		e := pdf.SetPasswd(p, passwd)
		if e != nil {
			log.Fatalf("在处理第%d个文件%s的时候发生错误:%v\t中断操作\n", i+1, p, err)
		}
	}
}
