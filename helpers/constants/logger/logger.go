package logger

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

func LogError(err error) {
	_, file, line, _ := runtime.Caller(1)

	paths := strings.Split(file, "/")
	file = paths[len(paths)-1]

	errResp := fmt.Sprintf("[%s:%d] %s", file, line, err)
	log.Println(errResp)
}

func LogString(err string) {
	_, file, line, _ := runtime.Caller(1)

	paths := strings.Split(file, "/")
	file = paths[len(paths)-1]

	errResp := fmt.Sprintf("[%s:%d] %s", file, line, err)
	log.Println(errResp)
}

func LogDebugger(data interface{}) {
	_, file, line, _ := runtime.Caller(1)

	paths := strings.Split(file, "/")
	file = paths[len(paths)-1]

	errResp := fmt.Sprintf("[%s:%d]", file, line)
	log.Print(errResp)
	fmt.Println(data)
}
