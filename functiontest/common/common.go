package common

import (
	"crypto/tls"
	"crypto/x509"
	//"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/auth/token"
	"github.com/gophercloud/gophercloud/openstack"
)

func AuthAKSK() (*gophercloud.ProviderClient, error) {
	akskOptions := aksk.AKSKOptions{

		IdentityEndpoint: "https://iam.cn-north-1.myhuaweicloud.com/v3",
		DomainID:         "",
		ProjectID:        "",
		AccessKey:        "",
		SecretKey:        "",
		Domain:           "myhuaweicloud.com",
		Region:           "cn-north-1",
	}

	provider, err := openstack.AuthenticatedClient(akskOptions)
	if err != nil {
		panic(err)
	}

	return provider, nil
}

func AuthToken() (*gophercloud.ProviderClient, error) {
	tokenOpts := token.TokenOptions{
		IdentityEndpoint: "",
		Username:         "",
		Password:         "",
		DomainID:         "",
		TenantID:         "",
		AllowReauth:      true,

	}
	provider, err := openstack.AuthenticatedClient(tokenOpts)

	if err != nil {
		fmt.Println("Failed to authenticate:", err)
		return nil, err
	}
	return provider, nil
}

func OpenstackHTTPClient(cacert string, insecure bool) (http.Client, error) {
	if cacert == "" {
		return http.Client{}, nil
	}
	caCertPool := x509.NewCertPool()
	caCert, err := ioutil.ReadFile(cacert)
	if err != nil {
		return http.Client{}, errors.New("Can't read certificate file")
	}
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		RootCAs:            caCertPool,
		InsecureSkipVerify: insecure,
	}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}

	return http.Client{Transport: transport}, nil
}
