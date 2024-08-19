package sniloader

import (
    "bufio"
    "os"
)

// LoadSNIs loads SNI values from a file.
func LoadSNIs(filename string) ([]string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var snis []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        snis = append(snis, scanner.Text())
    }
    return snis, scanner.Err()
}
