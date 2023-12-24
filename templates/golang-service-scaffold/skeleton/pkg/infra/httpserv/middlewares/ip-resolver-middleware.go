package middlewares

import (
	"net"
	"strings"

	"github.com/eser/go-service/pkg/infra/httpserv"
)

func GetClientIps(req *httpserv.Request) []string {
	// first check the X-Forwarded-For header
	requester := req.Header.Get("True-Client-IP")

	if len(requester) == 0 {
		requester = req.Header.Get("X-Forwarded-For")
	}

	// if empty, check the Real-IP header
	if len(requester) == 0 {
		requester = req.Header.Get("X-Real-IP")
	}

	// if the requester is still empty, use the hard-coded address from the socket
	if len(requester) == 0 {
		requester = req.RemoteAddr
	}

	// split comma delimited list into a slice
	// (this happens when proxied via elastic load balancer then again through nginx)
	return strings.Split(requester, ",")
}

func DetectLocalNetwork(requestIp string) bool {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	requestIpNet := net.ParseIP(requestIp)

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}

		if !ipNet.Contains(requestIpNet) {
			continue
		}

		if requestIpNet.IsLoopback() {
			return true
		}
	}

	return false
}

func IpResolverMiddleware() httpserv.HandlerFunc {
	return func(ctx *httpserv.Context) {
		// config, _ := config.GetConfig()
		ips := GetClientIps(ctx.Request)
		ctx.Set("ip", ips)

		if DetectLocalNetwork(ips[0]) {
			ctx.Set("ip-origin", "local")
			ctx.Header("X-Request-Origin", "local: "+strings.Join(ips, ", "))

			ctx.Next()

			return
		}

		// TODO(@eser) add ip allowlist and blocklist implementations

		ctx.Set("ip-origin", "remote")
		ctx.Header("X-Request-Origin", strings.Join(ips, ", "))
		ctx.Next()
	}
}
