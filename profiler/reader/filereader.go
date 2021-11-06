package reader

import (
	"io/ioutil"

	"gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/profiler/util"
)

type FileReader struct {
	dat []byte
}

func (fr *FileReader) ReadFile(filePath string) {
	var err error

	fr.dat, err = ioutil.ReadFile(filePath)
	util.ValidateErrorStatus(err)
}

func (fr *FileReader) GetFile() []byte {
	return fr.dat
}
