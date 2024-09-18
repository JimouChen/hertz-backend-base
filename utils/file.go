package utils

import (
	"fmt"
	"github.com/goccy/go-json"
	"log"
	"os"
	"path/filepath"
)

// FileUtils 用于文件操作的工具类
type FileUtils struct{}

// WriteToJSON 将字典或列表数据写入到JSON文件
func (fu *FileUtils) WriteToJSON(jsonPath string, data interface{}, mode string) error {
	err := fu.CheckDirPath(jsonPath)
	if err != nil {
		return err
	}
	var file *os.File

	switch mode {
	case "w":
		file, err = os.Create(jsonPath)
	case "w+", "a", "a+":
		file, err = os.OpenFile(jsonPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	default:
		return fmt.Errorf("unsupported mode: %s", mode)
	}

	if err != nil {
		return err
	}
	defer file.Close()

	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	_, err = file.Write(jsonData)
	return err
}

// LoadJSON 从JSON文件中读取并返回数据
func (fu *FileUtils) LoadJSON(jsonPath string) (interface{}, error) {
	content, err := os.ReadFile(jsonPath)
	if err != nil {
		return nil, err
	}

	if len(content) == 0 {
		log.Printf("Your json file in %s is empty!!!", jsonPath)
		return nil, nil
	}

	var data interface{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// CheckDirPath 检查目录路径是否存在，如果不存在则创建
func (fu *FileUtils) CheckDirPath(filePath string) error {
	// 提取文件路径的目录部分
	dirPath := filepath.Dir(filePath)

	// 检查目录是否存在
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// 如果目录不存在，则创建目录
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return err
		}
		log.Printf("mkdir path: %s", dirPath) // 使用 log 输出警告信息
	}
	return nil
}
