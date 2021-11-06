package raw_test

import (
	"strconv"
	"testing"
	"time"

	"gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/profiler/dataset/raw"
	"gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/profiler/util"
	test_helper "gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/tests"
)

func TestParse_Description(t *testing.T) {
	sampleData := []string{"1", "03/31/2021", "\"PROFILER\"", "12:27:47", "\"pjustas\""}

	testTime, err := time.Parse("01/02/2006T15:04:05", "03/31/2021"+"T"+"12:27:47") // build the time.Time type
	util.ValidateErrorStatus(err)

	descriptionData := new(raw.Descriptiondata)
	descriptionData.Parse(sampleData)

	test_helper.CheckStringValues("1", strconv.Itoa(descriptionData.GetVersion()), "Version", t)
	test_helper.CheckTimeTimeValues(testTime, descriptionData.GetSystemDateTime(), "SystemDateTime", t)
	test_helper.CheckStringValues("\"PROFILER\"", descriptionData.GetDescription(), "Description", t)
	test_helper.CheckStringValues("\"pjustas\"", descriptionData.GetUserId(), "UserId", t)
}
