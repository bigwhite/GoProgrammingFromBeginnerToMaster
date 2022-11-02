## 勘误表

由于作者能力有限，不可避免的会在书中出现各种错误。

本文记录了由读者发现的书中的一些错误，供大家参考，同时在后续可能的修订中会更正这些错误。

下表中打~~删除线~~的序号项表示是在某一版某次印刷时已修正的勘误，具体看该勘误项的**状态**一列。

|  序号   | 书名  |   位置 | 内容 | 提交者 | 错误类型 | 说明 | 状态 |
|  ----  | ----  | ----  | ----  | --- | ----  | ----  | ---- |
| ~~1~~  | Go语言精进之路vol1 | p84 | “无论底层的数组元素类型有多大，切片打开的窗口有多长”这一句感觉不通顺 | chenwenbo1988@outlook.com  | 表达不清 | 改为“无论底层的数组元素是什么类型，切片打开的窗口有多大” | 2022年8月第一版第2次印刷已改 |
| ~~2~~  | Go语言精进之路vol1 | p118 | “string(b)用在map类型的key中”这句话下面的代码m[[3]string{string(b),"key1","key2"}]="value"好像有语法错误| chenwenbo1988@outlook.com  | 语法错误 | m[[3]string{string(b),"key1","key2"}]="value" 应该删除 | 2022年8月第一版第2次印刷已改 |
| ~~3~~  | Go语言精进之路vol1 | p383 | type interface error | 大鹏(github.com/Degfy)  | 语法错误 | 应改为 type error interface | 2022年8月第一版第2次印刷已改 |
| ~~4~~  | Go语言精进之路vol1 | p227 | type MyInterface I和type Mystruct T | 1264644959@qq.com  | 语法错误 | 应改为 type MyInterface= I和type Mystruct = T  | 2022年8月第一版第2次印刷已改 |
| ~~5~~  | Go语言精进之路vol1 | p36 | 图5-1中PEADME.md | Yao-Shang Tseng(https://github.com/yakushou730) | typo | 应改为 README.md  | 2022年8月第一版第2次印刷已改 |
| 6 | Go语言精进之路vol2 | p165 | 图51-6中 非对称加密 中的“A的公钥”与“A的私钥”写反了 | 1264644959@qq.com | typo | 应将这两者调换一下  | 未改 |
| ~~7~~  | Go语言精进之路vol1 | p391 | 小结中“如果可以通过错误值类型的特征进行错误检视，那么尽量使用错误行为特征检视策略；在上述两种策略无法实施的情况下，再用“哨兵”策略和错误值类型检视策略” | AVOlili(https://github.com/AVOlili) | 表达不清| 应改为：“如果可以通过错误行为特征进行错误检视，那么尽量使用错误行为特征检视策略；在上述两种策略无法实施的情况下，再用“哨兵”策略和错误值类型检视策略；”  | 2022年8月第一版第2次印刷已改 |
| 8  | Go语言精进之路vol2(微信读书电子版) | 第47条 和 第49条，具体参见[issue6](https://github.com/bigwhite/GoProgrammingFromBeginnerToMaster/issues/6) | html字符被转义了，影响阅读体验 | AVOlili(https://github.com/AVOlili) | 编辑错误 | 等待微信读书渠道编辑去掉转义，恢复html字符原貌 | 未改 |
| 9  | Go语言精进之路vol1 | p97 | chapter3/sources/map_stable_iterate.go示例代码通过range map给切片赋值可能会给大家带去误解| feng zhao (ifenng2020@gmail.com) | 表达不清| 这个例子中用切片保存是第一次map迭代的元素order。我的原意并非一定是按照1, 2, 3的顺序保存，只是要保证后续的iterate order都与第一次相同即可。只是在我的机器上第一次iterate的order恰好是 1,2,3的顺序。不过这个例子的确会给大家带去困惑。后续如果再版，会在这处做出说明 | 未改 |
| 10  | Go语言精进之路vol1 | p161 | for_range_bench_test.go的输出结果与结论有悖 | 324127863(324127863@qq.com) | 内容错误 | for range数组性能好的原因与Go编译器根据数组元素大小进行的优化有关。可以参考一下[这篇文章](https://tonybai.com/2022/03/19/for-range-vs-classic-for-loop-when-iterating-large-array)  | 未改 |
| ~~11~~  | Go语言精进之路vol1 | p103 | 示例代码有误，具体参见[issue 9](https://github.com/bigwhite/GoProgrammingFromBeginnerToMaster/issues/9) | liansyyy(https://github.com/liansyyy) | typo | 应改为 p := &m[key] | 2022年8月第一版第2次印刷已改 |
| ~~12~~  | Go语言精进之路vol1 | p350 | “如何s[n]T或\*[n]T的数组类型，len(s)返回数组的长度n” | bin4tre(bin4re@foxmail.com) | typo | 应将“如何”改为 “如果” | 2022年8月第一版第2次印刷已改 |
| ~~13~~  | Go语言精进之路vol1 | p352 | trySend函数的default分支的返回语句拼写错误"etrun false" | bin4tre(bin4re@foxmail.com) | typo | 应改为 “return false” | 2022年8月第一版第2次印刷已改 |
| ~~14~~  | Go语言精进之路vol1 | p62, 具体参见[issue12](https://github.com/bigwhite/GoProgrammingFromBeginnerToMaster/issues/12) | 图8-1 变量声明语法错误 | https://github.com/jackbai233 | 内容错误 | 图中a :=(int32)17应改为a := int32(17) 另一个var a = (int32)17应改为 var a = int32(17) | 2022年8月第一版第2次印刷已改 |
| ~~15~~  | Go语言精进之路vol1 | p89, 具体参见[issue13](https://github.com/bigwhite/GoProgrammingFromBeginnerToMaster/issues/13) | slice_unbind_orig_array.go代码中数组u初始化有笔误 | https://github.com/jackbai233 | typo | 书中代码u := []int{11, 12, 13, 14, 15} 应该为u := [...]int{11, 12, 13, 14, 15}| 2022年8月第一版第2次印刷已改 |
| 16  | Go语言精进之路vol1 | p94, 具体参见[issue14](https://github.com/bigwhite/GoProgrammingFromBeginnerToMaster/issues/14) | map变量初始化有误 | https://github.com/jackbai233 | typo | m := map[string]int应该为m := map[string]int{}| 未改 |
| ~~17~~  | Go语言精进之路vol1 | p81, 具体参见[issue16](https://github.com/bigwhite/GoProgrammingFromBeginnerToMaster/issues/16) | Go 不存在 float 类型 | https://github.com/XQ-Gang | typo | fnumbers := [...]float{}应改为fnumbers := [...]float64{}，上面的注释中代码亦是| 2022年8月第一版第2次印刷已改 |
| ~~18~~  | Go语言精进之路vol1 | p103, 具体参见[issue17](https://github.com/bigwhite/GoProgrammingFromBeginnerToMaster/issues/17) | 关于Go不允许获取map中value的地址的示例代码错误| https://github.com/bravility | typo | p := m[key] 应该改为 p := &m[key]| 2022年8月第一版第2次印刷已改 |
| ~~19~~  | Go语言精进之路vol1 | p49, 具体参见[issue20](https://github.com/bigwhite/GoProgrammingFromBeginnerToMaster/issues/20) |  "小骆峰拼写法"的说法不常见| https://github.com/suica | typo | “小骆峰拼写法”应改为“小驼峰拼写法”| 2022年8月第一版第2次印刷已改 |
| ~~20~~  | Go语言精进之路vol1 | p29, 具体参见[issue21](https://github.com/bigwhite/GoProgrammingFromBeginnerToMaster/issues/20) |  "传递归递给下去"不通顺| https://github.com/suica | typo | 应改为“递归传递下去”| 2022年8月第一版第2次印刷已改 |
| ~~21~~  | Go语言精进之路vol1 | p27, 具体参见[issue22](https://github.com/bigwhite/GoProgrammingFromBeginnerToMaster/issues/22) |  素数定义不准确 | https://github.com/suica | 内容错误 | 素数定义中的“具有”应改为“仅有”，“除数”改为“因数”| 2022年8月第一版第2次印刷已改 |
| ~~22~~  | Go语言精进之路vol1 | p27, p57 中的代码，具体参见[issue22](https://github.com/bigwhite/GoProgrammingFromBeginnerToMaster/issues/22) |  代码的字体或许启用了连字, 两个字符合起来仅占用了一个字符的空间，并不是等宽的，阅读起来不够美观 | https://github.com/suica | 编辑错误 | 应关闭连字，使用等宽字体| 2022年8月第一版第2次印刷已改 |
| 23  | Go语言精进之路vol2 | p66 中的第一段代码，具体参见[issue23](https://github.com/bigwhite/GoProgrammingFromBeginnerToMaster/issues/23) |  示例代码中 b.RunParallel(...) 没有与之匹配的右括号，将造成错误 | https://github.com/XQ-Gang | typo | 倒数第二行代码的}右侧应补充一个右小括号 | 未改 |
| ~~24~~  | Go语言精进之路vol1 | p34，具体参见[issue24](https://github.com/bigwhite/GoProgrammingFromBeginnerToMaster/issues/24) |  "Go 1.13版本在src下面增加了go.mod和go.num"中的go.num 应改为go.sum | https://github.com/banana42 | typo | go.num应该为go.sum | 2022年8月第一版第2次印刷已改 |
| 25  | Go语言精进之路vol2 | P229 错别字问题 具体参见[issue25](https://github.com/bigwhite/GoProgrammingFromBeginnerToMaster/issues/25) |  最下面一段，私钥可以看“出“数对，应为看“成” | https://github.com/XQ-Gang | typo | 私钥可以看出数对 应改为 私钥可以看成数对  | 未改 |
| 26  | Go语言精进之路vol1 | p131 |  "mypkg1则指代的是chapter3-demo1/pkg/pkg1下面的包"中的路径有误| asd13878288822@qq.com | typo | "mypkg1则指代的是chapter3-demo1/pkg/pkg1下面的包"应该为"mypkg1则指代的是chapter3-demo2/pkg/pkg1下面的包" | 未改 |
| 27  | Go语言精进之路vol1 | p298，具体参见[issue27](https://github.com/bigwhite/GoProgrammingFromBeginnerToMaster/issues/27) | 这里简单的用 3600/3 = 1200 作为最终的最短耗时预期是不合理的 | https://github.com/MrBear2018 | 内容错误 | “效率应该会稳定在1200(3600/3)左右” 改为 “效率应该会稳定在1800(xRayCheckTmCost * 30 / 3)左右” | 未改 |
| 28  | Go语言精进之路vol1 | p324，具体参见[issue28](https://github.com/bigwhite/GoProgrammingFromBeginnerToMaster/issues/28) | p324源码go-concurrency-pattern-5.go中的 worker goroutine在收到quit后并未立即退出 | https://github.com/ptgeft | 内容错误 | p324源码go-concurrency-pattern-5.go中的源码quit <- "ok" 的下一行应增加return | 源码库已改，但纸版书中尚未修改 |
| 29  | Go语言精进之路vol1 | p50，具体参见[issue30](https://github.com/bigwhite/GoProgrammingFromBeginnerToMaster/issues/30) | 命令中的idents找不到 | https://github.com/huangxiangrongphper | 内容错误 | 脚注中需要添加indents命令的仓库地址：https://github.com/bigwhite/go/tree/master/cmd/indents  | 源码库已改，但纸版书中尚未修改 |
| 30  | Go语言精进之路vol2 | p180 | 图52-4 中下面的“编码”有误 | 白小九 | typo | 图52-4 中下面的“编码”应该为“解码” | 未改 |
