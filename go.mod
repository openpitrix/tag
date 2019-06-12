module openpitrix.io/tag

require (
	github.com/bitly/go-simplejson v0.5.0
	github.com/fatih/camelcase v1.0.0
	github.com/fatih/structs v1.1.0
	github.com/golang/protobuf v1.3.1
	github.com/google/gops v0.0.0-20180903072510-f341a40f99ec
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0
	github.com/grpc-ecosystem/grpc-gateway v1.9.0
	github.com/jinzhu/gorm v1.9.8
	github.com/koding/multiconfig v0.0.0-20171124222453-69c27309b2d7
	github.com/pkg/errors v0.8.1
	github.com/sony/sonyflake v0.0.0-20181109022403-6d5bd6181009
	github.com/speps/go-hashids v2.0.0+incompatible
	github.com/stretchr/testify v1.2.2
	golang.org/x/net v0.0.0-20181220203305-927f97764cc3
	google.golang.org/genproto v0.0.0-20190128161407-8ac453e89fca
	google.golang.org/grpc v1.19.0
	openpitrix.io/logger v0.1.0
	openpitrix.io/notification v0.2.0 // indirect
)

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.26.0
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190308221718-c2843e01d9a2
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190227174305-5b3e6a55c961
	golang.org/x/net => github.com/golang/net v0.0.0-20190503192946-f4e77d36d62c
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20180821212333-d2e6202438be
	golang.org/x/sync => github.com/golang/sync v0.0.0-20180314180146-1d60e4601c6f
	golang.org/x/sys => github.com/golang/sys v0.0.0-20180830151530-49385e6e1522
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/tools => github.com/golang/tools v0.0.0-20180828015842-6cd1fcedba52
	google.golang.org/appengine => github.com/golang/appengine v1.1.0
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20180817151627-c66870c02cf8
	google.golang.org/grpc => github.com/grpc/grpc-go v1.16.0
)
