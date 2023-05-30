# cwp (Current Weather Print) :HELP:
[![build](https://github.com/NakaokaTomoki/cwp/actions/workflows/build.yaml/badge.svg)](https://github.com/NakaokaTomoki/cwp/actions/workflows/build.yaml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
![Version](https://img.shields.io/badge/Version-0.2.10-blueviolet)

[![Coverage Status](https://coveralls.io/repos/github/NakaokaTomoki/cwp/badge.svg?branch=main)](https://coveralls.io/github/NakaokaTomoki/cwp?branch=main)
[![codebeat badge](https://codebeat.co/badges/f97b15e5-c079-49f6-b21b-d1b49ff863a9)](https://codebeat.co/projects/github-com-nakaokatomoki-cwp-main)
[![Go Report Card](https://goreportcard.com/badge/github.com/NakaokaTomoki/cwp)](https://goreportcard.com/report/github.com/NakaokaTomoki/cwp)

- cwp version 0.2.10

<!-- tagline -->
- 指定した場所の気象情報をCLI上に表示するコマンド


# Description
- OpenWeatherMap(https://openweathermap.org/) が提供しているAPI を使用して指定した場所の天気や気温を取得する
- API から取得した情報 CLI上に表示する
- 出力項目は，天気・気温
    - 出力形式は，デフォルトはプレーンテキスト
    - (ex) 天気： 晴れ☀️ 気温： 22度


# Usage
```
$ cwp [オプション]

オプション:
    -place:  天気予報を行う場所を指定(デフォルトは日本)
    -units: 測定単位
    -lang: 出力言語(デフォルトは日本語)
    -format: 出力形式を指定(デフォルトはプレーンテキスト)
    -version: cwpのバージョン情報を表示
    -help: cwpコマンドのバージョン情報および利用可能なオプションを表示
```

# License
MIT License


# About
## License


## Developers
- Nakaoka Tomoki


## Icon
![Icon](docs/static/images/weather_cat.png)
