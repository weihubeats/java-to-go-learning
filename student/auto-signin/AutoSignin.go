package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"sync"
)

type SigninTask struct {
	Name    string // 签到任务名
	CurlCmd string
}

func main() {
	tasks := []SigninTask{
		{Name: "音乐自动签到", CurlCmd: ""},
		{Name: "掘金签到", CurlCmd: ""},
	}

	var wg sync.WaitGroup

	for _, task := range tasks {
		wg.Add(1)
		defer wg.Done()
		result, err := executeCurl(task.CurlCmd)
		if err != nil {
			fmt.Printf("[%s]签到失败%v\n", task.Name, err)
		}
		fmt.Printf("[%s]签到成功%s\n", task.Name, result)
	}
	wg.Wait()
}

func executeCurl(curlCmd string) (string, error) {
	cmd := exec.Command("bash", "-c", curlCmd)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}
