package reader

import "gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/profiler/dataset/raw"

type RawLineParser interface {
	Parse(strList []string)
	AddMeToSet(set *raw.Full)
}
