# 12. Reflection
コンパイル時に変数の型がわかっていなくても, 実行時に変数を更新したり, 値を調べたり,.. といったことをやりたいことがある.
特に次の２つのモジュールが提供する API において reflection は重要.

- fmt.* : 文字列へのフォーマット
- encoding/* : プロトコルのエンコーディング
- あるいは text/template, html/template も内部では reflection を利用

### 12.1. Why Reflection?

- fmt.Sprintf の例
- switch/case と .type では限界がある
  - 未知の型を持つ値がどういう表現を持つのか調べられない

- 特に []float64, map[string][]string などは表現が無数に存在する
- また map[string][]string を基底型とする url.Values 型の値は map[string][]string とは別の型と認識されるので、map[string][]string との比較には意味がない
    - url.Values に一致するか調べるには url を import する必要がある

### 12.2. reflect.Type and reflect.Value

reflect.TypeOf で reflect.Type が返る.

```go
// TypeOf の引数は interface{}
t := reflect.TypeOf(3)
// 動的な型は int
fmt.Println(t) // "int"

var w io.Writer = os.Stdout
// 動的な型は *os.File
reflect.TypeOf(w) // "*os.File"

// 内部では TypeOf を利用している
fmt.Printf("%T\n", 3) // "int"
```

- reflect.ValueOf で reflect.Value が返る.
- 逆の操作は reflect.Value.Interface
  - ただし型は interface{} 型になる
  - 型アサーションをつかない限りほとんど何もできない
  - Value 型であればその内容を調べるためのメソッドが色々揃っている
  - Value.Kind() の値に応じて利用できるメソッドが異なる

```go
v := reflect.ValueOf(3)
// Value の String() は文字列でなければ型を返す
fmt.Println(v.String()) // "<int Value>"
// 内部で reflect.Value を特別に処理している
fmt.Println(v)
fmt.Printf("%v\n", v)

// reflect.Value.Interface
x := v.Interface()
i := x.(int)
fmt.Printf("%d\n", i) // "3"
```

Kind の値は有限なので switch で処理できる.
合成型、インタフェース、参照型を処理するには再帰などを利用する必要がある.

- reflect.Invalid
- reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64
- reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr
- reflect.Bool
- reflect.String
- reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map

### 12.3. Display, a Recursive Value Printer
コンポジット型（配列、スライス、マップ、構造体：４章）の処理.

```go
e, _ := eval.Parse("sqrt(A/pi)")
Display("e", e)
```

次のような表示にしたい.

```
Display e (eval.call):
e.fn = "sqrt"
e.args[0].type = eval.binary
e.args[0].value.op = 47
e.args[0].value.x.type = eval.Var
e.args[0].value.x.value = "A"
e.args[0].value.y.type = eval.Var
e.args[0].value.y.value = "pi"
```

方法：reflect.Value 中のメソッドを利用して e 内部の各要素を取り出し, 再帰処理を実施する.

- Kind に応じたメソッドしか呼び出せないので注意
  - 例えば構造体（Struct）に対してスライス（Slice）のメソッドである Index を呼び出すと panic になる. らしい
- スライスと配列：Len と Index(添字) を使う
- 構造体：NumField と Field(添字) を使う
- マップ：MapKeys と MaxIndex(キー) を使う
  - 演習 12.1 で拡張
- ポインタ：Elm と IsNil を使う
- インタフェース：Elm と IsNil を使う

reflect を使うと公開されていないフィールドも見ることができる.
例えば次のコードは非公開のフィールドを表示するが, 結果はプラットフォームに応じて変わる可能性がある.

```go
Display("rv", reflect.ValueOf(os.Stderr))
```

また次の２つは異なる結果を返す.

```go
i := 3
Display("i", i) // i = 3
Display("&i", &i)
// (*&i).type = int
// (*&i).value = 3
```

参照の循環があると終了しなくなる.

```go
type Cycle struct{Value int; Tail *Cycle}
var c Cycle
c = Cycle{42, &c}
Display("c", c)
```

循環に対して頑強な再帰処理を実装するのは難しい（辿ってきた参照の集まりを記録する必要がある）. 13.3 節の unsafe が使える. らしい.

fmt.Sprint は, 自身を含むスライスやマップを受け取らない限りは, あまり問題にならない出力を行う. 例えばポインタならアドレスだけ表示してそこで打ち切る.

### 12.4. Example: Encoding S-Expressions
S 式：Lisp をはじめとして広く使われているが方言だらけで標準化には至っていない.

Display の処理は Go オブジェクトの marshal (json への変換など) と大して変わらない.

12.4 （出力にインデントを追加するプリティプリンタの作成）はやりがいのある練習問題らしい..

- 練習問題 12.7 : ch4/github を参照のこと

### 12.5. Setting Variables with reflect.Value
TODO

### 12.6. Example: Decoding S-Expressions
TODO

### 12.7. Accessing Struct Field Tags
TODO

### 12.8. Displaying the Methods of a Type
TODO

### 12.9. A Word of Caution
TODO
