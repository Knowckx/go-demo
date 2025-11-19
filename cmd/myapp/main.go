package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

// 根命令
var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "这是一个专业的 Go 控制台程序示例",
	Run: func(cmd *cobra.Command, args []string) {
		// 获取上下文
		ctx := cmd.Context()
		runBusinessLogic(ctx)
	},
}

func main() {
	// 1. 设置结构化日志 (JSON格式适合生产环境，Text适合开发环境)
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// 2. 优雅退出 (Graceful Shutdown) 上下文设置
	// 监听 Ctrl+C (SIGINT) 和 终止信号 (SIGTERM)
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	// 3. 执行命令
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		slog.Error("程序执行失败", "error", err)
		os.Exit(1)
	}
}

// 模拟业务逻辑
func runBusinessLogic(ctx context.Context) {
	slog.Info("程序启动", "go_version", "1.25+")

	// 模拟一个长耗时任务
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			slog.Warn("接收到退出信号，正在清理资源...")
			// 可以在这里做关闭数据库、保存文件等操作
			time.Sleep(500 * time.Millisecond) // 模拟清理耗时
			slog.Info("安全退出")
			return
		case t := <-ticker.C:
			slog.Info("正在工作...", "time", t.Format(time.TimeOnly))
		}
	}
}