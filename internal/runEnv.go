package internal

import (
	"log/slog"
	"path/filepath"

	"github.com/Knowckx/infa"
)

var (
	RootPath = ""
	PicPath  = ""
)

// init 函数会在包被导入时自动执行，且执行顺序在 main() 之前
func init() {
	var err error
	RootPath, err = infa.FindProjectRoot()
	if err != nil {
		// 既然是核心路径，如果找不到，程序通常无法继续运行，panic 是可以接受的
		slog.Error("严重错误：无法定位项目根目录", "err", err)
		panic(err)
	}

	// 顺便初始化关联路径
	PicPath = filepath.Join(RootPath, "cache", "pictures")
	slog.Info("项目根目录初始化完成", RootPath)
}
