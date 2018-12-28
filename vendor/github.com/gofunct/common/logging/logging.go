
package logging

import (
	"github.com/dixonwille/wlog"
	kitlog "github.com/go-kit/kit/log"
	"github.com/kyokomi/emoji"
	"github.com/mattn/go-isatty"
	"log"
	"os"
)


var (
	noColor = os.Getenv("TERM") == "dumb" ||
		(!isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdout.Fd()))
)

var ui = wlog.New(os.Stdin, os.Stdout, os.Stderr)

var prefix = &wlog.PrefixUI{
	LogPrefix:     emoji.Sprint(":speech_balloon:"),
	OutputPrefix:   emoji.Sprint(":boom:"),
	SuccessPrefix:  emoji.Sprint(":white_check_mark:"),
	InfoPrefix:     emoji.Sprint(":wave:"),
	ErrorPrefix:    emoji.Sprint(":x:"),
	WarnPrefix:     emoji.Sprint(":warning:"),
	RunningPrefix:  emoji.Sprint(":zap:"),
	AskPrefix:      emoji.Sprint(":interrobang:"),
	UI:            ui,
}


type Logger struct{
	UI 		wlog.UI
	KitLog kitlog.Logger
}

func NewLogger() *Logger {
	logger := kitlog.NewJSONLogger(kitlog.NewSyncWriter(os.Stdout))
	l := &Logger{
		UI: ui,
		KitLog: logger,
	}

	log.SetOutput(kitlog.NewStdlibAdapter(l.KitLog))
	return l
}


func (l *Logger) AddColor(optionColor, questionColor, responseColor, errorColor wlog.Color) {
	if !noColor {
		l.UI = wlog.AddColor(questionColor, errorColor, wlog.None, wlog.None, optionColor, responseColor, wlog.None, wlog.None, wlog.None, l.UI)
	}
}
