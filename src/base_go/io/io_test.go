package io_test

import (
	"myproject/src/base_go/io"
	"testing"
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
