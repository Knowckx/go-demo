## 现代golang项目的标准目录结构

```
my-cli-app/
├── cmd/
│   └── myapp/           # 主程序入口
│       └── main.go      # 仅仅负责初始化和调用逻辑，尽量简短
├── internal/            # 私有业务逻辑（外部项目无法导入）
│   ├── app/             # 核心应用逻辑
│   └── config/          # 配置加载逻辑
├── pkg/                 # 公共库代码（可选，允许其他项目导入）
│   └── utils/
├── go.mod               # 依赖管理
├── go.sum
├── Makefile             # 构建脚本 (Windows下推荐使用 Taskfile)
└── README.md
```

## 核心技术栈

CLI 框架: Cobra (行业标准，Kubernetes、Docker 都在用)  
配置管理: Viper (通常与 Cobra 搭配)  
日志: log/slog (Go 1.21+ 内置标准库，强烈推荐，无需再引入 zap/logrus)  
UI/交互: Bubble Tea (如果你想做类似 TUI 的炫酷界面)