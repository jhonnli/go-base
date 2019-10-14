package initial

import (
	"fmt"
	"github.com/jhonnli/logs"
)

func InitLog() {
	logs.Async()
	logs.Async(1e4)
	logs.SetLogger(logs.AdapterConsole, fmt.Sprintf(`{"level":%d}`, logs.LevelError))
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
}
