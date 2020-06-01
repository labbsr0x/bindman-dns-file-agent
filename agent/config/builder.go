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
	agentConfigPath     = "agent-config-file"
	logLevel            = "log-level"
	port                = "port"
)

// Flags define the fields that will be passed via cmd
type Flags struct {
	DNSManagerAddr      string
	DNSReverseProxyAddr string
	AgentConfigPath     string
	LogLevel            string
	Port                string
}

// AgentBuilder defines the parametric information of a server instance
type AgentBuilder struct {
	*Flags
}

// AddFlags adds flags for Builder.
func AddFlags(flags *pflag.FlagSet) {
	flags.StringP(dnsManagerAddr, "d", "", "Bindman DNS Manager Address")
	flags.StringP(dnsReverseProxyAddr, "r", "", "Bindman DNS Reverse Proxy Address")
	flags.StringP(agentConfigPath, "c", "", "Bindman Agent Config Path")
	flags.StringP(logLevel, "l", "info", "[optional] Sets the Log Level to one of seven (trace, debug, info, warn, error, fatal, panic). Default: info")
	flags.StringP(port, "p", "7071", "[optional] Custom port for accessing bindman agent's services. Default: 7071")
}

// Init initializes the web server builder with properties retrieved from Viper.
func (b *AgentBuilder) Init(v *viper.Viper) *AgentBuilder {
	flags := new(Flags)
	flags.DNSManagerAddr = v.GetString(dnsManagerAddr)
	flags.DNSReverseProxyAddr = v.GetString(dnsReverseProxyAddr)
	flags.AgentConfigPath = v.GetString(agentConfigPath)
	flags.LogLevel = v.GetString(logLevel)
	flags.Port = v.GetString(port)
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
		{flags.AgentConfigPath, agentConfigPath},
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
