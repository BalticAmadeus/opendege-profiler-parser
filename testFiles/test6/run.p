PROFILER:ENABLED = TRUE.
PROFILER:DIRECTORY = "C:\work\progress-opentracing-profiler\profilerFiles".
PROFILER:FILE-NAME = "C:\work\progress-opentracing-profiler\profilerFiles\test6.prof".
PROFILER:LISTINGS = FALSE.
PROFILER:DESCRIPTION = "PROFILER".
PROFILER:PROFILING = TRUE.
PROFILER:TRACE-FILTER = "*".

run main.p.

def var testClass as class Test1 no-undo.

testClass = new Test1().
testClass:test().

PROFILER:ENABLED = FALSE.
PROFILER:PROFILING = FALSE.
PROFILER:WRITE-DATA().
