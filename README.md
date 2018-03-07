# sha256
sha256 with hash difficulty.

## 哈希生成算法
原生sha256 "crypto/sha256"

## 函数功能
``` go
func HashwithDifficulty(data []byte, d int) (result [Size]byte, nonce int64)

// 使用
sum, nonce := HashwithDifficulty([]byte("hello world"), 3)
```

## 哈希有效算法
- 方法一：通过Sprintf将[]byte转换成16进制表示的string，再按位计算是否为0，fmt有反射，效率可能会低。
```go
func difficulty(hash [Size]byte, d int) bool {
    var l int
    s := fmt.Sprintf("%x", hash)
    for l = 0 ; l < len(s) ; l++ {
        if s[l] != '0' {
            break
        }  
    }
    return l >= d
}
```
- 方法二：使用[]byte直接跟16进制数对比，判断是否为0。
```go
func difficulty(hash [Size]byte, d int) bool {
    dn := d/2
    sn := d%2
    for i :=0; i < dn;i++ {
        if hash[i] != 0x00 {
        return false
        }
    }
    if sn !=0 {
        if hash[dn*2+1] > 0x0f {
        return false
        }
    }
    return true
}
```

## 结果对比
哈希数据: []byte("hello world")+[]byte(nonce)

|    method  |nonce| difficulty |timecost|result  |
| ---------- | ----| ---------- | -----  |--------|
| 1          |160  |  2         |0.007   |0085155c6a2b306dcb8387dcbd7dd6c2fbcaf5b6735e3fd58c24914c7b909c13|
| 2          |160  |  2         | 0.006  |0085155c6a2b306dcb8387dcbd7dd6c2fbcaf5b6735e3fd58c24914c7b909c13|
| 1          |6742 |  3         |0.015   |000fd9024e22437d38075ad87a7ca2649e66384ee67943a66eef482f5fe437c7|
| 2          |6742 |  3         | 0.009  |000fd9024e22437d38075ad87a7ca2649e66384ee67943a66eef482f5fe437c7|
| 1          |49967|  4         |0.059   |00004921c6f7acd81acd24a477fd29d1effeb58ba6943007e63420a6d2b0e973|
| 2          |49967|  4         | 0.025  |00004921c6f7acd81acd24a477fd29d1effeb58ba6943007e63420a6d2b0e973|

Change Log
2018/02/28 增加停止计算hash方法
```
func StopHash() bool {
	return atomic.CompareAndSwapInt32(&stop, 0, 1)
}
```