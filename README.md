# lunar 阳历、农历相互转换

### Usage

```go
n := time.Now()

s := TimeToSolar(n)
printSolar(*s)

l := SolarToLunar(*s)
printLunar(*l)

```

### License

MIT License
