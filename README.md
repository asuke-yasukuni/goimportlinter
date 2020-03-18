# goimportlinter

warn if newline count in imports is greater than or equal to the number passed to -n

## Insall

```
go get -u github.com/asuke-yasukuni/goimportlinter
```

## How to use

```
$ goimportlinter sample.go -n 2

sample.go
フォーマットが不正のようです
imports (
  "flag"
  "fmt"
  "go/ast"
  "go/parser"

  "go/token"
  "io/ioutil"
-
  "log"
  "os"
  "path/filepath"
 )
```
