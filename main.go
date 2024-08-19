package main

import (
    "fmt"
    "github.com/copartner6412/ipscanner/iploader"
    "github.com/copartner6412/ipscanner/scanner"
    "github.com/copartner6412/ipscanner/sniloader"
)

func main() {
    ips, err := iploader.LoadIPs("cloudflare_ips.txt")
    if err != nil {
        panic(err)
    }

    snis, err := sniloader.LoadSNIs("sni_list.txt")
    if err != nil {
        panic(err)
    }

    s := scanner.Scanner{
        SNIs: snis,
        Port: "443", // or "80" for HTTP
    }

    for _, ip := range ips {
        result := s.ScanIP(ip)
        httpAvg, httpsAvg := result.AverageLatency()
        fmt.Printf("IP: %s, HTTP Avg: %v, HTTPS Avg: %v\n", result.IP, httpAvg, httpsAvg)
    }
}
