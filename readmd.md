golang测试项目说明：

## debug说明

### 普通文件debug
https://github.com/golang/vscode-go/blob/master/docs/debugging.md#snippets

### 测试文件debug
注意项目如果是符号链接到 go/src 文件夹，测试文件debug会找不到文件，创建不了breakpoint(通过配置launch.json的test模式加上log发现此问题)

### 普通多main文件
在 go/src 目录下打开，vscode会报错，这是golang语言特情，目前无合适方法
在 符号链接目录 下不报

### 测试文件注意
1. xxx_test.go 为文件名
2. 测试方法同目录下，不同文件也不能重名
3. debug test 和 run test输出内容不一致，debug较详细

### 改造原有联系过程
1. 直接改造 -> xxx_test.go 但package main还是报一堆错
2. 建origin目录，再迁移
3. 同一目录下文件package声明必须一致


### TIPS:
1. 所有的函数默认都是同目录/包可见的，不可重名，变量也是 -> TODO: 如何禁止?
2. 普通类库用Testing练习，只有客户、服务器需要创建比较多的测试目录
3. 关闭buildOnSave