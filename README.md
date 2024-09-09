# boxCtx


使用过多的context.WithValue会导致context 层数过多，从而导致内存占用过多和降低效率，采用此工具可以有效的降低context的层级


## 用法

boxCtx 也是 context.Context 所以context中的其他方法也适用，例如：  
``` golang
xCtx := context.WithValue(bCtx,"x","x")
```
或者  
``` golang
c, f := context.WithCancel(bCtx)
```

可以使用context.Value直接获取，box 也提供了一个范型的From方法，用于方便获取  
``` golang
x , ok := box.From[string](bCtx, "x")
```
box.From方法兼容context.Value

## 最佳实践

尽量将boxCtx放置在context树的最上层
