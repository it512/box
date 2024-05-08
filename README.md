# boxCtx


使用过多的context.WithValue会导致context 层数过多，从而导致内存占用过多和降低效率，采用此工具可以有效的降低context的层级


## 用法

创建和加入

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
box.WithValue 的返回值遵循以下规则
1. 首先在parent中查找是否存在boxCtx，如果存在，则增加key，val至boxCtx，返回**parent**
2. 如果parent中不存在boxCtx，则创建新的boxCtx，在其增加key，value，返回**新建的boxCtx**

``` golang
func WithValue(parent context.Context, key, val any) context.Context {
	if ctx := parent.Value(self); ctx != nil {
		b := ctx.(*boxCtx)
		b.put(key, val)
		return parent
	}

	b := new(parent)
	b.put(key, val)
	return b
}
```

获取

可以使用context.Value直接获取，box 也提供了一个范型的From方法,用于方便获取.  
``` golang
x:=box.From[string](bCtx, "x")
```

**注意**
如果使用box.From获取元素，当value为nil时，会导致panic

