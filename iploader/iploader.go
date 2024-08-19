package iploader

import (
	"bufio"
	"encoding/binary"
	"math/big"
	"crypto/rand"
	"net"
	"os"
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
        randomIP, _ := randomIPFromCIDR(ipnet)
        ips = append(ips, randomIP)
    }
    return ips, scanner.Err()
}

// randomIPFromCIDR generates a random IP address within the given CIDR range.
func randomIPFromCIDR(ipnet *net.IPNet) (net.IP, error) {
    ip := ipnet.IP.To4()

    ipInt := binary.BigEndian.Uint32(ip)

    // Calculate the number of addresses in the subnet
    mask := binary.BigEndian.Uint32(ipnet.Mask)
    networkSize := ^mask

    // Generate a random offset within the range
    offset, err := rand.Int(rand.Reader, big.NewInt(int64(networkSize)))
    if err != nil {
        return nil, err
    }

    // Add the offset to the base IP address
    randomIPInt := ipInt + uint32(offset.Int64())

    // Convert the result back to an IP address
    randomIP := make(net.IP, 4)
    binary.BigEndian.PutUint32(randomIP, randomIPInt)

    return randomIP, nil
}

