package v1

import (
	operatorv1 "github.com/openshift/api/operator/v1"
)

type ConsoleServerConfig struct {
	APIVersion    string                          `json:"apiVersion"`
	Kind          string                          `json:"kind"`
	ServingInfo   ConsoleServingInfo              `json:"servingInfo"`
	ClusterInfo   ConsoleClusterInfo              `json:"clusterInfo"`
	Auth          ConsoleAuth                     `json:"auth"`
	Customization operatorv1.ConsoleCustomization `json:"customization"`
}

// ServingInfo holds configuration for serving HTTP.
type ConsoleServingInfo struct {
	BindAddress string `json:"bindAddress"`
	CertFile    string `json:"certFile"`
	KeyFile     string `json:"keyFile"`

	// These fields are defined in `HTTPServingInfo`, but are not supported for console. Fail if any are specified.
	// https://github.com/openshift/api/blob/0cb4131a7636e1ada6b2769edc9118f0fe6844c8/config/v1/types.go#L7-L38
	BindNetwork           string        `json:"bindNetwork"`
	ClientCA              string        `json:"clientCA"`
	NamedCertificates     []interface{} `json:"namedCertificates"`
	MinTLSVersion         string        `json:"minTLSVersion"`
	CipherSuites          []string      `json:"cipherSuites"`
	MaxRequestsInFlight   int64         `json:"maxRequestsInFlight"`
	RequestTimeoutSeconds int64         `json:"requestTimeoutSeconds"`
}

// ClusterInfo holds information the about the cluster such as master public URL and console public URL.
type ConsoleClusterInfo struct {
	ConsoleBaseAddress string `json:"consoleBaseAddress"`
	ConsoleBasePath    string `json:"consoleBasePath"`
	MasterPublicURL    string `json:"masterPublicURL"`
}

// Auth holds configuration for authenticating with OpenShift. The auth method is assumed to be "openshift".
type ConsoleAuth struct {
	ClientID            string `json:"clientID"`
	ClientSecretFile    string `json:"clientSecretFile"`
	OAuthEndpointCAFile string `json:"oauthEndpointCAFile"`
	LogoutRedirect      string `json:"logoutRedirect"`
}
