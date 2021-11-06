package raw

import (
	"strconv"
	"time"

	"gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/profiler/util"
)

type Linesummarydata []Linesummaryline

type Linesummaryline struct {
	moduleId       int
	lineNo         int
	execCount      int
	actualTime     time.Duration
	cumulativeTime time.Duration
}

func (lsd *Linesummarydata) Parse(strList []string) {
	var line Linesummaryline

	intActualTime, err := strconv.ParseFloat(strList[3], 64)
	util.ValidateErrorStatus(err)

	intCumulativeTime, err := strconv.ParseFloat(strList[4], 64)
	util.ValidateErrorStatus(err)

	line.moduleId, err = strconv.Atoi(strList[0])
	util.ValidateErrorStatus(err)

	line.lineNo, err = strconv.Atoi(strList[1])
	util.ValidateErrorStatus(err)

	line.execCount, err = strconv.Atoi(strList[2])
	util.ValidateErrorStatus(err)

	line.actualTime = time.Duration(intActualTime * float64(time.Second))
	line.cumulativeTime = time.Duration(intCumulativeTime * float64(time.Second))

	*lsd = append(*lsd, line)
}

func (lsd *Linesummarydata) AddMeToSet(set *Full) {
	set.lineSummary = *lsd
}

func (lsl Linesummaryline) GetModuleId() int {
	return lsl.moduleId
}

func (lsl Linesummaryline) GetLineNo() int {
	return lsl.lineNo
}

func (lsl Linesummaryline) GetExecCount() int {
	return lsl.execCount
}

func (lsl Linesummaryline) GetActualTime() time.Duration {
	return lsl.actualTime
}

func (lsl Linesummaryline) GetCumulativeTime() time.Duration {
	return lsl.cumulativeTime
}
