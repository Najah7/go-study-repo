# Goroutine

## goroutineとは
goroutineはGo言語のランタイムによって管理される軽量なスレッド。

## goroutineの作り方
goroutineは`go`キーワードを使って作る。

```go
go 関数名()
```

## sync.WaitGroup
- goroutineが終了するまで待つためには、`sync.WaitGroup`を使う。
- wg.Add(1)で待つgoroutineの数を増やす。
- wg.Done()で待つgoroutineの数を減らす。

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        defer wg.Done()
        fmt.Println("Hello, World!")
    }()
    wg.Wait()
}
```

## チャンネル
- goroutine間でデータをやり取りするためには、チャンネルを使う。
- チャンネルは`make()`関数で作る。
- チャンネルは`chan`キーワードで宣言する。
- チャンネルは`<-`演算子でデータを送受信する。
- チャンネルは`close()`関数で閉じる。
- チャンネルはキューのように動作する。

```go
package main

import (
    "fmt"
)

func goroutine1(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum
}

func goroutine2(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v * v
    }
    c <- sum
}

func main() {
    s := []int{1, 2, 3, 4, 5}

    c1 := make(chan int)
    c2 := make(chan int)

    go goroutine1(s, c1)
    go goroutine2(s, c2)

    x := <-c1
    y := <-c2

    fmt.Println(x, y)
}
```

## Buffered Channels
- バッファ付きチャンネルは`make()`関数の第二引数にバッファのサイズを指定する。
- バッファ付きチャンネルはバッファが埋まるまでデータを送信できる。

```go
package main

import (
    "fmt"
)

func main() {
    c := make(chan int, 2)
    c <- 1
    c <- 2
    fmt.Println(<-c)
    fmt.Println(<-c)
}
```

## チャンネルのclose
- チャンネルは`close()`関数で閉じる。
- チャンネルが閉じられたかどうかは`ok`で判定する。
- channelはcloseしないとデッドロックになる。

```go
package main

import (
    "fmt"
)

func main() {
    c := make(chan int, 2)
    c <- 1
    c <- 2
    close(c)
    for i := range c {
        fmt.Println(i)
    }
}
```
## チャンネルとループ
- チャンネルは`range`でループできる。
- チャンネルが閉じられるまでループする。なので、事前にcolseは必須。
- チャンネルの型の書き方
    - 受信専用: `chan<- int`
    - 送信専用: `<-chan int`

```go
package main

import (
    "fmt"
)

func main() {
    c := make(chan int, 2)
    c <- 1
    c <- 2
    close(c)
    for i := range c {
        fmt.Println(i)
    }
}
```

## producerとconsumer
- producerはチャンネルにデータを送信する。
- consumerはチャンネルからデータを受信する。
- producerとconsumerは別のgoroutineで動作する。

```go
package main

import (
    "fmt"
    "sync"
)

func producer(c chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 10; i++ {
        c <- i
    }
    close(c)
}

func consumer(c chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := range c {
        fmt.Println(i)
    }
}

func main() {
    var wg sync.WaitGroup
    c := make(chan int)
    for i := 0; i < 2; i++ {
        wg.Add(1)
        go consumer(c, &wg)
    }

    go producer(c, &wg)
    wg.Wait()
}
```

## fun-out, fan-in
- fun-out: 複数のgoroutineで処理を分散すること。
- fan-in: 複数のgoroutineで処理した結果を集約すること。

```go
package main

import (
    "fmt"
    "sync"
)

func producer(nums []int, c chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for _, n := range nums {
        c <- n
    }
    close(c)
}

func square(c chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for n := range c {
        fmt.Println(n * n)
    }
}

func main() {
    var wg sync.WaitGroup
    c := make(chan int)
    nums := []int{1, 2, 3, 4, 5}
    wg.Add(1)
    go producer(nums, c, &wg)
    for i := 0; i < 2; i++ {
        wg.Add(1)
        go square(c, &wg)
    }
    wg.Wait()
}
```

## select
- selectは複数のチャンネルを待ち受ける。
- selectは最初に受信したチャンネルの処理を実行する。
- selectはdefault節を持つことができる。
- selectはforループの中で利用する。

```go
package main

import (
    "fmt"
    "time"
)

func goroutine1(ch chan string) {
    for {
        ch <- "packet from 1"
        time.Sleep(3 * time.Second)
    }
}

func goroutine2(ch chan string) {
    for {
        ch <- "packet from 2"
        time.Sleep(1 * time.Second)
    }
}

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    go goroutine1(ch1)
    go goroutine2(ch2)
    for {
        select {
        case msg1 := <-ch1:
            fmt.Println(msg1)
        case msg2 := <-ch2:
            fmt.Println(msg2)
        }
    }
}
```

## default selection
- selectはdefault節を持つことができる。
- default節はどのチャンネルも受信できない場合に実行される。

```go
package main

import (
    "fmt"
    "time"
)

func goroutine1(ch chan string) {
    for {
        ch <- "packet from 1"
        time.Sleep(3 * time.Second)
    }
}

func goroutine2(ch chan string) {
    for {
        ch <- "packet from 2"
        time.Sleep(1 * time.Second)
    }
}

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    go goroutine1(ch1)
    go goroutine2(ch2)
    for {
        select {
        case msg1 := <-ch1:
            fmt.Println(msg1)
        case msg2 := <-ch2:
            fmt.Println(msg2)
        default:
            fmt.Println("default")
        }
    }
}
```

### 公式サイトのサンプル
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

```

## selectとfor break

```go
package main

import (
    "fmt"
    "time"
)

func goroutine1(ch chan string) {
    for {
        ch <- "packet from 1"
        time.Sleep(3 * time.Second)
    }
}

func goroutine2(ch chan string) {
    for {
        ch <- "packet from 2"
        time.Sleep(1 * time.Second)
    }
}

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    go goroutine1(ch1)
    go goroutine2(ch2)
    OuterLoop:
        for {
            select {
            case msg1 := <-ch1:
                fmt.Println(msg1)
            case msg2 := <-ch2:
                fmt.Println(msg2)
                break OuterLoop
            default:
                fmt.Println("default")
            }
        }
}
```

## sync.Mutex
- sync.Mutexは排他制御を実現するための構造体。
- sync.MutexはLock()とUnlock()でロックとアンロックを行う。

```go
package main

import (
    "fmt"
    "sync"
)

type Counter struct {
    mu sync.Mutex
    x  int
}

func (c *Counter) Inc() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.x++
}

func (c *Counter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.x
}

func main() {
    var counter Counter
    for i := 0; i < 1000; i++ {
        go func() {
            counter.Inc()
        }()
    }
    fmt.Println(counter.Value())
}
```