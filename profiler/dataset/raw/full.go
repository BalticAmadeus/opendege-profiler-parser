package raw

type Full struct {
	description Descriptiondata
	module      Moduledata
	callTree    Calltreedata
	lineSummary Linesummarydata
	tracing     Tracingdata
}

func (f Full) GetDescription() Descriptiondata {
	return f.description
}

func (f Full) GetModule() Moduledata {
	return f.module
}

func (f Full) GetCallTree() Calltreedata {
	return f.callTree
}

func (f Full) GetLineSummary() Linesummarydata {
	return f.lineSummary
}

func (f Full) GetTracing() Tracingdata {
	return f.tracing
}
