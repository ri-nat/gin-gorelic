package gorelic

import (
	"time"

	"github.com/gin-gonic/gin"
	metrics "github.com/yvasiyarov/go-metrics"
	"github.com/yvasiyarov/gorelic"
)

// NewrelicAgentMiddleware inits gorelic's NewRelic object and returns handler function
func NewrelicAgentMiddleware(license string, appname string, verbose bool) gin.HandlerFunc {
	var agent *gorelic.Agent

	if license == "" {
		return nil
	}

	agent = gorelic.NewAgent()
	agent.NewrelicLicense = license

	agent.HTTPTimer = metrics.NewTimer()
	agent.CollectHTTPStat = true
	agent.Verbose = verbose

	agent.NewrelicName = appname
	agent.Run()

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()

		if agent != nil {
			agent.HTTPTimer.UpdateSince(startTime)
		}
	}
}
