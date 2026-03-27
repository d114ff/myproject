package io

// buffered.go - 带缓冲的文件写入实现
// 通过减少系统调用次数提升大量小数据写入的性能

import "os"

const (
	logText = "黄梅时节家家雨，青草池塘处处蛙。有约不来过夜半，闲敲棋子落灯花。\n"
)

// BufferedFileWriter 带缓冲的文件写入器，减少磁盘I/O次数
// 内部维护一个缓冲区，先将数据写入缓冲区，待缓冲区满或主动刷新时再一次性写入磁盘
// 适用于频繁写入少量数据的场景，可显著减少系统调用开销
type BufferedFileWriter struct {
	fout           *os.File // 目标文件句柄
	buffer         []byte   // 内存缓冲区，用于暂存待写入的数据
	bufferEndIndex int      // 缓冲区当前已占用的大小，即有效数据结束位置
}

// NewBufferedFileWriter 创建带缓冲的文件写入器
// 参数:
//   - fout: 已打开的可写入文件句柄
//   - bufferSize: 缓冲区大小，建议值为4096（页大小）或其倍数
//
// 返回: 初始化好的BufferedFileWriter指针
func NewBufferedFileWriter(fout *os.File, bufferSize int) *BufferedFileWriter {
	return &BufferedFileWriter{
		fout:           fout,
		buffer:         make([]byte, bufferSize),
		bufferEndIndex: 0,
	}
}

// Flush 将缓冲区中的数据刷新到磁盘
// 调用此方法会立即将缓冲区内的所有数据写入文件，并重置缓冲区索引
// 建议在写入完成后主动调用，或使用defer确保数据不会丢失
func (w *BufferedFileWriter) Flush() {
	w.fout.Write(w.buffer[0:w.bufferEndIndex])
	w.bufferEndIndex = 0
}

// Write 将数据写入缓冲区或直接写入文件
// 写入策略:
//   - 如果单次写入数据 >= 缓冲区大小，先刷新缓冲区，再直接写入磁盘
//   - 如果缓冲区剩余空间足够，将数据复制到缓冲区
//   - 如果缓冲区剩余空间不足，先刷新缓冲区，再将数据复制到缓冲区
func (w *BufferedFileWriter) Write(cont []byte) {
	if len(cont) >= len(w.buffer) {
		w.Flush()
		w.fout.Write(cont)
	} else {
		if w.bufferEndIndex+len(cont) > len(w.buffer) {
			w.Flush()
		}
		copy(w.buffer[w.bufferEndIndex:], cont)
		w.bufferEndIndex += len(cont)
	}

}

// WriteString 将字符串写入缓冲区
// 内部将string转换为[]byte后调用Write方法
func (w *BufferedFileWriter) WriteString(cont string) {
	w.Write([]byte(cont))
}

// WiteDirect 直接写入文件（无缓冲）
// 每次循环都调用一次文件写入操作，产生大量系统调用
// 用于与缓冲写入性能对比测试
func WiteDirect(outFile string) {
	fout, err := os.OpenFile(outFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer fout.Close()

	for i := 0; i < 100000; i++ {
		fout.WriteString(logText)
	}
}

// WriteWithBuffer 带缓冲写文件
// 使用BufferedFileWriter进行写入测试，缓冲区大小为4096字节
// 相比直接写入，大幅减少系统调用次数，提升写入性能
func WriteWithBuffer(outFile string) {
	fout, err := os.OpenFile(outFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer fout.Close()

	writer := NewBufferedFileWriter(fout, 4096)

	defer writer.Flush()
	for i := 0; i < 100000; i++ {
		//writer.WriteString(logText)
		writer.Write([]byte(logText))

	}
}
