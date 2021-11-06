package processor

import (
	"context"
	"fmt"
)

func (tp *TraceProcessor) reportToZipkin(sortedData Treedata) {
	if len(sortedData) == 0 {
		fmt.Println("WARNING: Nothing to report to Zipkin")
	} else {
		ctx := context.Background()

		tp.reportToZipkinByLevels(0, sortedData, ctx)
	}
}

func (tp *TraceProcessor) reportToZipkinByLevels(startIndex int, sortedData Treedata, ctx context.Context) int {
	for i := startIndex; i < len(sortedData); i++ {
		if sortedData[i].trace.GetLineNo() == 0 {
			startedSpanIndex := i

			span, ctx := tp.tracer.GetSpanAndContext(sortedData[i].moduleName, ctx, sortedData[i].trace.GetStartTime())
			i = tp.reportToZipkinByLevels(i+1, sortedData, ctx)

			span.FinishedWithDuration((sortedData[i].trace.GetStartTime() - 
				sortedData[startedSpanIndex].trace.GetStartTime() + 
 				sortedData[i].trace.GetActualTime()))
		} else {
			return i
		}
	}

	return len(sortedData) - 1
}
