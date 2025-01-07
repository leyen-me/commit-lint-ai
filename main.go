package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type RequestBody struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Choice struct {
	Message Message `json:"message"`
}

type Response struct {
	Choices []Choice `json:"choices"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Missing commit message file argument")
		os.Exit(1)
	}

	// 读取commit message文件内容
	commitMsg := []byte(os.Args[1])
	if commitMsg == nil {
		fmt.Println("Error: commit message is empty")
		os.Exit(1)
	}

	apiKey := os.Getenv("DEEPSEEK_API_KEY")
	if apiKey == "" {
		apiKey = "sk-"
	}

	// 构建请求体
	reqBody := RequestBody{
		Model: "deepseek-chat",
		Messages: []Message{
			{
				Role: "user",
				Content: fmt.Sprintf(`请检查以下 git commit message 是否符合规范，需要满足以下条件：
1. 必须以 feat:, fix:, docs:, style:, refactor:, test:, chore: 等类型开头
2. 描述必须清晰明了
3. 不能太长，建议不超过50个字符
4. commit message 必须用英文书写

Commit Message: '%s'

只需要回复 "通过" 或 "不通过：原因"。`, string(commitMsg)),
			},
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Printf("Error marshaling request: %v\n", err)
		os.Exit(1)
	}

	// 发送请求到API
	req, err := http.NewRequest("POST", "https://api.deepseek.com/v1/chat/completions", strings.NewReader(string(jsonData)))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		os.Exit(1)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		os.Exit(1)
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Printf("Error parsing response: %v\n", err)
		os.Exit(1)
	}

	if len(response.Choices) > 0 {
		result := response.Choices[0].Message.Content
		if strings.Contains(result, "不通过") {
			fmt.Println("❌ Commit message 格式检查失败：")
			fmt.Println(result)
			os.Exit(1)
		} else {
			fmt.Println("✅ Commit message 格式检查通过")
			os.Exit(0)
		}
	} else {
		fmt.Println("Error: Empty response from API")
		os.Exit(1)
	}
}
