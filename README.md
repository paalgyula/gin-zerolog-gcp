# gin-zerolog-gcp

go package to serve GCP (Google Cloud Platform) style logs with [zerolog](https://github.com/rs/zerolog) and with [gin-gonic](github.com/gin-gonic/gin) request router.

Sample usage with gin:

`go get github.com/paalgyula/gin-zerolog-gcp`

```go
    import gcp "github.com/paalgyula/gin-zerolog-gcp"

    ...

func main() {
    gke.SetupLogger(os.GetEnv("DEBUG") != "")

    api := gin.New()

    api.Engine.Use(gin.Recovery())
	api.Engine.Use(gcp.WithAccessLog())

    ...
}

```
