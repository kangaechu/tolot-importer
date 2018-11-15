tolot_importer
====

[![CircleCI](https://circleci.com/gh/kangaechu/tolot-importer.svg?style=svg)](https://circleci.com/gh/kangaechu/tolot-importer)

[tolot](https://tolot.com)の連絡帳インポーター

## Description

CSVファイルで作成した連絡帳情報を[tolot](https://tolot.com)の連絡帳に取り込みます。

## Download

https://github.com/kangaechu/tolot_importer/releases/

## Requirements

Google Chrome

## Usage

### バイナリから実行(Windows)

#### ダウンロード

https://github.com/kangaechu/tolot-importer/releases から最新の project_windows_amd64.exe をダウンロード

#### 設定ファイルの作成

settings.yaml.sampleをsettings.yamlにコピーする。
settings.yamlにログインのアカウント情報とインポートするCSV連絡帳のファイル名を指定する。
```
userID: "USERID"               # ログインのユーザ名
password: "PASSWORD"           # パスワード
addressFileName: "address.csv" # インポートするCSV連絡帳のファイル名
```

#### インポートするCSV連絡帳の作成

CSVファイルで連絡帳を作成する。デフォルトはaddress.csv。
UTF-8、改行コードはCRLFで作成。
フォーマットは以下の通り。ヘッダ行は不要。

```
[姓],[名],[連名],[郵便番号],[都道府県名],[市区町村名],[住所それ以降]
```

#### 実行

tolot_importer.exeをダブルクリック。
1度目の実行時にはWindows Defender Smartscreenのダイアログが表示されるので、詳細情報→実行を選択する。
Google Chromeが立ち上がり、実行が開始される。
認証ダイアログは適宜入力、閉じること。
実行が終わるとoutput.csvに30分ごとの電力量が出力される。

### ソースから実行

#### ダウンロード

```
$ go get github.com/kangaechu/tolot_importer
$ cd $GOPATH/github.com/kangaechu/tolot_importer
```

#### 設定ファイルの作成

settings.yaml.sampleをsettings.yamlにコピーする。
settings.yamlにログインのアカウント情報とインポートするCSV連絡帳のファイル名を指定する。
```
userID: "USERID"               # ログインのユーザ名
password: "PASSWORD"           # パスワード
addressFileName: "address.csv" # インポートするCSV連絡帳のファイル名
```

#### インポートするCSV連絡帳の作成

CSVファイルで連絡帳を作成する。デフォルトはaddress.csv。
UTF-8、改行コードはCRLFで作成。
フォーマットは以下の通り。ヘッダ行は不要。

```
[姓],[名],[連名],[郵便番号],[都道府県名],[市区町村名],[住所それ以降]
```

#### その他

実行時にGoogle Chromenについての画面が出力される場合は
https://github.com/rjeczalik/chromedp/commit/520a76514fd911e7544c72412d307fec5ae524ad を適用すると修正されます。

#### 実行
```
$ go run main.go
```

## Install

クロスビルドも可能です。
```
# Linux
$ GOOS=linux GOARCH=amd64 go build -o tolot_importer main.go
# OSX
$ GOOS=darwin GOARCH=amd64 go build -o tolot_importer main.go
# Windows
$ GOOS=windows GOARCH=amd64 go build -o tolot_importer.exe main.go
```

## Licence

[MIT](https://github.com/tcnksm/tool/blob/master/LICENCE)

## Author

[kangaechu](https://github.com/kangaechu)
