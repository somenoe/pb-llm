# POCKETBASE DOCS|2026-02-25|1 sections

## 1.Use as framework
# Use as framework
One of the main feature of PocketBase is that
it can be used as a framework which enables you to write your own custom app business
logic in
Go or JavaScript and still have a portable
backend at the end.
Choose Extend with Go if you are already familiar
with the language or have the time to learn it.
As the primary PocketBase language, the Go APIs are better documented and you&#39;ll be able to integrate with
the Go APIs are slightly more verbose and it may require some time to get used to, especially if this is your
first time working with Go.
Choose Extend with JavaScript
if you don&#39;t intend to write too much custom code and want a quick way to explore the PocketBase capabilities.
The embedded JavaScript engine is a pluggable wrapper around the existing Go APIs, so most of the time the
slight performance penalty will be negligible because it&#39;ll invoke the Go functions under the hood.
As a bonus, because the JS VM mirrors the Go APIs, you would be able migrate gradually without much code changes
from JS -&gt; Go at later stage in case you hit a bottleneck or want more control over the execution flow.
With both Go and JavaScript, you can:
-Register custom routes:
JavaScript
app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
e.Router.GET("/old/hello", func(c echo.Context) error {
return c.String(http.StatusOK, "Hello world!")
}, apis.ActivityLogger(app))
return nil
routerAdd("GET", "/hello", (c) => {
return c.string(200, "Hello world!")
}, $apis.activityLogger($app))
-Bind to event hooks and intercept responses:
JavaScript
app.OnRecordBeforeCreateRequest("posts").Add(func(e *core.RecordCreateEvent) error {
requestInfo := apis.RequestInfo(e.HttpContext)
// if not an admin, overwrite the newly submitted "posts" record status to pending
if requestInfo.Admin == nil {
e.Record.Set("status", "pending")
return nil
onRecordBeforeCreateRequest((e) => {
let requestInfo = $apis.requestInfo(e.httpContext)
// if not an admin, overwrite the newly submitted "posts" record status to pending
if (!requestInfo.admin) {
e.record.set("status", "pending")
}, "posts")
-Register custom console commands:
JavaScript
app.RootCmd.AddCommand(&amp;cobra.Command{
Use: "hello",
Run: func(cmd *cobra.Command, args []string) {
print("Hello world!")
$app.rootCmd.addCommand(new Command({
use: "hello",
run: (cmd, args) => {
console.log("Hello world!")
-and many more...
For further info, please check the related Extend with Go or
Extend with JavaScript guides.
