# 开发文档

### API URL: https://api.msg.launch.im/message

### Request Method: POST

### Request Parameter:
Field|Variable Name|Required|Type|Sample Value|Description
---|---|:---:|:---:|:---:|---
Push ID|push_id|True|string(6)|A1b2CZ|Obtain from App application configuration information
Nonce String|nonce|True|string(16)|0123456789abcdef|A-Za-z0-9
Timestamp|timestamp|True|int64|1620761112|Unix timestamp (seconds), valid within 1 minute of the current time
Signature|sign|True|string(64)| |sha256 check value, see the signature generation method below
Message to send|message|True|string(4000)| |json string, see message parameters for fields

### Message Parameter:
Field Name|Variable Name|Required|Type|Sample Value|Description
---|---|:---:|:---:|:---:|---
Title|title|True|string(100)|Memory Warning|
Type|msg_type|True|int(8)|0|5 types in total, 0-5
Content|content|True|string(4000)|Custom content|
Group|group|False|string(20)| developer group |Optional

### Signature generation method (sign field in request parameters):
**Generate initial signature string**

Sort the variable names in the request parameters according to the ASCII code from smallest to largest (dictionary order), and use the URL key-value pair format to splice into a string (connector &), if the parameter value is empty, it will not participate in the signature.

Suppose the request parameters are as follows:

Variable Name|Sample Value
---|---
push_id|A1b2Cz
nonce|0123456789abcdef
timestamp|1620761112
message|{"title": "test title", "msg_type": 0, "content": "test content", "group": "group name"}


#### Sort variable names and generate initial string in key=value format
```
firstString = "message={"title": "test title", "msg_type": 0, "content": "test content", "group": "group name"}&nonce=0123456789abcdef&push_id=A1b2CZ& timestamp=1620761112
```
#### Splicing secret key (check in App configuration information) Put secret=my secret value and splice it at the end of the initial string
```
secondString = "message={"title": "test title", "msg_type": 0, "content": "test content", "group": "group name"}&nonce=0123456789abcdef&push_id=A1b2CZ&timestamp=1620761112&secret=my secret value"
```

#### Sha256 signature on secondString
signString = sha256(secondString) // Get the signature value: c6c7c8c5f056360bd42343f4fa99905c823c4d3c482ca900c0b87df053af9b89

### The final parameters sent are as follows:

Variable Name|Sample Value
---|---
push_id|A1b2Cz
nonce|0123456789abcdef
timestamp|1620761112
message|{"title": "test title", "msg_type": 0, "content": "test content", "group": "group name"}
sign|c6c7c8c5f056360bd42343f4fa99905c823c4d3c482ca900c0b87df053af9b89

#### Json serialize the finally sent request parameters, and send it through the http post request to be OK。

### Response
http status code|Response Content(json string)
---|---
4xx|{"code": 4xx, "error": "error message detail"}
200|{"code": 200, "message": "success"}

### Remarks:
1. http responds with 429, the request is rate limited, this request will be discarded and no longer processed. The rate is up to 3 times within 1 minute. After the rate limit, you can wait 1 minute to send the request again.
0. The server does not save any message content, please do a good job of message backup management.
