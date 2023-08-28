## 单元测试

go语言单元测试不像java 有单独的测试包，然后需要引入三方依赖，本身goland 标准库就支持

## 规则
1. 一般测试类和要被测试的类在同一包下面
2. 测试类命名规定为 xxx_test.go
3. 测试方法名必须 Testxxx(t *testing.T)

## 例子

例子参考 [fooer_test.go](fooer_test.go)