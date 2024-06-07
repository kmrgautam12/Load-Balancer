package constantservice

import (
	utils "Pay-AI/financial-transaction-server/financial-transaction-server/Utils"
	"net/url"
)

const (
	PEM_PUBLIC_PATH  = "pem_public"
	PEM_PRIVATE_PATH = "pem_private"
)

type server struct {
	Url             *url.URL
	Weight          int
	CurrentWeight   int
	EffectiveWeight int
}

var RRServers = []server{
	{
		Url: utils.MustParseUrl("http://localhost:8081"),
	},
	{
		Url: utils.MustParseUrl("http://localhost:8082"),
	},
}

var WeightedRRServers = []server{
	{
		Url:    utils.MustParseUrl("http://localhost:8081"),
		Weight: 5,
	},
	{
		Url:    utils.MustParseUrl("http://localhost:8082"),
		Weight: 1,
	},
}
