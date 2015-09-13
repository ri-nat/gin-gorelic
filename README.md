GoRelic
=======

NewRelic middleware for gin-gonic framework.

## Installation
`go get github.com/ri-nat/gin-gorelic`

## Usage

```go
import(
	"github.com/gin-gonic/gin"
	"github.com/ri-nat/gin-gorelic"
)

func main(){
	g := gin.Default()

	handler := gorelic.NewrelicAgentMiddleware("NEWRELIC_LICENSE_KEY", "app name", true)
	g.Use(handler)

	g.Run()
}
```

## Authors

* [Rinat Shaykhutdinov](http://github.com/ri-nat)
* [Jason Waldrip](http://github.com/jwaldrip)
* [Yuriy Vasiyarov](http://github.com/yvasiyarov) [original martini gorelic]
