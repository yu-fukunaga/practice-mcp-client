package main

import (
	"context"
	"fmt"
	"log"

	my_mcp "practice-go-mcp-client/mcp"

	"github.com/mark3labs/mcp-go/mcp"
)

func main() {
	clients, err := my_mcp.StartMCPServers("config.json")
	if err != nil {
		log.Printf("MCPサーバー起動エラー: %v", err)
		return
	}

	ctx := context.Background()

	for name, client := range clients {
		if err := client.Close(); err != nil {
			// エラーハンドリングを追加
			log.Printf("Error closing client: %v", err)
		}

		log.Printf("サーバー '%s' のツール一覧を取得します", name)
		tools, err := client.ListTools(ctx, mcp.ListToolsRequest{})
		if err != nil {
			log.Printf("サーバー '%s' のツール一覧取得に失敗: %v", name, err)
			continue
		}

		var toolNames []string
		for _, tool := range tools.Tools {
			toolNames = append(toolNames, tool.Name)
		}
		fmt.Printf("サーバー '%s' の利用可能なツール: %v\n", name, toolNames)
	}
}

// TODO: MCPのToolsを利用できるようにしたい。
// func callOpenAIAPI(prompt string) (string, error) {
// 	client := openai.NewClient()

// 	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
// 		Messages: []openai.ChatCompletionMessageParamUnion{
// 			openai.SystemMessage("あなたは、与えられたツール一覧から実行すべきツール名を1つだけ返すアシスタントです。"),
// 			openai.UserMessage(prompt),
// 		},
// 		Model: openai.ChatModelGPT4o,
// 	})
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	return strings.TrimSpace(chatCompletion.Choices[0].Message.Content), nil
// }
