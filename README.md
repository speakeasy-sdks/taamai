# github.com/speakeasy-sdks/taamai

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://github.com/speakeasy-sdks/taamai.git/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/taamai/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
    
</div>

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/taamai
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	"github.com/speakeasy-sdks/taamai"
	"github.com/speakeasy-sdks/taamai/pkg/models/operations"
	"github.com/speakeasy-sdks/taamai/pkg/models/shared"
	"log"
)

func main() {
	s := taamai.New(
		taamai.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.WorkbookAndFolders.AddandremovefromfavDocument(ctx, operations.AddandremovefromfavDocumentRequest{
		ID:     6,
		Type:   "document",
		UserID: 1,
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.AddandremovefromfavDocument != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [.WorkbookAndFolders](docs/sdks/workbookandfolders/README.md)

* [AddandremovefromfavDocument](docs/sdks/workbookandfolders/README.md#addandremovefromfavdocument) - Add and remove from fav Document
* [Contentsinworkbook](docs/sdks/workbookandfolders/README.md#contentsinworkbook) - Contents in work book
* [CreateFolder](docs/sdks/workbookandfolders/README.md#createfolder) - Create Folder
* [CreateWorkbook](docs/sdks/workbookandfolders/README.md#createworkbook) - Create Workbook
* [DeleteWorkspace](docs/sdks/workbookandfolders/README.md#deleteworkspace) - Delete Workspace
* [Deleteallkindofdocuments](docs/sdks/workbookandfolders/README.md#deleteallkindofdocuments) - Delete all kind of documents
* [Joinworkbook](docs/sdks/workbookandfolders/README.md#joinworkbook) - Join workbook
* [Joinworkbookrequestr](docs/sdks/workbookandfolders/README.md#joinworkbookrequestr) - Join workbook requestr
* [PermanentDeletefolder](docs/sdks/workbookandfolders/README.md#permanentdeletefolder) - Permanent Delete folder
* [PermanentDeleteworkspace](docs/sdks/workbookandfolders/README.md#permanentdeleteworkspace) - Permanent Delete workspace
* [RejectWorkbook](docs/sdks/workbookandfolders/README.md#rejectworkbook) - Reject Workbook
* [RestoreDocuemntofalltypeContentVoiceoverImagesTranscriptCode](docs/sdks/workbookandfolders/README.md#restoredocuemntofalltypecontentvoiceoverimagestranscriptcode) - Restore Docuemnt of all type  (content.voiceover,images,transcript,code)
* [RestoreWorkspace](docs/sdks/workbookandfolders/README.md#restoreworkspace) - Restore Workspace
* [Setdefualtworkspace](docs/sdks/workbookandfolders/README.md#setdefualtworkspace) - Set defualt workspace
* [Trashedfolders](docs/sdks/workbookandfolders/README.md#trashedfolders) - Trashed folders
* [Trashedworkspaces](docs/sdks/workbookandfolders/README.md#trashedworkspaces) - Trashed workspaces
* [WorkbookDetail](docs/sdks/workbookandfolders/README.md#workbookdetail) - Workbook Detail
* [Workbookvoiceovers](docs/sdks/workbookandfolders/README.md#workbookvoiceovers) - Workbook voiceovers
* [Allworkbooks](docs/sdks/workbookandfolders/README.md#allworkbooks) - all workbooks
* [Deletefolder](docs/sdks/workbookandfolders/README.md#deletefolder) - delete folder
* [PermanentDeletedocument](docs/sdks/workbookandfolders/README.md#permanentdeletedocument) - permanent Delete document
* [Restorefolder](docs/sdks/workbookandfolders/README.md#restorefolder) - restore folder
* [Userchats](docs/sdks/workbookandfolders/README.md#userchats) - user chats
* [Workbookcodes](docs/sdks/workbookandfolders/README.md#workbookcodes) - workbook codes
* [Workbookimages](docs/sdks/workbookandfolders/README.md#workbookimages) - workbook images
* [Workbookpolicies](docs/sdks/workbookandfolders/README.md#workbookpolicies) - workbook policies
* [Workbooktranscripts](docs/sdks/workbookandfolders/README.md#workbooktranscripts) - workbook transcripts

### [.CustimTemplates](docs/sdks/custimtemplates/README.md)

* [CreateCustomTemplate](docs/sdks/custimtemplates/README.md#createcustomtemplate) - Create Custom Template
* [CustomTemplategenerate](docs/sdks/custimtemplates/README.md#customtemplategenerate) - Custom Template generate
* [CustomTemplates](docs/sdks/custimtemplates/README.md#customtemplates) - Custom Templates
* [DeleteCustomtemplate](docs/sdks/custimtemplates/README.md#deletecustomtemplate) - Delete Custom template
* [FavCustomTemplates](docs/sdks/custimtemplates/README.md#favcustomtemplates) - Fav Custom Templates
* [RestoreCustomtemplate](docs/sdks/custimtemplates/README.md#restorecustomtemplate) - Restore Custom template
* [TrashedCustomTemplates](docs/sdks/custimtemplates/README.md#trashedcustomtemplates) - Trashed Custom Templates

### [.ChatWithPdf](docs/sdks/chatwithpdf/README.md)

* [NewRequest](docs/sdks/chatwithpdf/README.md#newrequest) - New Request
* [Sendandgetmsgtochatpdf](docs/sdks/chatwithpdf/README.md#sendandgetmsgtochatpdf) - Send and get msg to chat pdf
* [Fileupload](docs/sdks/chatwithpdf/README.md#fileupload) - file upload
* [Pdftotext](docs/sdks/chatwithpdf/README.md#pdftotext) - pdf to text
* [Uploadfileforchatpdf](docs/sdks/chatwithpdf/README.md#uploadfileforchatpdf) - upload file for chat pdf

### [.Misc](docs/sdks/misc/README.md)

* [AllCategories](docs/sdks/misc/README.md#allcategories) - All Categories

### [.AddonFeatures](docs/sdks/addonfeatures/README.md)

* [GenerateCode](docs/sdks/addonfeatures/README.md#generatecode) - Generate Code
* [GenerateimagefromAI](docs/sdks/addonfeatures/README.md#generateimagefromai) - Generate image from AI
* [Generatespeechtotext](docs/sdks/addonfeatures/README.md#generatespeechtotext) - Generate speech to text
* [Savecodeinworkspace](docs/sdks/addonfeatures/README.md#savecodeinworkspace) - Save code in workspace
* [Savetranscript](docs/sdks/addonfeatures/README.md#savetranscript) - Save transcript

### [.Auth](docs/sdks/auth/README.md)

* [Register](docs/sdks/auth/README.md#register) - Register
* [Login](docs/sdks/auth/README.md#login) - login

### [.Product](docs/sdks/product/README.md)

* [CreateProduct](docs/sdks/product/README.md#createproduct) - Create Product
* [DeleteProduct](docs/sdks/product/README.md#deleteproduct) - Delete Product
* [RestoreProduct](docs/sdks/product/README.md#restoreproduct) - Restore Product
* [TrashedProducts](docs/sdks/product/README.md#trashedproducts) - Trashed Products
* [UpdateProduct](docs/sdks/product/README.md#updateproduct) - Update Product
* [UserProductd](docs/sdks/product/README.md#userproductd) - User Productd
* [ParmenentdeleteProduct](docs/sdks/product/README.md#parmenentdeleteproduct) - parmenent delete Product

### [.PromptTemplate](docs/sdks/prompttemplate/README.md)

* [Addandremovefrombookmarkprompttemplate](docs/sdks/prompttemplate/README.md#addandremovefrombookmarkprompttemplate) - Add and remove from bookmark prompt template
* [CreatePromptTemplate](docs/sdks/prompttemplate/README.md#createprompttemplate) - Create Prompt Template
* [Generateprompttemplate](docs/sdks/prompttemplate/README.md#generateprompttemplate) - Generate prompt template
* [ParmanentDeletePrompttemplate](docs/sdks/prompttemplate/README.md#parmanentdeleteprompttemplate) - Parmanent Delete Prompt template
* [PromptTemplates](docs/sdks/prompttemplate/README.md#prompttemplates) - Prompt Templates
* [RestorePromptTemplate](docs/sdks/prompttemplate/README.md#restoreprompttemplate) - Restore Prompt Template
* [TrashedPromptTemplates](docs/sdks/prompttemplate/README.md#trashedprompttemplates) - Trashed Prompt Templates
* [Deleteprmopttemplate](docs/sdks/prompttemplate/README.md#deleteprmopttemplate) - delete prmopt template
* [Prompttemplatelikeorremovefromlike](docs/sdks/prompttemplate/README.md#prompttemplatelikeorremovefromlike) - prompt template like or remove from like

### [.Templates](docs/sdks/templates/README.md)

* [AllTemplates](docs/sdks/templates/README.md#alltemplates) - All Templates
* [FavTemplates](docs/sdks/templates/README.md#favtemplates) - Fav Templates
* [GenerateTemplate](docs/sdks/templates/README.md#generatetemplate) - Generate Template
* [ProcessTemplate](docs/sdks/templates/README.md#processtemplate) - Process Template
* [TemplateDetail](docs/sdks/templates/README.md#templatedetail) - Template Detail
* [Templategroups](docs/sdks/templates/README.md#templategroups) - Template groups
<!-- End SDK Available Operations -->

<!-- Start Dev Containers -->

<!-- End Dev Containers -->

<!-- Start Pagination -->
# Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
<!-- End Pagination -->

<!-- Start Go Types -->

<!-- End Go Types -->



<!-- Start Error Handling -->
# Error Handling

Handling errors in your SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.


<!-- End Error Handling -->



<!-- Start Server Selection -->
# Server Selection

## Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://taam.ai/api` | None |
| 1 | `http://127.0.0.1:8000/api` | None |

For example:

```go
package main

import (
	"context"
	"github.com/speakeasy-sdks/taamai"
	"github.com/speakeasy-sdks/taamai/pkg/models/operations"
	"github.com/speakeasy-sdks/taamai/pkg/models/shared"
	"log"
)

func main() {
	s := taamai.New(
		taamai.WithServerIndex(1),
		taamai.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.WorkbookAndFolders.AddandremovefromfavDocument(ctx, operations.AddandremovefromfavDocumentRequest{
		ID:     6,
		Type:   "document",
		UserID: 1,
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.AddandremovefromfavDocument != nil {
		// handle response
	}
}

```


## Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:

```go
package main

import (
	"context"
	"github.com/speakeasy-sdks/taamai"
	"github.com/speakeasy-sdks/taamai/pkg/models/operations"
	"github.com/speakeasy-sdks/taamai/pkg/models/shared"
	"log"
)

func main() {
	s := taamai.New(
		taamai.WithServerURL("https://taam.ai/api"),
		taamai.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.WorkbookAndFolders.AddandremovefromfavDocument(ctx, operations.AddandremovefromfavDocumentRequest{
		ID:     6,
		Type:   "document",
		UserID: 1,
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.AddandremovefromfavDocument != nil {
		// handle response
	}
}

```

## Override Server URL Per-Operation

The server URL can also be overridden on a per-operation basis, provided a server list was specified for the operation. For example:

```go
package main

import (
	"context"
	"github.com/speakeasy-sdks/taamai"
	"github.com/speakeasy-sdks/taamai/pkg/models/operations"
	"github.com/speakeasy-sdks/taamai/pkg/models/shared"
	"log"
)

func main() {
	s := taamai.New(
		taamai.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.WorkbookAndFolders.AddandremovefromfavDocument(ctx, operations.AddandremovefromfavDocumentRequest{
		ID:     6,
		Type:   "document",
		UserID: 1,
	}, operations.WithServerURL("http://127.0.0.1:8000/api"))
	if err != nil {
		log.Fatal(err)
	}

	if res.AddandremovefromfavDocument != nil {
		// handle response
	}
}

```
<!-- End Server Selection -->



<!-- Start Custom HTTP Client -->
# Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client -->



<!-- Start Authentication -->

# Authentication

## Per-Client Security Schemes

Your SDK supports the following security scheme globally:

| Name        | Type        | Scheme      |
| ----------- | ----------- | ----------- |
| `Bearer`    | http        | HTTP Bearer |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:

```go
package main

import (
	"context"
	"github.com/speakeasy-sdks/taamai"
	"github.com/speakeasy-sdks/taamai/pkg/models/operations"
	"github.com/speakeasy-sdks/taamai/pkg/models/shared"
	"log"
)

func main() {
	s := taamai.New(
		taamai.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.WorkbookAndFolders.AddandremovefromfavDocument(ctx, operations.AddandremovefromfavDocumentRequest{
		ID:     6,
		Type:   "document",
		UserID: 1,
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.AddandremovefromfavDocument != nil {
		// handle response
	}
}

```
<!-- End Authentication -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
