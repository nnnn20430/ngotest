package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"net/http"
	"time"
)

func main() {
	priv, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}

	certTemplate := x509.Certificate{
		SerialNumber: big.NewInt(time.Now().Unix()),

		Subject: pkix.Name{
			Organization: []string{"Acme Co"},
			CommonName:   "localhost",
		},

		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(365 * 24 * time.Hour),

		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},

		BasicConstraintsValid: false,
		IsCA: true,

		IPAddresses: []net.IP{net.ParseIP("127.0.0.1")},
		DNSNames:    []string{"localhost"},
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &certTemplate, &certTemplate, priv.Public(), priv)
	if err != nil {
		panic(err)
	}

	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}

	srv := http.Server{
		Addr:    ":8443",
		Handler: http.FileServer(http.Dir(".")),
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{tlsCert},
			NextProtos:   []string{"h2"},
		},
	}

	fmt.Printf("https://localhost:8443/\n")
	panic(srv.ListenAndServeTLS("", ""))
}
