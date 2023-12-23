# ChatWithPdf
(*ChatWithPdf*)

### Available Operations

* [NewRequest](#newrequest) - New Request
* [Sendandgetmsgtochatpdf](#sendandgetmsgtochatpdf) - Send and get msg to chat pdf
* [Fileupload](#fileupload) - file upload
* [Pdftotext](#pdftotext) - pdf to text
* [Uploadfileforchatpdf](#uploadfileforchatpdf) - upload file for chat pdf

## NewRequest

New Request

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/taamai/pkg/models/shared"
	"github.com/speakeasy-sdks/taamai"
	"context"
	"github.com/speakeasy-sdks/taamai/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := taamai.New(
        taamai.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.ChatWithPdf.NewRequest(ctx, operations.NewRequestRequest{
        Path: "assets/pdf/64ef458eabc4e_Soban-Shahid (Laravel).pdf",
        Question: "what data this file contains",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.NewRequestRequest](../../pkg/models/operations/newrequestrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |


### Response

**[*operations.NewRequestResponse](../../pkg/models/operations/newrequestresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Sendandgetmsgtochatpdf

Send and get msg to chat pdf

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/taamai/pkg/models/shared"
	"github.com/speakeasy-sdks/taamai"
	"context"
	"github.com/speakeasy-sdks/taamai/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := taamai.New(
        taamai.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.ChatWithPdf.Sendandgetmsgtochatpdf(ctx, &operations.SendandgetmsgtochatpdfRequestBody{
        Question: "what is the education",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.SendandgetmsgtochatpdfRequestBody](../../pkg/models/operations/sendandgetmsgtochatpdfrequestbody.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |


### Response

**[*operations.SendandgetmsgtochatpdfResponse](../../pkg/models/operations/sendandgetmsgtochatpdfresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Fileupload

file upload

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/taamai/pkg/models/shared"
	"github.com/speakeasy-sdks/taamai"
	"context"
	"github.com/speakeasy-sdks/taamai/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := taamai.New(
        taamai.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.ChatWithPdf.Fileupload(ctx, &operations.FileuploadRequestBody{
        File: operations.File{
            Content: []byte("0xA40D81A486"),
            FileName: "martin_coleslaw_application.gif",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.FileuploadRequestBody](../../pkg/models/operations/fileuploadrequestbody.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |


### Response

**[*operations.FileuploadResponse](../../pkg/models/operations/fileuploadresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Pdftotext

pdf to text

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/taamai/pkg/models/shared"
	"github.com/speakeasy-sdks/taamai"
	"context"
	"github.com/speakeasy-sdks/taamai/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := taamai.New(
        taamai.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.ChatWithPdf.Pdftotext(ctx, &operations.PdftotextRequestBody{
        File: operations.PdftotextFile{
            Content: []byte("0x0F1b5e786D"),
            FileName: "woman_secured.pdf",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.PdftotextRequestBody](../../pkg/models/operations/pdftotextrequestbody.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |


### Response

**[*operations.PdftotextResponse](../../pkg/models/operations/pdftotextresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Uploadfileforchatpdf

upload file for chat pdf

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/taamai/pkg/models/shared"
	"github.com/speakeasy-sdks/taamai"
	"context"
	"github.com/speakeasy-sdks/taamai/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := taamai.New(
        taamai.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.ChatWithPdf.Uploadfileforchatpdf(ctx, &operations.UploadfileforchatpdfRequestBody{
        File: operations.UploadfileforchatpdfFile{
            Content: []byte("0x62cfbDe38b"),
            FileName: "radian.mp4v",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.UploadfileforchatpdfRequestBody](../../pkg/models/operations/uploadfileforchatpdfrequestbody.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |


### Response

**[*operations.UploadfileforchatpdfResponse](../../pkg/models/operations/uploadfileforchatpdfresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
