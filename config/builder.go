package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	dnsManagerAddr      = "dns-manager-addr"
	dnsReverseProxyAddr = "dns-reverse-proxy-addr"
	dnsConfigFile       = "dns-config-file"
)

// Flags define the fields that will be passed via cmd
type Flags struct {
	DNSManagerAddr      string
	DNSReverseProxyAddr string
	DNSConfigFile       string
}

// WebBuilder defines the parametric information of a server instance
type WebBuilder struct {
	*Flags
}

// AddFlags adds flags for Builder.
func AddFlags(flags *pflag.FlagSet) {
	flags.StringP(dnsManagerAddr, "d", "", "DNS Manager Address")
	flags.StringP(dnsReverseProxyAddr, "r", "", "DNS Reverse Proxy Address")
	flags.StringP(dnsConfigFile, "f", "", "DNS Config File")
}

// Init initializes the web server builder with properties retrieved from Viper.
func (b *WebBuilder) Init(v *viper.Viper) *WebBuilder {
	flags := new(Flags)
	flags.DNSManagerAddr = v.GetString(dnsManagerAddr)
	flags.DNSReverseProxyAddr = v.GetString(dnsReverseProxyAddr)
	flags.DNSConfigFile = v.GetString(dnsConfigFile)
	flags.check()

	b.Flags = flags

	return b
}

func (flags *Flags) check() {
	logrus.Infof("Flags: '%v'", flags)

	requiredFlags := []struct {
		value string
		name  string
	}{
		{flags.DNSManagerAddr, dnsManagerAddr},
		{flags.DNSReverseProxyAddr, dnsReverseProxyAddr},
		{flags.DNSConfigFile, dnsConfigFile},
	}

	var errMsg string

	for _, flag := range requiredFlags {
		if flag.value == "" {
			errMsg += fmt.Sprintf("\n\t%v", flag.name)
		}
	}

	if errMsg != "" {
		errMsg = "The following flags are missing: " + errMsg
		panic(errMsg)
	}
}
