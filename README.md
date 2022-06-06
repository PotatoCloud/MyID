# MyID

As said in the description, he is a good id generation service

### How to call the gRPC interface in MyID?

Don't worry, we provide an example 😂

```http request
GRPC YOUR_HOST/schemes.MyID/GenerateID

{
  "DC": 1,
  "worker_id": 2000,
  "request_id": 1
}
```

**JSON** is used here to make the `payload` look more intuitive

- As shown in the `payload`, you need three parameters to complete this request, they are: `DC`, `worker_id`, `request_id`
- Among them, `request_id` is the most important, it can help you confirm which **request** a **response** corresponds to

#### type of data
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

#### interface
```protobuf
service MyID {
  rpc GenerateID(IDRequest) returns (IDReply) {}
}
```