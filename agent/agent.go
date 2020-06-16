package agent

import (
	"net/http"
	"sync"
	"time"

	hook "github.com/labbsr0x/bindman-dns-webhook/src/client"

	"github.com/gin-gonic/gin"
	"github.com/labbsr0x/bindman-dns-file-agent/agent/config"
	"github.com/labbsr0x/bindman-dns-file-agent/file"
	"github.com/sirupsen/logrus"
)

// Agent struct owns items needed to start the service
type Agent struct {
	*config.AgentBuilder
	file          file.File
	app           *gin.Engine
	SyncLock      *sync.RWMutex
	WebhookClient *hook.DNSWebhookClient
}

// InitFromAgentBuilder builds a Server instance
func (a *Agent) InitFromAgentBuilder(agentBuilder *config.AgentBuilder) *Agent {
	a.AgentBuilder = agentBuilder
	a.app = gin.Default()
	a.file = file.GetFile(a.AgentConfigPath)
	a.SyncLock = new(sync.RWMutex)
	// a.WebhookClient = hook.New(new(hook.BindmanHTTPHelper))

	logLevel, err := logrus.ParseLevel(a.AgentBuilder.LogLevel)
	if err != nil {
		logrus.Errorf("Not able to parse log level string. Setting default level: info.")
		logLevel = logrus.InfoLevel
	}
	logrus.SetLevel(logLevel)

	return a
}

// Run a Server instance
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

// Sync defines a routine for syncing the dns records present in the docker swarm and being managed by the bindman dns manager
func (a *Agent) Sync() {
	var maxTries uint = 100
	var leftTries uint = 100

	for leftTries > 0 {
		func() {
			defer a.SyncLock.Unlock()
			a.SyncLock.Lock()

			for _, record := range a.file.Records {
				logrus.Infof("%v", record)

				// bs, err := a.WebhookClient.GetRecord(record.Name, "A")
				// if err != nil { // means record was not found on manager; so we create it
				// a.delegate("create", a.file.Domain, record)
				logrus.Infof("Create %v", a.file.Domain)
				// }

				// if bs.Name != record.Name || bs.Value != a.DNSReverseProxyAddr || bs.Type != "A" { // if true, record exists and needs to be update
				// a.delegate("update", a.file.Domain, record)
				logrus.Infof("Update %v", a.file.Domain)
				// }
			}
		}()
		backoffWait(maxTries, leftTries, time.Minute) // wait time increases exponentially
		leftTries--
	}
}

// delegate appropriately calls the dns manager to handle the addition or removal of a DNS rule
// func (a *Agent) delegate(action string, domain string, record file.Records) {
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
