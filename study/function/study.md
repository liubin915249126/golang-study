## go 函数

##### 可变参数
```go
    func sum(vals ...int) int {
        total := 0
        for _, val := range vals {
            total += val
        }
        return total
    }
```

#### defer
- 执行顺序与声名顺序相反
- defer语句中的函数会在return语句更新返回值变量后再执行，
- 又因为在函数中定义的匿名函数可以访问该函数包括返回值变量在内的所有变量，所以，对匿名函数采用defer机制，可以使其观察函数的返回值。
func double(x int) (result int) {
    defer func() { fmt.Printf("double(%d) = %d\n", x,result) }()
    return x + x
}
_ = double(4)
- defer表达式中变量的值在defer表达式被定义时就已经明确