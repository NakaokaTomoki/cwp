package main

import (
	"os"
	"fmt"
        // "time"
	"path/filepath"

	flag "github.com/spf13/pflag"
	"github.com/NakaokaTomoki/cwp"
)

const VERSION = "0.2.41"

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
func helpMessage(args []string, flags *flag.FlagSet) string {
	prog := "cwp"
	if len(args) > 0 {
		prog = filepath.Base(args[0])
	}
	return fmt.Sprintf(`%s [OPTIONS] [PLACEs...]
OPTIONS:
%s
ARGUMENT:
    PLACE      天気予報を行う場所を指定(デフォルトはtokyo(東京))`, prog, flags.FlagUsages())
}

type CwpError struct {
	statusCode int
	message    string
}

func (e CwpError) Error() string {
	return e.message
}

type flags struct {
	helpFlag      bool
	versionFlag   bool
}

type runOpts struct {
	token  string
	config string
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

/*
Define the options and return the pointer to the options and the pointer to the flagset.
*/
func buildOptions(args []string) (*options, *flag.FlagSet) {
	opts := newOptions()
        completions := false

	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage(args, flags)) }
	flags.StringVarP(&opts.runOpt.token, "token", "t", "", "OpenWeatherMap API の使用に必要となるトークンを指定(必須)")
	flags.BoolVarP(&opts.flagSet.helpFlag, "help", "h", false, "cwpコマンドのバージョン情報および利用可能なオプションを表示")
	flags.BoolVarP(&opts.flagSet.versionFlag, "version", "v", false, "cwpのバージョン情報を表示")

	flags.BoolVarP(&completions, "generate-completions", "", false, "generate completions")
	flags.MarkHidden("generate-completions")

	return opts, flags
}

/*
parseOptions parses options from the given command line arguments.
*/
func parseOptions(args []string) (*options, []string, *CwpError) {
	opts, flags := buildOptions(args)
	flags.Parse(args[1:])

	if value, _ := flags.GetBool("generate-completions"); value {
		err := GenerateCompletion(flags)
		if err != nil {
			return nil, nil, &CwpError{statusCode: 1, message: err.Error()}
		}
		return nil, nil, &CwpError{statusCode: 0, message: "generate completions"}
	}

	if opts.flagSet.helpFlag {
		fmt.Println(helpMessage(args, flags))
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

func printResult(result *cwp.OpenWeather) {
	// fmt.Println(result.Place)
	// fmt.Println(result.Weather[0].WeatherDescription)
	// fmt.Println(result.Temp["temp"])
	// fmt.Println(time.Unix(result.DateTime, 0))

	fmt.Printf("[場所] %s, [天気] %s, [気温] %.2f\n", result.Place, result.Weather[0].WeatherDescription, result.Temp["temp"])
}

func getWeatherEach(openweathermap *cwp.OpenWeatherMap, place string, config *cwp.Config) error {
	result, err := openweathermap.GetWeather(place, config)
	if err != nil {
		return err
	}
	// fmt.Println(result)
        printResult(result)

	return nil
}

func perform(opts *options, args []string) *CwpError {
	openweathermap := cwp.NewOpenWeatherMap()
	config := cwp.NewConfig(opts.runOpt.config)
	config.Token = opts.runOpt.token

	return performImpl(args, func(place string) error {
		return getWeatherEach(openweathermap, place, config)
	})
	return nil
}

func performImpl(args []string, executor func(place string) error) *CwpError {
	args = parseArgs(args)
	for _, place := range args {
		err := executor(place)
		if err != nil {
			return makeError(err, 3)
		}
	}
	return nil
}

func parseArgs(args []string) []string {
	default_place := "tokyo"

	if len(args) == 0 {
		args = append(args, default_place)
	}
	return args
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
