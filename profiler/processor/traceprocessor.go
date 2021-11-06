package processor

import (
	"fmt"

	"gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/config"
	"gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/profiler/dataset/raw"
	"gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/profiler/util"
	"gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/tracingsystem"
)

type TraceProcessor struct {
	rawData       *raw.Full
	sortedTracing raw.Tracingdata
	callTree      raw.Calltreedata
	module        raw.Moduledata
	treeData      Treedata
	tracer        tracingsystem.ZipkinProperties
}

func (tp *TraceProcessor) ParseTrace() {
	tp.GetTracingData()
	tp.SetCallTreeAndModule()

	fmt.Println("Start 1: calculating call tree")

	tp.GenerateCallTreeFromTrace()
	tp.SortTreeData()
}

func (tp *TraceProcessor) GenerateCallTreeFromTrace() {
	for i := 0; i < len(tp.sortedTracing); i++ {
		tp.startOrEndTreeIfNeeded(i)

		if tp.sortedTracing[i].GetLineNo() == 0 && tp.sortedTracing[i-1].GetLineNo() != 0 {
			tp.pushStartNode(i)
			tp.pushEndNode(i-1)
		}
	}
}

func (tp *TraceProcessor) startOrEndTreeIfNeeded(i int) {
	lastElementIndex := len(tp.sortedTracing) - 1

	if i == 0 {
		tp.treeData.StartTree(tp.module.GetModuleNameById(tp.sortedTracing[i].GetModuleId()), tp.sortedTracing[i])
	}
	if i == lastElementIndex && tp.sortedTracing[i].GetLineNo() != 0 && tp.sortedTracing[i].GetModuleId() == tp.treeData[0].trace.GetModuleId(){
		tp.pushEndNode(i)
	}
}

func (tp *TraceProcessor) pushStartNode(i int) {
	tp.treeData.Push(tp.module.GetModuleNameById(tp.sortedTracing[i].GetModuleId()), tp.sortedTracing[i])
}

func (tp *TraceProcessor) pushEndNode(i int) {
	currElementStartTime := tp.sortedTracing[i].GetStartTime()

	// in some cases we need to fix element start time 
	for j := 0; j < len(tp.treeData); j++ {
		if tp.treeData[j].trace.GetStartTime() > currElementStartTime {
			currElementStartTime = tp.treeData[j].trace.GetStartTime()
		}
	}
	tp.sortedTracing[i].SetStartTime(currElementStartTime)

	tp.treeData.Push(tp.module.GetModuleNameById(tp.sortedTracing[i].GetModuleId()), tp.sortedTracing[i])
}

func (tp *TraceProcessor) ReportDataToZipkin() {
	fmt.Println("Start 2: sending data to zipkin")

	tp.tracer.GetNewTracer(config.Cfg.Zipkin.ServiceName)
	defer func() {
		err := tp.tracer.GetReporter().Close()
		util.ValidateErrorStatus(err)
	}()

	tp.reportToZipkin(tp.treeData)
}

func (tp *TraceProcessor) GetTracingData() {
	tp.sortedTracing = tp.rawData.GetTracing()
	tp.sortedTracing.SetNodeEndsForClassCreations()
}

func (tp *TraceProcessor) SortTreeData() {
	tp.treeData.SortTree()
}

func (tp *TraceProcessor) SetCallTreeAndModule() {
	tp.setCallTree()
	tp.setModule()
}

func (tp *TraceProcessor) setCallTree() {
	tp.callTree = tp.rawData.GetCallTree()
}

func (tp *TraceProcessor) setModule() {
	tp.module = tp.rawData.GetModule()
}

func (tp *TraceProcessor) GetTreeData() Treedata {
	return tp.treeData
}

func (tp *TraceProcessor) SetRawData(parsedRawData *raw.Full) {
	tp.rawData = parsedRawData
}
