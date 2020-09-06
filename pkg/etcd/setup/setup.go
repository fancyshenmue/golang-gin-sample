package setup

import (
	"os"
	"strings"
)

var (
	// ETCDConnectInfo | ETCD Connect Info
	ETCDConnectInfo     = strings.Split(os.Getenv("ETCD_CONNECT_INFO"), ",")
	ETCDConnectUser     = os.Getenv("ETCD_CONNECT_USER")
	ETCDConnectPassword = os.Getenv("ETCD_CONNECT_PASSWD")
	// ETCDConnectInfoMap | ETCD Connection Info
	ETCDConnectInfoMap = map[string][]string{
		"SIT": {
			"http://etcd01.pf.sit:2379",
			"http://etcd02.pf.sit:2379",
			"http://etcd03.pf.sit:2379",
		},
		"STG": {
			"http://etcd01.pf.stg:2379",
			"http://etcd02.pf.stg:2379",
			"http://etcd03.pf.stg:2379",
		},
		"UAT": {
			"http://etcd01.pf.uat:2379",
			"http://etcd02.pf.uat:2379",
			"http://etcd03.pf.uat:2379",
		},
		"PROD": {
			"http://etcd01.pf:2379",
			"http://etcd02.pf:2379",
			"http://etcd03.pf:2379",
		},
	}
)
