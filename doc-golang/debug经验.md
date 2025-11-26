

### go test 编译失败

```
slog.Info("启动服务", "port") // 不是成对的k-v
```
这行代码理论上是可以运行的。因为go test会不能通过。
报错: FAIL	github.com/Knowckx/go-demo/internal [build failed]
同时坑点在于他不会提示报错信息

手动去目录中执行go test  会有报错
.\run_env_test.go:10:2: call to slog.Info missing a final value

原因是go的vet工具很严格