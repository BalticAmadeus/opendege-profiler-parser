package reader

func ReadProfilerFile(profilerFileToParse string) FileReader {
	fileReader := new(FileReader)
	fileReader.ReadFile(profilerFileToParse)

	return *fileReader
}

func TransformProfilerData(fileReader FileReader) ProfDataReader {
	profDataReader := new(ProfDataReader)
	profDataReader.ParseProf(fileReader.GetFile())

	return *profDataReader
}
