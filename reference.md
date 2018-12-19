# 東北ずん子語リファレンスマニュアル

t_zungoは、[東北]と[ずんだ]と[太もも]の３つの単語を使って処理を書くプログラミンング言語です。

## 実例

[hello_world.znk](https://github.com/135yshr/t_zungo/blob/master/hello_world.znk)は、コンソール画面に`Hello, world!`と表示するサンプルプログラムです。

```
./zungo hello_world.znk
```

## 文法

IMP (Instruction Modification Parameter)、コマンド、パラメータの3つ組で命令を表現します。

| IMP | 対象 |
| ---- | ---- |
| [東北] | スタック操作 |
| [太もも] [東北] | 演算 |
| [太もも] [太もも] | ヒープアクセス |
| [ずんだ] | フロー制御 |
| [太もも] [ずんだ] | 入出力 |

* 数値は二進記数法で表現する。[東北]が0で、[太もも]が1で、[ずんだ]が終端記号になります。
   また、最初が[太もも]の場合は、マイナスの値として扱います。

| パターン | 数値 |
| ------- | ---- |
| [東北][東北][ずんだ] | 0 |
| [東北][太もも][ずんだ] | 1 |
| [東北][太もも][東北][太もも][ずんだ] | 5 |
| [太もも][太もも][ずんだ] | -1 |
| [太もも][太もも][東北][太もも][ずんだ] | -5 |

### スタック操作

| 命令 | パラメータ |　対象 |
| ---- | --- | --- |
| [東北] | 数値 | 数値をスタックに積む |
| [ずんだ] [東北] |  | スタックトップを複製する |
| [ずんだ] [太もも] |  | スタックトップと２番目を交換する |
| [ずんだ] [ずんだ] |  | スタックトップを捨てる |
| [太もも] [東北] | 数値 | スタックのn番目をコピーして一番上に積む |
| [太もも] [ずんだ] | 数値 | スタックのn番目を移動して一番上に積む |


### 演算

| 命令 | パラメータ |　対象 |
| ---- | --- | --- |
| [東北] [東北] |  | スタックの上から２つを加算 |
| [東北] [太もも] |  | スタックの上から２つで減算 |
| [東北] [ずんだ] |  | スタックの上から２つで乗算 |
| [太もも] [東北] |  | スタックの上から２つで除算 |
| [太もも] [太もも] |  | スタックの上から２つで剰余 |


### ヒープアクセス

| 命令 | パラメータ |　対象 |
| ---- | --- | --- |
| [東北] |  | スタックの値をアドレスに格納 |
| [太もも] |  | アドレスから値をスタックにコピーする |

### フロー制御

| 命令 | パラメータ |　対象 |
| ---- | --- | --- |
| [東北] [東北] | Label | Label 定義 |
| [東北] [太もも] | Label | サブルーチンを呼び出す |
| [東北] [ずんだ] | Label | 無条件で Label にジャンプする |
| [太もも] [東北] | Label | スタックトップがゼロなら Label にジャンプ  |
| [太もも] [太もも] | Label | スタックトップが負なら Label にジャンプ |
| [太もも] [ずんだ] |  | サブルーチンを終了 |
| [ずんだ] [ずんだ] |  | プログラム終了 |

### 入出力

| 命令 | パラメータ |　対象 |
| ---- | --- | --- |
| [東北] [東北] |  | スタックトップの値を文字として出力 |
| [東北] [太もも] |  | スタックトップの値を数値として出力 |
| [太もも] [東北] |  | 文字を読み込みアドレスに格納 |
| [太もも] [太もも] |  | 数値を読み込みアドレスに格納  |