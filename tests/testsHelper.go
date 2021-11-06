package test_helper

import (
	"strconv"
	"testing"
	"time"

	"gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/profiler/util"
)

func CheckStringValues(expectedValue string, gotValue string, fieldName string, t *testing.T) {
	if expectedValue != gotValue {
		t.Errorf("Expected "+fieldName+" = %v got = %v", expectedValue, gotValue)
	}
}

func CheckTimeTimeValues(expectedValue time.Time, gotValue time.Time, fieldName string, t *testing.T) {
	if expectedValue != gotValue {
		t.Errorf("Expected "+fieldName+" = %v got = %v", expectedValue, gotValue)
	}
}

func CheckTimeDurationValues(expectedValue time.Duration, gotValue time.Duration, fieldName string, t *testing.T) {
	if expectedValue != gotValue {
		t.Errorf("Expected "+fieldName+" = %v got = %v", expectedValue, gotValue)
	}
}

func GetTimeFloat(stringTime string) float64 {
	floatTime, err := strconv.ParseFloat(stringTime, 64)
	util.ValidateErrorStatus(err)

	return floatTime
}

func GetTimeDuration(floatTime float64) time.Duration {
	return time.Duration(floatTime * float64(time.Second))
}
