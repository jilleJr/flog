package loglevel

import "strings"

type Level int

const (
	Undefined Level = 0
	Unknown   Level = 1 << (iota - 1)
	Trace
	Debug
	Information
	Warning
	Error
	Critical
	Fatal
	Panic
)

var singularLevels = []Level{
	Unknown,
	Trace,
	Debug,
	Information,
	Warning,
	Error,
	Critical,
	Fatal,
	Panic,
}

func (lvl Level) String() string {
	var b = strings.Builder{}
	for _, singularLevel := range singularLevels {
		if lvl&singularLevel != Undefined {
			if b.Len() > 0 {
				b.WriteRune('|')
			}
			b.WriteString(singularLevelString(singularLevel))
		}
	}
	if b.Len() == 0 {
		return "Undefined"
	}
	return b.String()
}

func singularLevelString(lvl Level) string {
	switch lvl {
	case Unknown:
		return "Unknown"
	case Trace:
		return "Trace"
	case Debug:
		return "Debug"
	case Information:
		return "Information"
	case Warning:
		return "Warning"
	case Error:
		return "Error"
	case Critical:
		return "Critical"
	case Fatal:
		return "Fatal"
	case Panic:
		return "Panic"
	}
	return "Undefined"
}

func ParseLevel(s string) Level {
	switch strings.ToLower(s) {
	case "trc", "trce", "trac", "trace":
		return Trace

	case "dbg", "debu", "dbug", "debg", "debug":
		return Debug

	case "inf", "info", "information":
		return Information

	case "warn", "warning":
		return Warning

	case "err", "erro", "error", "fail":
		return Error

	case "crit", "critical":
		return Critical

	case "fata", "fatal":
		return Fatal

	case "panic":
		return Panic
	}

	return Unknown
}
