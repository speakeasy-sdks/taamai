# Templates
(*.Templates*)

### Available Operations

* [AllTemplates](#alltemplates) - All Templates
* [FavTemplates](#favtemplates) - Fav Templates
* [GenerateTemplate](#generatetemplate) - Generate Template
* [ProcessTemplate](#processtemplate) - Process Template
* [TemplateDetail](#templatedetail) - Template Detail
* [Templategroups](#templategroups) - Template groups

## AllTemplates

All Templates

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
        taamai.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Templates.AllTemplates(ctx, operations.AllTemplatesRequest{
        UserID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AllTemplates != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.AllTemplatesRequest](../../models/operations/alltemplatesrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |


### Response

**[*operations.AllTemplatesResponse](../../models/operations/alltemplatesresponse.md), error**


## FavTemplates

Fav Templates

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
        taamai.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Templates.FavTemplates(ctx, operations.FavTemplatesRequest{
        UserID: 12,
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
| `request`                                                                        | [operations.FavTemplatesRequest](../../models/operations/favtemplatesrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |


### Response

**[*operations.FavTemplatesResponse](../../models/operations/favtemplatesresponse.md), error**


## GenerateTemplate

Generate Template

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
        taamai.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Templates.GenerateTemplate(ctx, operations.GenerateTemplateRequest{
        Creativity: 0.5,
        Description: "tell me about code",
        FolderID: 1,
        Language: "en-US",
        MaxResults: 1,
        TemplateCode: "EEKZF",
        Title: "Tsmokeshope",
        UserID: 40,
        Words: 100,
        WorkbookID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GenerateTemplate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GenerateTemplateRequest](../../models/operations/generatetemplaterequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |


### Response

**[*operations.GenerateTemplateResponse](../../models/operations/generatetemplateresponse.md), error**


## ProcessTemplate

Process Template

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
        taamai.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Templates.ProcessTemplate(ctx, operations.ProcessTemplateRequest{
        ContentID: 286,
        MaxResults: 1,
        MaxWords: 100,
        Temperature: 0.5,
        UserID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ProcessTemplate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ProcessTemplateRequest](../../models/operations/processtemplaterequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |


### Response

**[*operations.ProcessTemplateResponse](../../models/operations/processtemplateresponse.md), error**


## TemplateDetail

Template Detail

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
        taamai.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Templates.TemplateDetail(ctx, operations.TemplateDetailRequest{
        TemplateID: 1,
        UserID: 40,
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
| `request`                                                                            | [operations.TemplateDetailRequest](../../models/operations/templatedetailrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |


### Response

**[*operations.TemplateDetailResponse](../../models/operations/templatedetailresponse.md), error**


## Templategroups

Template groups

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/taamai"
	"github.com/speakeasy-sdks/taamai/pkg/models/shared"
)

func main() {
    s := taamai.New(
        taamai.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Templates.Templategroups(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |


### Response

**[*operations.TemplategroupsResponse](../../models/operations/templategroupsresponse.md), error**

