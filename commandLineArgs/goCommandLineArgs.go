package commandLineArgs

import "os"

func PlayWithCommandLineArgs() {
	var someString, separator string
	separator = "\n"
	var args = os.Args
	for i := 0; i < len(args); i++ {
		someString = someString + separator + args[i]
	}
}
