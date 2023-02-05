# runedance_douyin

第五届青训营-抖音项目

## 项目结构

### 目录结构

暂时如下：
```go,
.
|__idl
| |__relation.thrift
|__cmd
| |__api
| | |__main.go
|__go.mod
|__pkg
| |__errnos
| | |__errnos.go
| |__consts
| | |__constants.go
|__sql
| |__create
| | |__relation.sql
| | |__database.sql

```

### 说明

cmd里放不同的微服务，kitex生成的代码都集中放到kitex_gen目录下。

除了cmd，其他目录下一些package都是公用的。尽量复用，而不是重新生成或造轮子。如需修改可以在群里说声。

## 编写微服务的几个步骤

> 从编写idl到启动微服务

*kitex生成代码命令*
```bash
kitex -module example -service example echo.thrift
```

1. 编写idl（参考relation.idl)
2. 假设服务是xx，则新建cmd/xx目录，并在目录下调用kitex代码生成工具。（注意版本需要为v0.4.4.）
3. 把生成的kitex_gen目录里的内容**移动**到最外层的kitex_gen目录中，其他内容保留在当前目录中即可。
4. cmd/xx目录中的main.go就是我们的服务端启动的入口。(如果是goland编译器，可以直接点三角号)
5. 在handler.go中实现业务逻辑即可。
