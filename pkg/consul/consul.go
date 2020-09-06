package consul

import (
	"os"

	"github.com/hashicorp/consul/api"
)

var (
	ConsulAddress = os.Getenv("CONSUL_CONNECT_INFO")
	ConsulToken   = os.Getenv("CONSUL_TOKEN")
)

func ConnectConsulAPI(address, token string) *api.Client {
	var Config = &api.Config{
		Address: address,
		Token:   token,
	}

	client, err := api.NewClient(Config)
	if err != nil {
		panic(err)
	}

	return client
}
