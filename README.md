# hertz-backend-base
A backend basic web framework implemented by herzt

## some call demo for common utils
- file

```go

package main

import (
	"fmt"
	"herrz-backend-base/utils"
	"log"
)

func main() {
	fu := &utils.FileUtils{}

	// 写入JSON文件
	data := map[string]interface{}{
		"name": "Go",
		"type": "Programming Language",
		"lst":  []int{21, 3, 4, 5, 1},
		"map": map[string]interface{}{
			"a": "A",
			"b": "B",
		},
	}
	err := fu.WriteToJSON("./data/example.json", data, "w")
	if err != nil {
		log.Fatalf("Error writing to JSON: %v", err)
	}

	// 读取JSON文件
	loadedData, err := fu.LoadJSON("./data/example.json")
	if err != nil {
		log.Fatalf("Error loading JSON: %v", err)
	}
	fmt.Printf("Loaded data: %v\n", loadedData)
}

```

## usage demo of ctx

- mock demo

```go
package main

import (
	"context"
	"time"
	"log"
	"github.com/cloudwego/hertz/pkg/app"
)

func SearchAll(ctx context.Context, c *app.RequestContext) {
	// 创建带有 3 秒超时的 context
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel() // 在函数结束时取消超时 context

	// 模拟长时间的操作
	done := make(chan bool)
	go func() {
		// 假设这里是一些需要时间的操作，比如数据库查询
		time.Sleep(5 * time.Second) // 模拟 5 秒的操作
		done <- true
	}()

	// 处理超时或正常完成
	select {
	case <-ctx.Done(): // ctx 超时或取消时会进入这里
		log.Println("Request timed out")
		c.JSON(504, map[string]string{"error": "Request timed out"})
		return
	case <-done: // 操作完成
		log.Println("Operation completed")
		c.JSON(200, map[string]string{"message": "Operation completed successfully"})
	}
}

```