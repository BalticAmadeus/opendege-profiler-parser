package raw

import (
	"strconv"
	"time"

	"gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/profiler/util"
)

type Tracingdata []Tracingline

type Tracingline struct {
	moduleId   int
	lineNo     int
	actualTime time.Duration
	startTime  time.Duration
}

func (td *Tracingdata) Parse(strList []string) {
	var line Tracingline

	intActualTime, err := strconv.ParseFloat(strList[2], 64)
	util.ValidateErrorStatus(err)

	intStartTime, err := strconv.ParseFloat(strList[3], 64)
	util.ValidateErrorStatus(err)

	line.moduleId, err = strconv.Atoi(strList[0])
	util.ValidateErrorStatus(err)

	line.lineNo, err = strconv.Atoi(strList[1])
	util.ValidateErrorStatus(err)

	line.actualTime = time.Duration(intActualTime * float64(time.Second))
	line.startTime = time.Duration(intStartTime * float64(time.Second))

	*td = append(*td, line)
}

func (td *Tracingdata) AddMeToSet(set *Full) {
	set.tracing = *td
}

func (td Tracingdata) GetTracing(callerId int, lineNo int) Tracingline {
	tracing := Tracingline{}

	for i := range td {
		if td[i].moduleId == callerId && td[i].lineNo == lineNo {
			tracing = td[i]
			break
		}
	}

	return tracing
}

func (td Tracingdata) SetNodeEndsForClassCreations() {
	nodesToDelete := td.SetNodesToDelete()
	td.deleteEmptyNodes(nodesToDelete)
}

func (td Tracingdata) SetNodesToDelete() []int {
	var nodesToDelete []int

	for i := 0; i < len(td); i++ {
		if td[i].GetLineNo() == 0 {
			if td.ClassInstatiation(td[i].GetModuleId()) {
				nodesToDelete = append(nodesToDelete, i)
			}
		}
	}

	return nodesToDelete
}

func (td Tracingdata) deleteEmptyNodes(nodesToDelete []int) {
	for i := len(nodesToDelete) - 1; i >= 0; i-- {
		td.RemoveIndex(nodesToDelete[i])
	}
}

func (td Tracingdata) RemoveIndex(index int) []Tracingline {
	return append(td[:index], td[index+1:]...)
}

func (td Tracingdata) ClassInstatiation(moduleId int) bool {
	for i := 0; i < len(td); i++ {
		if td[i].GetModuleId() == moduleId && td[i].GetLineNo() != 0 {
			return false
		}
	}

	return true
}

func (tl Tracingline) GetModuleId() int {
	return tl.moduleId
}

func (tl *Tracingline) SetModuleId(moduleId int) {
	tl.moduleId = moduleId
}

func (tl Tracingline) GetLineNo() int {
	return tl.lineNo
}

func (tl *Tracingline) SetLineNo(lineNo int) {
	tl.lineNo = lineNo
}

func (tl Tracingline) GetActualTime() time.Duration {
	return tl.actualTime
}

func (tl *Tracingline) SetActualTime(actualTime time.Duration) {
	tl.actualTime = actualTime
}

func (tl Tracingline) GetStartTime() time.Duration {
	return tl.startTime
}

func (tl *Tracingline) SetStartTime(startTime time.Duration) {
	tl.startTime = startTime
}
