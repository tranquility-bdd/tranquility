<img src="logo.jpg" height="300px" />

[![GoDoc](https://godoc.org/github.com/tranquility-bdd/tranquility?status.svg)](https://godoc.org/github.com/tranquility-bdd/tranquility)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/5a8c1957079146108d37c090167f1f58)](https://app.codacy.com/app/filfreire/tranquility-bdd?utm_source=github.com&utm_medium=referral&utm_content=tranquility-bdd/tranquility-bdd&utm_campaign=Badge_Grade_Dashboard)
[![CircleCI](https://circleci.com/gh/tranquility-bdd/tranquility.svg?style=svg)](https://circleci.com/gh/tranquility-bdd/tranquility)

a lightweight dependency/bloat-free alternative to serenity-rest-assured and serenity-core, a bit more maintainable than a regular Postman collection.

## The mission

- **Maintainable**: be more maintainable than postman - with little clutter - but equally as fast to develop full-length check suites as postman
- **Featherweight**: little dependencies (contrasting to serenity 170+ dependencies, or newmans 140+ dependencies)
- **Flexible**: do not lock the user to a "BDD" approach where user is forced to use Gherkin in order to use `tranquility`

## Requirements

Make sure you have [go](link) installed. Latest version we've tried on was `go1.12.7 darwin/amd64`


## How to use

Let's say you want to do an automated check for a `POST` http request.

This is what a `curl` of that request would look like:

```
curl --location \
	--header "Content-Type: application/json" \
	--header "Accept-Language: es-ES" \
	--request POST "https://postman-echo.com/post?status=active" \
	--data '{"title":"Finding Tranquility"}'
```

And you expect the http response status should be `200 OK` the response body should, among other data, include the `json` data you sent:
```
{"args":{"status":"active"},"data":{"title":"Finding Tranquility."},"files":{},"form":{},"headers":{"x-forwarded-proto":"https","host":"postman-echo.com","content-length":"32","accept-encoding":"gzip","accept-language":"es-ES","content-type":"application/json","user-agent":"Go-http-client/1.1","x-forwarded-port":"443"},"json":{"title":"Finding Tranquility."},"url":"https://postman-echo.com/post?status=active"}
```

In a new folder, create a file `example.go`, and in this file write, first import tranquility:
```
package main

import (
	"fmt"
	"github.com/tranquility-bdd/tranquility"
)
```

Then to define the above single `Request` where you'll code:
1) it's `PreAction`,
2) it's `Action`,
3) and it's `Test` phases,

```
TODO Request snippet
```

You can as well as setup your `Env` which will influence the execution of the `Request` (when calling `yourRequest.Run()`);

```
TODO Env snippet
```

You can join up multiple `Request`'s using a `Collection`, and they'll run sequentially by calling `yourCollection.Run()`.

```
TODO Collection snippet
```

## More examples

You can check out a full example of Tranquility [here](link) and another one that uses Cucumber [here](link).


## Feature roadmap (2 June 2019)
- *(MVP)* Users can code their actions and manipulate environment (which influences actions) ✅
- *(MVP)* Create example repository that uses MVP version of tranquility with `gucumber` ✅
- Users can use PreAction and Test template/abstractions as part of their tests, if they opt for a unit approach ✅
- Users can define "Lego piece" of the 3 main components - PreAction, Action and Test ✅
- Users can import Postman environment Json files to pre-fill `Env` with the config data ✅
- (Optional/Doubt/Nice-to-have) Be able to export Scenarios built with Lego pieces to JSONs that obey postman schema

## Usage
In case you area already using [Go modules](https://github.com/golang/go/wiki/Modules) you can just include `"github.com/tranquility-bdd/tranquility"` in your import statement.

Otherwise you can use the command `go get "github.com/tranquility-bdd/tranquility"` to download the package and then import it as before.
