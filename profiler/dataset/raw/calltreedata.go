package raw

import (
	"strconv"

	"gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/profiler/util"
)

type Calltreedata []Calltreeline

type Calltreeline struct {
	callerId     int
	callerLineNo int
	calleeId     int
	callCount    int
}

func (ctd *Calltreedata) Parse(strList []string) {
	var err error
	var line Calltreeline

	line.callerId, err = strconv.Atoi(strList[0])
	util.ValidateErrorStatus(err)

	line.callerLineNo, err = strconv.Atoi(strList[1])
	util.ValidateErrorStatus(err)

	line.calleeId, err = strconv.Atoi(strList[2])
	util.ValidateErrorStatus(err)

	line.callCount, err = strconv.Atoi(strList[3])
	util.ValidateErrorStatus(err)

	*ctd = append(*ctd, line)
}

func (ctd *Calltreedata) AddMeToSet(set *Full) {
	set.callTree = *ctd
}

func (ctl Calltreeline) GetCallerId() int {
	return ctl.callerId
}

func (ctl Calltreeline) GetCallerLineNo() int {
	return ctl.callerLineNo
}

func (ctl Calltreeline) GetCalleeId() int {
	return ctl.calleeId
}

func (ctl Calltreeline) GetCallCount() int {
	return ctl.callCount
}
