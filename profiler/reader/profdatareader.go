package reader

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"

	"gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/profiler/dataset/raw"
	"gitlab.baltic-amadeus.lt/progress/progress-opentracing-profiler/profiler/util"
)

const (
	dataSeparator = "."
	supportedSections = 5
)

type ProfDataReader struct {
	rawData *raw.Full
}

func (pdr *ProfDataReader) ParseProf(data []byte) {
	pdr.rawData = new(raw.Full)
	separatorCounter := 1
	scanner := bufio.NewScanner(strings.NewReader(string(data)))

	lineParser, err := pdr.GetRawLineParser(separatorCounter)
	util.ValidateErrorStatus(err)

	for scanner.Scan() {
		if scanner.Text() == dataSeparator {
			lineParser.AddMeToSet(pdr.rawData)

			separatorCounter++
			if separatorCounter > supportedSections {
				return
			}

			lineParser, err = pdr.GetRawLineParser(separatorCounter)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			s := scanner.Text()
			r := regexp.MustCompile(`[^\s"']+|"([^"]*)"`)

			arr := r.FindAllString(s, -1)
			if arr == nil {
				panic("No match has been found in string and defined pattern")
			}
			lineParser.Parse(arr)
		}
	}
}

func (pdr *ProfDataReader) GetRaw() *raw.Full {
	return pdr.rawData
}

func (pdr *ProfDataReader) GetRawLineParser(separatorCounter int) (RawLineParser, error) {
	switch separatorCounter {
	case 1:
		return new(raw.Descriptiondata), nil
	case 2:
		return new(raw.Moduledata), nil
	case 3:
		return new(raw.Calltreedata), nil
	case 4:
		return new(raw.Linesummarydata), nil
	case 5:
		return new(raw.Tracingdata), nil
	default:
		return nil, fmt.Errorf("parsing of section '%d' not supported", separatorCounter)
	}
}
