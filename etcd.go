package main

import (
	"github.com/sirupsen/logrus"
	"go.etcd.io/etcd/server/v3/embed"
	"io/ioutil"
	"os"
	"path/filepath"
	"sigs.k8s.io/yaml"
)

type ETCDConfig struct {
	InitialOptions      `json:",inline"`
	Name                string      `json:"name,omitempty"`
	ListenClientURLs    string      `json:"listen-client-urls,omitempty"`
	ListenMetricsURLs   string      `json:"listen-metrics-urls,omitempty"`
	ListenPeerURLs      string      `json:"listen-peer-urls,omitempty"`
	AdvertiseClientURLs string      `json:"advertise-client-urls,omitempty"`
	DataDir             string      `json:"data-dir,omitempty"`
	SnapshotCount       int         `json:"snapshot-count,omitempty"`
	ServerTrust         ServerTrust `json:"client-transport-security"`
	PeerTrust           PeerTrust   `json:"peer-transport-security"`
	ForceNewCluster     bool        `json:"force-new-cluster,omitempty"`
	HeartbeatInterval   int         `json:"heartbeat-interval"`
	ElectionTimeout     int         `json:"election-timeout"`
	Logger              string      `json:"logger"`
	LogOutputs          []string    `json:"log-outputs"`
}

type ServerTrust struct {
	CertFile       string `json:"cert-file"`
	KeyFile        string `json:"key-file"`
	ClientCertAuth bool   `json:"client-cert-auth"`
	TrustedCAFile  string `json:"trusted-ca-file"`
}

type PeerTrust struct {
	CertFile       string `json:"cert-file"`
	KeyFile        string `json:"key-file"`
	ClientCertAuth bool   `json:"client-cert-auth"`
	TrustedCAFile  string `json:"trusted-ca-file"`
}

type InitialOptions struct {
	AdvertisePeerURL string `json:"initial-advertise-peer-urls,omitempty"`
	Cluster          string `json:"initial-cluster,omitempty"`
	State            string `json:"initial-cluster-state,omitempty"`
}

func ETCD(args ETCDConfig) error {
	cfg := embed.NewConfig()
	cfg.Dir = "./runtime"
	etcd, err := embed.StartEtcd(cfg)
	if err != nil {
		logrus.Errorf("start etcd err: %v", err)
		return nil
	}

	//go func() {
		select {
		case <-etcd.Server.StopNotify():
			logrus.Fatalf("etcd stopped - if this node was removed from the cluster, you must backup and delete %s before rejoining", args.DataDir)
		case err := <-etcd.Err():
			logrus.Fatalf("etcd exited: %v", err)
		}
	//}()
	return nil
}

func (e ETCDConfig) ToConfigFile() (string, error) {
	confFile := filepath.Join(e.DataDir, "config")
	bytes, err := yaml.Marshal(&e)
	if err != nil {
		return "", err
	}

	if err := os.MkdirAll(e.DataDir, 0700); err != nil {
		return "", err
	}
	return confFile, ioutil.WriteFile(confFile, bytes, 0600)
}
