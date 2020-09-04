package tlsloader

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"

	"github.com/pkg/errors"
)

func Load(certFile, keyFile string) (cert *tls.Certificate, err error) {
	var b []byte
	if b, err = ioutil.ReadFile(certFile); err != nil {
		return
	}
	block, _ := pem.Decode(b)
	if b == nil {
		return nil, errors.Wrap(err, "cannot parse cert PEM")
	}
	var crt *x509.Certificate
	if crt, err = x509.ParseCertificate(block.Bytes); err != nil {
		return nil, errors.Wrap(err, "parse certificate failed")
	}

	if b, err = ioutil.ReadFile(keyFile); err != nil {
		return
	}

	block, _ = pem.Decode(b)
	if b == nil {
		return nil, errors.New("cannot parse key PEM")
	}
	k, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.Wrap(err, "cannot parse key")
	}

	cert = &tls.Certificate{
		Certificate: [][]byte{crt.Raw},
		PrivateKey:  k,
		Leaf:        crt,
	}
	return
}
