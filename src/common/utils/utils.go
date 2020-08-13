package utils

import (
	"crypto/rand"
	"fmt"
	"net"
	"time"
)

// TestTCPConn tests TCP connection
// timeout: the total time before returning if something is wrong
// with the connection, in second
// interval: the interval time for retring after failure, in second
func TestTCPConn(addr string, timeout, interval int) error {
	success := make(chan int)
	cancel := make(chan int)

	go func() {
		n := 1

	loop:
		for {
			select {
			case <-cancel:
				break loop
			default:
				conn, err := net.DialTimeout("tcp", addr, time.Duration(n)*time.Second)
				if err != nil {
					fmt.Errorf("failed to connect to tcp://%s, retry after %d seconds :%v",
						addr, interval, err)
					n = n * 2
					time.Sleep(time.Duration(interval) * time.Second)
					continue
				}
				if err = conn.Close(); err != nil {
					fmt.Errorf("failed to close the connection: %v", err)
				}
				success <- 1
				break loop
			}
		}
	}()

	select {
	case <-success:
		return nil
	case <-time.After(time.Duration(timeout) * time.Second):
		cancel <- 1
		return fmt.Errorf("failed to connect to tcp:%s after %d seconds", addr, timeout)
	}
}

// GenerateRandomStringWithLen generates a random string with length
func GenerateRandomStringWithLen(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	l := len(chars)
	result := make([]byte, length)
	_, err := rand.Read(result)
	if err != nil {
		fmt.Errorf("Error reading random bytes: %v", err)
	}
	for i := 0; i < length; i++ {
		result[i] = chars[int(result[i])%l]
	}
	return string(result)
}

// GenerateRandomString generate a random string with 32 byte length
func GenerateRandomString() string {
	return GenerateRandomStringWithLen(32)
}
