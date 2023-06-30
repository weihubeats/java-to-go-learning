package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	url := "https://open.feishu.cn/open-apis/bot/v2/hook/bd8a9a5d-9657-4d38-be60-bbcf36a2b058"
	ss := "{\n    \"msg_type\": \"interactive\",\n    \"card\":{\n  \"elements\": [\n    {\n      \"tag\": \"markdown\",\n      \"content\": \"sourceTopic: %s\"\n    },\n    {\n      \"tag\": \"div\",\n      \"text\": {\n        \"content\": \"targetTopic: %s\",\n        \"tag\": \"plain_text\"\n      }\n    },\n    {\n      \"tag\": \"div\",\n      \"text\": {\n        \"content\": \"sourceNum: %s\",\n        \"tag\": \"plain_text\"\n      }\n    },\n    {\n      \"tag\": \"div\",\n      \"text\": {\n        \"content\": \"targetNum: %s\",\n        \"tag\": \"plain_text\"\n      }\n    }\n  ],\n  \"header\": {\n    \"template\": \"blue\",\n    \"title\": {\n      \"content\": \"mq消息同步\",\n      \"tag\": \"plain_text\"\n    }\n  }\n}\n}"
	sprintf := fmt.Sprintf(ss, "test", "test1", strconv.Itoa(1), strconv.Itoa(13))

	req, err := http.NewRequest("GET", url, bytes.NewBuffer([]byte(sprintf)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:")
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	fmt.Println(buf.String())

}
