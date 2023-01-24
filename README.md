# runedance_douyin
第五届青训营-抖音项目

## 项目结构

### 目录结构

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


### 说明

cmd里放不同的微服务，kitex生成的代码都集中放到kitex_gen目录下。

除了cmd，其他目录下一些package都是公用的。不需要重新生成或造轮子了。

