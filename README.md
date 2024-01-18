# jsonc-to-json (go)

A `go` implementation of [strip-json-comments-cli](https://github.com/sindresorhus/strip-json-comments-cli) (NodeJS).

`jsonc-to-json` converts JSONC to JSON and is released as a single static binary.

## Features

* single static binary
* read JSONC input from a file or stdin/pipe
* output JSON result to stdout
* (*optional*): validate the JSON result
* (*optional*): pretty print the JSON result

## Arguments & Flags

|||
|:------------------|--------------------------------------------------------|
| `<no argument>`   | print usage instructions                               |
| `-file "${path}"` | read input from given file, use `-` to read from stdin |
| `-validate`       | validate JSON result                                   |
| `-pretty`         | pretty print JSON result                               |

## Development

<details>
<summary>Instructions</summary>

```bash
# print usage
go run main.go

# convert a JSONC file to JSON
go run main.go -file "${json_file_path}"

# read JSONC body from stdin
echo -e '{"foo":"bar"\n//foo\n}' | go run main.go -file -

# format the source code
gofmt -w main.go

# build the binary
# -w disables DWARF debugging information generation (debug_info)
# -s strip / omit the symbol table
# -trimpath - see: https://go.dev/doc/go1.13#go-command
go build -ldflags "-extldflags '-static' -s -w" -o ./strip-json-comments
# build with external linker i.e. gcc
go build -ldflags "-linkmode 'external' -extldflags '-static'" -o ./strip-json-comments
```

</details>
