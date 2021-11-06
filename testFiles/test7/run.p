PROFILER:ENABLED = TRUE.
PROFILER:DIRECTORY = "C:\work\progress-opentracing-profiler\profilerFiles".
PROFILER:FILE-NAME = "C:\work\progress-opentracing-profiler\profilerFiles\test7.prof".
PROFILER:LISTINGS = FALSE.
PROFILER:DESCRIPTION = "PROFILER".
PROFILER:PROFILING = TRUE.
PROFILER:TRACE-FILTER = "*".

def var mainClass as class Main no-undo.

mainClass = new Main().
mainClass:test5().

PROFILER:ENABLED = FALSE.
PROFILER:PROFILING = FALSE.
PROFILER:WRITE-DATA().
