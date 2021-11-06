package processor_test

import (
	"bufio"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"

	"gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/profiler/processor"
	"gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/profiler/reader"
	"gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/profiler/util"
	test_helper "gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/tests"
)

func TestProcessProfilerData(t *testing.T) {
	fileReader := reader.ReadProfilerFile("../testData/testFullTrace.txt")
	profDataReader := reader.TransformProfilerData(fileReader)
	processedData := processor.ProcessProfilerData(profDataReader.GetRaw())

	callTreeTraceFromFile := getExpectedResultString()
	callTreeTraceFromParser := getCallTreeString(processedData)

	test_helper.CheckStringValues(callTreeTraceFromFile, callTreeTraceFromParser, "CallTree", t)
}

func getExpectedResultString() string {
	file, err := ioutil.ReadFile("../testData/testGenerateCallTree.txt")
	util.ValidateErrorStatus(err)

	callTreeTraceFromFile := ""
	scanner := bufio.NewScanner(strings.NewReader(string(file)))
	for scanner.Scan() {
		s := scanner.Text()
		callTreeTraceFromFile += s
	}

	return callTreeTraceFromFile
}

func getCallTreeString(tp *processor.TraceProcessor) string {
	var callTreeString string

	for i := 0; i < len(tp.GetTreeData()); i++ {
		callTreeString += tp.GetTreeData()[i].GetModuleName() + " " + strconv.Itoa(tp.GetTreeData()[i].GetTrace().GetModuleId()) + " " + strconv.Itoa(tp.GetTreeData()[i].GetTrace().GetLineNo()) + " " + tp.GetTreeData()[i].GetTrace().GetActualTime().String() + " " + tp.GetTreeData()[i].GetTrace().GetStartTime().String()
	}

	return callTreeString
}
