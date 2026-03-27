<!-- -*- encoding: utf-8 -*- -->

siiview
====

標準入力から受け取った画像データをウインドウ表示するGo言語製簡易ビューワ

Description / 説明
------------------

* 標準入力から受け取った画像をウインドウ表示する。
* fyneを使ってウインドウ表示をしている。
* png/bmp/jpeg/gif/tiff に対応。
 
* ESC / Q key で終了

Requirement / 依存関係
----------------------

* Go 1.25.7 64bit
* fyne.io/fyne/v2 v2.7.3
* golang.org/x/image v0.38.0
* WinLibs (gcc.exe)

Environment / 動作確認環境
--------------------------

* Windows11 x64 25H2

Usage / 使い方
--------------

```
cat image.png | siiview

gmic sp lena o -.png | siiview
```

ビルド手順
----------

main.go を作成後、以下を打った。

```
go mod init siiview
go mod tidy
```

依存関係が go.mod として作成された。

GUIライブラリ fyne を使うので gcc が必要になる。
今回は WinLibs を使用した。

[WinLibs - GCC+MinGW-w64 compiler for Windows](https://winlibs.com/)

go build でビルド。siiview.exe が生成される。

License / ライセンス
--------------------

[MIT](https://opensource.org/licenses/MIT)


Author / 作者名
---------------

[mieki256](https://github.com/mieki256)

