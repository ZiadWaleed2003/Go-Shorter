package service

import (
	"net"
)


var encoded_urls map[string]string = make(map[string]string)

func ShortenURL(url string) string {


	port := ":8080"
	encoded_url := Encoder(len(encoded_urls))
	shortened_url := "http://" + getLocalIP() + ":" + port + "/" + encoded_url
	
	encoded_urls[encoded_url] = url

	return shortened_url
}

func Encoder(num int) string {
	// Base62 encoding
	const base62 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var encoded string
	num += 1

	for num > 0{
	    remainder := num % 62
	    encoded = string(base62[remainder]) + encoded
	    num /= 62
	}

	return encoded
}

func getLocalIP() string {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        return "localhost"
    }

    for _, addr := range addrs {
        var ip net.IP
        switch v := addr.(type) {
        case *net.IPNet:
            ip = v.IP
        case *net.IPAddr:
            ip = v.IP
        }
        if ip == nil || ip.IsLoopback() {
            continue
        }
        ip = ip.To4()
        if ip == nil {
            continue // not an ipv4 address
        }
 
        return ip.String()
    }
    return "localhost"
}
