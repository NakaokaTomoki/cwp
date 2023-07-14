# cwp (Current Weather Print) :worried:
[![build](https://github.com/NakaokaTomoki/cwp/actions/workflows/build.yml/badge.svg)](https://github.com/NakaokaTomoki/cwp/actions/workflows/build.yml)
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

- 出力項目は，場所・天気・気温
    - 出力形式は，プレーンテキスト
    - (ex)
        - [場所] 東京都, [天気] 曇りがち, [気温] 28.90


# Usage
```
$ cwp [オプション] [引数]

オプション:
    -t, --token <TOKEN>    OpenWeatherMap API の使用に必要となるトークンを指定(必須)
    -h, --help             cwpコマンドのバージョン情報および利用可能なオプションを表示
    -v, --version          cwpのバージョン情報を表示
引数:
    PLACEs:                 天気予報を行う場所を指定(デフォルトはtokyo(東京))
```

# Installation

## :beer: Homebrew
```sh
brew install NakaokaTomoki/brew/cwp
```

## :whale: Docker
```sh
docker run -it --rm NakaokaTomoki/cwp:latest -t <token> <place...>
```

# License
MIT License


# About
## License
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)


## Developers
- Nakaoka Tomoki


## Icon
![Icon](docs/static/images/weather_cat.png)
