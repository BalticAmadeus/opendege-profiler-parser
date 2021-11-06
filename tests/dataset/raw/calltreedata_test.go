package raw_test

import (
	"strconv"
	"testing"

	"gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/profiler/dataset/raw"
	test_helper "gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/tests"
)

func TestParse_TreeData(t *testing.T) {
	callTreeData := *new(raw.Calltreedata)

	sampleData := []string{"0", "0", "1", "1"} // TODO: move sampleData and callTreeData assignments, code duplication
	callTreeData.Parse(sampleData)

	test_helper.CheckStringValues("0", strconv.Itoa(callTreeData[0].GetCallerId()), "CallerId", t)
	test_helper.CheckStringValues("0", strconv.Itoa(callTreeData[0].GetCallerLineNo()), "CallerLineNo", t)
	test_helper.CheckStringValues("1", strconv.Itoa(callTreeData[0].GetCalleeId()), "CalleeId", t)
	test_helper.CheckStringValues("1", strconv.Itoa(callTreeData[0].GetCallCount()), "CallCount", t)
}
