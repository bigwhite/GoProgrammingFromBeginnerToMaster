## 勘误表

由于作者能力有限，不可避免的会在书中出现各种错误。

本文记录了由读者发现的书中的一些错误，供大家参考，同时在后续可能的修订中会更正这些错误。

|  序号   | 书名  |   位置 | 内容 | 提交者 | 错误类型 | 说明 |
|  ----  | ----  | ----  | ----  | --- | ----  | ----  |
| 1  | Go语言精进之路vol1 | p84 | “无论底层的数组元素类型有多大，切片打开的窗口有多长”这一句感觉不通顺 | chenwenbo1988@outlook.com  | 表达不清 | 改为“无论底层的数组元素是什么类型，切片打开的窗口有多大” |
| 2  | Go语言精进之路vol1 | p118 | “string(b)用在map类型的key中”这句话下面的代码m[[3]string{string(b),"key1","key2"}]="value"好像有语法错误| chenwenbo1988@outlook.com  | 语法错误 | m[[3]string{string(b),"key1","key2"}]="value" 应该删除 |
| 3  | Go语言精进之路vol1 | p383 | type interface error | 大鹏(github.com/Degfy)  | 语法错误 | 应改为 type error interface  |


