package goflexer

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

// Database creates and returns a MongoDB database client
func (c *Context) Database(name string) (*mgo.Database, error) {
	key := fmt.Sprintf("_nflexdb_%s", name)
	connstring, ok := c.Secrets[key]
	if !ok {
		return nil, fmt.Errorf("database secret not found")
	}

	// work around https://github.com/go-mgo/mgo/issues/84
	connstring = strings.Replace(connstring, "&ssl=true", "", 1)

	dialinfo, err := mgo.ParseURL(connstring)
	if err != nil {
		return nil, err
	}

	cert, err := ioutil.ReadFile("/nflex-root.pem")
	if err != nil {
		return nil, fmt.Errorf("failed to read SSL certificate")
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(cert)

	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}

	dialinfo.Timeout = 10 * time.Second
	dialinfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		if err != nil {
			log.Errorf("dial %s => %+v\n", addr.String(), err)
		}
		return conn, err
	}

	s, err := mgo.DialWithInfo(dialinfo)
	if err != nil {
		return nil, err
	}

	return s.DB(""), nil
}
