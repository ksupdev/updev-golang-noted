```powershell
go get github.com/gin-gonic/gin
go get gopkg.in/h2non/gentleman.v2
go get go.opentelemetry.io/otel
go get go.opentelemetry.io/otel/sdk
go get go.opentelemetry.io/otel/semconv/v1.7.0
go get go.opentelemetry.io/otel/exporters/jaeger

go get go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin


go get go.opentelemetry.io/otel/semconv

go get go.opentelemetry.io/otel/exporters/jaeger
go get go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp

```

- go mod init labs.test/opl-jager02
- docker-compose up jaeger 
- Jaeger search ui ``http://localhost:16686/search``


- https://github.com/open-telemetry/opentelemetry-go-contrib/tree/main/instrumentation

- Build command  go build main.go

### Test Step

- run docker compose for start jaeger ``docker-compose up -d``
- Star main_serv by ``go run .\main_serv\main.go``
- Star micro_serv01 by ``go run .\micro_serv01\main.go``
- test with ``\docs\dev.http``
