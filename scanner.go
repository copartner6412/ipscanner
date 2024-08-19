package scanner

import (
    "net/http"
    "time"
    "crypto/tls"
    "context"
    "fmt"
    "net"
    "ipscanner/internal/results"
)

type Scanner struct {
    SNIs []string
    Port string
}

// ScanIP performs the scan for a given IP across all SNIs.
func (s *Scanner) ScanIP(ip net.IP) results.ScanResult {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    scanResult := results.ScanResult{
        IP: ip.String(),
        Results: make(map[string]results.SNIResult),
    }

    for _, sni := range s.SNIs {
        httpLatency := s.testConnection(ctx, "http", ip, sni)
        httpsLatency := s.testConnection(ctx, "https", ip, sni)

        scanResult.Results[sni] = results.SNIResult{
            HttpLatency:  httpLatency,
            HttpsLatency: httpsLatency,
        }
    }

    return scanResult
}

func (s *Scanner) testConnection(ctx context.Context, protocol string, ip net.IP, sni string) time.Duration {
    var url string
    if protocol == "https" {
        url = fmt.Sprintf("https://%s:%s", ip.String(), s.Port)
    } else {
        url = fmt.Sprintf("http://%s:%s", ip.String(), s.Port)
    }

    start := time.Now()

    req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
    if protocol == "https" {
        req.Host = sni
        transport := &http.Transport{
            TLSClientConfig: &tls.Config{
                ServerName: sni,
            },
        }
        client := &http.Client{Transport: transport}
        _, err := client.Do(req)
        if err != nil {
            return -1
        }
    } else {
        client := &http.Client{}
        _, err := client.Do(req)
        if err != nil {
            return -1
        }
    }

    latency := time.Since(start)
    return latency
}
