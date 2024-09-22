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