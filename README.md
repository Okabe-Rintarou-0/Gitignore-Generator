# Gitignore Generator

有这个项目的原因是因为每次起项目都要写 `gitignore`，我习惯从 [github/gitignore](https://github.com/github/gitignore) 上拉对应的文件，
但是每次都这样很麻烦，所以干脆就写一个工具。

根据 [github/gitignore](https://github.com/github/gitignore) 生成对应的 `.gitignore` 文件。

## 下载

见 [releases](https://github.com/Okabe-Rintarou-0/Gitignore-Generator/releases) 。

## 食用方式

```
Usage:
  igen [flags]
  igen [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  ls          List all available languages

Flags:
  -h, --help            help for igen
  -o, --out string      output dir (default "./")
  -t, --target string   target language

Use "igen [command] --help" for more information about a command.
```

例子：

创建 `go` 语言的 `.gitignore` 文件：
```shell
igen -t go 
```