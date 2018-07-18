# gomaphore
gomaphore is simple lock-free semaphore library written in golang

## Requirement

Go (>=1.8)

## Installation

```shell
go get github.com/hlts2/gomaphore
```

## Example

```go
wg := new(sync.WaitGroup)

for i := 0; i < size; i++ {
    wg.Add(1)

    go func(i int) {
        defer wg.Done()

        semaphore.Wait()

        cnt++

        semaphore.Signal()
    }(i)
}

wg.Wait()

```
