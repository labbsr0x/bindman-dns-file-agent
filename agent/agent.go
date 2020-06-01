package agent

import (
	"fmt"
	"net/http"
	"sync"

	hook "github.com/labbsr0x/bindman-dns-webhook/src/client"

	"github.com/gin-gonic/gin"
	"github.com/labbsr0x/bindman-dns-file-agent/agent/config"
	"github.com/labbsr0x/bindman-dns-file-agent/file"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

type Agent struct {
	*config.AgentBuilder
	*file.FileBuilder
	app           *gin.Engine
	SyncLock      *sync.RWMutex
	WebhookClient *hook.DNSWebhookClient
}

// InitFromAgentBuilder builds a Server instance
func (a *Agent) InitFromAgentBuilder(agentBuilder *config.AgentBuilder) *Agent {
	a.AgentBuilder = agentBuilder
	a.app = gin.Default()
	a.FileBuilder = a.FileBuilder.Init(a.AgentConfigPath)

	logLevel, err := logrus.ParseLevel(a.AgentBuilder.LogLevel)
	if err != nil {
		logrus.Errorf("Not able to parse log level string. Setting default level: info.")
		logLevel = logrus.InfoLevel
	}
	logrus.SetLevel(logLevel)

	return a
}

func (a *Agent) Run() {
	consumerGroup := a.app.Group("/")
	{
		consumerGroup.GET("/", index)
	}
	a.app.Run("0.0.0.0:" + a.AgentBuilder.Port)
}

func index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Welcome to Bindman DNS File Agent")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func startSync() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 * * * * *", func() { fmt.Println("Every minute") })
	c.AddFunc("*/5 * * * * *", func() { fmt.Println("Every 5 second") })
	c.Start()
}

// Sync defines a routine for syncing the dns records present in the docker swarm and being managed by the bindman dns manager
// func (a *Agent) Sync() {
// 	var maxTries uint = 100
// 	var leftTries uint = 100

// 	for leftTries > 0 {
// 		func() {
// 			defer a.SyncLock.Unlock()
// 			a.SyncLock.Lock()

// 			for _, record := range a.File.Records {
// 				logrus.Infof("%v", record)

// 				bs, err := a.WebhookClient.GetRecord(record.Name, "A")
// 				if err != nil { // means record was not found on manager; so we create it
// 					a.delegate("create", a.File.Domain, record)
// 				}

// 				if bs.Name != record.Name || bs.Value != a.DNSReverseProxyAddr || bs.Type != "A" { // if true, record exists and needs to be update
// 					a.delegate("update", a.File.Domain, record)
// 				}
// 			}
// 		}()
// 		backoffWait(maxTries, leftTries, time.Minute) // wait time increases exponentially
// 		leftTries--
// 	}
// }

// delegate appropriately calls the dns manager to handle the addition or removal of a DNS rule
// func (a *Agent) delegate(action string, domain string, record *file.Records) {
// 	var ok bool
// 	var err error

// 	if action == "update" {
// 		ok, err = a.WebhookClient.UpdateRecord(&hookTypes.DNSRecord{Name: record.Name, Type: record.Type, Value: record.Value})
// 	}

// 	if action == "create" {
// 		ok, err = a.WebhookClient.AddRecord(record.Name, record.Type, record.Value) // adds to the dns manager
// 	}

// 	if !ok {
// 		logrus.Errorf("Error to %v the Domain '%v' from the service '%v': %v", action, domain, record.Name, err)
// 	}

// }
