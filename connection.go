//
// Copyright (c) 2017 Joey <majunjiev@gmail.com>.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Some codes of this file is from https://github.com/CpuID/ovirt-engine-sdk-go/blob/master/sdk/http/http.go.
// And I made some bug fixes, Thanks to @CpuID

package ovirtsdk4

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
	"time"
)

// Connection represents an HTTP connection to the engine server.
// It is intended as the entry point for the SDK, and it provides access to the `system` service and, from there,
// to the rest of the services provided by the API.
type Connection struct {
	url      *url.URL
	username string
	password string
	token    string
	insecure bool
	caFile   string
	kerberos bool
	timeout  time.Duration
	compress bool
	//
	client *http.Client
}

// URL returns the base URL of this connection.
func (c *Connection) URL() string {
	return c.url.String()
}

// SystemService returns a reference to the root of the services tree.
func (c *Connection) SystemService() *systemService {
	return NewSystemService(c, "")
}

// NewConnectionBuilder creates the `ConnectionBuilder struct instance
func NewConnectionBuilder() *ConnectionBuilder {
	return &ConnectionBuilder{conn: &Connection{}, err: nil}
}

// ConnectionBuilder represents a builder for the `Connection` struct
type ConnectionBuilder struct {
	conn *Connection
	err  error
}

// URL sets the url field for `Connection` instance
func (connBuilder *ConnectionBuilder) URL(inputRawURL string) *ConnectionBuilder {
	// If already has errors, just return
	if connBuilder.err != nil {
		return connBuilder
	}

	// Save the URL:
	useURL, err := url.Parse(inputRawURL)
	if err != nil {
		connBuilder.err = err
		return connBuilder
	}
	connBuilder.conn.url = useURL
	return connBuilder
}

// Username sets the username field for `Connection` instance
func (connBuilder *ConnectionBuilder) Username(username string) *ConnectionBuilder {
	// If already has errors, just return
	if connBuilder.err != nil {
		return connBuilder
	}

	connBuilder.conn.username = username
	return connBuilder
}

// Password sets the password field for `Connection` instance
func (connBuilder *ConnectionBuilder) Password(password string) *ConnectionBuilder {
	// If already has errors, just return
	if connBuilder.err != nil {
		return connBuilder
	}

	connBuilder.conn.password = password
	return connBuilder
}

// Insecure sets the insecure field for `Connection` instance
func (connBuilder *ConnectionBuilder) Insecure(insecure bool) *ConnectionBuilder {
	// If already has errors, just return
	if connBuilder.err != nil {
		return connBuilder
	}
	connBuilder.conn.insecure = insecure
	return connBuilder
}

// Timeout sets the timeout field for `Connection` instance
func (connBuilder *ConnectionBuilder) Timeout(timeout time.Duration) *ConnectionBuilder {
	// If already has errors, just return
	if connBuilder.err != nil {
		return connBuilder
	}
	connBuilder.conn.timeout = timeout
	return connBuilder
}

// CAFile sets the caFile field for `Connection` instance
func (connBuilder *ConnectionBuilder) CAFile(caFilePath string) *ConnectionBuilder {
	// If already has errors, just return
	if connBuilder.err != nil {
		return connBuilder
	}
	connBuilder.conn.caFile = caFilePath
	return connBuilder
}

// Kerberos sets the kerberos field for `Connection` instance
func (connBuilder *ConnectionBuilder) Kerberos(kerbros bool) *ConnectionBuilder {
	// If already has errors, just return
	if connBuilder.err != nil {
		return connBuilder
	}
	// TODO: kerbros==true is not implemented
	if kerbros == true {
		connBuilder.err = errors.New("Kerberos is not currently implemented")
		return connBuilder
	}
	connBuilder.conn.kerberos = kerbros
	return connBuilder
}

// Compress sets the compress field for `Connection` instance
func (connBuilder *ConnectionBuilder) Compress(compress bool) *ConnectionBuilder {
	// If already has errors, just return
	if connBuilder.err != nil {
		return connBuilder
	}
	connBuilder.conn.compress = compress
	return connBuilder
}

