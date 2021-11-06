package raw

import (
	"strconv"
	"strings"

	"gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/profiler/util"
)

type Moduledata []Moduleline

type Moduleline struct {
	moduleId         int
	moduleName       string
	debugListingFile string
	crcVal           int
}

func (md *Moduledata) Parse(strList []string) {
	var err error
	var line Moduleline

	line.moduleId, err = strconv.Atoi(strList[0])
	util.ValidateErrorStatus(err)

	line.moduleName = strings.Trim(strList[1], "\"")
	line.debugListingFile = strList[2]
	line.crcVal, err = strconv.Atoi(strList[3])
	util.ValidateErrorStatus(err)

	*md = append(*md, line)
}

func (md *Moduledata) AddMeToSet(set *Full) {
	set.module = *md
}

func (md Moduledata) GetModuleById(moduleId int) Moduleline {
	module := Moduleline{}

	for i := range md {
		if md[i].moduleId == moduleId {
			module = md[i]
			break
		}
	}

	return module
}

func (md Moduledata) GetModuleNameById(moduleId int) string {
	moduleName := ""

	for i := range md {
		if md[i].moduleId == moduleId {
			moduleName = md[i].moduleName
			break
		}
	}

	return moduleName
}

func (ml Moduleline) GetModuleId() int {
	return ml.moduleId
}

func (ml Moduleline) GetModuleName() string {
	return ml.moduleName
}

func (ml Moduleline) GetDebugListingFile() string {
	return ml.debugListingFile
}

func (ml Moduleline) GetCrcVal() int {
	return ml.crcVal
}
