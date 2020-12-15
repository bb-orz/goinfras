module goinfras

go 1.14

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.47.0
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.4
	github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
	go.uber.org/atomic => github.com/uber-go/atomic v1.5.0
	go.uber.org/multierr => github.com/uber-go/multierr v1.4.0
	go.uber.org/tools => github.com/uber-go/tools v0.0.0-20190618225709-2cfd321de3ee
	go.uber.org/zap => github.com/uber-go/zap v1.12.0
	// 使用vendor时，本地开发包目录替换
	// github.com/bb-orz/goinfras => /Users/fun/Code/MyProject/goinfras
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190325154230-a5d413f7728c
	golang.org/x/exp => github.com/golang/exp v0.0.0-20191030013958-a1ab85dbe136
	golang.org/x/image => github.com/golang/image v0.0.0-20191009234506-e7c1f5e7dbb8
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190930215403-16217165b5de
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20191031020345-0945064e013a
	golang.org/x/mod => github.com/golang/mod v0.1.0
	golang.org/x/net => github.com/golang/net v0.0.0-20190327025741-74e053c68e29
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190604053449-0f29369cfe45
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190227155943-e225da77a7e6
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190322080309-f49334f85ddc
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/time => github.com/golang/time v0.0.0-20191024005414-555d28b269f0
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190330180304-aef51cc3777c
	golang.org/x/xerrors => github.com/golang/xerrors v0.0.0-20191011141410-1b5146add898
	google.golang.org/api => github.com/googleapis/google-api-go-client v0.13.0
	google.golang.org/appengine => github.com/golang/appengine v1.6.5
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20191028173616-919d9bdd9fe6
	google.golang.org/grpc => github.com/grpc/grpc-go v1.24.0

)

require (
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.296
	github.com/aliyun/aliyun-oss-go-sdk v2.0.8+incompatible
	github.com/baiyubin/aliyun-sts-go-sdk v0.0.0-20180326062324-cfa1a18b161f // indirect
	github.com/bmizerany/pq v0.0.0-20131128184720-da2b95e392c1
	github.com/coreos/bbolt v1.3.5 // indirect
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/go-systemd/v22 v22.1.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/didi/gendry v1.3.2
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/elastic/go-elasticsearch/v8 v8.0.0-20201202142044-1e78b5bf06b1
	github.com/fastly/go-utils v0.0.0-20180712184237-d95a45783239 // indirect
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/locales v0.13.0
	github.com/go-playground/universal-translator v0.17.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.4.3
	github.com/gomodule/redigo v1.8.3
	github.com/google/uuid v1.1.2 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20200217142428-fce0ec30dd00 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/imroc/req v0.3.0
	github.com/jehiah/go-strftime v0.0.0-20171201141054-1d33003b3869 // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/json-iterator/go v1.1.10
	github.com/kataras/iris v0.0.2 // indirect
	github.com/labstack/echo v3.3.10+incompatible // indirect
	github.com/labstack/echo/v4 v4.1.17
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.3.0+incompatible
	github.com/lestrrat-go/strftime v1.0.1 // indirect
	github.com/nats-io/nats-server/v2 v2.1.7 // indirect
	github.com/nats-io/nats.go v1.10.0
	github.com/olivere/elastic/v7 v7.0.22
	github.com/prometheus/client_golang v1.8.0 // indirect
	github.com/prometheus/common v0.15.0
	github.com/qiniu/api.v7/v7 v7.4.1
	github.com/robfig/cron/v3 v3.0.1
	github.com/sirupsen/logrus v1.7.0 // indirect
	github.com/smartystreets/goconvey v1.6.4
	github.com/spf13/cobra v1.1.1 // indirect
	github.com/spf13/viper v1.7.1
	github.com/tebeka/strftime v0.1.4 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20200427203606-3cfed13b9966 // indirect
	go.etcd.io/bbolt v1.3.5 // indirect
	go.etcd.io/etcd v3.3.25+incompatible
	go.mongodb.org/mongo-driver v1.3.2
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0
	golang.org/x/crypto v0.0.0-20201208171446-5f87f3452ae9 // indirect
	golang.org/x/lint v0.0.0-20201208152925-83fdc39ff7b5 // indirect
	golang.org/x/net v0.0.0-20201209123823-ac852fbbde11
	golang.org/x/sys v0.0.0-20201211090839-8ad439b19e0f // indirect
	golang.org/x/text v0.3.4 // indirect
	golang.org/x/time v0.0.0-20201208040808-7e3f01d25324 // indirect
	golang.org/x/tools v0.0.0-20201211025543-abf6a1d87e11 // indirect
	google.golang.org/genproto v0.0.0-20201211151036-40ec1c210f7a // indirect
	google.golang.org/grpc v1.34.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/mysql v1.0.3
	gorm.io/driver/postgres v1.0.5
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.20.8
	honnef.co/go/tools v0.0.1-2020.1.6 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)
