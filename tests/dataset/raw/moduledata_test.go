package raw_test

import (
	"strconv"
	"testing"

	"gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/profiler/dataset/raw"
	"gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/profiler/util"
	test_helper "gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/tests"
)

func TestParse_Module(t *testing.T) {
	moduleData := new(raw.Moduledata)

	sampleData := []string{"6", "\"oop.TestOop2\"", "\"\"", "27477", "0", "\"\""}
	moduleData.Parse(sampleData)

	moduleId, err := strconv.Atoi("6")
	util.ValidateErrorStatus(err)

	test_helper.CheckStringValues("6", strconv.Itoa(moduleData.GetModuleById(moduleId).GetModuleId()), "ModuleId", t)
	test_helper.CheckStringValues("oop.TestOop2", moduleData.GetModuleById(moduleId).GetModuleName(), "ModuleName", t)
	test_helper.CheckStringValues("\"\"", moduleData.GetModuleById(moduleId).GetDebugListingFile(), "DebugListingFile", t)
	test_helper.CheckStringValues("27477", strconv.Itoa(moduleData.GetModuleById(moduleId).GetCrcVal()), "CrcVal", t)
}
