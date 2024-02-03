package main

import (
	"flag"
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
	ip := flag.String("ip", "0.0.0.0", "a string")
	portRange := flag.String("port-range", "{1111-9999}", "a string")

	flag.Parse()

	proxies := generateProxyList(fmt.Sprintf("%s:%s", *ip, *portRange))

	fmt.Println(getRandomProxy(proxies))
}
