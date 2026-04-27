package io_test

import (
	"fmt"
	"myproject/src/base_go/io"
	"testing"
	"time"
)

func TestWriteFile(t *testing.T) {
	io.WriteFile()
}
func TestReadFile(t *testing.T) {
	io.ReadFile()
}

func TestReadFileWithBuffer(t *testing.T) {
	io.ReadFileWithBuffer()
}

func TestWriteFileWithBuffer(t *testing.T) {
	io.WiteFileWithBuffer()
}

func TestBufferedFileWriter(t *testing.T) {
	t1 := time.Now()
	io.WiteDirect("../data/no_buffer.txt")
	t2 := time.Now()
	io.WriteWithBuffer("../data/with_buffer.txt")
	t3 := time.Now()
	fmt.Printf("不用缓冲耗时%dms，用缓冲耗时%dms\n", t2.Sub(t1).Milliseconds(), t3.Sub(t2).Milliseconds())
}

func TestCreateFile(t *testing.T) {
	io.CreateFile("../data/pmdpa.txt")
}

func TestWalkDir(t *testing.T) {
	io.WalkDir("../data/sys")

}

func TestSplitFile(t *testing.T) {
	io.SplitFile("../img/大乔乔好课.png", "../img/图像分割", 4)
}

func TestMergeFile(t *testing.T) {
	io.MergeFile("../img/图像分割", "../img/图像合并.png")
}

func TestLimitReader(t *testing.T) {
	io.LimitReader()
}

func TestMultiReader(t *testing.T) {
	io.MultiReader()
}
func TestMultiWriter(t *testing.T) {
	io.MultiWriter()
}

func TestTeeReader(t *testing.T) {
	io.TeeReader()
}

func TestPipeIO(t *testing.T) {
	io.PipeIO()
}

func TestCopy(t *testing.T) {
	io.Copy("../img/大乔乔好课.png", "../img/大乔乔好课2.png")
}

/*
	func TestCompress(t *testing.T) {
		io.Compress("../img/大乔乔好课.png", "../img/大乔乔好课.png.gzp")
	}

	func TestDecompress(t *testing.T) {
		io.Compress("../img/大乔乔好课.png", "../img/大乔乔好课.png.gzp")
		io.Decompress("../img/大乔乔好课.png.gzp", "../data/大乔乔好课.png")
	}
*/
func TestJsonSerialize(t *testing.T) {
	io.JsonSerialize()
}

func TestNewLoger(t *testing.T) {
	logger := io.NewLoger("../data/biz.log")
	io.Log(logger)
}

func TestSlog(t *testing.T) {
	logger := io.NewSlogger("../data/siz.log")
	io.Slog(logger)
}

func TestSysCall(t *testing.T) {
	io.SysCall()
}

func TestUseRegex(t *testing.T) {
	io.UseRegex()
}
