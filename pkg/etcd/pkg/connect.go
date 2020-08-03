package pkg

import (
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
)

// ConnectETCD | Connect to ETCD
func ConnectETCD(etcdServer []string) (clientv3.KV, *clientv3.Client) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   etcdServer,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}

	kv := clientv3.NewKV(cli)

	return kv, cli
}
