package reader_test

import (
	"bufio"
	"strconv"
	"strings"
	"testing"
	"time"

	"gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/profiler/reader"
	"gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/profiler/util"
	test_helper "gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/tests"
)

func TestValidateFile(t *testing.T) {
	util.ValidateFile("../testData/testDescription.txt")
}

func TestReaderHelper_ReadProfilerFile(t *testing.T) {
	fileReader := reader.ReadProfilerFile("../testData/testDescription.txt")
	expectedData := "1 03/31/2021 \"PROFILER\" 12:27:47 \"pjustas\""

	scanner := bufio.NewScanner(strings.NewReader(string(fileReader.GetFile())))
	for scanner.Scan() {
		s := scanner.Text()
		if s != expectedData {
			t.Errorf("Expected: " + expectedData + " got: " + s)
		}
	}
}

func TestTransformProfilerData(t *testing.T) { // TODO: Do we need to check transformation for all dataset types?
	fileReader := reader.ReadProfilerFile("../testData/testFullTrace.txt")
	profDataReader := reader.TransformProfilerData(fileReader)
	gotDescription := profDataReader.GetRaw().GetDescription()

	testTime, err := time.Parse("01/02/2006T15:04:05", "03/31/2021"+"T"+"12:27:47") // build the time.Time type
	util.ValidateErrorStatus(err)

	test_helper.CheckStringValues("1", strconv.Itoa(gotDescription.GetVersion()), "Version", t)
	test_helper.CheckTimeTimeValues(testTime, gotDescription.GetSystemDateTime(), "SystemDateTime", t)
	test_helper.CheckStringValues("\"PROFILER\"", gotDescription.GetDescription(), "Description", t)
	test_helper.CheckStringValues("\"pjustas\"", gotDescription.GetUserId(), "UserId", t)
}
