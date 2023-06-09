package main

import "testing"

func _Example_Main() {
	goMain([]string{"./cwp", "-t", "token"})
	// Output:
	// Hello World
}

func Example_Help() {
	goMain([]string{"./cwp", "--help"})
	// Output:
	// cwp [OPTIONS] [URLs...]
	// OPTIONS
	//     -t, --token <TOKEN>      specify the token for the service. This option is mandatory.
	//     -q, --qrcode <FILE>      include QR-code of the URL in the output.
	//     -c, --config <CONFIG>    specify the configuration file.
	//     -g, --group <GROUP>      specify the group name for the service. Default is "cwp"
	//     -d, --delete             delete the specified shorten URL.
	//     -h, --help               print this mesasge and exit.
	//     -v, --version            print the version and exit.
	// ARGUMENT
	//     URL     specify the url for shortening. this arguments accept multiple values.
	//             if no arguments were specified, cwp prints the list of available shorten urls.
}

func Test_Main(t *testing.T) {
	if status := goMain([]string{"./cwp", "-v"}); status != 0 {
		t.Error("Expected 0, got ", status)
	}
}
