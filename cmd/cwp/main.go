package main

import (
    "fmt"
    "os"
    "path/filepath"

    flag "github.com/spf13/pflag"
    "github.com/NakaokaTomoki/cwp"
)

const VERSION = "0.2.11"

func versionString(args []string) string {
    prog := "cwp"
    if len(args) > 0 {
        prog = filepath.Base(args[0])
    }
    return fmt.Sprintf("%s version %s", prog, VERSION)
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
    // return fmt.Sprintf(`%s [OPTIONS] [URLs...]
    return fmt.Sprintf(`%s [OPTIONS] [PLACEs...]
オプション
    -units: 測定単位
    -lang: 出力言語(デフォルトは日本語)
    -format: 出力形式を指定(デフォルトはプレーンテキスト)
    -version: cwpのバージョン情報を表示
    -help: cwpコマンドのバージョン情報および利用可能なオプションを表示

引数
    場所:  天気予報を行う場所を指定(デフォルトは日本)` , prog)
}

type cwpError struct {
    statusCode int
    message    string
}

func (e cwpError) Error() string {
    return e.message
}

type flags struct {
    deleteFlag    bool
    listGroupFlag bool
    helpFlag      bool
    versionFlag   bool
}

type runOpts struct {
    unit    string
    lang    string
    format  string
    version string
    help    string
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

func (opts *options) mode(args []string) cwp.Mode {
    switch {
    case opts.flagSet.listGroupFlag:
        return cwp.ListGroup
    case len(args) == 0:
        return cwp.List
    case opts.flagSet.deleteFlag:
        return cwp.Delete
    // case opts.runOpt.qrcode != "":
    //     return cwp.QRCode
    default:
        return cwp.Shorten
    }
}

/*
Define the options and return the pointer to the options and the pointer to the flagset.
*/
func buildOptions(args []string) (*options, *flag.FlagSet) {
    opts := newOptions()
    flags := flag.NewFlagSet(args[0], flag.ContinueOnError)

    flags.Usage = func() { fmt.Println(helpMessage(args)) }

    flags.StringVarP(&opts.runOpt.unit, "units", "u", "", "測定単位")
    flags.StringVarP(&opts.runOpt.qrcode, "lang", "l", "", "出力言語(デフォルトは日本語)")
    flags.StringVarP(&opts.runOpt.config, "format", "f", "", "出力形式を指定(デフォルトはプレーンテキスト)")
    flags.StringVarP(&opts.runOpt.group, "version", "v", "", "cwpのバージョン情報を表示(デフォルトは \"cwp\")")
    flags.BoolVarP(&opts.flagSet.helpFlag, "help", "h", false, "cwpコマンドのバージョン情報および利用可能なオプションを表示")
    // flags.BoolVarP(&opts.flagSet.listGroupFlag, "list-group", "L", false, "list the groups. This is hidden option.")
    // flags.BoolVarP(&opts.flagSet.deleteFlag, "delete", "d", false, "delete the specified shorten URL.")
    // flags.BoolVarP(&opts.flagSet.versionFlag, "version", "v", false, "print the version and exit.")

    return opts, flags
}

/*
parseOptions parses options from the given command line arguments.
*/
func parseOptions(args []string) (*options, []string, *cwpError) {
    opts, flags := buildOptions(args)
    flags.Parse(args[1:])

    if opts.flagSet.helpFlag {
        fmt.Println(helpMessage(args))
        return nil, nil, &cwpError{statusCode: 0, message: ""}
    }
    if opts.flagSet.versionFlag {
        fmt.Println(versionString(args))
        return nil, nil, &cwpError{statusCode: 0, message: ""}
    }

    // if opts.runOpt.token == "" {
    //     return nil, nil, &cwpError{statusCode: 3, message: "no token was given"}
    // }
    return opts, flags.Args(), nil
}

func shortenEach(bitly *cwp.Bitly, config *cwp.Config, url string) error {
    result, err := bitly.Shorten(config, url)
    if err != nil {
        return err
    }
    fmt.Println(result)
    return nil
}

func deleteEach(bitly *cwp.Bitly, config *cwp.Config, url string) error {
    return bitly.Delete(config, url)
}

func listUrls(bitly *cwp.Bitly, config *cwp.Config) error {
    urls, err := bitly.List(config)
    if err != nil {
        return err
    }
    for _, url := range urls {
        fmt.Println(url)
    }
    return nil
}

func listGroups(bitly *cwp.Bitly, config *cwp.Config) error {
    groups, err := bitly.Groups(config)
    if err != nil {
        return err
    }
    for i, group := range groups {
        fmt.Printf("GUID[%d] %s\n", i, group.Guid)
    }
    return nil
}

func performImpl(args []string, executor func(url string) error) *cwpError {
    for _, url := range args {
        err := executor(url)
        if err != nil {
            return makeError(err, 3)
        }
    }
    return nil
}

func perform(opts *options, args []string) *cwpError {
    bitly := cwp.NewBitly(opts.runOpt.group)
    config := cwp.NewConfig(opts.runOpt.config, opts.mode(args))
    config.Token = opts.runOpt.token
    switch config.RunMode {
    case cwp.List:
        err := listUrls(bitly, config)
        return makeError(err, 1)
    case cwp.ListGroup:
        err := listGroups(bitly, config)
        return makeError(err, 2)
    case cwp.Delete:
        return performImpl(args, func(url string) error {
            return deleteEach(bitly, config, url)
        })
    case cwp.Shorten:
        return performImpl(args, func(url string) error {
            return shortenEach(bitly, config, url)
        })
    }
    return nil
}

func makeError(err error, status int) *cwpError {
    if err == nil {
        return nil
    }
    ue, ok := err.(*cwpError)
    if ok {
        return ue
    }
    return &cwpError{statusCode: status, message: err.Error()}
}

func goMain(args []string) int {
    opts, args, err := parseOptions(args)
    if err != nil {
        if err.statusCode != 0 {
            fmt.Println(err.Error())
        }
        return err.statusCode
    }
    if err := perform(opts, args); err != nil {
        fmt.Println(err.Error())
        return err.statusCode
    }
    return 0
}

func main() {
    status := goMain(os.Args)
    os.Exit(status)
}
