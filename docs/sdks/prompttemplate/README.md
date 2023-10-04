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

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.AddandremovefrombookmarkprompttemplateRequest](../../models/operations/addandremovefrombookmarkprompttemplaterequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |
| `opts`                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                             | :heavy_minus_sign:                                                                                                                   | The options for this request.                                                                                                        |


### Response

**[*operations.AddandremovefrombookmarkprompttemplateResponse](../../models/operations/addandremovefrombookmarkprompttemplateresponse.md), error**


## CreatePromptTemplate

Create Prompt Template

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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreatePromptTemplateRequest](../../models/operations/createprompttemplaterequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |


### Response

**[*operations.CreatePromptTemplateResponse](../../models/operations/createprompttemplateresponse.md), error**


## Generateprompttemplate

Generate prompt template

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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GenerateprompttemplateRequest](../../models/operations/generateprompttemplaterequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |


### Response

**[*operations.GenerateprompttemplateResponse](../../models/operations/generateprompttemplateresponse.md), error**


## ParmanentDeletePrompttemplate

Parmanent Delete Prompt template

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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.ParmanentDeletePrompttemplateRequest](../../models/operations/parmanentdeleteprompttemplaterequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |


### Response

**[*operations.ParmanentDeletePrompttemplateResponse](../../models/operations/parmanentdeleteprompttemplateresponse.md), error**


## PromptTemplates

Prompt Templates

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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.PromptTemplatesRequest](../../models/operations/prompttemplatesrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |


### Response

**[*operations.PromptTemplatesResponse](../../models/operations/prompttemplatesresponse.md), error**


## RestorePromptTemplate

Restore Prompt Template

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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.RestorePromptTemplateRequest](../../models/operations/restoreprompttemplaterequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |


### Response

**[*operations.RestorePromptTemplateResponse](../../models/operations/restoreprompttemplateresponse.md), error**


## TrashedPromptTemplates

Trashed Prompt Templates

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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.TrashedPromptTemplatesRequest](../../models/operations/trashedprompttemplatesrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |


### Response

**[*operations.TrashedPromptTemplatesResponse](../../models/operations/trashedprompttemplatesresponse.md), error**


## Deleteprmopttemplate

delete prmopt template

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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.DeleteprmopttemplateRequest](../../models/operations/deleteprmopttemplaterequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |


### Response

**[*operations.DeleteprmopttemplateResponse](../../models/operations/deleteprmopttemplateresponse.md), error**


## Prompttemplatelikeorremovefromlike

prompt template like or remove from like

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

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.PrompttemplatelikeorremovefromlikeRequest](../../models/operations/prompttemplatelikeorremovefromlikerequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |


### Response

**[*operations.PrompttemplatelikeorremovefromlikeResponse](../../models/operations/prompttemplatelikeorremovefromlikeresponse.md), error**

