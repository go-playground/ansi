package ansi

// EscSeq is a predefined ANSI escape sequence
type EscSeq string

// ANSI escape sequences
// NOTE: in a standard xterm terminal the light colors will appear BOLD instead of the light variant
const (
	Reset        EscSeq = "\x1b[0m"
	Underline           = "\x1b[4m"
	Blink               = "\x1b[5m"
	Inverse             = "\x1b[7m"
	UnderlineOff        = "\x1b[24m"
	BlinkOff            = "\x1b[25m"
	Black               = "\x1b[30m"
	DarkGray            = "\x1b[30;1m"
	Red                 = "\x1b[31m"
	LightRed            = "\x1b[31;1m"
	Green               = "\x1b[32m"
	LightGreen          = "\x1b[32;1m"
	Brown               = "\x1b[33m"
	Yellow              = "\x1b[33;1m"
	Blue                = "\x1b[34m"
	LightBlue           = "\x1b[34;1m"
	Magenta             = "\x1b[35m"
	LightMagenta        = "\x1b[35;1m"
	Cyan                = "\x1b[36m"
	LightCyan           = "\x1b[36;1m"
	LightGray           = "\x1b[37m"
	White               = "\x1b[37;1m"
	DefaultFG           = "\x1b[39m"
	DefaultBG           = "\x1b[49m"
)
