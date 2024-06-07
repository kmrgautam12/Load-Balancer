package loadbalancer

import (
	"Pay-AI/financial-transaction-server/financial-transaction-server/constantservice"
	"fmt"
	"net/http/httputil"
	"net/url"
	"sync"

	"github.com/gin-gonic/gin"
)

type RoundRobinLoadBalancer struct {
	Servers []*url.URL
	Index   int
	mutex   sync.Mutex
}

func (lb RoundRobinLoadBalancer) HandleRequest(c *gin.Context) {

	lb.mutex.Lock()
	defer lb.mutex.Lock()
	serverUrl := lb.Servers[lb.Index]

	lb.Index = (lb.Index + 1) % len(lb.Servers)
	proxy := httputil.NewSingleHostReverseProxy(serverUrl)
	proxy.ServeHTTP(c.Writer, c.Request)
}

func GetServer() *RoundRobinLoadBalancer {
	backendServers := constantservice.RRServers
	servers := []*url.URL{}
	for _, u := range backendServers {
		servers = append(servers, u.Url)
	}

	return &RoundRobinLoadBalancer{
		Servers: servers,
		Index:   0,
	}
}

func ServeRequestWithProxy() {
	fmt.Println("Inside serve proxy request")
	GetServer()
}
