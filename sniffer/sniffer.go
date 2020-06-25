package sniffer

import (
	"net"
)

type Sniffer struct {
	EnableTLS  bool
	EnableHTTP bool
}

func NewSniffer(enableTLS, enableHTTP bool) *Sniffer {
	return &Sniffer{enableTLS, enableHTTP}
}

func (s *Sniffer) SnifferTLSSNI(conn net.Conn) (serverName string, buf []byte, err error) {
	if s.EnableTLS {
		buf = make([]byte, 1024)
		n := 0
		n, err = conn.Read(buf)
		buf = buf[:n]
		if err != nil {
			return
		}
		serverName, err = parseServerNameFromSNI(buf)
		if err != nil {
			return
		}
		return
	}
	return
}

func (s *Sniffer) SnifferHTTPHost(data []byte) (string, error) {
	if s.EnableHTTP {
		return "", nil
	}
	return "", nil
}