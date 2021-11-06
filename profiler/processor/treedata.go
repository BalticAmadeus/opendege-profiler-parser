package processor

import (
	"time"
	"sort"

	"gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/profiler/dataset/raw"
)

type Treedata []Tree

type Tree struct {
	moduleName string
	trace      raw.Tracingline
}

func (td *Treedata) StartTree(moduleName string, sortedTracing raw.Tracingline) {
	tracingLine := new(raw.Tracingline)
	tracingLine.SetModuleId(sortedTracing.GetModuleId())
	tracingLine.SetLineNo(0)
	tracingLine.SetActualTime(0)
	tracingLine.SetStartTime(time.Duration(sortedTracing.GetStartTime().Nanoseconds()) - 1)

	*td = append(*td, Tree{moduleName, *tracingLine})
}

func (td *Treedata) Push(moduleName string, sortedTracing raw.Tracingline) {
	*td = append(*td, Tree{moduleName, sortedTracing})
}

// Sort by startTime which is in nanoseconds, which is too abstract in some cases so if action happened in the same nanoseconds - 
// sort by lineNo for the same modules, and by the order of trace lines for different modules
func (td Treedata) SortTree() {	
	sort.SliceStable(td, func(i, j int) bool {
		if td[i].trace.GetStartTime().Nanoseconds() < td[j].trace.GetStartTime().Nanoseconds() {
			return true
		}
		if td[i].trace.GetStartTime().Nanoseconds() > td[j].trace.GetStartTime().Nanoseconds() {
			return false
		}
		
		if td[i].trace.GetModuleId() == td[j].trace.GetModuleId() {
			return td[i].trace.GetLineNo() < td[j].trace.GetLineNo()
		}

		return td[i].trace.GetLineNo() == 0
	})
}

func (t Tree) GetModuleName() string {
	return t.moduleName
}

func (t Tree) GetTrace() raw.Tracingline {
	return t.trace
}
