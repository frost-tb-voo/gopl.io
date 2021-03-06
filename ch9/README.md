# 共有された変数による並行性（concurrency）
## 9.6 競合検出器
go run/build/test に -race flag をつけてみよう

- -race flag は, 共有された変数へのアクセス、
- go func, channel 操作, sync.Mutex, sync.WaitGroup などの同期イベントを記録するようプログラムを修正

修正されたプログラムは, 直近で他の go func から書き込みがあった共有された変数に対して, ある go func が読み書きする事例がないか, すなわちデータ強豪の事例を検査. 以下をレポート

- その変数がどれか
- 関係する go routine の stack

実行時に発生した競合状態のみが検査対象であり（動的解析）, すべての可能性を検査するわけではない

## 9.7 例：並行的で待たされないキャッシュ (concurrent non-blocking cache)
メモ化（memorizing）により, 一度だけしか計算する必要がないように関数の結果をキャッシュするテクニック.
以下の関数 memorizing を実現する. ReadAll の返り値 byte[] の型がシグニチャでは interface により隠蔽されていることに注意.

```go
func httpGetBody(url string) (interface, error) {
  resp, err := http.Get(url)
  if err != nil {
      return nil, err
  }
  defer resp.Body.Close()
  return ioutil.ReadAll(resp.Body)
}
```

gopl.io/ch9/memo1

以下省略

## 9.7

## 9.8 ごルーチンとスレッド
### 9.8.1 伸長可能なスタック
個々の OS においてスレッドはたいてい 2MB 程度のメモリをスタックとしてもっている.

- WaitGroup で単に待ってからチャネルを閉じるだけのごルーチンでは不要なサイズ
- 数十万個のごルーチンを同時に作成することも稀ではないため, その場合スタックの深さに対してサイズ不足

ごルーチンは 2KB の小さなスタックで開始できるようになっている.
スタックサイズが呼び出し時に固定される OS のスレッドと異なり, ごルーチンのスタックは最大 1GB 程度まで伸縮が可能.

### 9.8.2 ごルーチンのスケジュール
OS スレッドのスケジューリングはカーネルによりスケジュールされる.
スケジューラと呼ばれるカーネル関数を, ハードウェアタイマの割り込みにより数ミリ秒ごと定期的に起動することで実現している.

- スレッドを一時停止
- レジスタをメモリへ退避
- スレッド一覧から次のスレッドを選んで
- 次に実施するスレッドのレジスタ内容をメモリから回復し
- 次のスレッドを再開する

完全な context-switch による実現で局所性に乏しくメモリアクセス数も多いので時間がかかる.
（ハードウェアの進化によって？）メモリアクセスに必要な CPU サイクル数は増加しつつあるので, この影響はより大きくなっている.

Go runtime は m 個のごルーチンと n 個のスレッドで多重化を行っており, この技法は m:n スケジューリングと呼ばれている.
OS のスレッドとは異なり, Go runtime が扱うのは単一の go program のみである.

go のスケジューリングは OS のスケジューリングのような定期的なハードウェアタイマによる割り込みで実現されているわけではない.
go routine による time.Sleep の呼び出しやチャネルや mutex の呼び出しで待機に入ったタイミングで別のごルーチンへの切り替えを行う.
カーネルコンテキストへの切り替えが不要なので低コスト.

### 9.8.3 GOMAXPROCS
省略
