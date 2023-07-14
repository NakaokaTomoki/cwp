package main

import "testing"

func Example_Help() {
	goMain([]string{"./cwp", "--help"})
	// Output:
	// cwp [OPTIONS] [PLACEs...]
	// OPTIONS:
	// 	-h, --help           cwpコマンドのバージョン情報および利用可能なオプションを表示
	// 	-t, --token string   OpenWeatherMap API の使用に必要となるトークンを指定(必須)
	// 	-v, --version        cwpのバージョン情報を表示

	// RGUMENT:
	// 	PLACE      天気予報を行う場所を指定(デフォルトはtokyo(東京))
}

func Test_Main(t *testing.T) {
	if status := goMain([]string{"./cwp", "-v"}); status != 0 {
		t.Error("Expected 0, got ", status)
	}
}
