package loadbalancer

import (
	"Pay-AI/financial-transaction-server/constantservice"
	"net/url"
	"sync"
)

type WeightedLoadBalancer struct {
	Server []ServerUrl
	Mutex  sync.Mutex
}

type ServerUrl struct {
	Url             *url.URL
	Weight          int
	CurrentWeight   int
	EffectiveWeight int
}

// In each weighted server, each server is assigned a specific weight
// It represents power with respect to other server
// Higher weight should revieve higher traffic proportionately

func (lb *WeightedLoadBalancer) NewWeightedLoadBalancer() *ServerUrl {

	totalWeight := 0
	var best *ServerUrl
	for _, server := range lb.Server {
		totalWeight += server.EffectiveWeight
		server.CurrentWeight += server.EffectiveWeight
		if best == nil || server.EffectiveWeight > best.EffectiveWeight {
			best = &ServerUrl{
				Url:             server.Url,
				Weight:          server.Weight,
				EffectiveWeight: server.EffectiveWeight,
				CurrentWeight:   server.CurrentWeight,
			}
		}
	}
	best.CurrentWeight -= totalWeight
	return &ServerUrl{
		Url:             best.Url,
		EffectiveWeight: best.EffectiveWeight,
	}

}

func NextWeightedServer() *WeightedLoadBalancer {
	weightedServer := constantservice.WeightedRRServers

	var servers []ServerUrl

	for _, itr := range weightedServer {
		servers = append(servers,
			ServerUrl{
				Url:             itr.Url,
				EffectiveWeight: itr.EffectiveWeight,
				Weight:          itr.Weight,
				CurrentWeight:   itr.CurrentWeight,
			})
	}
	return &WeightedLoadBalancer{
		Server: servers,
	}

}
