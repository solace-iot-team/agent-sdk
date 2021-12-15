package config

import (
	"github.com/Axway/agent-sdk/pkg/util/log"
	"io/ioutil"
	"net/url"
	"os"
	"time"

	"github.com/Axway/agent-sdk/pkg/util/exception"
)

const tokenEndpoint = "/protocol/openid-connect/token"

// AuthConfig - Interface for service account config
type AuthConfig interface {
	GetTokenURL() string
	GetRealm() string
	GetAudience() string
	GetClientID() string
	GetPrivateKey() string
	GetPublicKey() string
	GetKeyPassword() string
	GetTimeout() time.Duration
	validate()
}

// AuthConfiguration -
type AuthConfiguration struct {
	AuthConfig
	URL            string        `config:"url"`
	Realm          string        `config:"realm"`
	ClientID       string        `config:"clientId"`
	PrivateKey     string        `config:"privateKey"`
	PublicKey      string        `config:"publicKey"`
	PrivateKeyData string        `config:"privateKeyData"`
	PublicKeyData  string        `config:"publicKeyData"`
	KeyPwd         string        `config:"keyPassword"`
	Timeout        time.Duration `config:"timeout"`
}

func newAuthConfig() AuthConfig {
	return &AuthConfiguration{
		Timeout: 30 * time.Second,
	}
}

func (a *AuthConfiguration) validate() {
	if a.URL == "" {
		exception.Throw(ErrBadConfig.FormatError(pathAuthURL))
	} else if _, err := url.ParseRequestURI(a.URL); err != nil {
		exception.Throw(ErrBadConfig.FormatError(pathAuthURL))
	}

	if a.GetRealm() == "" {
		exception.Throw(ErrBadConfig.FormatError(pathAuthRealm))
	}

	if a.GetClientID() == "" {
		exception.Throw(ErrBadConfig.FormatError(pathAuthClientID))
	}

	a.validatePrivateKey()
	a.validatePublicKey()
}

func (a *AuthConfiguration) validatePrivateKey() {
	log.Tracef("validating PrivateKey Setting [CENTRAL_AUTH_PRIVATEKEY:%s]", a.GetPrivateKey())
	if a.GetPrivateKey() == "" {
		log.Warn("CENTRAL_AUTH_PRIVATEKEY not defined or empty")
		exception.Throw(ErrBadConfig.FormatError(pathAuthPrivateKey))
	} else {
		if !fileExists(a.GetPrivateKey()) {
			log.Tracef("CENTRAL_AUTH_PRIVATEKEY file does not exist: %s", a.GetPrivateKey())
			privateKeyData := os.Getenv("CENTRAL_AUTH_PRIVATEKEY_DATA")
			if privateKeyData == "" {
				//todo JT REMOVE
				log.Warn("CENTRAL_AUTH_PRIVATEKEY_DATA is empty")
				exception.Throw(ErrBadConfig.FormatError(pathAuthPrivateKey))
			}
			if err := saveKeyData(a.GetPrivateKey(), privateKeyData); err != nil {
				// todo JT REMOVE
				log.Errorf("Can not write private key to file location %s %s", a.GetPrivateKey(), err)
				exception.Throw(ErrReadingKeyFile.FormatError("private key", a.GetPrivateKey()))
			}

		}
		// Validate that the file is readable
		if _, err := os.Open(a.GetPrivateKey()); err != nil {
			// todo JT REMOVE
			log.Errorf("CENTRAL_AUTH_PRIVATEKEY:%s file is not readable %s", a.GetPrivateKey(), err)
			exception.Throw(ErrReadingKeyFile.FormatError("private key", a.GetPrivateKey()))
		}
	}
}

func (a *AuthConfiguration) validatePublicKey() {
	log.Tracef("validating PublicKey Setting [CENTRAL_AUTH_PUBLICKEY:%s]", a.GetPublicKey())
	if a.GetPublicKey() == "" {
		log.Warn("CENTRAL_AUTH_PUBLICKEY not defined or empty")
		exception.Throw(ErrBadConfig.FormatError(pathAuthPublicKey))
	} else {
		if !fileExists(a.GetPublicKey()) {
			publicKeyData := os.Getenv("CENTRAL_AUTH_PUBLICKEY_DATA")
			if publicKeyData == "" {
				//todo JT REMOVE
				log.Warn("CENTRAL_AUTH_PUBLICKEY_DATA is empty")
				exception.Throw(ErrBadConfig.FormatError(pathAuthPublicKey))
			}
			if err := saveKeyData(a.GetPublicKey(), publicKeyData); err != nil {
				log.Errorf("Can not write public key to file location %s %s", a.GetPublicKey(), err)
				exception.Throw(ErrReadingKeyFile.FormatError("public key", a.GetPublicKey()))
			}
		}
		// Validate that the file is readable
		if _, err := os.Open(a.GetPublicKey()); err != nil {
			// todo JT REMOVE
			log.Errorf("CENTRAL_AUTH_PUBLICKEY:%s file is not readable %s", a.GetPublicKey(), err)
			exception.Throw(ErrReadingKeyFile.FormatError("public key", a.GetPublicKey()))
		}
	}
}

// GetTokenURL - Returns the token URL
func (a *AuthConfiguration) GetTokenURL() string {
	if a.URL == "" || a.Realm == "" {
		return ""
	}
	return a.URL + "/realms/" + a.Realm + tokenEndpoint
}

// GetRealm - Returns the token audience URL
func (a *AuthConfiguration) GetRealm() string {
	return a.Realm
}

// GetAudience - Returns the token audience URL
func (a *AuthConfiguration) GetAudience() string {
	if a.URL == "" || a.Realm == "" {
		return ""
	}
	return a.URL + "/realms/" + a.Realm
}

// GetClientID - Returns the token audience URL
func (a *AuthConfiguration) GetClientID() string {
	return a.ClientID
}

// GetPrivateKey - Returns the private key file path
func (a *AuthConfiguration) GetPrivateKey() string {
	return a.PrivateKey
}

// GetPublicKey - Returns the public key file path
func (a *AuthConfiguration) GetPublicKey() string {
	return a.PublicKey
}

// GetKeyPassword - Returns the token audience URL
func (a *AuthConfiguration) GetKeyPassword() string {
	return a.KeyPwd
}

// GetTimeout - Returns the token audience URL
func (a *AuthConfiguration) GetTimeout() time.Duration {
	return a.Timeout
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func saveKeyData(filename string, data string) error {
	dataBytes := []byte(data)
	return ioutil.WriteFile(filename, dataBytes, 0600)
}
