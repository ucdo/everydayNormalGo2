## 前置条件
1. 安装go
2. 安装MingW 64bit
3. 设置 ```go env -w GOBIN=\path\to\goinstallpath\bin```
4. 执行 ```go install fyne.io/fyne/v2/cmd/fyne@latest```

## 运行
1. 定义 一个app  appx:= app.New() . 额不能叫app因为和导入的包同名了
2. 执行时 window.ShowAndRun()  appx.Run()
3. 定义窗口 window := appx.NewWindow("window name")
4. 往窗口里写内容 window.SetContent(widget.NewLabel("hello"))
5. window.Show()展示窗口，放在appx.Run() 的前面
6. 设置PC端窗口的大小 window.Resize(fyne.NewSize(100,100)) // 移动端不搞这个，因为仅仅以全屏展示
7. 经过测试发现 widget.NewButton().Resize() 不起作用
8. 