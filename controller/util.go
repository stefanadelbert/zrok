package controller

import (
	"crypto/sha512"
	"crypto/x509"
	"encoding/hex"
	errors2 "github.com/go-openapi/errors"
	"github.com/jaevor/go-nanoid"
	"github.com/openziti-test-kitchen/zrok/rest_model_zrok"
	"github.com/openziti/edge/rest_management_api_client"
	"github.com/openziti/edge/rest_util"
	"net/http"
	"strings"
)

type zrokAuthenticator struct {
	cfg *Config
}

func newZrokAuthenticator(cfg *Config) *zrokAuthenticator {
	return &zrokAuthenticator{cfg}
}

func (za *zrokAuthenticator) authenticate(token string) (*rest_model_zrok.Principal, error) {
	tx, err := str.Begin()
	if err != nil {
		return nil, err
	}
	defer func() { _ = tx.Rollback() }()

	if a, err := str.FindAccountWithToken(token, tx); err == nil {
		principal := &rest_model_zrok.Principal{
			ID:    int64(a.Id),
			Token: a.Token,
			Email: a.Email,
		}
		return principal, nil
	} else {
		// check for admin secret
		if cfg.Admin != nil {
			for _, secret := range cfg.Admin.Secrets {
				if token == secret {
					principal := &rest_model_zrok.Principal{
						ID:    int64(-1),
						Admin: true,
					}
					return principal, nil
				}
			}
		}

		// no match
		return nil, errors2.New(401, "invalid api key")
	}
}

func edgeClient() (*rest_management_api_client.ZitiEdgeManagement, error) {
	caCerts, err := rest_util.GetControllerWellKnownCas(cfg.Ziti.ApiEndpoint)
	if err != nil {
		return nil, err
	}
	caPool := x509.NewCertPool()
	for _, ca := range caCerts {
		caPool.AddCert(ca)
	}
	return rest_util.NewEdgeManagementClientWithUpdb(cfg.Ziti.Username, cfg.Ziti.Password, cfg.Ziti.ApiEndpoint, caPool)
}

func createServiceToken() (string, error) {
	gen, err := nanoid.CustomASCII("abcdefghijklmnopqrstuvwxyz0123456789", 12)
	if err != nil {
		return "", err
	}
	return gen(), nil
}

func createToken() (string, error) {
	gen, err := nanoid.CustomASCII("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", 12)
	if err != nil {
		return "", err
	}
	return gen(), nil
}

func hashPassword(raw string) string {
	hash := sha512.New()
	hash.Write([]byte(raw))
	return hex.EncodeToString(hash.Sum(nil))
}

func realRemoteAddress(req *http.Request) string {
	ip := strings.Split(req.RemoteAddr, ":")[0]
	fwdAddress := req.Header.Get("X-Forwarded-For")
	if fwdAddress != "" {
		ip = fwdAddress

		ips := strings.Split(fwdAddress, ", ")
		if len(ips) > 1 {
			ip = ips[0]
		}
	}
	return ip
}

func proxyUrl(svcToken string) string {
	return strings.Replace(cfg.Proxy.UrlTemplate, "{svcToken}", svcToken, -1)
}