// Build contructs the `Connection` instance
func (connBuilder *ConnectionBuilder) Build() (*Connection, error) {
	// If already has errors, just return
	if connBuilder.err != nil {
		return nil, connBuilder.err
	}

	// Check parameters
	if connBuilder.conn.url == nil {
		return nil, errors.New("The URL must not be empty")
	}
	if len(connBuilder.conn.username) == 0 {
		return nil, errors.New("The Username must not be empty")
	}
	if len(connBuilder.conn.password) == 0 {
		return nil, errors.New("The Password must not be empty")
	}

	// Construct http.Client
	var tlsConfig *tls.Config
	if connBuilder.conn.url.Scheme == "https" {
		tlsConfig = &tls.Config{
			InsecureSkipVerify: connBuilder.conn.insecure,
		}
		if len(connBuilder.conn.caFile) > 0 {
			// Check if the CA File specified exists.
			if _, err := os.Stat(connBuilder.conn.caFile); os.IsNotExist(err) {
				return nil, fmt.Errorf("The ca file '%s' doesn't exist", connBuilder.conn.caFile)
			}
			pool := x509.NewCertPool()
			caCerts, err := ioutil.ReadFile(connBuilder.conn.caFile)
			if err != nil {
				return nil, err
			}
			if ok := pool.AppendCertsFromPEM(caCerts); ok == false {
				return nil, fmt.Errorf("Failed to parse CA Certificate in file '%s'", connBuilder.conn.caFile)
			}
			tlsConfig.RootCAs = pool
		}
	}
	connBuilder.conn.client = &http.Client{
		Timeout: connBuilder.conn.timeout,
		Transport: &http.Transport{
			// Close the http connection after calling resp.Body.Close()
			DisableKeepAlives:  true,
			DisableCompression: !connBuilder.conn.compress,
			TLSClientConfig:    tlsConfig,
		},
	}
	return connBuilder.conn, nil
}

// Test tests the connectivity with the server. If connectivity works correctly it returns a nil error. If there is any
// connectivity problem it will return an error containing the reason as the message.
func (c *Connection) Test() error {
	return nil
}

func (c *Connection) getHref(object Href) (string, bool) {
	return object.Href()
}

// IsLink indicates if the given object is a link.
// An object is a link if it has an `href` attribute.
func (c *Connection) IsLink(object Href) bool {
	_, ok := c.getHref(object)
	return ok
}

// FollowLink follows the `href` attribute of the given object, retrieves the target object and returns it.
func (c *Connection) FollowLink(object Href) (interface{}, error) {
	if !c.IsLink(object) {
		return nil, errors.New("Can't follow link because object don't have any")
	}
	href, ok := c.getHref(object)
	if !ok {
		return nil, errors.New("Can't follow link because the 'href' attribute does't have a value")
	}
	useURL, err := url.Parse(c.URL())
	if err != nil {
		return nil, errors.New("Failed to parse connection url")
	}
	prefix := useURL.Path
	if !strings.HasSuffix(prefix, "/") {
		prefix = prefix + "/"
	}
	if !strings.HasPrefix(href, prefix) {
		return nil, fmt.Errorf("The URL '%v' isn't compatible with the base URL of the connection", href)
	}
	path := href[len(prefix):]
	service, err := NewSystemService(c, "").Service(path)
	if err != nil {
		return nil, err
	}

	serviceValue := reflect.ValueOf(service)
	// `object` is ptr, so use Elem() to get struct value
	hrefObjectValue := reflect.ValueOf(object).Elem()
	var requestCaller reflect.Value
	// If it's TypeStructSlice (list)
	if hrefObjectValue.FieldByName("slice").IsValid() {
		// Call List() method
		requestCaller = serviceValue.MethodByName("List").Call([]reflect.Value{})[0]
	} else {
		requestCaller = serviceValue.MethodByName("Get").Call([]reflect.Value{})[0]
	}
	callerResponse := requestCaller.MethodByName("Send").Call([]reflect.Value{})[0]
	// Method 0 could retrieve the data
	returnedValues := callerResponse.Method(0).Call([]reflect.Value{})
	result, ok := returnedValues[0].Interface(), returnedValues[1].Bool()
	if !ok {
		return nil, errors.New("The data retrieved not exists")
	}
	return result, nil
}

// Close releases the resources used by this connection.
func (c *Connection) Close() {
}
