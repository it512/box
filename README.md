# box


实用过多的context.WithValue会导致context 层数过多，从而导致内存占用过多和降低效率，采用此工具可以有效的降低context的层级


## 用法

创建和加入
``` golang
bCtx:=box.New(context.Background())
bCtx.Add("x","x")
```
boxCtx 也是 context.Context 所以context中的其他方法也适用，例如：
``` golang
xCtx := context.WithValue(bCtx,"x","x")
```
或者
``` golang
c,f:=context.WithCancel(bCtx)
```

获取

可以使用context.Value直接获取，box 也提供了一个范型的From方法,用于方便获取
``` golang
x:=box.From[string](bCtx, "x")
```

获取自身

一些情况下，需要获取boxCtx自身，并添加其他的元素，box提供了Self方法用于获取自身
``` golang
boxCtx:=box.Self(x) // 从x中获取boxCtx
```
