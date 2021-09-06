# Proto FileDescriptor Demo

A little demo to show how `.proto` files are encoded/decoded with
FileDescriptors.

`make protos` generates `hello.pb` from `hello.proto`. `hello.pb` is the proto
binary encoded FileDescriptor of `hello.proto`.

`make run` reads `hello.pb` into a FileDescriptor struct and outputs it as JSON:

```json
{
  "file": [
    {
      "name": "hello.proto",
      "package": "echo",
      "messageType": [
        {
          "name": "HelloRequest",
          "field": [
            {
              "name": "message",
              "number": 1,
              "label": "LABEL_OPTIONAL",
              "type": "TYPE_STRING",
              "jsonName": "message"
            }
          ]
        },
        {
          "name": "HelloResponse",
          "field": [
            {
              "name": "response",
              "number": 1,
              "label": "LABEL_OPTIONAL",
              "type": "TYPE_STRING",
              "jsonName": "response"
            }
          ]
        }
      ],
      "service": [
        {
          "name": "HelloService",
          "method": [
            {
              "name": "Hello",
              "inputType": ".echo.HelloRequest",
              "outputType": ".echo.HelloResponse"
            }
          ]
        }
      ],
      "syntax": "proto3"
    }
  ]
}
```
