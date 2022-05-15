# Playground for golang

[Go 编程语言规范](https://go.dev/ref/spec)

## Order of evaluation 执行顺序
在包级别，初始化依赖关系确定变量声明中各个初始化表达式的评估顺序。 否则，在计算表达式、赋值或返回语句的操作数时，所有函数调用、方法调用和通信操作都按词法从左到右的顺序计算。

来源: [知乎](https://www.zhihu.com/question/65916025/answer/2484976011)