# POCKETBASE DOCS|2026-02-25|1 sections

## 1.Extend with Go - Routing
GET /hello/:name
# Extend with Go - Routing
`GET /hello/:name`
PocketBase uses
echo (v5)
for routing. The internal router is exposed in the app.OnBeforeServe event hook and you can register
your own custom endpoints and/or middlewares the same way as using directly echo.
### Routes
##### Registering new routes
Each route consists of at least a path and a handler function. For example, the below code registers
GET /hello/:name route that responds with a json body:
import (
"log"
"net/http"
"github.com/labstack/echo/v5"
"github.com/pocketbase/pocketbase"
"github.com/pocketbase/pocketbase/core"
app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
e.Router.GET("/old/hello/:name", func(c echo.Context) error {
name := c.PathParam("name")
return c.JSON(http.StatusOK, map[string]string{"message": "Hello " + name})
}, /* optional middlewares */)
return nil
To avoid collisions with future internal routes you should avoid using the /api/...
base path or consider combining it with a unique prefix like /api/myapp/....
There are several Router methods available, but the most common ones are:
e.Router.GET(path, handler, [middlewares...])
e.Router.POST(path, handler, [middlewares...])
e.Router.PUT(path, handler, [middlewares...])
e.Router.PATCH(path, handler, [middlewares...])
e.Router.DELETE(path, handler, [middlewares...])
e.Router.OPTIONS(path, handler, [middlewares...])
e.Router.HEAD(path, handler, [middlewares...])
Each handler function receives a echo request context argument (usually named
c).
The request context is also accessible in the event request hooks under the HttpContext
field.
Below you can find common request context operations.
##### Request context store
The request context comes with a local store that you can use to share data related only to the current
request between routes and middlewares.
// store for the duration of the request
c.Set("someKey", 123)
// retrieve later
val := c.Get("someKey").(int) // 123
##### Retrieving the current auth state
We also use the store to manage the current auth state with the admin and
authRecord special keys.
admin, _ := c.Get(apis.ContextAdminKey).(*models.Admin)
record, _ := c.Get(apis.ContextAuthRecordKey).(*models.Record)
// alternatively, you can also read the auth state from the cached request info
info   := apis.RequestInfo(c)
admin  := info.Admin      // nil if not authenticated as admin
record := info.AuthRecord // nil if not authenticated as regular auth record
isGuest := admin == nil &amp;&amp; record == nil
##### Reading path parameters
c.PathParam(&quot;paramName&quot;).
`id := c.PathParam("id")`
##### Reading query parameters
search := c.QueryParam("search")
// or via the cached request object
search := apis.RequestInfo(c).Query["search"]
##### Reading request headers
token := c.Request().Header.Get("Some-Header")
// or via the cached request object (the header value is always normalized)
token := apis.RequestInfo(c).Headers["some_header"]
##### Writing response headers
`c.Response().Header().Set("Some-Header", "123")`
##### Reading request body
// read the body via the cached request object
// (this method is commonly used in hook handlers because it allows reading the body more than once)
data := apis.RequestInfo(c).Data
title := data["title"]
// read/scan the request body fields into a typed struct
// (note that a body cannot be read twice with Bind because it is a stream)
data := struct {
Title       string `json:"title" form:"title"`
Description string `json:"description" form:"description"`
Public      bool   `json:"public" form:"public"`
if err := c.Bind(&amp;data); err != nil {
return apis.NewBadRequestError("Failed to read request data", err)
// read single multipart/form-data field
title := c.FormValue("title")
// read single multipart/form-data file
doc, err := c.FormFile("document")
##### Writing response body
// send response with json body
c.JSON(200, map[string]any{"name": "John"})
// send response with string body
// send response with html body
// (check also the "Rendering templates" section)
c.HTML(200, "&lt;h1>Hello!&lt;/h1>")
// redirect
// send response with no body
c.NoContent(204)
### Middlewares
Middlewares could be used to apply a shared behavior or to intercept and modify route requests.
Middlewares can be registered both to a single route (by passing them after the handler) and globally usually
by using e.Router.Use(someMiddlereFunc).
app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
// attach a middleware globally to all routes
e.Router.Use(someMiddlereFunc)
// attach multiple middlewares to a single route
// each route will execute their own middlewares + the global ones
e.Router.GET("/old/hello", func(c echo.Context) error {
return c.String(200, "Hello world!")
}, apis.ActivityLogger(app), apis.RequireAdminAuth())
##### Builtin middlewares
// logs the request in the Admin UI > Logs
apis.ActivityLogger(app)
// requires the request client to be unauthenticated, aka. guest
apis.RequireGuestOnly()
// requires the request client to be authenticated as an auth record
apis.RequireRecordAuth(optCollectionNames...)
// require the request client to be authenticated as admin
apis.RequireAdminAuth()
// require the request client to be authenticated as admin OR auth record
apis.RequireAdminOrRecordAuth(optCollectionNames...)
// require the request client to be authenticated as admin OR auth record
// that matches the ownerIdParam path parameter
apis.RequireAdminOrOwnerAuth(ownerIdParam = "id")
##### Custom middlewares
return func(c echo.Context) error {
// eg. inspect some header value before processing the request
header := c.Request().Header.Get("Some-Header")
if header == "" {
return apis.NewBadRequestError("Invalid request", nil)
app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
e.Router.Use(myCustomMiddleware)
return nil
### Error response
PocketBase has a global error handler and every returned or thrown Error from a route or
middleware will be safely converted by default to a generic HTTP 400 error to avoid accidentally leaking
sensitive information (the original error will be visible only in the Admin UI &gt; Logs or when in
--dev mode).
To make it easier returning formatted json error responses, PocketBase provides
apis.ApiError constructor that can be instantiated directly or using the builtin factories.
validation.Error items.
import (
validation "github.com/go-ozzo/ozzo-validation/v4"
"github.com/pocketbase/pocketbase/apis"
// construct ApiError with custom status code and validation data error
return apis.NewApiError(500, "something went wrong", map[string]validation.Error{
"title": validation.NewError("invalid_title", "Invalid or missing title"),
// if message is empty string, a default one will be set
return apis.NewBadRequestError(optMessage, optData)   // 400 ApiError
return apis.NewUnauthorizedError(optMessage, optData) // 401 ApiError
return apis.NewForbiddenError(optMessage, optData)    // 403 ApiError
return apis.NewNotFoundError(optMessage, optData)     // 404 ApiError
### Helpers
The
apis package
expose several helpers you can use as part of your route hooks.
##### Auth response
apis.RecordAuthResponse() writes standardised json record auth response (aka. token + record data)
into the specified request context. Could be used as a return result from a custom auth route.
app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
e.Router.GET("/old/phone-login", func(c echo.Context) error {
data := struct {
Phone    string `json:"phone" form:"phone"`
Password string `json:"password" form:"password"`
if err := c.Bind(&amp;data); err != nil {
return apis.NewBadRequestError("Failed to read request data", err)
record, err := app.Dao().FindFirstRecordByData("users", "phone", data.Phone)
if err != nil || !record.ValidatePassword(data.Password) {
// return generic 400 error to prevent phones enumeration
return apis.NewBadRequestError("Invalid credentials", err)
return apis.RecordAuthResponse(app, c, record, nil)
}, /* optional middlewares */)
return nil
##### Enrich record(s)
apis.EnrichRecord() and apis.EnrichRecords() helpers parses the request context and
enrich the provided record(s) by:
-expands relations (if defaultExpands and/or ?expand query parameter is set)
-ensures that the emails of the auth record and its expanded auth relations are visible only for the
current logged admin, record owner or record with manage access
app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
e.Router.GET("/old/custom-article", func(c echo.Context) error {
records, err := app.Dao().FindRecordsByFilter("article", "status = 'active'", "-created", 40)
if err != nil {
return apis.NewNotFoundError("No active articles", err)
// enrich the records with the "categories" relation as default expand
apis.EnrichRecords(c, app.Dao(), records, "categories")
return c.JSON(http.StatusOK, records)
}, /* optional middlewares */)
return nil
##### Serving static files
app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
// serves static files from the provided dir (if exists)
e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("/path/to/public"), false))
return nil
### Sending request to custom routes using the SDKs
The official PocketBase SDKs expose the internal send() method that could be used to send requests
to your custom route(s).
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
await pb.send("/old/hello", {
// for all possible options check
// https://developer.mozilla.org/en-US/docs/Web/API/fetch#options
query: { "abc": 123 },
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
await pb.send("/old/hello", query: { "abc": 123 })
