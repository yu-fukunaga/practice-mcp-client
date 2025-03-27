package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"
)

type MCPServerConfig struct {
	Command string            `json:"command"`
	Args    []string          `json:"args"`
	Env     map[string]string `json:"env"`
}

func (m *MCPServerConfig) GetEnv() []string {
	var envs []string
	for k := range m.Env {
		envs = append(envs, k) // fmt.Sprintfを削除
	}
	return envs
}

type Config struct {
	MCPServers map[string]MCPServerConfig `json:"mcpServers"`
}

type ServerProcess struct {
	Cmd    *exec.Cmd
	Stdin  io.WriteCloser
	Stdout io.ReadCloser
}

func StartMCPServers(configPath string) (map[string]*client.StdioMCPClient, error) {
	ctx := context.Background()
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("設定ファイルの読み込みに失敗しました: %w", err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("JSONのパースに失敗しました: %w", err)
	}

	clients := make(map[string]*client.StdioMCPClient)
	for name, serverConfig := range config.MCPServers {
		log.Printf("サーバー '%s' を起動します", name)
		client, err := client.NewStdioMCPClient(serverConfig.Command, serverConfig.GetEnv(), serverConfig.Args...)
		if err != nil {
			log.Printf("サーバー %s の起動に失敗: %v", name, err)
			continue
		}
		if _, err := client.Initialize(ctx, mcp.InitializeRequest{}); err != nil {
			log.Printf("サーバー '%s' の初期化に失敗: %v", name, err)
			continue
		}
		clients[name] = client
	}
	return clients, nil
}
