package util

const (
	EMERGENCY = iota
	ALERT
	CRITICAL
	ERROR
	WARNING
	NOTICE
	INFO
	DEBUG
	TRACE
)

var LOG_LEVELS = map[string]int{
	"EMERGENCY": EMERGENCY,
	"ALERT":     ALERT,
	"CRITICAL":  CRITICAL,
	"ERROR":     ERROR,
	"WARNING":   WARNING,
	"NOTICE":    NOTICE,
	"INFO":      INFO,
	"DEBUG":     DEBUG,
	"TRACE":     TRACE,
}

// Placeholder functions for the exported Perl functions
// You'll need to implement these yourself
func CmpDeeply()        {}
func DetectionAliases() {}
func DetectionMethods() {}
func DumpOneLine()      {}
func LogLevelAliases()  {}
