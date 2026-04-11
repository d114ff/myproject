package io

import (
	"compress/gzip"
	"compress/zlib"
	"fmt"
	"io"
	"os"
)

// 拷贝文件

func Copy(inFile, outFile string) {
	fin, err := os.Open(inFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	fout, err := os.OpenFile(outFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	/*
		bs := make([]byte, 1024)
		for {
			n, err := fin.Read(bs) // 内容读到bs里面去
			if err != nil {
				if err == io.EOF {
					if n > 0 {
						fout.Write(bs[:n]) //把bs前n个字节写到fout里面去
					}
				} else {
					fmt.Println(err)
				}
				break
			} else {
				fout.Write(bs[:n])
			}
		}
	*/
	io.Copy(fout, fin) //fin拷贝到fout里面去
	fout.Close()
	fin.Close()
}

// 文件压缩
func Compress(inFile, outFile string, compressAlgo int) {
	fin, err := os.Open(inFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	stat, _ := fin.Stat()
	fmt.Printf("压缩前文件大小 %dB\n", stat.Size())

	fout, err := os.OpenFile(outFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	//writer := gzip.NewWriter(fout)
	//writer := zlib.NewWriter(fout)
	var writer io.WriteCloser
	switch compressAlgo {
	case GZIP:
		writer = gzip.NewWriter(fout)
	case ZLIB:
		writer = zlib.NewWriter(fout)
	}
	io.Copy(writer, fin)
	writer.Close()
	fout.Close()
	fin.Close()
}

const (
	_    = iota //0, 如果担心0值滥用，可以不用0值
	GZIP        //1
	ZLIB        //2
)

// 解压
func Decompress(inFile, outFile string, compressAlgo int) {
	fin, err := os.Open(inFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	stat, _ := fin.Stat()
	fmt.Printf("压缩后文件大小 %dB\n", stat.Size())

	//reader, _ := gzip.NewReader(fin)
	//reader, _ := zlib.NewReader(fin)
	var reader io.ReadCloser
	switch compressAlgo {
	case GZIP:
		reader, _ = gzip.NewReader(fin)
	case ZLIB:
		reader, _ = zlib.NewReader(fin)
	}

	fout, err := os.OpenFile(outFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	io.Copy(fout, reader)
	reader.Close()
	fout.Close()
	fin.Close()
}
