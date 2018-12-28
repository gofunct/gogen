
package logging

import (
	"github.com/dixonwille/wlog"
	kitlog "github.com/go-kit/kit/log"
	"log"
	"os"
)

type Logger struct{
	UI 		wlog.UI
	KitLog kitlog.Logger
}

func NewLogger() *Logger {
	logger := kitlog.NewJSONLogger(kitlog.NewSyncWriter(os.Stdout))
	l := &Logger{
		UI: &wlog.PrefixUI{
			LogPrefix:     ":speech_balloon:",
			OutputPrefix:  ":boom:",
			SuccessPrefix: ":white_check_mark:",
			InfoPrefix:    ":wave:",
			ErrorPrefix:   ":x:",
			WarnPrefix:    ":warning:",
			RunningPrefix: ":zap:",
			AskPrefix:     ":interrobang:",
			UI:            wlog.New(os.Stdin, os.Stdout, os.Stderr),
		},
		KitLog: logger,
	}

	log.SetOutput(kitlog.NewStdlibAdapter(l.KitLog))
	return l
}
