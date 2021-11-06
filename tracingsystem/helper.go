package tracingsystem

import (
	"context"
	"time"

	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/reporter"
	httpreporter "github.com/openzipkin/zipkin-go/reporter/http"
	"gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/config"
	"gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/profiler/util"
)

type ZipkinProperties struct {
	reporter  reporter.Reporter
	tracer    *zipkin.Tracer
	startTime time.Time
}

func (zp *ZipkinProperties) GetNewTracer(serviceName string) {
	endpointURL := config.Cfg.Zipkin.Url + config.Cfg.Zipkin.Endpoint
	zp.reporter = httpreporter.NewReporter(endpointURL)

	endpoint, err := zipkin.NewEndpoint(serviceName, config.Cfg.Zipkin.HostPort)
	util.ValidateErrorStatus(err)

	// set-up our sampling strategy
	sampler, err := zipkin.NewCountingSampler(1)
	util.ValidateErrorStatus(err)

	tracer, err := zipkin.NewTracer(
		zp.reporter,
		zipkin.WithLocalEndpoint(endpoint),
		zipkin.WithSampler(sampler),
	)
	util.ValidateErrorStatus(err)

	zp.tracer = tracer
	zp.startTime = time.Now()
}

func (zp *ZipkinProperties) GetSpanAndContext(moduleName string, ctx context.Context, timeToAdd time.Duration) (zipkin.Span, context.Context) {
	var span zipkin.Span
	startTime := zipkin.StartTime(zp.startTime.Add(timeToAdd))

	if ctx == context.Background() {
		span = zp.tracer.StartSpan(moduleName, startTime)
		ctx = zipkin.NewContext(ctx, span)
	} else {
		span, ctx = zp.tracer.StartSpanFromContext(ctx, moduleName, startTime)
	}

	return span, ctx
}

func (zp ZipkinProperties) GetTracer() *zipkin.Tracer {
	return zp.tracer
}

func (zp ZipkinProperties) GetReporter() reporter.Reporter {
	return zp.reporter
}
