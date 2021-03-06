# 関数
ウェブクローラを題材として Go に特有な関数の特徴を見ていく.

## 5.1 関数宣言

- 名前、パラメータリスト、結果リスト、本体で構成される.
- 結果にも名前をつけることができる
- （ブランク識別子をパラメータに指定することの意味は何なのだろう）
- 関数の型をシグニチャと呼ぶことがある
  - パラメータと結果の列の型が同じであればシグニチャは同じとみなすらしい. 名前は見ない
  - （界隈によっては関数名などを含む場合もあるが..型チェックでの話をしているのだろう）
- デフォルトパラメータ値という概念はないのでパラメータに対する引数は必ず提供せねばならない
  - by value で渡される（ローカル変数として扱う）
  - ポインタやスライスなどの間接的な参照があると呼び出し元でも影響を受ける
- アセンブリなど他の言語で実装されている関数を参照する場合には本体（`{}` ブラケットの中身）を書かない

## 5.2 再帰

- golang.org には非標準パッケージがおいてある
  - 今回の題材である html.Parse も非標準
- 多くの言語ではスタックが固定長であるので再帰呼び出しの深さに制限があり、引っかかると stack overflow error となり, security 上のリスクをもたらすことがある
  - Go は可変長スタック
  - （ということはメモリを食いつぶすまで無限に再起できてしまう..？）

## 5.3 複数戻り値

- 多値関数の呼び出し結果は、複数のパラメータを持つ関数呼び出しの引数として、変数を経由することなくそのまま突っ込むことができる
- 関数の宣言で結果に名前をつけてあれば空リターンを使うことができる
  - コードの重複は減らせるが理解を容易にはしないので多用すべきではない

## 5.4 エラー

- エラーはパッケージの API やアプリケーションのユーザインタフェースの重要な一部であり、失敗は複数の予期される振る舞いのうちの一つにすぎません
- 単に実行の失敗を示すために、関数の戻り値の最後に追加の結果をブーリアンとして記述することが多い
- I/O のように様々な原因がありうる場合、すなわち原因に対する説明が必要な場合には error を返す
  - 成功すれば error の代わりに nil を返す
  - error は fmt.Println や fmt.Printf("%v", err) で表示可能
  - 詳細は７章
- 原因のハンドリングのために通常は error だけを見ればよいが error だけでは情報が不十分なときもある
  - Read は読み出しの途中でエラーが起こると、読み込むことができたバイト数も返す（バイト数を見ることでどのデータの読み出しで中断したかがわかる）
- Go では error を exception （特別なもの）として扱わずに普通の制御フロー（if や return）で制御する
  - （実際 catch の中で if を使った制御を書いたりするので catch がさらにコードを複雑にするいう主張はわからんでもない）

### 5.4.1 エラー処理戦略
５つの戦略が書かれている.

#### エラーの伝播

- fmt.Errorf でエラーメッセージへ情報の追加が可能
  - Sprintf と同じ感覚で呼び出し元エラー情報に連鎖して書ける

#### 再試（retry）

#### 呼び出し元で制御して停止

- これが適用できるのは main 関数くらい
- log package を使えば接頭辞の設定や時刻の表示抑制が可能

#### エラーを記録するにとどめて制限された機能で処理を続行
オフラインで処理を続行するなど

#### エラーを無視
/tmp 配下の削除処理など

### 5.4.2 ファイルの終わり
EOF (end-of-file) はエラーとして返ってくるが通常のエラーとは違う対処が必要. 詳細は 7.11

## 5.5 関数値

- 関数値 function value はファーストクラス値 first-classs value
  - 関数定義そのものを変数に代入したり引数として別の関数に渡したりできる
- 関数型のゼロ値は nil
- 関数値は比較不可で map key にもなれない
- 振る舞いをカスタマイズするために関数値をパラメータとして使う例として strings.Map などがある

## 5.6 無名関数
