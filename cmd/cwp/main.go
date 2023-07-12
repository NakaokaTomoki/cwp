package main

import (
	"fmt"
	"os"
	"path/filepath"

	flag "github.com/spf13/pflag"
	"github.com/NakaokaTomoki/cwp"
)

const VERSION = "0.2.12"

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
	return fmt.Sprintf(`%s [OPTIONS] [Places...]
OPTIONS
    -t, --token <TOKEN>      specify the token for the service. This option is mandatory.`, prog)
}

type cwpError struct {
	statusCode int
	message    string
}

func (e cwpError) Error() string {
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
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage(args)) }
	flags.StringVarP(&opts.runOpt.token, "token", "t", "", "specify the token for the service. This option is mandatory.")
	// flags.StringVarP(&opts.runOpt.qrcode, "qrcode", "q", "", "include QR-code of the URL in the output.")
	// flags.StringVarP(&opts.runOpt.config, "config", "c", "", "specify the configuration file.")
	// flags.StringVarP(&opts.runOpt.group, "group", "g", "", "specify the group name for the service. Default is \"cwp\"")
	// flags.BoolVarP(&opts.flagSet.listGroupFlag, "list-group", "L", false, "list the groups. This is hidden option.")
	// flags.BoolVarP(&opts.flagSet.deleteFlag, "delete", "d", false, "delete the specified shorten URL.")
	flags.BoolVarP(&opts.flagSet.helpFlag, "help", "h", false, "print this mesasge and exit.")
	flags.BoolVarP(&opts.flagSet.versionFlag, "version", "v", false, "print the version and exit.")
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
	if opts.runOpt.token == "" {
		return nil, nil, &cwpError{statusCode: 3, message: "no token was given"}
	}
	return opts, flags.Args(), nil
}

func getWeatherEach(openweathermap *cwp.OpenWeatherMap, place string, config *cwp.Config) error {
	result, err := openweathermap.GetWeather(place, config)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

func performImpl(args []string, executor func(place string) error) *cwpError {
	for _, place := range args {
		err := executor(place)
		if err != nil {
			return makeError(err, 3)
		}
	}
	return nil
}

func perform(opts *options, args []string) *cwpError {
	openweathermap := cwp.NewOpenWeatherMap()
	config := cwp.NewConfig(opts.runOpt.config)
	config.Token = opts.runOpt.token

	return performImpl(args, func(place string) error {
		return getWeatherEach(openweathermap, place, config)
	})
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
