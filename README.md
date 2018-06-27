# scf-go-lib


用于 SCF GO 环境的库及工具


## GO 环境开发说明

```
package main

import (
	"context"
	"fmt"

	"github.com/tencentyun/scf-go-lib/cloudfunction"
)

type DefineEvent struct {
	// test event define
	Key1 string `json:"key1"`
	Key2 string `json:"key2"`
	Key3 string `json:"key3"`
}

func hello(ctx context.Context, event DefineEvent) (string, error) {
	fmt.Println("key1:", event.Key1)
	fmt.Println("key2:", event.Key2)
	fmt.Println("key3:", event.Key3)
	return fmt.Sprintf("Hello %s!", event.Key1), nil
}

func main() {
	// Make the handler available for Remote Procedure Call by Cloud Function
	cloudfunction.Start(hello)
}

```

1. 需要使用 package main 包含 main 函数。
2. 引用 `github.com/tencentyun/scf-go-lib/cloudfunction` 库。
3. 入口函数入参可选 0~2 参数，如包含参数，需 context 在前，event 在后，入参组合有 （），（event），（context），（context，event）。
4. 入口函数返回值可选 0~2 参数，如包含参数，需 返回内容在前，error 错误信息在否，返回值组合有 （），（ret），（error），（ret，error）。
5. 入参 event ，和返回值 ret，均需要能够兼容 `encoding/json` 标准库，可以进行 Marshal、Unmarshal。
6. 在 main 函数中使用包内的 Start 函数启动入口函数


## GO 环境编译说明

使用指定 OS 及 ARCH 即可跨平台编译为二进制，随后通过 zip 工具打包二进制，生成可以上传的代码包。

在 Linux 或 MacOS 下可使用如下命令编译及打包

```
GOOS=linux GOARCH=amd64 go build -o main main.go
zip main.zip main
```

在 Windows 下可使用如下命令编译，然后使用打包工具对输出的二进制文件进行打包，二进制文件需要在 zip 包根目录。

```
set GOOS=linux
set GOARCH=amd64
go build -o main main.go
```

## 云函数创建说明

创建云函数时，GO 环境仅支持 zip 包上传。其中入口函数位置填写 *二进制文件名* 即可，无需按文件名.函数名格式填写。


