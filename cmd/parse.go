package main

import (
	"flag"
	"fmt"
	"time"

	config "gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/config"
	"gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/profiler/processor"
	"gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/profiler/reader"
	"gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/profiler/util"
)

const separatorLine = "------------------------------------------"

func main() {
	profilerFileToParse, configFile := parseInputParameters()
	util.ValidateFile(profilerFileToParse)
	util.ValidateFile(configFile)
	config.GetConfig(configFile)

	fmt.Println("Reading profiler file: " + profilerFileToParse)
	fileReader := reader.ReadProfilerFile(profilerFileToParse)

	fmt.Println("Start transforming data")
	profDataReader := reader.TransformProfilerData(fileReader)
	fmt.Println(separatorLine)

	start := time.Now()
	traceProcessor := processor.ProcessProfilerData(profDataReader.GetRaw())
	duration := time.Since(start)

	traceProcessor.ReportDataToZipkin()
	printResults(duration)
}

func parseInputParameters() (string, string) {
	flag.Parse()

	return flag.Arg(0), flag.Arg(1)
}

func printResults(duration time.Duration) {
	fmt.Println("Parsing took: ", duration)
	fmt.Println(separatorLine)
	fmt.Println("Done. Check zipkin at " + config.Cfg.Zipkin.Url + "/" + "\n")
}
