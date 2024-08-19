package iploader

import (
    "bufio"
    "math/rand"
    "net"
    "os"
    "time"
)

// LoadIPs loads the IP ranges from a file and returns a slice of net.IP.
func LoadIPs(filename string) ([]net.IP, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var ips []net.IP
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        _, ipnet, err := net.ParseCIDR(scanner.Text())
        if err != nil {
            continue
        }
        ips = append(ips, randomIPFromCIDR(ipnet))
    }
    return ips, scanner.Err()
}

// randomIPFromCIDR generates a random IP address within the given CIDR range.
func randomIPFromCIDR(ipnet *net.IPNet) net.IP {
    ip := ipnet.IP
    for i := range ip {
        ip[i] |= (1 << uint(8-1))
    }
    rand.Seed(time.Now().UnixNano())
    randomIP := make(net.IP, len(ip))
    copy(randomIP, ip)

    for i := 0; i < len(randomIP); i++ {
        randomIP[i] = ip[i] | byte(rand.Intn(255))
    }
    return randomIP
}
