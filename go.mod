module github.com/o-big/oconfig

go 1.16

require (
	github.com/sirupsen/logrus v1.8.1
	github.com/urfave/cli v1.22.5
	go.etcd.io/etcd/server/v3 v3.5.0
	go.uber.org/zap v1.18.1 // indirect
	golang.org/x/time v0.0.0-20210611083556-38a9dc6acbc6 // indirect
	sigs.k8s.io/yaml v1.2.0
)

replace github.com/coreos/bbolt v1.3.6 => go.etcd.io/bbolt v1.3.6

replace go.etcd.io/bbolt v1.3.6 => github.com/coreos/bbolt v1.3.6

replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
