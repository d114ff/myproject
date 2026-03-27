package io

import (
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

func SplitFile(infile string, outDir string, n int) { // 切分文件
	fin, err := os.Open(infile) // 打开文件
	if err != nil {
		log.Panic(err) // 打开文件失败
	}
	defer fin.Close() // 关闭文件

	stat, err := fin.Stat() // 获取文件信息
	if err != nil {
		log.Panic(err) // 获取文件信息失败
	}
	fileSize := stat.Size()      // 获取文件大小，单位为字节字节
	chunk := fileSize / int64(n) // 计算每个切分的大小，单位为字节字节

	if chunk <= 0 {
		panic("file is too small or n is too large") // 文件过小或n过大
	}

	for i := 0; i < n; i++ {
		fout, err := os.OpenFile(path.Join(outDir, strconv.Itoa(i)+"_"+path.Base(infile)), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
		if err != nil {
			log.Panic(err) // 打开文件失败
		}
		defer fout.Close() // 关闭文件
		need := int(chunk) // 计算每个切分的大小，单位为字节字节
		if i == n-1 {
			need = int(fileSize) - (n-1)*int(chunk)
		}
		buffer := make([]byte, need) // 读取文件
		_, err = fin.Read(buffer)
		if err != nil {
			log.Panic(err)
		}
		_, err = fout.Write(buffer) // 写入文件
		if err != nil {
			log.Panic(err)
		}
		fout.Close() // 关闭文件
	}

}

func AppendFile(fout *os.File, infile string) {
	fin, err := os.Open(infile)
	if err != nil {
		log.Panic(err)
	}
	defer fin.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := fin.Read(buffer)
		if err != nil {
			if err == io.EOF {
				if n > 0 {
					fout.Write(buffer[:n])
				}
			} else {
				log.Println(err)
			}
			break
		} else {
			fout.Write(buffer[:n])
		}
	}
}

// 把dir这个目录下的所有文件合并到mergedFile里去
func MergeFile(dir string, mergedFile string) {
	fout, err := os.OpenFile(mergedFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Panic(err)
	}
	defer fout.Close()
	if fileInfos, err := os.ReadDir(dir); err != nil {
		log.Panic(err)
	} else {
		for _, fileInfo := range fileInfos {
			if fileInfo.Type().IsRegular() {
				infile := filepath.Join(dir, fileInfo.Name())
				AppendFile(fout, infile)
			}
		}
	}
}
