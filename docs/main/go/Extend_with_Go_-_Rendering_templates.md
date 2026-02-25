# POCKETBASE DOCS|2026-02-25|1 sections

## 1.Extend with Go - Rendering templates
Response 200:
{{template "placeholderName" .}}
Response 200:
{{block "placeholderName" .}}default...{{end}}
Response 200:
{{define "placeholderName"}}custom...{{end}}
# Extend with Go - Rendering templates
### Overview
A common task when creating custom routes or emails is the need of generating HTML output.
There are plenty of Go template-engines available that you can use for this, but often for simple cases
the Go standard library html/template package should work just fine.
To make it slightly easier to load template files concurrently and on the fly, PocketBase also provides a
thin wrapper around the standard library in the
github.com/pocketbase/pocketbase/tools/template
utility package.
import "github.com/pocketbase/pocketbase/tools/template"
data := map[string]any{"name": "John"}
html, err := template.NewRegistry().LoadFiles(
"views/base.html",
"views/partial1.html",
"views/partial2.html",
).Render(data)
The general flow when working with composed and nested templates is that you create &quot;base&quot; template(s)
The dot object (.) in the above represents the data passed to the templates
via the Render(data) method.
By default the templates apply contextual (HTML, JS, CSS, URI) auto escaping so the generated template
For more information about the template syntax please refer to the
html/template
and
text/template
package godocs.
Another great resource is also the Hashicorp&#39;s
Learn Go Template Syntax
tutorial.
### Example HTML page with layout
Consider the following app directory structure:
myapp/
views/
layout.html
hello.html
main.go
We define the content for layout.html as:
&lt;!DOCTYPE html>
&lt;html lang="en">
&lt;head>
&lt;title>{{block "title" .}}Default app title{{end}}&lt;/title>
&lt;/head>
&lt;body>
Header...
{{block "body" .}}
Default app body...
{{end}}
&lt;/body>
&lt;/html>
We define the content for hello.html as:
{{define "title"}}
Page 1
{{end}}
{{define "body"}}
&lt;p>Hello from {{.name}}&lt;/p>
{{end}}
Then to output the final page, we&#39;ll register a custom /hello/:name route:
// main.go
package main
import (
"log"
"net/http"
"github.com/labstack/echo/v5"
"github.com/pocketbase/pocketbase"
"github.com/pocketbase/pocketbase/apis"
"github.com/pocketbase/pocketbase/core"
"github.com/pocketbase/pocketbase/tools/template"
func main() {
app := pocketbase.New()
app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
// this is safe to be used by multiple goroutines
// (it acts as store for the parsed templates)
registry := template.NewRegistry()
e.Router.GET("/old/hello/:name", func(c echo.Context) error {
name := c.PathParam("name")
html, err := registry.LoadFiles(
"views/layout.html",
"views/hello.html",
).Render(map[string]any{
"name": name,
if err != nil {
// or redirect to a dedicated 404 HTML page
return apis.NewNotFoundError("", err)
return c.HTML(http.StatusOK, html)
return nil
if err := app.Start(); err != nil {
log.Fatal(err)
