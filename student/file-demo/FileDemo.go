package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	file, err := os.ReadFile("test.json")
	fileName := filepath.Base("https://picx.zhimg.com/80/v2-827365267ea950b8039af0f1259f25ac_1440w.webp?source=1940ef5c")
	if err != nil {
		fmt.Println("读取文件错误:", err)
		return
	}

	configString := string(file)
	path := getOutputFilePath("https://picx.zhimg.com/80/v2-827365267ea950b8039af0f1259f25ac_1440w.webp?source=1940ef5c", "test")
	fmt.Println("配置文件内容:", configString)
	fmt.Println("文件名字 ", fileName)
	fmt.Println(path)
}

func getOutputFilePath(imageUrl, outputFormat string) string {
	fileName := filepath.Base(imageUrl)
	fileNameWithoutExt := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	return fileNameWithoutExt + "." + outputFormat

}
