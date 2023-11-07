# Product
(*.Product*)

### Available Operations

* [CreateProduct](#createproduct) - Create Product
* [DeleteProduct](#deleteproduct) - Delete Product
* [RestoreProduct](#restoreproduct) - Restore Product
* [TrashedProducts](#trashedproducts) - Trashed Products
* [UpdateProduct](#updateproduct) - Update Product
* [UserProductd](#userproductd) - User Productd
* [ParmenentdeleteProduct](#parmenentdeleteproduct) - parmenent delete Product

## CreateProduct

Create Product

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
    res, err := s.Product.CreateProduct(ctx, &operations.CreateProductRequestBody{
        BulletPoints: "friend",
        Description: "THis is testing",
        Name: "New Product",
        Tags0: "hello",
        Tags1: "friend",
        TargetAudience: "males",
        UserID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateProduct != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateProductRequestBody](../../models/operations/createproductrequestbody.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |


### Response

**[*operations.CreateProductResponse](../../models/operations/createproductresponse.md), error**


## DeleteProduct

Delete Product

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
    res, err := s.Product.DeleteProduct(ctx, operations.DeleteProductRequest{
        ProductID: 10,
        UserID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteProduct != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.DeleteProductRequest](../../models/operations/deleteproductrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |


### Response

**[*operations.DeleteProductResponse](../../models/operations/deleteproductresponse.md), error**


## RestoreProduct

Restore Product

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
    res, err := s.Product.RestoreProduct(ctx, operations.RestoreProductRequest{
        ProductID: 10,
        UserID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RestoreProduct != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.RestoreProductRequest](../../models/operations/restoreproductrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |


### Response

**[*operations.RestoreProductResponse](../../models/operations/restoreproductresponse.md), error**


## TrashedProducts

Trashed Products

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
    res, err := s.Product.TrashedProducts(ctx, operations.TrashedProductsRequest{
        UserID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TrashedProducts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.TrashedProductsRequest](../../models/operations/trashedproductsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |


### Response

**[*operations.TrashedProductsResponse](../../models/operations/trashedproductsresponse.md), error**


## UpdateProduct

Update Product

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
    res, err := s.Product.UpdateProduct(ctx, operations.UpdateProductRequest{
        BulletPoints: 11,
        Description: "updated",
        Name: "updated",
        ProductID: 11,
        Tags0: "new",
        TargetAudience: "females",
        UserID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateProduct != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UpdateProductRequest](../../models/operations/updateproductrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |


### Response

**[*operations.UpdateProductResponse](../../models/operations/updateproductresponse.md), error**


## UserProductd

User Productd

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
    res, err := s.Product.UserProductd(ctx, operations.UserProductdRequest{
        UserID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UserProductd != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.UserProductdRequest](../../models/operations/userproductdrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |


### Response

**[*operations.UserProductdResponse](../../models/operations/userproductdresponse.md), error**


## ParmenentdeleteProduct

parmenent delete Product

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
    res, err := s.Product.ParmenentdeleteProduct(ctx, operations.ParmenentdeleteProductRequest{
        ProductID: 10,
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ParmenentdeleteProductRequest](../../models/operations/parmenentdeleteproductrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |


### Response

**[*operations.ParmenentdeleteProductResponse](../../models/operations/parmenentdeleteproductresponse.md), error**

