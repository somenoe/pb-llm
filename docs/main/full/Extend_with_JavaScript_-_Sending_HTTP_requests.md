# POCKETBASE DOCS|2026-02-25|1 sections

## 1.Extend with JavaScript - Sending HTTP requests
# Extend with JavaScript - Sending HTTP requests
### Overview
You can use the global $http.send(config) helper to send HTTP requests to external services.
This could be used for example to retrieve data from external data sources, to make custom requests to a payment
provider API, etc.
Below is a list with all currently supported config options and their defaults.
// throws on timeout or network connectivity error
const res = $http.send({
url:     "",
method:  "GET",
body:    "", // ex. JSON.stringify({"test": 123}) or new FormData()
headers: {"content-type": "application/json"},
timeout: 120, // in seconds
console.log(res.headers)    // the response headers (ex. res.headers['X-Custom'][0])
console.log(res.cookies)    // the response cookies (ex. res.cookies.sessionId.value)
console.log(res.statusCode) // the response HTTP status code
console.log(res.json)       // the response body as parsed json array or map
Here is an example that will enrich a single book record with some data based on its ISBN details from
openlibrary.org.
onRecordBeforeCreateRequest((e) => {
const isbn = e.record.get("isbn");
// try to update with the published date from the openlibrary API
try {
const res = $http.send({
url: "https://openlibrary.org/isbn/" + isbn + ".json",
if (res.statusCode == 200) {
e.record.set("published", res.json.publish_date)
} catch (err) {
console.log("request failed", err);
}, "books")
##### multipart/form-data requests
In order to send multipart/form-data requests (ex. uploading files) the request
body must be a FormData instance.
PocketBase JSVM&#39;s FormData has the same APIs as its
browser equivalent
with the main difference that for file values instead of Blob it accepts
$filesystem.File.
const formData = new FormData();
formData.append("title", "Hello world!")
formData.append("documents", $filesystem.fileFromBytes("doc1", "doc1.txt"))
formData.append("documents", $filesystem.fileFromBytes("doc2", "doc2.txt"))
const res = $http.send({
url:    "https://...",
method: "POST",
body:   formData,
console.log(res.statusCode)
