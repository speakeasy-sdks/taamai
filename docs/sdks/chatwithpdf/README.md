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
	"context"
	"log"
	"github.com/speakeasy-sdks/taamai"
	"github.com/speakeasy-sdks/taamai/pkg/models/shared"
	"github.com/speakeasy-sdks/taamai/pkg/models/operations"
)

func main() {
    s := taamai.New(
        taamai.WithSecurity(shared.Security{
            Bearer: "",
        }),
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.NewRequestRequest](../../models/operations/newrequestrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |


### Response

**[*operations.NewRequestResponse](../../models/operations/newrequestresponse.md), error**


## Sendandgetmsgtochatpdf

Send and get msg to chat pdf

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/taamai"
	"github.com/speakeasy-sdks/taamai/pkg/models/shared"
	"github.com/speakeasy-sdks/taamai/pkg/models/operations"
)

func main() {
    s := taamai.New(
        taamai.WithSecurity(shared.Security{
            Bearer: "",
        }),
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.SendandgetmsgtochatpdfRequestBody](../../models/operations/sendandgetmsgtochatpdfrequestbody.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |


### Response

**[*operations.SendandgetmsgtochatpdfResponse](../../models/operations/sendandgetmsgtochatpdfresponse.md), error**


## Fileupload

file upload

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/taamai"
	"github.com/speakeasy-sdks/taamai/pkg/models/shared"
	"github.com/speakeasy-sdks/taamai/pkg/models/operations"
)

func main() {
    s := taamai.New(
        taamai.WithSecurity(shared.Security{
            Bearer: "",
        }),
    )

    ctx := context.Background()
    res, err := s.ChatWithPdf.Fileupload(ctx, &operations.FileuploadRequestBody{
        File: operations.FileuploadRequestBodyFile{
            Content: []byte("d5#rF'h3C;"),
            File: "Martin coleslaw application",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.FileuploadRequestBody](../../models/operations/fileuploadrequestbody.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |


### Response

**[*operations.FileuploadResponse](../../models/operations/fileuploadresponse.md), error**


## Pdftotext

pdf to text

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/taamai"
	"github.com/speakeasy-sdks/taamai/pkg/models/shared"
	"github.com/speakeasy-sdks/taamai/pkg/models/operations"
)

func main() {
    s := taamai.New(
        taamai.WithSecurity(shared.Security{
            Bearer: "",
        }),
    )

    ctx := context.Background()
    res, err := s.ChatWithPdf.Pdftotext(ctx, &operations.PdftotextRequestBody{
        File: operations.PdftotextRequestBodyFile{
            Content: []byte("!|%P7_AE=r"),
            File: "woman secured",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.PdftotextRequestBody](../../models/operations/pdftotextrequestbody.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |


### Response

**[*operations.PdftotextResponse](../../models/operations/pdftotextresponse.md), error**


## Uploadfileforchatpdf

upload file for chat pdf

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/taamai"
	"github.com/speakeasy-sdks/taamai/pkg/models/shared"
	"github.com/speakeasy-sdks/taamai/pkg/models/operations"
)

func main() {
    s := taamai.New(
        taamai.WithSecurity(shared.Security{
            Bearer: "",
        }),
    )

    ctx := context.Background()
    res, err := s.ChatWithPdf.Uploadfileforchatpdf(ctx, &operations.UploadfileforchatpdfRequestBody{
        File: operations.UploadfileforchatpdfRequestBodyFile{
            Content: []byte(";-SdSu^1BO"),
            File: "radian",
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.UploadfileforchatpdfRequestBody](../../models/operations/uploadfileforchatpdfrequestbody.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |


### Response

**[*operations.UploadfileforchatpdfResponse](../../models/operations/uploadfileforchatpdfresponse.md), error**

