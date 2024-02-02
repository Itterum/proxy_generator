package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
)

func getRandomProxy(proxyList []string) string {
	var formatProxies []string

	for proxy := 0; proxy < len(proxyList); proxy++ {
		formatProxies = append(formatProxies, fmt.Sprintf("http://%s", proxyList[proxy]))
	}

	randProxy := rand.Intn(len(formatProxies))

	return formatProxies[randProxy]
}

func generateProxyList(rangeProxy string) []string {
	var proxyList []string

	rePORT := regexp.MustCompile(`(\d+)-(\d+)`)
	reIP := regexp.MustCompile(`^\d+.+:`)

	port := rePORT.FindStringSubmatch(rangeProxy)
	ip := reIP.FindStringSubmatch(rangeProxy)[0]

	minPort, _ := strconv.Atoi(port[1])
	maxPort, _ := strconv.Atoi(port[2])

	for port := minPort; port < maxPort; port++ {
		proxyList = append(proxyList, fmt.Sprint(ip, port))
	}
	
	return proxyList
}

func main() {
	proxies := generateProxyList("0.0.0.0:{1000-2000}")

	fmt.Println(getRandomProxy(proxies))
}
