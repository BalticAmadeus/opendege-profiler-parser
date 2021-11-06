# openedge-profiler-parser

Progress OpenEdge Profiler data parsing to OpenTracing format.

# Prerequisites
## In order to RUN you will be enough with
Docker: https://www.docker.com/products/docker-desktop
Zipkin: https://hub.docker.com/r/openzipkin/zipkin

To start zipkin docker image just run this command while having docker deamon active:
```sh
docker run -d -p 9411:9411 openzipkin/zipkin
```
NOTE: if you are going to change port, update information in code accordingly. At the moment there's no single place/properties file for setup of constants.

### Run on Jaeger
Jaeger docker image automatically exposes Zipkin compatible REST API, so all you need to do is start docker image using this command:
```sh
docker run -d -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 -p 16686:16686 -p 9411:9411 jaegertracing/all-in-one
```

### Run on Grafana
To run on Grafana, you will need to create docker-compose.yml file with services for Zipkin and Grafana (example in docker-compose-sample.yml). Then run:
```sh
docker-compose up -d
```
After this, you have to add Zipkin data source in Grafana with URL zipkin:9411.

## In order to DEVELOP you will need
Some kind of modern IDE that supports GO (we are using VS Code: https://code.visualstudio.com/download)

Golang engine: https://golang.org/dl/

# Usage of binary (compiled) utility
To run binary file on Unix systems use these commands:
```sh
chmod +x profiler-opentracing
profiler-opentracing /path/to/profiler/output/file -config=./config.yaml
```

To run binary file on Windows systems run this command:
```sh
profiler-opentracing.exe \path\to\profiler\output\file -config=./config.yamll
```

`-config` Flag passes given configuration. If this flag is not provided then default values will be used.

# Usage for project developers
NOTE: unix-windows slashes should be directed different ways. Keep this in mind based on your system
To run parser run:
```sh
cd progress-opentracing-profiler
go run .\cmd\parse.go .\profilerFiles\simple_oop.prof .\config\config.yaml
```

To run tests:
```sh
cd progress-opentracing-profiler
go test ./tests/...
```

Or if you want to test specific module, then (-v flag is not mandatory, but it gives the details of the tests):
```sh
cd progress-opentracing-profiler
go test -v ./tests/<module_name>
```

Run linter:
```sh
golangci-lint run
```

# Custom configuration
As for now custom configuration can be provided in one of 2 ways:
- Setting environment variables
- Providing config.yaml file in the root directory of the repository or alongside binary file. (config-sample.yaml provided in repository)

Supported configuration values:

All configuration values are under zipkin tag.

| Name        | Default value         | Description                                                                                                                                |
| ----------- | :-------------------: | -----------------------------------------------------------------------------------------------------------------------------------------  |
| url         | http://localhost:9411 | Default URL of Zipkin container with port 9411. If Zipkin configuration has been changed, the configuration values have to be changed too. |
| endpoint    | /api/v2/spans         | Endpoint of Zipkin span collector. It should NOT be changed.                                                                               |
| hostPort    | localhost:80          | It's used for Zipkin tag, in this case, it shows the address and port of traced application's endpoint. It can and should be changed.      |
| serviceName | TestService2          | It's used for naming the trace, can and should be changed.                                                                                 |
