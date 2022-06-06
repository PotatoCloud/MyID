# MyID

正如描述中所说，他是一个很好的ID生成服务

### 我该怎么通过gRPC去调用MyID?

不要担心，我们已经为你准备好了一个请求示例
```http request
GRPC YOUR_HOST/schemes.MyID/GenerateID

{
  "DC": 1,
  "worker_id": 2000,
  "request_id": 1
}
```

此处使用 **JSON** 以使`payload`看起来更直观

- 如`payload`所示，完成这个请求需要三个参数，分别是：`DC`、`worker_id`、`request_id`
- 其中，`request_id`是最重要的，它可以帮助你确认一个**response**对应哪个**request**

#### 数据类型
```protobuf
message IDRequest {
  uint32 DC = 1;
  uint64 worker_id = 2;
  uint64 request_id = 3;
}

message IDReply {
  int64 id = 1;
  uint64 timestamp = 2;
  uint64 reply_id = 3;
}
```

#### 接口
```protobuf
service MyID {
  rpc GenerateID(IDRequest) returns (IDReply) {}
}
```