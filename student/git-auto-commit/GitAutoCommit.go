package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const (
	defaultPath     = "/Users/xiaozou/Desktop/sofe/java/weihubeats"
	defaultInterval = "10s"
	defaultJavaFile = "MQ.java"
	commitMsg       = "add test"
)

func main() {

	var projectPath, intervalStr string
	flag.StringVar(&projectPath, "path", defaultPath, "Path to the project")
	flag.StringVar(&intervalStr, "interval", defaultInterval, "Interval between commits (e.g., 30m, 2h)")
	flag.Parse()

	interval, err := time.ParseDuration(intervalStr)
	if err != nil {
		fmt.Println("Invalid interval format:", err)
		return
	}

	filePath := filepath.Join(projectPath, defaultJavaFile)

	if !checkFileExists(filePath) {
		err := os.WriteFile(filePath, []byte{}, 0644)
		if err != nil {
			fmt.Printf("Failed to create file %s: error: %s\n", defaultJavaFile, err)
			return
		}
		fmt.Printf("create file %s success/n", defaultJavaFile)

		/*if err := exec.Command("git", "-C", projectPath, "add", filePath).Run(); err != nil {
			fmt.Println("Failed to add HelloWord.java to Git:", err)
			return
		}*/
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	addSpace := true

	for range ticker.C {
		if err := modifyFile(filePath, addSpace); err != nil {
			fmt.Println("Failed to modify file:", err)
			return
		}
		if addSpace {
			fmt.Println("add space success")

		} else {
			fmt.Println("remove space success")
		}

		if err := commitChanges(projectPath, commitMsg); err != nil {
			fmt.Println("Failed to commit changes:", err)
			return
		}
		fmt.Println("push git success")
		addSpace = !addSpace

	}

}

func checkFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)

}

/*
*
追加空格
*/
func modifyFile(filePath string, addSpace bool) error {

	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var newContent string
	if addSpace {
		newContent = string(content) + " "
	} else {
		newContent = strings.TrimRight(string(content), " ")
	}
	return os.WriteFile(filePath, []byte(newContent), 0644)
}

func commitChanges(projectPath, message string) error {

	cmd := exec.Command("git", "add", ".")
	cmd.Dir = projectPath
	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("git", "commit", "-m", message)
	cmd.Dir = projectPath
	if err := cmd.Run(); err != nil {
		return err
	}
	fmt.Printf("git commit success message: %s\n", message)
	// 将master替换为您的远程分支名称
	cmd = exec.Command("git", "push", "origin", "master")
	cmd.Dir = projectPath
	return cmd.Run()
}
