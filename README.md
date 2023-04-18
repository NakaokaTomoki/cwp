# cwp (Current Weather Print)
- cwp version 1.0.0
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
