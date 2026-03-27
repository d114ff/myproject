package io

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func CreateFile(fileName string) {
	os.Remove(fileName) //先删除，不去理会Remove可能返回的错误error
	if file, err := os.Create(fileName); err != nil {
		fmt.Printf("创建文件%s失败，错误信息：%s\n", fileName, err)
	} else {
		file.Chmod(0o666)                                      //设置权限
		fmt.Printf("创建文件%s成功，文件描述符：%d\n", fileName, file.Fd()) //获取文件描述符file descriptor，这是一个整数
		file.WriteString("问君能有几多愁\n")                          //写入字符串
		info, _ := file.Stat()                                 //获取文件信息
		fmt.Printf("is dir %t\n", info.IsDir())                //判断是否是目录
		fmt.Printf("modify time %s\n", info.ModTime())         //获取文件的修改时间
		fmt.Printf("mode %v\n", info.Mode())                   //获取文件的权限
		fmt.Printf("file name %s\n", info.Name())              //获取文件的名字
		fmt.Printf("size %dB\n", info.Size())                  //获取文件的大小，单位是字节
		file.Close()                                           //关闭文件
	}
	os.Mkdir("../data/sys", os.ModePerm)          //创建目录并设置权限
	os.MkdirAll("../data/sys/a/b/c", os.ModePerm) //增强版Mkdir，沿途的目录不存在时会一并创建
	//os.Rename("../data/sys/a", "../data/sys/p")       //给文件或目录重命名
	//os.Rename("../data/sys/p/b/c", "../data/sys/p/c") //Rename还可以实现move的功能
	//os.Remove("../data/sys")                          //删除文件或目录，目录不为空时才能删除成功
	//os.RemoveAll("../data/sys")                       // WalkDir 遍历指定目录，打印所有子目录和文件的信息
	// 参数：
	//   path: 要遍历的目录路径
	// 返回值：
	//   error: 遍历过程中遇到的错误，如果没有错误则返回 nil
	// 使用 filepath.Walk 递归遍历目录
}

func WalkDir(path string) error {

	filepath.Walk(path, func(subPath string, info fs.FileInfo, err error) error {
		// 处理遍历过程中遇到的错误
		if err != nil {
			return err
		} else if info.Mode().IsDir() && subPath != path {
			// 如果是子目录（排除起始目录），打印目录路径
			fmt.Printf("path is dir %s\n", subPath)
		} else if info.Mode().IsRegular() {
			// 如果是普通文件，打印文件路径和文件名
			fmt.Printf("path is file %s basename %s\n", subPath, info.Name())
		}
		// 返回 nil 继续遍历
		return nil
	})
	// 遍历完成，返回 nil 表示没有错误
	return nil
}
