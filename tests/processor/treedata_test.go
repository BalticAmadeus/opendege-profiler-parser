package processor_test

import (
	"strconv"
	"testing"

	"gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/profiler/dataset/raw"
	"gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/profiler/processor"
	test_helper "gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/tests"
)

func TestStartTree(t *testing.T) {
	tracingData := *new(raw.Tracingdata)

	sampleData := []string{"2", "7", "0.000000", "0.778602"}
	tracingData.Parse(sampleData)

	traceProcessor := new(processor.TraceProcessor)
	treeData := traceProcessor.GetTreeData()
	treeData.StartTree("oop/testOop.p", tracingData[0])

	intActualTime := test_helper.GetTimeFloat("0.000000")
	intStartTime := test_helper.GetTimeFloat("0.778602")
	currentActualTime := test_helper.GetTimeDuration(intActualTime)
	currentStartTime := test_helper.GetTimeDuration(intStartTime) - 1

	test_helper.CheckStringValues("oop/testOop.p", treeData[0].GetModuleName(), "ModuleName", t)
	test_helper.CheckStringValues("2", strconv.Itoa(treeData[0].GetTrace().GetModuleId()), "ModuleId", t)
	test_helper.CheckStringValues("0", strconv.Itoa(treeData[0].GetTrace().GetLineNo()), "LineNo", t)
	test_helper.CheckTimeDurationValues(currentActualTime, treeData[0].GetTrace().GetActualTime(), "ActualTime", t)
	test_helper.CheckTimeDurationValues(currentStartTime, treeData[0].GetTrace().GetStartTime(), "StartTime", t)
}

func TestPush_TreeData(t *testing.T) {
	tracingData := *new(raw.Tracingdata)

	sampleData := []string{"2", "7", "0.000000", "0.778602"}
	tracingData.Parse(sampleData)
	sampleData = []string{"2", "11", "0.000002", "0.779219"}
	tracingData.Parse(sampleData)

	traceProcessor := new(processor.TraceProcessor)
	treeData := traceProcessor.GetTreeData()
	treeData.StartTree("oop/testOop.p", tracingData[0])
	treeData.Push("oop/testOop.p", tracingData[1])

	intActualTime := test_helper.GetTimeFloat("0.000002")
	intStartTime := test_helper.GetTimeFloat("0.779219")
	currentActualTime := test_helper.GetTimeDuration(intActualTime)
	currentStartTime := test_helper.GetTimeDuration(intStartTime)

	test_helper.CheckStringValues("oop/testOop.p", treeData[1].GetModuleName(), "ModuleName", t)
	test_helper.CheckStringValues("2", strconv.Itoa(treeData[1].GetTrace().GetModuleId()), "ModuleId", t)
	test_helper.CheckStringValues("11", strconv.Itoa(treeData[1].GetTrace().GetLineNo()), "LineNo", t)
	test_helper.CheckTimeDurationValues(currentActualTime, treeData[1].GetTrace().GetActualTime(), "ActualTime", t)
	test_helper.CheckTimeDurationValues(currentStartTime, treeData[1].GetTrace().GetStartTime(), "StartTime", t)

}
