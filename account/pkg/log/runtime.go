package log

import (
	"runtime"
	"strings"
)

type Caller struct {
	FunctionName string
	FilePath     string
	Line         int
}

// GetFunctionNameAtRuntime returns the name of the function that called it.
// It is used for logging purposes.
func GetFunctionNameAtRuntime(preStep int) Caller {
	functionName, filePath, line, _ := runtime.Caller(preStep)
	ls := strings.Split(getLastPathString(runtime.FuncForPC(functionName).Name()), "/")
	name := ls[len(ls)-1]
	lastIndex := strings.Index(name, ".")
	return Caller{
		FunctionName: name[lastIndex+1:],
		FilePath:     getLastPathString(filePath),
		Line:         line,
	}
}

// get string from last two / to end
// ex: root/data/file.mp3 -> data/file.mp3
func getLastPathString(a string) string {
	ls := strings.Split(a, "/")
	if len(ls) < 3 {
		return strings.Join(ls, "/")
	}

	return strings.Join(ls[len(ls)-3:], "/")
}
