module runedance_douyin

go 1.18

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0

require (
	github.com/apache/rocketmq-client-go/v2 v2.1.1
	github.com/apache/thrift v0.17.0
	github.com/astaxie/beego v1.12.3
	github.com/cloudwego/hertz v0.5.2
	github.com/cloudwego/hertz-examples/bizdemo/hertz_jwt v0.0.0-20230204112917-78c8960383b4
	github.com/cloudwego/kitex v0.4.4
	github.com/golang-jwt/jwt/v4 v4.4.3
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/google/uuid v1.3.0
	github.com/hibiken/asynq v0.24.0
	github.com/importcjj/sensitive v0.0.0-20200106142752-42d1c505be7b
	github.com/juju/ratelimit v1.0.2
	github.com/kitex-contrib/registry-etcd v0.1.0
	github.com/pkg/errors v0.9.1
	github.com/redis/go-redis/v9 v9.0.2
	github.com/sirupsen/logrus v1.8.1
	github.com/tencentyun/cos-go-sdk-v5 v0.7.41
	go.mongodb.org/mongo-driver v1.11.1
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gorm.io/driver/mysql v1.4.4
	gorm.io/gorm v1.24.3
	gorm.io/plugin/opentelemetry v0.1.0
)

require (
	github.com/bytedance/go-tagexpr/v2 v2.9.2 // indirect
	github.com/bytedance/gopkg v0.0.0-20220531084716-665b4f21126f // indirect
	github.com/bytedance/sonic v1.7.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/chenzhuoyu/iasm v0.0.0-20220818063314-28c361dae733 // indirect
	github.com/choleraehyq/pid v0.0.15 // indirect
	github.com/clbanning/mxj v1.8.4 // indirect
	github.com/cloudwego/fastpb v0.0.3 // indirect
	github.com/cloudwego/frugal v0.1.3 // indirect
	github.com/cloudwego/netpoll v0.3.1 // indirect
	github.com/cloudwego/thriftgo v0.2.4 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/emirpasic/gods v1.12.0 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-redis/redis/v8 v8.11.5 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/google/pprof v0.0.0-20220608213341-c488b8fa1db3 // indirect
	github.com/henrylee2cn/ameda v1.4.10 // indirect
	github.com/henrylee2cn/goutil v0.0.0-20210127050712-89660552f6f8 // indirect
	github.com/hertz-contrib/jwt v1.0.0 // indirect
	github.com/jhump/protoreflect v1.8.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/klauspost/cpuid/v2 v2.1.0 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe // indirect
	github.com/mozillazg/go-httpheader v0.2.1 // indirect
	github.com/nyaruka/phonenumbers v1.0.55 // indirect
	github.com/oleiade/lane v1.0.1 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/rogpeppe/go-internal v1.8.0 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/stathat/consistent v1.0.0 // indirect
	github.com/tidwall/gjson v1.13.0 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.1 // indirect
	github.com/xdg-go/stringprep v1.0.3 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	go.etcd.io/etcd/api/v3 v3.5.5 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.5 // indirect
	go.etcd.io/etcd/client/v3 v3.5.5 // indirect
	go.opentelemetry.io/otel v1.8.0 // indirect
	go.opentelemetry.io/otel/metric v0.31.0 // indirect
	go.opentelemetry.io/otel/trace v1.8.0 // indirect
	go.uber.org/atomic v1.8.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.17.0 // indirect
	golang.org/x/arch v0.0.0-20220722155209-00200b7164a7 // indirect
	golang.org/x/crypto v0.1.0 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/text v0.6.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20211118181313-81c1377c94b1 // indirect
	google.golang.org/grpc v1.46.2 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
