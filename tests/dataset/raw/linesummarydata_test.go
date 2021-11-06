package raw_test

import (
	"strconv"
	"testing"

	"gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/profiler/dataset/raw"
	test_helper "gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/tests"
)

func TestParse_LineSummaryData(t *testing.T) {
	lineSummaryData := *new(raw.Linesummarydata)

	sampleData := []string{"0", "0", "1", "0.000000", "0.001782"}
	lineSummaryData.Parse(sampleData)

	intActualTime := test_helper.GetTimeFloat("0.000000")
	intCumulativeTime := test_helper.GetTimeFloat("0.001782")
	currentActualTime := test_helper.GetTimeDuration(intActualTime)
	currentCumulativeTime := test_helper.GetTimeDuration(intCumulativeTime)

	test_helper.CheckStringValues("0", strconv.Itoa(lineSummaryData[0].GetModuleId()), "ModuleId", t)
	test_helper.CheckStringValues("0", strconv.Itoa(lineSummaryData[0].GetLineNo()), "LineNo", t)
	test_helper.CheckStringValues("1", strconv.Itoa(lineSummaryData[0].GetExecCount()), "ExecCount", t)
	test_helper.CheckTimeDurationValues(currentActualTime, lineSummaryData[0].GetActualTime(), "ActualTime", t)
	test_helper.CheckTimeDurationValues(currentCumulativeTime, lineSummaryData[0].GetCumulativeTime(), "CumulativeTime", t)

}
