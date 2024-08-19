package results

import "time"

// SNIResult holds the latency results for a given SNI.
type SNIResult struct {
    HttpLatency  time.Duration
    HttpsLatency time.Duration
}

// ScanResult holds the result of scanning a single IP.
type ScanResult struct {
    IP      string
    Results map[string]SNIResult
}

// AverageLatency calculates the average latency for HTTP and HTTPS.
func (sr *ScanResult) AverageLatency() (httpAvg, httpsAvg time.Duration) {
    var httpTotal, httpsTotal time.Duration
    var httpCount, httpsCount int

    for _, result := range sr.Results {
        if result.HttpLatency > 0 {
            httpTotal += result.HttpLatency
            httpCount++
        }
        if result.HttpsLatency > 0 {
            httpsTotal += result.HttpsLatency
            httpsCount++
        }
    }

    if httpCount > 0 {
        httpAvg = httpTotal / time.Duration(httpCount)
    }
    if httpsCount > 0 {
        httpsAvg = httpsTotal / time.Duration(httpsCount)
    }

    return httpAvg, httpsAvg
}
