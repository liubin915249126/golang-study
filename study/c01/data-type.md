## go数据类型

#### 数组
固定长度，每个数组元素都是完全相同的类型
```go
    var m [3]int = [3]int{1, 2, 3}
    // 改变作为参数的数组
    func ArrIsArgs(arr *[4]int) {
      arr[0] = 20
    }
    m:= [...]int{1, 2, 3, 4}
    ArrIsArgs(&m)
    // & 是取地址符号 , 即取得某个变量的地址 ,如: &a
    // * 是指针运算符 , 可以表示一个变量是指针类型 , 也可以表示一个指针变量所指向的存储单元 , 也就是这个地址所存储的值.
```

#### slice
切片 代表变长的序列,序列中每个元素都有相同的类型
```go
// 切片
// 其中0 ≤ i≤ j≤ cap(s),用于创建一个新的slice,引用s的从第i个元素开始到第j-1个元素的子序列。新的slice将只有j-i个元素。如果i位置的索引被省略 的话将使用0代替,如果j位置的索引被省略的话将使用len(s)代替。 如果切片操作超出cap(s)的上限将导致一个panic异常,但是超出len(s)则是意味着扩展了lice,因为新slice的长度会变大. 
   s[i:j]
```
```go
//    基于数组创建
   arrVar := [4]int{1, 2, 3，4}
   sliceVar := arrVar[1:3]
//    直接创建
   var p *[]int = new([]int)  //分配slice结构内存
   var ｍ []int = make([]int,100) //m指向一个新分配的有100个整数的数组
    slice1 := make([]int,5)//创建一个元素个数5的slice,cap也是5
    slice2 := make([]int,5,10)//创建一个元素个数5的slice，cap是10
    slice3 := []int{1,2,3,4,5}//创建一个元素个数为5的slice，cap是5
    var slice []int //创建一个空的slice，cap和len都是0
//    equal
   func equal(x, y []string) bool {
    if len(x) != len(y) {
            return false
    }
    for i := range x {
            if x[i] != y[i] {
                return false
            }
    }
    return true
    }
```
#### Map 
映射 map[K]V  map中所有的key都有相同的类型,所有的value也有着相同的类型,但是key和value之间可以是不同的数据类型
```go
   var m map[string] string
//    创建
   m := make(map[string]int) 
   m := map[string]int{
    "keke": 001,
    "jame": 002,
   }
```
#### 结构体
```js
   type Rectangle struct {
      width float64
      length float64
    }
```