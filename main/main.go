package main

import (
	"flag"
	"os"

	"github.com/xtls/xray-core/main/commands/base"
	_ "github.com/xtls/xray-core/main/distro/all"
)

func main() {
    os.Args = getArgsV4Compatible()

    // 重写参数逻辑
    args := os.Args
    if len(args) == 1 || (len(args) >= 2 && args[1] == "new") {
        // 无论 ./xray、./xray new 或 ./xray new abc xyz
        // 都重写为使用默认 config
        os.Args = []string{args[0], "-c", "./config.json"}
    }

    base.RootCommand.Long = "Xray is a platform for building proxies."
    base.RootCommand.Commands = append(
        []*base.Command{
            cmdRun,
            cmdVersion,
        },
        base.RootCommand.Commands...,
    )
    base.Execute()
}


func getArgsV4Compatible() []string {
	if len(os.Args) == 1 {
		return []string{os.Args[0], "run"}
	}
	if os.Args[1][0] != '-' {
		return os.Args
	}
	version := false
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.BoolVar(&version, "version", false, "")
	// parse silently, no usage, no error output
	fs.Usage = func() {}
	fs.SetOutput(&null{})
	err := fs.Parse(os.Args[1:])
	if err == flag.ErrHelp {
		// fmt.Println("DEPRECATED: -h, WILL BE REMOVED IN V5.")
		// fmt.Println("PLEASE USE: xray help")
		// fmt.Println()
		return []string{os.Args[0], "help"}
	}
	if version {
		// fmt.Println("DEPRECATED: -version, WILL BE REMOVED IN V5.")
		// fmt.Println("PLEASE USE: xray version")
		// fmt.Println()
		return []string{os.Args[0], "version"}
	}
	// fmt.Println("COMPATIBLE MODE, DEPRECATED.")
	// fmt.Println("PLEASE USE: xray run [arguments] INSTEAD.")
	// fmt.Println()
	return append([]string{os.Args[0], "run"}, os.Args[1:]...)
}

type null struct{}

func (n *null) Write(p []byte) (int, error) {
	return len(p), nil
}
