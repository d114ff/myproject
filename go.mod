module myproject

go 1.25.4

// go get github.com/bytedance/sonic 把文件源代码拉到本地，放在gonpasth/pkg/mod 目录下
// go mod tidy 把需要依赖下载下来，不需要删掉

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/cloudwego/kitex v0.16.1 // indirect
	github.com/go-sql-driver/mysql v1.9.3 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	google.golang.org/genproto v0.0.0-20210513213006-bf773b8c8384 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)
