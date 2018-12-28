
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
	WarnPrefix:     emoji.Sprint(":grimacing:"),
	RunningPrefix:  emoji.Sprint(":fire:"),
	AskPrefix:      emoji.Sprint(":question:"),
	UI:            ui,
}


type Logger struct{
	UI 		wlog.UI
	KitLog kitlog.Logger
}

func NewLogger() *Logger {
	logger := kitlog.NewJSONLogger(kitlog.NewSyncWriter(os.Stdout))
	l := &Logger{
		UI: prefix,
		KitLog: logger,
	}

	log.SetOutput(kitlog.NewStdlibAdapter(l.KitLog))
	return l
}


func (l *Logger) AddColor() {
		l.UI = wlog.AddColor(wlog.Green, wlog.Red, wlog.BrightBlue, wlog.Blue, wlog.Yellow, wlog.BrightMagenta, wlog.Yellow, wlog.BrightGreen, wlog.BrightRed, l.UI)
}
