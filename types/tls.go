package types

// TLSVersion is a strongly-typed int used for relative comparison
// (minimum, maximum, greater than, less than) of TLS versions
type TLSVersion int

const (
	// Error value, excluded from lookup maps
	TLSVersionInvalid TLSVersion = iota - 1

	// Explictly allow implementation to select TLS version
	// May be useful to supercede defaults specified at a higher layer
	TLSVersionAuto

	_ // Placeholder for SSLv3, hopefully we won't have to add this

	// TLS versions
	TLSv1_0
	TLSv1_1
	TLSv1_2
	TLSv1_3
)

var (
	TLSVersions = map[string]TLSVersion{
		"TLS_AUTO": TLSVersionAuto,
		"TLSv1_0":  TLSv1_0,
		"TLSv1_1":  TLSv1_1,
		"TLSv1_2":  TLSv1_2,
		"TLSv1_3":  TLSv1_3,
	}
	// NOTE: This interface is deprecated in favor of TLSVersions
	// and should be eventually removed in a future release.
	DeprecatedAgentTLSVersions = map[string]TLSVersion{
		"":      TLSVersionAuto,
		"tls10": TLSv1_0,
		"tls11": TLSv1_1,
		"tls12": TLSv1_2,
		"tls13": TLSv1_3,
	}
	HumanTLSVersionStrings = map[TLSVersion]string{
		TLSVersionAuto: "Allow implementation to select TLS version",
		TLSv1_0:        "TLS 1.0",
		TLSv1_1:        "TLS 1.1",
		TLSv1_2:        "TLS 1.2",
		TLSv1_3:        "TLS 1.3",
	}
	EnvoyTLSVersionStrings = map[TLSVersion]string{
		TLSVersionAuto: "TLS_AUTO",
		TLSv1_0:        "TLSv1_0",
		TLSv1_1:        "TLSv1_1",
		TLSv1_2:        "TLSv1_2",
		TLSv1_3:        "TLSv1_3",
	}
)

func (v TLSVersion) String() string {
	return HumanTLSVersionStrings[v]
}

func (v TLSVersion) EnvoyString() string {
	return EnvoyTLSVersionStrings[v]
}

// IANA cipher suite constants and values as defined at
// https://www.iana.org/assignments/tls-parameters/tls-parameters.xhtml
// This is the total list of TLS 1.2-style cipher suites
// which are currently supported by either Envoy 1.21 or the Consul agent
// via Go, and may change as some older suites are removed in future
// Envoy releases and Consul drops support for older Envoy versions,
// and as supported cipher suites in the Go runtime change.
//
// The naming convention for cipher suites changed in TLS 1.3
// but constant values should still be globally unqiue
// Handling validation on a subset of TLSCipherSuite constants
// would be a future exercise if cipher suites for TLS 1.3 ever
// become configurable in BoringSSL, Envoy, or other implementation
type TLSCipherSuite uint16

const (
	// Envoy cipher suites also used by Consul agent
	TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256       TLSCipherSuite = 0xc02b
	TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256                = 0xcca9 // Not used by Consul agent yet
	TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256                        = 0xc02f
	TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256                  = 0xcca8 // Not used by Consul agent yet
	TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA                         = 0xc009
	TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA                           = 0xc013
	TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384                      = 0xc02c
	TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384                        = 0xc030
	TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA                         = 0xc00a
	TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA                           = 0xc014

	// Older cipher suites not supported for Consul agent TLS, will eventually be removed from Envoy defaults
	TLS_RSA_WITH_AES_128_GCM_SHA256 = 0x009c
	TLS_RSA_WITH_AES_128_CBC_SHA    = 0x002f
	TLS_RSA_WITH_AES_256_GCM_SHA384 = 0x009d
	TLS_RSA_WITH_AES_256_CBC_SHA    = 0x0035

	// Additional cipher suites used by Consul agent but not Envoy
	// TODO: these are both explicitly listed as insecure and disabled in the Go source, should they be removed?
	// https://cs.opensource.google/go/go/+/refs/tags/go1.17.3:src/crypto/tls/cipher_suites.go;l=329-330
	TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256 = 0x0023
	TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256   = 0xc027
)

var (
	TLSCipherSuites = map[string]TLSCipherSuite{
		"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256": TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256,
		"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA":          TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
		"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256":       TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA":          TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
		"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384":       TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256":   TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256,
		"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA":            TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
		"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256":         TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA":            TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
		"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384":         TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,

		"TLS_RSA_WITH_AES_128_GCM_SHA256": TLS_RSA_WITH_AES_128_GCM_SHA256,
		"TLS_RSA_WITH_AES_128_CBC_SHA":    TLS_RSA_WITH_AES_128_CBC_SHA,
		"TLS_RSA_WITH_AES_256_GCM_SHA384": TLS_RSA_WITH_AES_256_GCM_SHA384,
		"TLS_RSA_WITH_AES_256_CBC_SHA":    TLS_RSA_WITH_AES_256_CBC_SHA,

		"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256": TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,
		"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256":   TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
	}
	HumanTLSCipherSuiteStrings = map[TLSCipherSuite]string{
		TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256: "TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256",
		TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA:          "TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA",
		TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256:       "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
		TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA:          "TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA",
		TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384:       "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
		TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256:   "TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256",
		TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA:            "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA",
		TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256:         "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
		TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA:            "TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA",
		TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384:         "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",

		TLS_RSA_WITH_AES_128_GCM_SHA256: "TLS_RSA_WITH_AES_128_GCM_SHA256",
		TLS_RSA_WITH_AES_128_CBC_SHA:    "TLS_RSA_WITH_AES_128_CBC_SHA",
		TLS_RSA_WITH_AES_256_GCM_SHA384: "TLS_RSA_WITH_AES_256_GCM_SHA384",
		TLS_RSA_WITH_AES_256_CBC_SHA:    "TLS_RSA_WITH_AES_256_CBC_SHA",

		TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256: "TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256",
		TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256:   "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256",
	}
	EnvoyTLSCipherSuiteStrings = map[TLSCipherSuite]string{
		TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256:       "ECDHE-ECDSA-AES128-GCM-SHA256",
		TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256: "ECDHE-ECDSA-CHACHA20-POLY1305",
		TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256:         "ECDHE-RSA-AES128-GCM-SHA256",
		TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256:   "ECDHE-RSA-CHACHA20-POLY1305",
		TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA:          "ECDHE-ECDSA-AES128-SHA",
		TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA:            "ECDHE-RSA-AES128-SHA",
		TLS_RSA_WITH_AES_128_GCM_SHA256:               "AES128-GCM-SHA256",
		TLS_RSA_WITH_AES_128_CBC_SHA:                  "AES128-SHA",
		TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384:       "ECDHE-ECDSA-AES256-GCM-SHA384",
		TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384:         "ECDHE-RSA-AES256-GCM-SHA384",
		TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA:          "ECDHE-ECDSA-AES256-SHA",
		TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA:            "ECDHE-RSA-AES256-SHA",
		TLS_RSA_WITH_AES_256_GCM_SHA384:               "AES256-GCM-SHA384",
		TLS_RSA_WITH_AES_256_CBC_SHA:                  "AES256-SHA",
	}
)