# 开发文档

### 接口地址: ht<span>tps://api.msg.launch.im/message

### 请求方法: POST

### 请求参数:
字段名|变量名|必填|类型|示例值|描述
---|---|:---:|:---:|:---:|---
推送ID|push_id|是|string(6)|A1b2CZ|在App应用配置信息中获取
随机字符串D|nonce|是|string(16)|0123456789abcdef|A-Za-z0-9
时间戳|timestamp|是|int64|1620761112|Unix时间戳(秒), 当前时间1分钟内有效
签名|sign|是|string(64)| |sha256校验值, 参见下方的签名生成方法
要发送的消息|message|是|string(4000)| |json字符串, 字段参见消息参数

### 消息参数:
字段名|变量名|必填|类型|示例值|描述
---|---|:---:|:---:|:---:|---
主题|title|是|string(100)|内存告警|
类型|msg_type|是|int(8)|0|共5种类型, 0-5
内容|content|是|string(4000)|自定义内容|
消息组|group|否|string(20)| 开发组 |选填

### 签名生成方法(请求参数里面的sign字段):
**生成初始签名字符串**

将请求参数里面的变量名按ASCII码从小到大排序(字典序), 使用URL键值对格式拼接成字符串(连接符&), 如果参数值为空不参与签名。

假设请求参数如下:

变量名|示例值
---|---
push_id|A1b2Cz
nonce|0123456789abcdef
timestamp|1620761112
message|{"title": "test title", "msg_type": 0, "content": "test content", "group": "group name"}


#### 对变量名排序, key=value格式生成初始字符串
```
firstString = "message={"title": "test title", "msg_type": 0, "content": "test content", "group": "group name"}&nonce=0123456789abcdef&push_id=A1b2CZ& timestamp=1620761112
```
#### 拼接secret密钥(在App应用配置信息处查看), 把secret=我的secret值, 拼接在初始字符串最后
```
secondString = "message={"title": "test title", "msg_type": 0, "content": "test content", "group": "group name"}&nonce=0123456789abcdef&push_id=A1b2CZ&timestamp=1620761112&secret=我的secret值"
```

#### 对secondString进行sha256签名
signString = sha256(secondString) // 得到签名值: c6c7c8c5f056360bd42343f4fa99905c823c4d3c482ca900c0b87df053af9b89

### 最终发送的参数如下:

变量名|示例值
---|---
push_id|A1b2Cz
nonce|0123456789abcdef
timestamp|1620761112
message|{"title": "test title", "msg_type": 0, "content": "test content", "group": "group name"}
sign|c6c7c8c5f056360bd42343f4fa99905c823c4d3c482ca900c0b87df053af9b89

#### 对最终发送的请求参数进行Json序列化, 通过http post请求发送出去就OK了。

### 请求响应
http状态码|响应内容(json字符串)
---|---
4xx|{"code": 4xx, "error": "错误信息详情"}
200|{"code": 200, "message": "success"}

### 备注:
1. http响应429，请求被限速，此请求将抛弃不再处理。 速率为1分钟内最多3次，限速后等1分钟就可以再发请求了
0. 服务端不保存任何消息内容，请做好消息备份管理