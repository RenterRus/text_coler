package notes

import "fmt"

const (
	ForegStd           = "\033[39m"
	ForegBlack         = "\033[30m"
	ForegDarkRed       = "\033[31m"
	ForegDarkGreen     = "\033[32m"
	ForegDarkYellow    = "\033[33m"
	ForegDarkBlue      = "\033[34m"
	ForegDarkPurple    = "\033[35m"
	ForegDarkLightBlue = "\033[36m"
	ForegLightGrey     = "\033[37m"
	ForegDarkGrey      = "\033[90m"
	ForegRed           = "\033[91m"
	ForegGreen         = "\033[92m"
	ForegOrange        = "\033[93m"
	ForegBlue          = "\033[94m"
	ForegPurple        = "\033[95m"
	ForegLightBlue     = "\033[96m"
	ForegWhite         = "\033[97m"

	BackStd           = "\033[49m"
	BackBlack         = "\033[40m"
	BackDarkRed       = "\033[41m"
	BackDarkGreen     = "\033[42m"
	BackDarkYellow    = "\033[43m"
	BackDarkBlue      = "\033[44m"
	BackDarkPurple    = "\033[45m"
	BackDarkLightBlue = "\033[46m"
	BackLightGrey     = "\033[47m"
	BackDarkGrey      = "\033[100m"
	BackRed           = "\033[101m"
	BackGreen         = "\033[102m"
	BackOrange        = "\033[103m"
	BackBlue          = "\033[104m"
	BackPurple        = "\033[105m"
	BackLightBlue     = "\033[106m"
	BackWhite         = "\033[107m"

	Reset = "\033[0m"
)

func ColorPrintln(foreground, back, message string, a ...any) {
	fmt.Print(foreground, back)
	fmt.Printf(message, a...)
	fmt.Println(Reset)
}

func ColorPrint(foreground, back, message string, a ...any) {
	fmt.Print(foreground, back)
	fmt.Printf(message, a...)
	fmt.Print(Reset)
}
