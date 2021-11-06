package processor

import "gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/profiler/dataset/raw"

func ProcessProfilerData(rawDataset *raw.Full) *TraceProcessor {
	traceProcessor := new(TraceProcessor)
	traceProcessor.SetRawData(rawDataset)
	traceProcessor.ParseTrace()

	return traceProcessor
}
