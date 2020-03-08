### 关键信息

1. 缓存共享
2. 内存对齐


#### 缓存共享

**多核CPU多级缓存/一致性协议MESI**

    缓存行（Cache line）:缓存存储数据的单元，一般为64 byte。

主存（main memory） --> 高速缓存（cache） --> CPU

由于在多核的场景下，不同核心有自己的缓存，就需要MESI协议来同步不同核心缓存行的状态；

|状态|描述|监听任务|
|:--|:--|:--|
|M|修改 (Modified) |该Cache line有效，数据被修改了，和内存中的数据不一致，数据只存在于本Cache中。 	缓存行必须时刻监听所有试图读该缓存行相对就主存的操作，这种操作必须在缓存将该缓存行写回主存并将状态变成S（共享）状态之前被延迟执行。|
|E|独享、互斥 (Exclusive) |该Cache line有效，数据和内存中的数据一致，数据只存在于本Cache中。 	缓存行也必须监听其它缓存读主存中该缓存行的操作，一旦有这种操作，该缓存行需要变成S（共享）状态。|
|S|共享 (Shared) |该Cache line有效，数据和内存中的数据一致，数据存在于很多Cache中。 	缓存行也必须监听其它缓存使该缓存行无效或者独享该缓存行的请求，并将该缓存行变成无效（Invalid）。|
|I|无效 (Invalid) |该Cache line无效。|



**并发场景下的缓存共享问题**

真共享：不同核心需要访问和修改同一缓存行，数据就需要在不同的缓存间进行同步，导致性能问题；
假共享：在真共享的场景中，特定的缓存行可能包含了不同的数据，而实际上不同核心访问的正式这不同的数据，它们之间实际上并不存在共享，但同样带来了同步的性能问题；

代码能够解决的正是“假共享”的问题，通过合理的设计数据结构，来避免“假共享”问题的发生；

```go
// // go test -test.bench=".*" -test.cpu=1,2,4,8 ./cache_test.go > cache.bench

goos: windows
goarch: amd64
BenchmarkPad       	89478446	        13.5 ns/op
BenchmarkPad-2     	35056761	        32.4 ns/op
BenchmarkPad-4     	41921542	        25.6 ns/op
BenchmarkPad-8     	40855096	        26.8 ns/op
BenchmarkNoPad     	90115528	        13.3 ns/op
BenchmarkNoPad-2   	24092563	        49.6 ns/op
BenchmarkNoPad-4   	20330884	        59.2 ns/op
BenchmarkNoPad-8   	20054715	        59.8 ns/op
PASS
ok  	command-line-arguments	11.402s

```


#### 字节对齐

是为了解决主存加载到缓存的性能问题：

1.主存的读取是按块进行的（内存访问粒度），一般为2,4,8,16...
2.数据的大小也是不确定的（2^n）

假设内存访问粒度为8byte,而数据的大小是12byte(结构如下),假设地址都从0开始：

```go
type T struct {
	a uint32
	b uint64
}
```

**不进行对齐**

1.在访问T.b字段时，需要访问主存两次(0~7,8~15)

**进行对齐**

1.在访问T.b字段时，只需要访问主存一次(8~15)


但是，我们同样也发现，如果对结构体进行以下调整，就不存在对齐的问题：

```go
type T struct {
	b uint64
	a uint32
}
```

所以,在会字节对齐的前提下，调整字段顺序也是一种内存优化的方案；

在golang中，数据结构对齐的长度可通过`unsafe.Alignof`来获取,针对不同的类型有不同的对齐长度：

```go
func (s *StdSizes) Alignof(T Type) int64 {
	// For arrays and structs, alignment is defined in terms
	// of alignment of the elements and fields, respectively.
	switch t := T.Underlying().(type) {
	case *Array:
		// spec: "For a variable x of array type: unsafe.Alignof(x)
		// is the same as unsafe.Alignof(x[0]), but at least 1."
		return s.Alignof(t.elem)
	case *Struct:
		// spec: "For a variable x of struct type: unsafe.Alignof(x)
		// is the largest of the values unsafe.Alignof(x.f) for each
		// field f of x, but at least 1."
		max := int64(1)
		for _, f := range t.fields {
			if a := s.Alignof(f.typ); a > max {
				max = a
			}
		}
		return max
	case *Slice, *Interface:
		// Multiword data structures are effectively structs
		// in which each element has size WordSize.
		return s.WordSize
	case *Basic:
		// Strings are like slices and interfaces.
		if t.Info()&IsString != 0 {
			return s.WordSize
		}
	}
	a := s.Sizeof(T) // may be 0
	// spec: "For a variable x of any type: unsafe.Alignof(x) is at least 1."
	if a < 1 {
		return 1
	}
	// complex{64,128} are aligned like [2]float{32,64}.
	if isComplex(T) {
		a /= 2
	}
	if a > s.MaxAlign {
		return s.MaxAlign
	}
	return a
}
```


### 参考

https://changkun.de/golang/zh-cn/part1basic/ch02parallel/cache/
https://developer.ibm.com/articles/pa-dalign/
https://wizardforcel.gitbooks.io/gopl-zh/content/ch13/ch13-01.html