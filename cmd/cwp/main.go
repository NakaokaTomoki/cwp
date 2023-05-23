package main

import (
    "fmt"
    "os"
    "path/filepath"

    flag "github.com/spf13/pflag"
    // "github.com/NakaokaTomoki/cwp"
)

const VERSION = "0.1.16"


func versionString(args []string) string {
    prog := "cwp"
    if len(args) > 0 {
        prog = filepath.Base(args[0])
    }
    return fmt.Sprintf("%s version %s", prog, VERSION)
}

type CwpError struct {
    statusCode int
    message string
}

func (e CwpError) Error() string {
    return e.message
}

/*
helpMessage prints the help message.
This function is used in the small tests, so it may be called with a zero-length slice.
*/
func helpMessage(args []string) string {
    prog := "cwp"
    if len(args) > 0 {
        prog = filepath.Base(args[0])
    }
    return fmt.Sprintf(`%s [OPTIONS] [URLs...]
OPTIONS
    -t, --token <TOKEN>      specify the token for the service. This option is mandatory.
    -q, --qrcode <FILE>      include QR-code of the URL in the output.
    -c, --config <CONFIG>    specify the configuration file.
    -d, --delete             delete the specified shorten URL.
    -h, --help               print this mesasge and exit.
    -v, --version            print the version and exit.
    `, prog)
}

type flags struct {
    deleteFlag    bool
    listGroupFlag bool
    helpFlag      bool
    versionFlag   bool
}

type runOpts struct {
    token  string
    qrcode string
    config string
    group  string
}

/*
This struct holds the values of the options.
*/
type options struct {
    runOpt  *runOpts
    flagSet *flags
}

func newOptions() *options {
    return &options{runOpt: &runOpts{}, flagSet: &flags{}}
}

// func (opts *options) mode(args []string) cwp.Mode {
//     switch {
//     case opts.flagSet.listGroupFlag:
//         return cwp.ListGroup
//     case len(args) == 0:
//         return cwp.List
//     case opts.flagSet.deleteFlag:
//         return cwp.Delete
//     case opts.runOpt.qrcode != "":
//         return cwp.QRCode
//     default:
//         return cwp.Shorten
//     }
// }

func buildOptions(args []string) (*options, *flag.FlagSet) {
    opts := &options{}
    flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
    flags.Usage = func() { fmt.Println(helpMessage(args)) }
    flags.StringVarP(&opts.runOpt.token, "token", "t", "", "specify the")

    return opts, flags
}

func perform(opts *options, args []string) *CwpError {
    fmt.Println("Hello World")
    return nil
}

// func parseOptions(args []string) (*options, []string, *CwpError) {
//     opts, flags := buildOptions(args)
//
//     flags.Parse(args[1:])
//
//     if opts.help {
//         fmt.Println(helpMessage(args[0]))
//     }
//     if opts.token == "" {
//         return nil, nil, &CwpError{statusCode: 3, message: "no token was given"}
//     }
//     return opts, flags.Args(), nil
// }
/*
parseOptions parses options from the given command line arguments.
*/
func parseOptions(args []string) (*options, []string, *CwpError) {
    opts, flags := buildOptions(args)
    flags.Parse(args[1:])
    if opts.flagSet.helpFlag {
        fmt.Println(helpMessage(args))
        return nil, nil, &CwpError{statusCode: 0, message: ""}
    }
    if opts.flagSet.versionFlag {
        fmt.Println(versionString(args))
        return nil, nil, &CwpError{statusCode: 0, message: ""}
    }
    if opts.runOpt.token == "" {
        return nil, nil, &CwpError{statusCode: 3, message: "no token was given"}
    }
    return opts, flags.Args(), nil
}

func makeError(err error, status int) *CwpError {
    if err == nil {
        return nil
    }
    ue, ok := err.(*CwpError)
    if ok {
        return ue
    }
    return &CwpError{statusCode: status, message: err.Error()}
}

func goMain(args []string) int {
    opts, args, err := parseOptions(args)

    if err != nil {
        if err.statusCode != 0 {
            // fmt.Println(err.Error())
            fmt.Println(err)
        }
        return err.statusCode
    }
    if err := perform(opts, args); err != nil {
        // fmt.Println(err.Error())
        fmt.Println(err)
        return err.statusCode
    }
    return 0
}

func main() {
    status := goMain(os.Args)
    os.Exit(status)
}
