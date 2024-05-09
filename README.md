# boxCtx


使用过多的context.WithValue会导致context 层数过多，从而导致内存占用过多和降低效率，采用此工具可以有效的降低context的层级


## 用法

保持和标准库风格一致，box也采用WithValue方法加入元素  
``` golang
bCtx:=box.WithWalue(context.Background(), "a", "a")
```
boxCtx 也是 context.Context 所以context中的其他方法也适用，例如：  
``` golang
xCtx := context.WithValue(bCtx,"x","x")
```
或者  
``` golang
c, f := context.WithCancel(bCtx)
```

**注意**
box.WithValue 方法定义为：在已经存在的boxCtx中增加key，value，如果boxCtx不存在，则创建

``` golang
func WithValue(parent context.Context, key, val any) context.Context {
	if box, ok := From[*boxCtx](parent, self); ok {
		box.put(key, val)
		return parent
	}

	b := new(parent)
	b.put(key, val)
	return b
}
```

可以使用context.Value直接获取，box 也提供了一个范型的From方法,用于方便获取.  
``` golang
x , ok := box.From[string](bCtx, "x")
```

