---
title: Netools-API v1.0.0
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.5"

---

# Netools-API

> v1.0.0

# 文本

## POST 添加文本

POST /text/add

> Body 请求参数

```yaml
content: test

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» content|body|string| 是 |none|

> 返回示例

> 成功

```json
{
  "retcode": 0,
  "message": "success",
  "data": {
    "uuid": "29a05309-00d8-48e1-aa5b-60a910713253",
    "timestamp": 1648214600,
    "content": "test"
  },
  "success": true
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» retcode|integer|true|none|none|
|» message|string|true|none|none|
|» data|object|true|none|none|
|»» uuid|string|true|none|none|
|»» timestamp|integer|true|none|none|
|»» content|string|true|none|none|
|» success|boolean|true|none|none|

## GET 获取全部文本

GET /text/all

> 返回示例

> 成功

```json
{
  "retcode": 0,
  "message": "success",
  "data": [
    {
      "uuid": "9d587a1b-2453-4752-a483-22444f9a4946",
      "timestamp": 1648214658,
      "content": "test"
    },
    {
      "uuid": "29a05309-00d8-48e1-aa5b-60a910713253",
      "timestamp": 1648214600,
      "content": "test"
    }
  ],
  "success": true
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» retcode|integer|true|none|none|
|» message|string|true|none|none|
|» data|[object]¦null|true|none|none|
|»» uuid|string|true|none|none|
|»» timestamp|integer|true|none|none|
|»» content|string|true|none|none|
|» success|boolean|true|none|none|

# 文件

## GET 获取全部文件

GET /file/all

> 返回示例

> 成功

```json
{
  "retcode": 0,
  "message": "success",
  "data": [
    {
      "uuid": "4ecd2823-d6f7-452c-a341-ca3c1530048d",
      "timestamp": 1648214996,
      "filename": "1.txt",
      "size": 7
    }
  ],
  "success": true
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» retcode|integer|true|none|none|
|» message|string|true|none|none|
|» data|[object]¦null|true|none|none|
|»» uuid|string|false|none|none|
|»» timestamp|integer|false|none|none|
|»» filename|string|false|none|none|
|»» size|integer|false|none|none|
|» success|boolean|true|none|none|

## GET 下载文件

GET /file/download

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|uuid|query|string| 是 |none|

> 返回示例

> 成功

```json
12345
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|number|

## POST 文件上传

POST /file/upload

> Body 请求参数

```yaml
file: file://C:\Users\PassbyA\Desktop\1.txt

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» file|body|string| 是 |none|

> 返回示例

> 成功

```json
{
  "retcode": 0,
  "message": "success",
  "data": {
    "uuid": "4ecd2823-d6f7-452c-a341-ca3c1530048d",
    "timestamp": 1648214996,
    "filename": "1.txt",
    "size": 7
  },
  "success": true
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» retcode|integer|true|none|none|
|» message|string|true|none|none|
|» data|object|true|none|none|
|»» uuid|string|true|none|none|
|»» timestamp|integer|true|none|none|
|»» filename|string|true|none|none|
|»» size|integer|true|none|none|
|» success|boolean|true|none|none|

# 书签

## POST 新添书签

POST /bookmark/create

> Body 请求参数

```yaml
name: bookmark name
content: bookmark content
type: bookmark type

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» name|body|string| 是 |none|
|» content|body|string| 是 |none|
|» type|body|string| 是 |none|

> 返回示例

> 成功

```json
{
  "retcode": 0,
  "message": "success",
  "data": null,
  "success": true
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» retcode|integer|true|none|none|
|» message|string|true|none|none|
|» data|null|true|none|none|
|» success|boolean|true|none|none|

## GET 删除书签

GET /api/bookmark/delete

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|query|string| 是 |none|

> 返回示例

> 成功

```json
{
  "retcode": 0,
  "message": "success",
  "data": null,
  "success": true
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» retcode|integer|true|none|none|
|» message|string|true|none|none|
|» data|null|true|none|none|
|» success|boolean|true|none|none|

## GET 获取全部书签

GET /api/bookmark/all

> 返回示例

> 成功

```json
{
  "retcode": 0,
  "message": "success",
  "data": [
    {
      "ID": 2,
      "CreatedAt": "2022-03-22T20:36:46.2083141+08:00",
      "UpdatedAt": "2022-03-22T20:36:46.2083141+08:00",
      "DeletedAt": null,
      "name": "test2",
      "content": "test test test 123 测试",
      "type": "其他"
    }
  ],
  "success": true
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» retcode|integer|true|none|none|
|» message|string|true|none|none|
|» data|[object]¦null|true|none|none|
|»» ID|integer|false|none|none|
|»» CreatedAt|string|false|none|none|
|»» UpdatedAt|string|false|none|none|
|»» DeletedAt|null|false|none|none|
|»» name|string|false|none|none|
|»» content|string|false|none|none|
|»» type|string|false|none|none|
|» success|boolean|true|none|none|

## GET 通过ID获取书签

GET /bookmark/by

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|query|string| 是 |none|

> 返回示例

> 成功

```json
{
  "retcode": 0,
  "message": "success",
  "data": {
    "ID": 2,
    "CreatedAt": "2022-03-22T20:36:46.2083141+08:00",
    "UpdatedAt": "2022-03-22T20:36:46.2083141+08:00",
    "DeletedAt": null,
    "name": "test2",
    "content": "test test test 123 测试",
    "type": "其他"
  },
  "success": true
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» retcode|integer|true|none|none|
|» message|string|true|none|none|
|» data|object|true|none|none|
|»» ID|integer|true|none|none|
|»» CreatedAt|string|true|none|none|
|»» UpdatedAt|string|true|none|none|
|»» DeletedAt|null|true|none|none|
|»» name|string|true|none|none|
|»» content|string|true|none|none|
|»» type|string|true|none|none|
|» success|boolean|true|none|none|

# 数据模型

<h2 id="tocS_File">File</h2>

<a id="schemafile"></a>
<a id="schema_File"></a>
<a id="tocSfile"></a>
<a id="tocsfile"></a>

```json
{
  "uuid": "string",
  "timestamp": 0,
  "filename": "string",
  "size": 0
}

```

### 属性

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|uuid|string|true|none|none|
|timestamp|number|true|none|none|
|filename|string|true|none|none|
|size|number|true|none|kb|

<h2 id="tocS_Bookmark">Bookmark</h2>

<a id="schemabookmark"></a>
<a id="schema_Bookmark"></a>
<a id="tocSbookmark"></a>
<a id="tocsbookmark"></a>

```json
{
  "ID": 0,
  "CreatedAt": "string",
  "UpdatedAt": "string",
  "DeletedAt": "string",
  "name": "string",
  "content": "string",
  "type": "string"
}

```

### 属性

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|ID|number|true|none|none|
|CreatedAt|string|true|none|none|
|UpdatedAt|string|true|none|none|
|DeletedAt|string|true|none|none|
|name|string|true|none|none|
|content|string|true|none|none|
|type|string|true|none|none|

<h2 id="tocS_Text">Text</h2>

<a id="schematext"></a>
<a id="schema_Text"></a>
<a id="tocStext"></a>
<a id="tocstext"></a>

```json
{
  "uuid": "string",
  "timestamp": 0,
  "content": "string"
}

```

### 属性

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|uuid|string|true|none|none|
|timestamp|number|true|none|none|
|content|string|true|none|none|

