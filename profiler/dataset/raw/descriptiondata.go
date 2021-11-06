package raw

import (
	"strconv"
	"time"

	"gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/profiler/util"
)

type Descriptiondata struct {
	version        int
	systemDateTime time.Time
	description    string
	userId         string
}

func (dd *Descriptiondata) Parse(strList []string) {
	var err error

	dd.version, err = strconv.Atoi(strList[0])
	util.ValidateErrorStatus(err)

	dd.systemDateTime, err = time.Parse("01/02/2006T15:04:05", strList[1]+"T"+strList[3])
	util.ValidateErrorStatus(err)

	dd.description = strList[2]
	dd.userId = strList[4]
}

func (dd *Descriptiondata) AddMeToSet(set *Full) {
	set.description = *dd
}

func (dd Descriptiondata) GetVersion() int {
	return dd.version
}

func (dd Descriptiondata) GetSystemDateTime() time.Time {
	return dd.systemDateTime
}

func (dd Descriptiondata) GetDescription() string {
	return dd.description
}

func (dd Descriptiondata) GetUserId() string {
	return dd.userId
}
