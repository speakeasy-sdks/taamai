# Auth
(*Auth*)

### Available Operations

* [Register](#register) - Register
* [Login](#login) - login

## Register

Register

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/taamai/pkg/models/shared"
	"github.com/speakeasy-sdks/taamai"
	"context"
	"github.com/speakeasy-sdks/taamai/pkg/models/operations"
	"log"
)

func main() {
    s := taamai.New(
        taamai.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Auth.Register(ctx, operations.RegisterRequest{
        Country: "pakistan",
        Email: "sobanshahid47@gmail.com",
        Name: "soban",
        Password: "ali112233",
        Subdomain: "friend",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Examplewithdefaultplansubdomainandsubsctpion != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.RegisterRequest](../../pkg/models/operations/registerrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                 | :heavy_minus_sign:                                                           | The options for this request.                                                |


### Response

**[*operations.RegisterResponse](../../pkg/models/operations/registerresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Login

login

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/taamai/pkg/models/shared"
	"github.com/speakeasy-sdks/taamai"
	"context"
	"github.com/speakeasy-sdks/taamai/pkg/models/operations"
	"log"
)

func main() {
    s := taamai.New(
        taamai.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Auth.Login(ctx, &operations.LoginRequestBody{
        Email: "sobanshahid47@gmail.com",
        Password: "ali112233",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Login != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.LoginRequestBody](../../pkg/models/operations/loginrequestbody.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                   | :heavy_minus_sign:                                                             | The options for this request.                                                  |


### Response

**[*operations.LoginResponse](../../pkg/models/operations/loginresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
