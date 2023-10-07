# AddonFeatures
(*AddonFeatures*)

### Available Operations

* [GenerateCode](#generatecode) - Generate Code
* [GenerateimagefromAI](#generateimagefromai) - Generate image from AI
* [Generatespeechtotext](#generatespeechtotext) - Generate speech to text
* [Savecodeinworkspace](#savecodeinworkspace) - Save code in workspace
* [Savetranscript](#savetranscript) - Save transcript

## GenerateCode

Generate Code

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
    res, err := s.AddonFeatures.GenerateCode(ctx, operations.GenerateCodeRequest{
        Document: "new checking",
        Instructions: "generate a code to store image",
        Language: "php",
        UserID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GenerateCode != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GenerateCodeRequest](../../models/operations/generatecoderequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |


### Response

**[*operations.GenerateCodeResponse](../../models/operations/generatecoderesponse.md), error**


## GenerateimagefromAI

Generate image from AI

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
    res, err := s.AddonFeatures.GenerateimagefromAI(ctx, operations.GenerateimagefromAIRequest{
        MaxResults: 1,
        Name: "sample checking",
        Resolution: "256x256",
        Title: "need a eagle image",
        UserID: 1,
        WorkbookFolderID: 1,
        WorkbookID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GenerateimagefromAI != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GenerateimagefromAIRequest](../../models/operations/generateimagefromairequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |


### Response

**[*operations.GenerateimagefromAIResponse](../../models/operations/generateimagefromairesponse.md), error**


## Generatespeechtotext

Generate speech to text

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
    res, err := s.AddonFeatures.Generatespeechtotext(ctx, &operations.GeneratespeechtotextRequestBody{
        AudioFile: operations.GeneratespeechtotextRequestBodyAudioFile{
            AudioFile: "Dinar Bicycle",
            Content: []byte(",3S}7YH8}T"),
        },
        Document: "new",
        Language: "en",
        Task: "transcribe",
        UserID: 1,
        WorkbookFolderID: 1,
        WorkbookID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Generatespeechtotext != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GeneratespeechtotextRequestBody](../../models/operations/generatespeechtotextrequestbody.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |


### Response

**[*operations.GeneratespeechtotextResponse](../../models/operations/generatespeechtotextresponse.md), error**


## Savecodeinworkspace

Save code in workspace

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
    res, err := s.AddonFeatures.Savecodeinworkspace(ctx, operations.SavecodeinworkspaceRequest{
        CodeID: 32,
        FolderID: 1,
        UserID: 1,
        WorkbookID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Savecodeinworkspace != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.SavecodeinworkspaceRequest](../../models/operations/savecodeinworkspacerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |


### Response

**[*operations.SavecodeinworkspaceResponse](../../models/operations/savecodeinworkspaceresponse.md), error**


## Savetranscript

Save transcript

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
    res, err := s.AddonFeatures.Savetranscript(ctx, operations.SavetranscriptRequest{
        Text: "Marhaban, ana al-mutahaddithi al-iftiradi min imza'i al-jawda. Da'ani ulqi al-tahiyyata ala jumhourik wa u'arrifahom ala muntajatik abra wasilatin min akthar al-wasaili al-taswiqiya, tashwiqan wa mut'a.\",
        \"status\": \"success",
        Title: "New task",
        TranscriptID: 43,
        UserID: 1,
        WorkbookFolderID: 1,
        WorkbookID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Savetranscript != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.SavetranscriptRequest](../../models/operations/savetranscriptrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |


### Response

**[*operations.SavetranscriptResponse](../../models/operations/savetranscriptresponse.md), error**

