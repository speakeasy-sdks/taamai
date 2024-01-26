# PromptTemplate
(*PromptTemplate*)

### Available Operations

* [Addandremovefrombookmarkprompttemplate](#addandremovefrombookmarkprompttemplate) - Add and remove from bookmark prompt template
* [CreatePromptTemplate](#createprompttemplate) - Create Prompt Template
* [Generateprompttemplate](#generateprompttemplate) - Generate prompt template
* [ParmanentDeletePrompttemplate](#parmanentdeleteprompttemplate) - Parmanent Delete Prompt template
* [PromptTemplates](#prompttemplates) - Prompt Templates
* [RestorePromptTemplate](#restoreprompttemplate) - Restore Prompt Template
* [TrashedPromptTemplates](#trashedprompttemplates) - Trashed Prompt Templates
* [Deleteprmopttemplate](#deleteprmopttemplate) - delete prmopt template
* [Prompttemplatelikeorremovefromlike](#prompttemplatelikeorremovefromlike) - prompt template like or remove from like

## Addandremovefrombookmarkprompttemplate

Add and remove from bookmark prompt template

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
    res, err := s.PromptTemplate.Addandremovefrombookmarkprompttemplate(ctx, operations.AddandremovefrombookmarkprompttemplateRequest{
        TemplateID: 1,
        UserID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.NewRequest1 != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `request`                                                                                                                                | [operations.AddandremovefrombookmarkprompttemplateRequest](../../pkg/models/operations/addandremovefrombookmarkprompttemplaterequest.md) | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |
| `opts`                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                             | :heavy_minus_sign:                                                                                                                       | The options for this request.                                                                                                            |


### Response

**[*operations.AddandremovefrombookmarkprompttemplateResponse](../../pkg/models/operations/addandremovefrombookmarkprompttemplateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreatePromptTemplate

Create Prompt Template

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
    res, err := s.PromptTemplate.CreatePromptTemplate(ctx, operations.CreatePromptTemplateRequest{
        Activate: 1,
        Category: "text",
        Code0: "input-field-1",
        Edit: 1,
        InputField0: "input",
        Language: "en-US",
        Name: "New checking",
        Names0: "new",
        Package: "all",
        Placeholders0: "Enter relavent information",
        Prompt: "Want to create vlog",
        Public: 1,
        Tone: 1,
        UserID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreatePromptTemplate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreatePromptTemplateRequest](../../pkg/models/operations/createprompttemplaterequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |


### Response

**[*operations.CreatePromptTemplateResponse](../../pkg/models/operations/createprompttemplateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Generateprompttemplate

Generate prompt template

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
    res, err := s.PromptTemplate.Generateprompttemplate(ctx, operations.GenerateprompttemplateRequest{
        Creativity: 0.5,
        Description: "something new",
        FolderID: 1,
        Language: "en-US",
        MaxResults: 1,
        TemplateCode: "SXHY9",
        Text1: "hello",
        Title: "new",
        UserID: 1,
        Words: 10000,
        WorkbookID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Generateprompttemplate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GenerateprompttemplateRequest](../../pkg/models/operations/generateprompttemplaterequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |


### Response

**[*operations.GenerateprompttemplateResponse](../../pkg/models/operations/generateprompttemplateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ParmanentDeletePrompttemplate

Parmanent Delete Prompt template

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
    res, err := s.PromptTemplate.ParmanentDeletePrompttemplate(ctx, operations.ParmanentDeletePrompttemplateRequest{
        TemplateID: 3,
        UserID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ParmanentDeletePrompttemplate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.ParmanentDeletePrompttemplateRequest](../../pkg/models/operations/parmanentdeleteprompttemplaterequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |


### Response

**[*operations.ParmanentDeletePrompttemplateResponse](../../pkg/models/operations/parmanentdeleteprompttemplateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PromptTemplates

Prompt Templates

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
    res, err := s.PromptTemplate.PromptTemplates(ctx, operations.PromptTemplatesRequest{
        Cat: "publc",
        SubCat: "all",
        UserID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PromptTemplates != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.PromptTemplatesRequest](../../pkg/models/operations/prompttemplatesrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |


### Response

**[*operations.PromptTemplatesResponse](../../pkg/models/operations/prompttemplatesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## RestorePromptTemplate

Restore Prompt Template

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
    res, err := s.PromptTemplate.RestorePromptTemplate(ctx, operations.RestorePromptTemplateRequest{
        TemplateID: 3,
        UserID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RestorePromptTemplate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.RestorePromptTemplateRequest](../../pkg/models/operations/restoreprompttemplaterequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |


### Response

**[*operations.RestorePromptTemplateResponse](../../pkg/models/operations/restoreprompttemplateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## TrashedPromptTemplates

Trashed Prompt Templates

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
    res, err := s.PromptTemplate.TrashedPromptTemplates(ctx, operations.TrashedPromptTemplatesRequest{
        UserID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TrashedPromptTemplates != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.TrashedPromptTemplatesRequest](../../pkg/models/operations/trashedprompttemplatesrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |


### Response

**[*operations.TrashedPromptTemplatesResponse](../../pkg/models/operations/trashedprompttemplatesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Deleteprmopttemplate

delete prmopt template

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
    res, err := s.PromptTemplate.Deleteprmopttemplate(ctx, operations.DeleteprmopttemplateRequest{
        TemplateID: 3,
        UserID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Deleteprmopttemplate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.DeleteprmopttemplateRequest](../../pkg/models/operations/deleteprmopttemplaterequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |


### Response

**[*operations.DeleteprmopttemplateResponse](../../pkg/models/operations/deleteprmopttemplateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Prompttemplatelikeorremovefromlike

prompt template like or remove from like

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
    res, err := s.PromptTemplate.Prompttemplatelikeorremovefromlike(ctx, operations.PrompttemplatelikeorremovefromlikeRequest{
        TemplateID: 1,
        UserID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Prompttemplatelikeorremovefromlike != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.PrompttemplatelikeorremovefromlikeRequest](../../pkg/models/operations/prompttemplatelikeorremovefromlikerequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `opts`                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |


### Response

**[*operations.PrompttemplatelikeorremovefromlikeResponse](../../pkg/models/operations/prompttemplatelikeorremovefromlikeresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
