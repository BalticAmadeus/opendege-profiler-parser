package raw_test

import (
	"strconv"
	"testing"

	"gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/profiler/dataset/raw"
	test_helper "gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/tests"
)

func TestParse_TracingData(t *testing.T) {
	tracingData := *new(raw.Tracingdata)

	sampleData := []string{"2", "7", "0.000000", "0.778602"}
	tracingData.Parse(sampleData)

	intActualTime := test_helper.GetTimeFloat("0.000000")
	intStartTime := test_helper.GetTimeFloat("0.778602")
	currentActualTime := test_helper.GetTimeDuration(intActualTime)
	currentStartTime := test_helper.GetTimeDuration(intStartTime)

	test_helper.CheckStringValues("2", strconv.Itoa(tracingData[0].GetModuleId()), "ModuleId", t)
	test_helper.CheckStringValues("7", strconv.Itoa(tracingData[0].GetLineNo()), "LineNo", t)
	test_helper.CheckTimeDurationValues(currentActualTime, tracingData[0].GetActualTime(), "ActualTime", t)
	test_helper.CheckTimeDurationValues(currentStartTime, tracingData[0].GetStartTime(), "StartTime", t)

}

func TestGetTracing(t *testing.T) {
	tracingData := *new(raw.Tracingdata)

	sampleData := []string{"2", "7", "0.000000", "0.778602"}
	tracingData.Parse(sampleData)
	getTracing := tracingData.GetTracing(2, 7)

	test_helper.CheckStringValues(strconv.Itoa(tracingData[0].GetModuleId()), strconv.Itoa(getTracing.GetModuleId()), "ModuleId", t)
	test_helper.CheckStringValues(strconv.Itoa(tracingData[0].GetLineNo()), strconv.Itoa(getTracing.GetLineNo()), "LineNo", t)
	test_helper.CheckTimeDurationValues(tracingData[0].GetActualTime(), getTracing.GetActualTime(), "ActualTime", t)
	test_helper.CheckTimeDurationValues(tracingData[0].GetStartTime(), getTracing.GetStartTime(), "StartTime", t)

}

func TestClassInstatiation(t *testing.T) {
	tracingData := *new(raw.Tracingdata)

	sampleData := []string{"2", "7", "0.000000", "0.778602"}
	tracingData.Parse(sampleData)

	if tracingData.ClassInstatiation(2) {
		t.Errorf("Should return false, since LineNo != 0.")
	}
}

func TestSetNodesToDelete(t *testing.T) {
	tracingData := *new(raw.Tracingdata)

	sampleData := []string{"2", "7", "0.000000", "0.778602"}
	tracingData.Parse(sampleData)
	sampleData = []string{"3", "0", "0.000000", "0.778602"}
	tracingData.Parse(sampleData)

	if (len(tracingData.SetNodesToDelete())) == 0 {
		t.Errorf("Second tracing data element should be added to the array. Array size should be equal to 1, not 0.")
	}

}

func TestRemoveIndex(t *testing.T) {
	tracingData := *new(raw.Tracingdata)

	sampleData := []string{"2", "7", "0.000000", "1.778602"}
	tracingData.Parse(sampleData)
	sampleData = []string{"2", "2", "0.000000", "0.778602"}
	tracingData.Parse(sampleData)
	sampleData = []string{"4", "8", "0.000000", "2.778602"}
	tracingData.Parse(sampleData)
	tracingData.RemoveIndex(0)

	intActualTime := test_helper.GetTimeFloat("0.000000")
	intStartTime := test_helper.GetTimeFloat("0.778602")
	currentActualTime := test_helper.GetTimeDuration(intActualTime)
	currentStartTime := test_helper.GetTimeDuration(intStartTime)

	test_helper.CheckStringValues("2", strconv.Itoa(tracingData[0].GetModuleId()), "ModuleId", t)
	test_helper.CheckStringValues("2", strconv.Itoa(tracingData[0].GetLineNo()), "LineNo", t)
	test_helper.CheckTimeDurationValues(currentActualTime, tracingData[0].GetActualTime(), "ActualTime", t)
	test_helper.CheckTimeDurationValues(currentStartTime, tracingData[0].GetStartTime(), "StartTime", t)

}
