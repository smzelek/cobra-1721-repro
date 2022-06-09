# cobra-1721-repro

`go run main.go level-2 level-4` prints 

```
invalid argument "level-4" for "top-level level-2"

Did you mean this?
        level-3
```

During the evaluation of `OnlyValidArgs`, 
```go
...
for _, v := range args {
    if !stringInSlice(v, validArgs) {
        return fmt.Errorf("invalid argument %q for %q%s", v, cmd.CommandPath(), findSuggestions(cmd, args[0]))
    }
}
...
```

```
args = ["level-4"]
args[0] = "level-4"
v = "level-4"
```

