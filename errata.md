## 勘误表

由于作者能力有限，不可避免的会在书中出现各种错误。

本文记录了由读者发现的书中的一些错误，供大家参考，同时在后续可能的修订中会更正这些错误。

|  序号   | 书名  |   位置 | 内容 | 提交者 | 错误类型 | 说明 |
|  ----  | ----  | ----  | ----  | --- | ----  | ----  |
| 1  | Go语言精进之路vol1 | p84 | “无论底层的数组元素类型有多大，切片打开的窗口有多长”这一句感觉不通顺 | chenwenbo1988@outlook.com  | 表达不清 | 改为“无论底层的数组元素是什么类型，切片打开的窗口有多大” |
| 2  | Go语言精进之路vol1 | p118 | “string(b)用在map类型的key中”这句话下面的代码m[[3]string{string(b),"key1","key2"}]="value"好像有语法错误| chenwenbo1988@outlook.com  | 语法错误 | m[[3]string{string(b),"key1","key2"}]="value" 应该删除 |
| 3  | Go语言精进之路vol1 | p383 | type interface error | 大鹏(github.com/Degfy)  | 语法错误 | 应改为 type error interface  |
| 4  | Go语言精进之路vol1 | p227 | type MyInterface I和type Mystruct T | 1264644959@qq.com  | 语法错误 | 应改为 type MyInterface= I和type Mystruct = T  |
| 5  | Go语言精进之路vol1 | p36 | 图5-1中PEADME.md | Yao-Shang Tseng(https://github.com/yakushou730) | typo | 应改为 README.md  |
| 6  | Go语言精进之路vol2 | p165 | 图51-6中 非对称加密 中的“A的公钥”与“A的私钥”写反了 | 1264644959@qq.com | typo | 应将这两者调换一下  |
| 7  | Go语言精进之路vol1 | p391 | 小结中“如果可以通过错误值类型的特征进行错误检视，那么尽量使用错误行为特征检视策略；在上述两种策略无法实施的情况下，再用“哨兵”策略和错误值类型检视策略” | AVOlili(https://github.com/AVOlili) | 表达不清| 应改为：“如果可以通过错误行为特征进行错误检视，那么尽量使用错误行为特征检视策略；在上述两种策略无法实施的情况下，再用“哨兵”策略和错误值类型检视策略；”  |
| 8  | Go语言精进之路vol2(微信读书电子版) | 第47条 和 第49条，具体参见[issue6](https://github.com/bigwhite/GoProgrammingFromBeginnerToMaster/issues/6) | html字符被转义了，影响阅读体验 | AVOlili(https://github.com/AVOlili) | 编辑错误 | 等待微信读书渠道编辑去掉转义，恢复html字符原貌 |
| 9  | Go语言精进之路vol1 | p97 | chapter3/sources/map_stable_iterate.go示例代码通过range map给切片赋值可能会给大家带去误解| feng zhao (ifenng2020@gmail.com) | 表达不清| 这个例子中用切片保存是第一次map迭代的元素order。我的原意并非一定是按照1, 2, 3的顺序保存，只是要保证后续的iterate order都与第一次相同即可。只是在我的机器上第一次iterate的order恰好是 1,2,3的顺序。不过这个例子的确会给大家带去困惑。后续如果再版，会在这处做出说明 | 
| 10  | Go语言精进之路vol1 | p161 | for_range_bench_test.go的输出结果与结论有悖 | 324127863(324127863@qq.com) | 内容错误 | for range数组性能好的原因与Go编译器根据数组元素大小进行的优化有关。可以参考一下[这篇文章](https://tonybai.com/2022/03/19/for-range-vs-classic-for-loop-when-iterating-large-array)  |
| 11  | Go语言精进之路vol1 | p103 | 示例代码有误，具体参见[issue 9](https://github.com/bigwhite/GoProgrammingFromBeginnerToMaster/issues/9) | liansyyy(https://github.com/liansyyy) | typo | 应改为 p := &m[key] |
| 12  | Go语言精进之路vol1 | p350 | “如何s[n]T或\*[n]T的数组类型，len(s)返回数组的长度n” | bin4tre(bin4re@foxmail.com) | typo | 应将“如何”改为 “如果” |
| 13  | Go语言精进之路vol1 | p352 | trySend函数的default分支的返回语句拼写错误"etrun false" | bin4tre(bin4re@foxmail.com) | typo | 应改为 “return false” |
| 14  | Go语言精进之路vol1 | p62, 具体参见[issue12](https://github.com/bigwhite/GoProgrammingFromBeginnerToMaster/issues/12) | 图8-1 变量声明语法错误 | https://github.com/jackbai233 | 内容错误 | 图中a :=(int32)17应改为a := int32(17) 另一个var a = (int32)17应改为 var a = int32(17) |
| 15  | Go语言精进之路vol1 | p89, 具体参见[issue13](https://github.com/bigwhite/GoProgrammingFromBeginnerToMaster/issues/13) | slice_unbind_orig_array.go代码中数组u初始化有笔误 | https://github.com/jackbai233 | typo | 书中代码u := []int{11, 12, 13, 14, 15} 应该为u := [...]int{11, 12, 13, 14, 15}|
| 16  | Go语言精进之路vol1 | p94, 具体参见[issue14](https://github.com/bigwhite/GoProgrammingFromBeginnerToMaster/issues/14) | map变量初始化有误 | https://github.com/jackbai233 | typo | m := map[string]int应该为m := map[string]int{}|
| 17  | Go语言精进之路vol1 | p81, 具体参见[issue16](https://github.com/bigwhite/GoProgrammingFromBeginnerToMaster/issues/16) | Go 不存在 float 类型 | https://github.com/XQ-Gang | typo | fnumbers := [...]float{}应改为fnumbers := [...]float64{}，上面的注释中代码亦是|
| 18  | Go语言精进之路vol1 | p103, 具体参见[issue17](https://github.com/bigwhite/GoProgrammingFromBeginnerToMaster/issues/17) | 关于Go不允许获取map中value的地址的示例代码错误| https://github.com/bravility | typo | p := m[key] 应该改为 p := &m[key]|
| 19  | Go语言精进之路vol1 | p49, 具体参见[issue20](https://github.com/bigwhite/GoProgrammingFromBeginnerToMaster/issues/20) |  "小骆峰拼写法"的说法不常见| https://github.com/suica | typo | “小骆峰拼写法”应改为“小驼峰拼写法”|
| 20  | Go语言精进之路vol1 | p29, 具体参见[issue21](https://github.com/bigwhite/GoProgrammingFromBeginnerToMaster/issues/20) |  "传递归递给下去"不通顺| https://github.com/suica | typo | 应改为“递归传递下去”|
