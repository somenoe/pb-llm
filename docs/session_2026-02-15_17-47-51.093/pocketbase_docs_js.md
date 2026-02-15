# POCKETBASE DOCS|2026-02-15|39 sections

## 1.Introduction
Please keep in mind that PocketBase is still under active development and full backward
compatibility is not guaranteed before reaching v1.0.0. PocketBase is not recommended for
production critical applications yet, unless you are fine with reading the
and applying some manual migration steps from time to time.
PocketBase is an open source backend consisting of embedded database (SQLite) with realtime subscriptions,
built-in auth management, convenient dashboard UI and simple REST-ish API.
(~14MB zip)
(~14MB zip)
(~14MB zip)
(~13MB zip)
(~13MB zip)
(~13MB zip)
See the
for other platforms and more details.
Once you&#39;ve extracted the archive, you could start the application by running
./pocketbase serve in the extracted directory.
And that&#39;s it! A web server will be started with the following routes:
-http://127.0.0.1:8090
- if pb_public directory exists, serves the static content from it (html, css, images,
etc.)
-http://127.0.0.1:8090/_/
- Admin dashboard UI
-http://127.0.0.1:8090/api/
- REST API
The first time, when you access the Admin dashboard UI, it will prompt you to create your first admin
account (email and pass).
The prebuilt PocketBase executable will automatically create and manage 2 new directories alongside the
executable:
-pb_data - stores your application data, uploaded files, etc. (usually should be added in
.gitignore).
-pb_migrations - contains JS migration files with your collection changes (can be safely
committed in your repository).
You can even write custom migration scripts. For more info check the
JS migrations docs.
You could find all available commands and their options by running
./pocketbase --help or
./pocketbase [command] --help

## 2.Collections
text("", "example"):
editor("", "<p>example</p>"):
number(0, -1, 1, 1.5):+ (add), - (subtract)
bool(false, true):
email("", "test@example.com"):
url("", "https://example.com"):
date("", "2022-01-01 00:00:00.000Z"):
select (single)("", "optionA"):
select (multiple)([], ["optionA", "optionB"]):+ (append), - (remove)
relation (single)("", "JJ2YRU30FBG8MqX"):
relation (multiple)([], ["JJ2YRU30FBG8MqX", "eP2jCr1h3NGtsbz"]):+ (append), - (remove)
file (single)("", "example123_Ab24ZjL.png"):
file (multiple)([],
["file1_Ab24ZjL.png", "file2_Frq24ZjL.txt"]):- (remove)
json(any json value):
- **text** ("", "example"):
- **editor** ("", "<p>example</p>"):
- **number** (0, -1, 1, 1.5): + (add), - (subtract)
- **bool** (false, true):
- **date** ("", "2022-01-01 00:00:00.000Z"):
- **select (single)** ("", "optionA"):
- **select (multiple)** ([], ["optionA", "optionB"]): + (append), - (remove)
- **relation (single)** ("", "JJ2YRU30FBG8MqX"):
- **relation (multiple)** ([], ["JJ2YRU30FBG8MqX", "eP2jCr1h3NGtsbz"]): + (append), - (remove)
- **file (single)** ("", "example123_Ab24ZjL.png"):
- **file (multiple)** ([],
["file1_Ab24ZjL.png", "file2_Frq24ZjL.txt"]): - (remove)
- **json** (any json value):
### Overview
Collections represents your application data.
Under the hood they are plain SQLite tables that are generated automatically with the collection
name
and
fields (aka. columns).
Single entry of a collection is called
record - aka. a single row in the SQL table.
PocketBase comes with all sort of fields that you could use:
Field
Example values
Supported modifiers
`text`
`&quot;&quot;`, `&quot;example&quot;`
`editor`
`&quot;&quot;`, `&quot;&lt;p>example&lt;/p>&quot;`
`number`
`0`, `-1`, `1`, `1.5`
`+` (add), `-` (subtract)
`bool`
`false`, `true`
`email`
`url`
`date`
`&quot;&quot;`, `&quot;2022-01-01 00:00:00.000Z&quot;`
`select` (single)
`&quot;&quot;`, `&quot;optionA&quot;`
`select` (multiple)
`[]`, `[&quot;optionA&quot;, &quot;optionB&quot;]`
`+` (append), `-` (remove)
`relation` (single)
`&quot;&quot;`, `&quot;JJ2YRU30FBG8MqX&quot;`
`relation` (multiple)
`[]`, `[&quot;JJ2YRU30FBG8MqX&quot;, &quot;eP2jCr1h3NGtsbz&quot;]`
`+` (append), `-` (remove)
`file` (single)
`&quot;&quot;`, `&quot;example123_Ab24ZjL.png&quot;`
`file` (multiple)
`[&quot;file1_Ab24ZjL.png&quot;, &quot;file2_Frq24ZjL.txt&quot;]`
`-` (remove)
`json`
any json value
You could create collections and records from the Admin UI or the
Web API.
Usually you&#39;ll create your collections from the Admin UI and manage your records with the API using
the
client-side SDKs.
Here is what the collection panel looks like:
Currently there are 3 collection types: Base, View and
Auth.
### Base collection
Base collection is the default collection type and it could be used to store any application
data (eg. articles, products, posts, etc.).
It comes with 3 default system fields that are always available and automatically populated:
id, created, updated.
Only the id can be explicitly set (15 characters string).
### View collection
View collection is a read-only collection type where the data is populated from a plain
SQL SELECT statement, allowing users to perform aggregations or any other custom queries in
general.
For example, the following query will create a read-only collection with 3 posts
fields - id, name and totalComments:
SELECT
posts.id,
posts.name,
count(comments.id) as totalComments
FROM posts
LEFT JOIN comments on comments.postId = posts.id
GROUP BY posts.id
View collections don&#39;t receive realtime events because they don&#39;t have create/update/delete
operations.
### Auth collection
Auth collection has everything from the Base collection but with some additional
special fields to help you manage your app users providing various authentication options.
Each Auth collection comes with the following system fields:
id, created, updated,
username, email, emailVisibility, verified.
You can have as many Auth collections as you want (eg. users, managers, staffs, members, clients, etc.)
each with their own set of fields, separate login (email/username + password or OAuth2) and models
managing endpoints.
You can create all sort of different access controls:
-Role (Group)
For example, you could attach a &quot;role&quot; select field to your Auth collection with the
following options: &quot;regularUser&quot; and &quot;superUser&quot;. And then in some of your other collections you
could define the following rule to allow only &quot;superUsers&quot;:
@request.auth.role = &quot;superUser&quot;
-Relation (Ownership)
Let&#39;s say that you have 2 collections - &quot;posts&quot; base collection and &quot;users&quot; auth collection. In
your &quot;posts&quot; collection you can create &quot;author&quot;
relation field pointing to the &quot;users&quot; collection. To allow access to only the
&quot;author&quot; of the record(s), you could use a rule like:
@request.auth.id != &quot;&quot; &amp;&amp; author = @request.auth.id
Nested relation fields look ups are also supported, eg:
someRelField.anotherRelField.author = @request.auth.id
-Managed
In addition to the default &quot;List&quot;, &quot;View&quot;, &quot;Create&quot;, &quot;Update&quot;, &quot;Delete&quot; API rules, Auth
collections have also a special &quot;Manage&quot; API rule that could be used to allow one user (it could
be even from a different collection) to be able to fully manage the data of another user (eg.
changing their email, password, etc.).
-Mixed
You can build a mixed approach based on your unique use-case. Multiple rules can be grouped with
parenthesis () and combined with &amp;&amp;
(AND) and || (OR) operators:
@request.auth.id != &quot;&quot; &amp;&amp; (@request.auth.role = &quot;superUser&quot; || author = @request.auth.id)

## 3.API rules and filters
# API rules and filters
### API rules
API Rules are your collection access controls and data filters.
Each collection has 5 rules, corresponding to the specific API action:
-listRule
-viewRule
-createRule
-updateRule
-deleteRule
Auth collections has an additional options.manageRule used to allow one user (it could be even
from a different collection) to be able to fully manage the data of another user (ex. changing their email,
password, etc.).
Each rule could be set to:
-&quot;locked&quot; - aka. null, which means that the action could be performed
only by an authorized admin
(this is the default)
-Empty string - anyone will be able to perform the action (admins, authorized users and
guests)
-Non-empty string - only users (authorized or not) that satisfy the rule filter expression
will be able to perform this action
PocketBase API Rules act also as records filter!
Or in other words, you could for example allow listing only the &quot;active&quot; records of your collection,
by using a simple filter expression such as:
status = &quot;active&quot;
(where &quot;status&quot; is a field defined in your Collection).
Because of the above, the API will return 200 empty items response in case a request doesn&#39;t
satisfy a listRule, 400 for unsatisfied createRule and 404 for
unsatisfied viewRule, updateRule and deleteRule.
All rules will return 403 in case they were &quot;locked&quot; (aka. admin only) and the request client is not
an admin.
The API Rules are ignored when the action is performed by an authorized admin (admins can access everything)!
### Filters syntax
You could find information about the supported fields in your collection API rules tab:
There is autocomplete to help you guide you while typing the rule filter expression, but in general, you
have access to
3 groups of fields:
-Your Collection schema fields
This also include all nested relations fields too, ex.
someRelField.status != &quot;pending&quot;
-@request.*
Used to access the current request data, such as query parameters, body/form data, authorized user state,
etc.
@request.context - the context where the rule is used (ex.
@request.context != &quot;oauth2&quot;)
The currently supported context values are default, oauth2,
realtime, protectedFile.
-@request.method - the HTTP request method (ex.
@request.method = &quot;GET&quot;)
-@request.headers.* - the request headers as string values (ex.
@request.headers.x_token = &quot;test&quot;)
Note: All header keys are normalized to lowercase and &quot;-&quot; is replaced with &quot;_&quot; (for
example &quot;X-Token&quot; is &quot;x_token&quot;).
-@request.query.* - the request query parameters as string values (ex.
@request.query.page = &quot;1&quot;)
-@request.auth.* - the current authenticated model (ex.
@request.auth.id != &quot;&quot;)
-@request.data.* - the submitted body parameters (ex.
@request.data.title != &quot;&quot;)
Note: Uploaded files are not part of the @request.data
because they are evaluated separately (this behavior may change in the future).
-@collection.*
This filter could be used to target other collections that are not directly related to the current
one (aka. there is no relation field pointing to it) but both shares a common field value, like
for example a category id:
@collection.news.categoryId ?= categoryId &amp;&amp; @collection.news.author ?= @request.auth.id
In case you want to join the same collection multiple times but based on different criteria, you
can define an alias by appending :alias suffix to the collection name.
// see https://github.com/pocketbase/pocketbase/discussions/3805#discussioncomment-7634791
@request.auth.id != "" &amp;&amp;
@collection.courseRegistrations.user ?= id &amp;&amp;
@collection.courseRegistrations:auth.user ?= @request.auth.id &amp;&amp;
@collection.courseRegistrations.courseGroup ?= @collection.courseRegistrations:auth.courseGroup
The syntax basically follows the format
OPERAND
OPERATOR
OPERAND, where:
-OPERAND - could be any of the above field literal, string (single or double
quoted), number, null, true, false
-OPERATOR - is one of:
Equal
NOT equal
Greater than
Greater than or equal
-&lt;
Less than
-&lt;=
Less than or equal
Like/Contains (if not specified auto wraps the right string OPERAND in a &quot;%&quot; for wildcard
match)
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a &quot;%&quot; for
wildcard match)
Any/At least one of
Equal
Any/At least one of
NOT equal
Any/At least one of
Greater than
Any/At least one of
Greater than or equal
-?&lt;
Any/At least one of
Less than
-?&lt;=
Any/At least one of
Less than or equal
Any/At least one of
Like/Contains (if not specified auto wraps the right string OPERAND in a &quot;%&quot; for wildcard
match)
-?!~
Any/At least one of
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a &quot;%&quot; for
wildcard match)
To group and combine several expressions you could use parenthesis
(...), &amp;&amp; (AND) and || (OR) tokens.
Single line comments are also supported: // Example comment.
### Special identifiers and modifiers
##### @ macros
The following datetime macros are available and can be used as part of the filter expression:
// all macros are UTC based
@now        - the current datetime as string
@second     - @now second number (0-59)
@minute     - @now minute number (0-59)
@hour       - @now hour number (0-23)
@weekday    - @now weekday number (0-6)
@day        - @now day number
@month      - @now month number
@year       - @now year number
@todayStart - beginning of the current day as datetime string
@todayEnd   - end of the current day as datetime string
@monthStart - beginning of the current month as datetime string
@monthEnd   - end of the current month as datetime string
@yearStart  - beginning of the current year as datetime string
@yearEnd    - end of the current year as datetime string
For example:
`@request.data.publicDate >= @now`
##### :isset modifier
The :isset field modifier is available only for the @request.* fields and can be
used to check whether the client submitted a specific data with the request. Here is for example a rule that
disallows changing a &quot;role&quot; field:
`@request.data.role:isset = false`
Note that @request.data.*:isset at the moment doesn&#39;t support checking for
new uploaded files because they are evaluated separately and cannot be serialized (this behavior may change in the future).
##### :length modifier
The :length field modifier could be used to check the number of items in an array field
(multiple file, select, relation).
Could be used with both the collection schema fields and the @request.data.* fields. For example:
// check example submitted data: {"someSelectField": ["val1", "val2"]}
@request.data.someSelectField:length > 1
// check existing record field length
someRelationField:length = 2
Note that @request.data.*:length at the moment doesn&#39;t support checking
for new uploaded files because they are evaluated separately and cannot be serialized (this behavior may change in the future).
##### :each modifier
The :each field modifier works only with multiple select, file and
relation
type fields. It could be used to apply a condition on each item from the field array. For example:
// check if all submitted select options contain the "create" text
@request.data.someSelectField:each ~ "create"
// check if all existing someSelectField has "pb_" prefix
someSelectField:each ~ "pb_%"
Note that @request.data.*:each at the moment doesn&#39;t support checking for
new uploaded files because they are evaluated separately and cannot be serialized (this behavior may change in the future).
-Allow only registered users:
@request.auth.id != ""
-Allow only registered users and return records that are either &quot;active&quot; or &quot;pending&quot;:
@request.auth.id != "" &amp;&amp; (status = "active" || status = "pending")
-Allow only registered users who are listed in an allowed_users multi-relation field value:
@request.auth.id != "" &amp;&amp; allowed_users.id ?= @request.auth.id
-Allow access by anyone and return only the records where the title field value starts with
title ~ "Lorem%"

## 4.Client-side SDKs
# Client-side SDKs
The easiest way to interact with the PocketBase API is to use one of the official SDK clients:
-JavaScript SDK
(browser and node)
-Dart SDK
(web, mobile, desktop and cli)
You could find usage examples in each API section, but most of the time you will work with the
Records,
Files and
Realtime APIs.
PocketBase also generates individual documentation for your Collection Records with examples that you
could copy-paste by clicking on the API Preview button from the Admin UI:

## 5.Authentication
The PocketBase API uses JWT for authentication via the Authorization HTTP header:
Authorization: TOKEN.
You can also use the dedicated auth SDK helpers as shown in the examples below.
### Authenticate as admin
You can authenticate as admin using an email and password.
Admins can access everything and API rules don&#39;t apply to them.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
// after the above you can also access the auth data from the authStore
console.log(pb.authStore.isValid);
console.log(pb.authStore.token);
console.log(pb.authStore.model.id);
// "logout" the last authenticated account
pb.authStore.clear();
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
// after the above you can also access the auth data from the authStore
print(pb.authStore.isValid);
print(pb.authStore.token);
print(pb.authStore.model.id);
// "logout" the last authenticated account
pb.authStore.clear();
### Authenticate as app user
The easiest way to authenticate your app users is with their username/email and password.
You can customize the supported authentication options from your Auth collection configuration
(including disabling all auth options).
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('https://pocketbase.io');
const authData = await pb.collection('users').authWithPassword('YOUR_USERNAME_OR_EMAIL', '1234567890');
// after the above you can also access the auth data from the authStore
console.log(pb.authStore.isValid);
console.log(pb.authStore.token);
console.log(pb.authStore.model.id);
// "logout" the last authenticated model
pb.authStore.clear();
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('https://pocketbase.io');
final authData = await pb.collection('users').authWithPassword('YOUR_USERNAME_OR_EMAIL', '1234567890');
// after the above you can also access the auth data from the authStore
print(pb.authStore.isValid);
print(pb.authStore.token);
print(pb.authStore.model.id);
// "logout" the last authenticated model
pb.authStore.clear();
You can also authenticate your users with an OAuth2 provider (Google, GitHub, Microsoft, etc.). See the
section below for an example OAuth2 web integration.
### OAuth2 integration
Before starting, you&#39;ll need to create an OAuth2 app in the provider&#39;s dashboard in order to get a
Client Id and Client Secret, and register a redirect URL
Once you have obtained the Client Id and Client Secret, you can
enable and configure the provider from your PocketBase admin settings page.
This method handles everything within a single call without having to define custom redirects,
deeplinks or even page reload.
When creating your OAuth2 app, for a callback/redirect URL you have to use the
https://yourdomain.com/api/oauth2-redirect
(or when testing locally - http://127.0.0.1:8090/api/oauth2-redirect
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('https://pocketbase.io');
// This method initializes a one-off realtime subscription and will
// open a popup window with the OAuth2 vendor page to authenticate.
// Once the external OAuth2 sign-in/sign-up flow is completed, the popup
// window will be automatically closed and the OAuth2 data sent back
// If the popup is being blocked on Safari, you can try the suggestion from:
// https://github.com/pocketbase/pocketbase/discussions/2429#discussioncomment-5943061.
const authData = await pb.collection('users').authWithOAuth2({ provider: 'google' });
// after the above you can also access the auth data from the authStore
console.log(pb.authStore.isValid);
console.log(pb.authStore.token);
console.log(pb.authStore.model.id);
// "logout" the last authenticated model
pb.authStore.clear();
import 'package:pocketbase/pocketbase.dart';
import 'package:url_launcher/url_launcher.dart';
final pb = PocketBase('https://pocketbase.io');
// This method initializes a one-off realtime subscription and will
// call the provided urlCallback with the OAuth2 vendor url to authenticate.
// Once the external OAuth2 sign-in/sign-up flow is completed, the browser
// window will be automatically closed and the OAuth2 data sent back
final authData = await pb.collection('users').authWithOAuth2('google', (url) async {
// or use something like flutter_custom_tabs to make the transitions between native and web content more seamless
await launchUrl(url);
// after the above you can also access the auth data from the authStore
print(pb.authStore.isValid);
print(pb.authStore.token);
print(pb.authStore.model.id);
// "logout" the last authenticated model
pb.authStore.clear();
When authenticating manually with OAuth2 code you&#39;ll need 2 endpoints:
-somewhere to show the &quot;Login with ...&quot; links
-somewhere to handle the provider&#39;s redirect in order to exchange the auth code for token
Here is a simple web example:
-Links page
(eg. https://127.0.0.1:8090 serving pb_public/index.html):
&lt;!DOCTYPE html>
&lt;html>
&lt;head>
&lt;meta charset="utf-8" />
&lt;meta name="viewport" content="width=device-width, initial-scale=1" />
&lt;title>OAuth2 links page&lt;/title>
&lt;script src="https://code.jquery.com/jquery-3.6.0.slim.min.js">&lt;/script>
&lt;/head>
&lt;body>
&lt;ul id="list">
&lt;li>Loading OAuth2 providers...&lt;/li>
&lt;/ul>
&lt;script src="https://cdn.jsdelivr.net/gh/pocketbase/js-sdk@master/dist/pocketbase.umd.js">&lt;/script>
&lt;script type="text/javascript">
const pb = new PocketBase('http://127.0.0.1:8090');
const redirectUrl = 'http://127.0.0.1:8090/redirect.html';
async function loadLinks() {
const authMethods = await pb.collection('users').listAuthMethods();
const listItems = [];
for (const provider of authMethods.authProviders) {
const $li = $(`&lt;li>&lt;a>Login with ${provider.name}&lt;/a>&lt;/li>`);
$li.find('a')
.attr('href', provider.authUrl + redirectUrl)
.data('provider', provider)
.click(function () {
// store provider's data on click for verification in the redirect page
localStorage.setItem('provider', JSON.stringify($(this).data('provider')));
listItems.push($li);
$('#list').html(listItems.length ? listItems : '&lt;li>No OAuth2 providers.&lt;/li>');
loadLinks();
&lt;/script>
&lt;/body>
&lt;/html>
-Redirect handler page
(eg. https://127.0.0.1:8090/redirect.html serving
pb_public/redirect.html):
&lt;!DOCTYPE html>
&lt;html>
&lt;head>
&lt;meta charset="utf-8">
&lt;title>OAuth2 redirect page&lt;/title>
&lt;/head>
&lt;body>
&lt;pre id="content">Authenticating...&lt;/pre>
&lt;script src="https://cdn.jsdelivr.net/gh/pocketbase/js-sdk@master/dist/pocketbase.umd.js">&lt;/script>
&lt;script type="text/javascript">
const pb = new PocketBase("http://127.0.0.1:8090");
const redirectUrl = 'http://127.0.0.1:8090/redirect.html';
// parse the query parameters from the redirected url
const params = (new URL(window.location)).searchParams;
const provider = JSON.parse(localStorage.getItem('provider'))
// compare the redirect's state param and the stored provider's one
if (provider.state !== params.get('state')) {
throw "State parameters don't match.";
// authenticate
pb.collection('users').authWithOAuth2Code(
provider.name,
params.get('code'),
provider.codeVerifier,
redirectUrl,
// pass optional user create data
emailVisibility: false,
).then((authData) => {
document.getElementById('content').innerText = JSON.stringify(authData, null, 2);
}).catch((err) => {
document.getElementById('content').innerText = "Failed to exchange code.\n" + err;
&lt;/script>
&lt;/body>
&lt;/html>
When using the &quot;Manual code exchange&quot; flow for sign-in with Apple your redirect
handler must accept POST requests in order to receive the name and the
email of the Apple user. If you just need the Apple user id, you can keep the redirect
hanldler GET but you&#39;ll need to replace in the Apple authorization url
response_mode=form_post with response_mode=query.

## 6.Files upload and handling
### Uploading files
To upload files, you must first add a file field to your collection:
Once added, you could create/update a Record and upload &quot;documents&quot; files by sending a
multipart/form-data request using the Records create/update APIs.
The client SDK makes things a little easier and auto detect the request content type based on the
parameters that you provide. Here is an example how to create a new Record and upload multiple files to
the example file field &quot;documents&quot; using the SDKs:
Dart
// Example HTML:
// &lt;input type="file" id="fileInput" />
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const formData = new FormData();
const fileInput = document.getElementById('fileInput');
// listen to file input changes and add the selected files to the form data
fileInput.addEventListener('change', function () {
for (let file of fileInput.files) {
formData.append('documents', file);
// set some other regular text field value
formData.append('title', 'Hello world!');
// upload and create new record
const createdRecord = await pb.collection('example').create(formData);
import 'package:pocketbase/pocketbase.dart';
import 'package:http/http.dart' as http;
final pb = PocketBase('http://127.0.0.1:8090');
// create a new record and upload multiple files
final record = await pb.collection('example').create(
body: {
'title': 'Hello world!', // some regular text field
files: [
http.MultipartFile.fromString(
'document',
'example content 1...',
filename: 'file1.txt',
http.MultipartFile.fromString(
'document',
'example content 2...',
filename: 'file2.txt',
Each uploaded file will be stored with the original filename (sanitized) and suffixed with a
random (10 chars) part (eg. test_52iWbGinWd.png).
When you upload a new file to a single file upload field (aka.
file (if any) and upload the new one in its place.
When you upload a new file to a multiple file upload field (aka.
Max Files option is &gt; 1) the new file will be appended to the existing field
values (as long as the Max Files limit is not reached, otherwise an error will be
thrown).
### Deleting files
To delete uploaded file(s), you could either edit the Record from the admin UI, or use the API and set the
file field to a zero-value  (null, [], empty string, etc.).
If you want to delete individual file(s) from a multiple file upload field, you could
suffix the field name with - and specify the filename(s) you want to delete. Here are some examples
using the SDKs:
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
// delete all "documents" files
await pb.collection('example').update('RECORD_ID', {
'documents': null,
// delete individual files
await pb.collection('example').update('RECORD_ID', {
'documents-': ["file1.pdf", "file2.txt"],
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
// delete all "documents" files
await pb.collection('example').update('RECORD_ID', body: {
'documents': null,
// delete individual files
await pb.collection('example').update('RECORD_ID', body: {
'documents-': ["file1.pdf", "file2.txt"],
The above examples use the JSON object data format, but you could also use FormData instance
for multipart/form-data requests. If using
FormData set the file field to an empty string.
### File URL
Each uploaded file could be accessed by requesting its file url:
http://127.0.0.1:8090/api/files/COLLECTION_ID_OR_NAME/RECORD_ID/FILENAME
If your file field has the Thumb sizes option, you can get a thumb of the image file
(currently limited to jpg, png, and partially gif – its first frame) by adding the thumb
query parameter to the url like this:
http://127.0.0.1:8090/api/files/COLLECTION_ID_OR_NAME/RECORD_ID/FILENAME?thumb=100x300
The following thumb formats are currently supported:
-WxH
(eg. 100x300) - crop to WxH viewbox (from center)
-WxHt
(eg. 100x300t) - crop to WxH viewbox (from top)
-WxHb
(eg. 100x300b) - crop to WxH viewbox (from bottom)
-WxHf
(eg. 100x300f) - fit inside a WxH viewbox (without cropping)
-0xH
(eg. 0x300) - resize to H height preserving the aspect ratio
-Wx0
(eg. 100x0) - resize to W width preserving the aspect ratio
The original file would be returned, if the requested thumb size is not found or the file is not an image!
If you already have a Record model instance, the SDKs provide a convenient method to generate a file url
by its name.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const record = await pb.collection('example').getOne('RECORD_ID');
// get only the first filename from "documents"
// note:
// "documents" is an array of filenames because
// the "documents" field was created with "Max Files" option > 1;
// if "Max Files" was 1, then the result property would be just a string
const firstFilename = record.documents[0];
// returns something like:
// http://127.0.0.1:8090/api/files/example/kfzjt5oy8r34hvn/test_52iWbGinWd.png?thumb=100x250
const url = pb.files.getUrl(record, firstFilename, {'thumb': '100x250'});
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
final record = await pb.collection('example').getOne('RECORD_ID');
// get only the first filename from "documents"
// note:
// "documents" is an array of filenames because
// the "documents" field was created with "Max Files" option > 1;
// if "Max Files" was 1, then the result property would be just a string
final firstFilename = record.getListValue&lt;String>('documents')[0];
// returns something like:
// http://127.0.0.1:8090/api/files/example/kfzjt5oy8r34hvn/test_52iWbGinWd.png?thumb=100x250
final url = pb.files.getUrl(record, firstFilename, thumb: '100x250');
### Protected files
By default all files are public accessible if you know their full url.
For most applications this is fine since all files have a random part, but in some cases you may want an
extra security to prevent unauthorized access to sensitive files like ID card or Passport copies,
contracts, etc.
To do this you can mark the file field as Protected and then request the file with a
special short-lived file token.
Only requests that satisfy the View API rule of the record collection will be able
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
// authenticate
// generate a file token
const fileToken = await pb.files.getToken();
// retrieve an example protected file url (will be valid ~2min)
const record = await pb.collection('example').getOne('RECORD_ID');
const url = pb.files.getUrl(record, record.myPrivateFile, {'token': fileToken});
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
// authenticate
// generate a file token
final fileToken = await pb.files.getToken();
// retrieve an example protected file url (will be valid ~2min)
final record = await pb.collection('example').getOne('RECORD_ID');
final url = pb.files.getUrl(record, record.getStringValue('myPrivateFile'), token: fileToken);
### Storage options
By default PocketBase uses the local file system to store uploaded files (in the
pb_data/storage directory).
If you have limited disk space, you could use a S3 compatible storage (AWS S3, MinIO, Wasabi, DigitalOcean
Spaces, Vultr Object Storage, etc.). The easiest way to setup the connection settings is from the admin UI
(Settings &gt; Files storage):

## 7.Working with relations
# Working with relations
### Overview
Let&#39;s assume that we have the following collections structure:
The relation fields follow the same rules as any other collection field and can be set/modified
by directly updating the field value - with a record id or array of ids, in case a multiple relation is used.
Below is an example that shows creating a new posts record with 2 assigned tags.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const post = await pb.collection('posts').create({
'tags':  ['TAG_ID1', 'TAG_ID2'],
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
final post = await pb.collection('posts').create(body: {
'tags':  ['TAG_ID1', 'TAG_ID2'],
### Append to multiple relation
To append a single or multiple relation id(s) to an existing value you can use the
+ field modifier:
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const post = await pb.collection('posts').update('POST_ID', {
// append single tag
'tags+': 'TAG_ID1',
// append multiple tags at once
'tags+': ['TAG_ID1', 'TAG_ID2'],
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
final post = await pb.collection('posts').update('POST_ID', body: {
// append single tag
'tags+': 'TAG_ID1',
// append multiple tags at once
'tags+': ['TAG_ID1', 'TAG_ID2'],
### Remove from multiple relation
To remove a single or multiple relation id(s) from an existing value you can use the
- field modifier:
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const post = await pb.collection('posts').update('POST_ID', {
// remove single tag
'tags-': 'TAG_ID1',
// remove multiple tags at once
'tags-': ['TAG_ID1', 'TAG_ID2'],
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
final post = await pb.collection('posts').update('POST_ID', body: {
// remove single tag
'tags-': 'TAG_ID1',
// remove multiple tags at once
'tags-': ['TAG_ID1', 'TAG_ID2'],
### Expanding relations
You can also expand record relation fields directly in the returned response without making additional
requests by using the expand query parameter, eg. ?expand=user,post.tags
Only the relations that the request client can View (aka. satisfies the relation
collection&#39;s View API Rule) will be expanded.
Nested relation references in expand, filter or sort are supported
via dot-notation and up to 6-levels depth.
For example, to list all comments with their user relation expanded, we can
do the following:
Dart
`await pb.collection("comments").getList(1, 30, { expand: "user" })`
`await pb.collection("comments").getList(perPage: 30, expand: "user")`
"page": 1,
"perPage": 30,
"totalPages": 1,
"totalItems": 20,
"items": [
"id": "lmPJt4Z9CkLW36z",
"collectionId": "BHKW36mJl3ZPt6z",
"collectionName": "comments",
"created": "2022-01-01 01:00:00.456Z",
"updated": "2022-01-01 02:15:00.456Z",
"post": "WyAw4bDrvws6gGl",
"user": "FtHAW9feB5rze7D",
"message": "Example message...",
"expand": {
"user": {
"id": "FtHAW9feB5rze7D",
"collectionId": "srmAo0hLxEqYF7F",
"collectionName": "users",
"created": "2022-01-01 00:00:00.000Z",
"updated": "2022-01-01 00:00:00.000Z",
"username": "users54126",
"verified": false,
"emailVisibility": false,
"name": "John Doe"
### Back-relations
PocketBase supports also filter, sort and expand for
back-relations
- relations where the associated relation field is not in the main collection.
The following notation is used: referenceCollection_via_relField (ex.
comments_via_post).
For example, lets list the posts that has at least one comments record
containing the word &quot;hello&quot;:
Dart
await pb.collection("posts").getList(1, 30, {
filter: "comments_via_post.message ?~ 'hello'"
expand: "comments_via_post.user",
await pb.collection("posts").getList(
perPage: 30,
filter: "comments_via_post.message ?~ 'hello'"
expand: "comments_via_post.user",
"page": 1,
"perPage": 30,
"totalPages": 2,
"totalItems": 45,
"items": [
"id": "WyAw4bDrvws6gGl",
"collectionName": "posts",
"created": "2022-01-01 01:00:00.456Z",
"updated": "2022-01-01 02:15:00.456Z",
"expand": {
"comments_via_post": [
"id": "lmPJt4Z9CkLW36z",
"collectionId": "BHKW36mJl3ZPt6z",
"collectionName": "comments",
"created": "2022-01-01 01:00:00.456Z",
"updated": "2022-01-01 02:15:00.456Z",
"post": "WyAw4bDrvws6gGl",
"user": "FtHAW9feB5rze7D",
"expand": {
"user": {
"id": "FtHAW9feB5rze7D",
"collectionId": "srmAo0hLxEqYF7F",
"collectionName": "users",
"created": "2022-01-01 00:00:00.000Z",
"updated": "2022-01-01 00:00:00.000Z",
"username": "users54126",
"verified": false,
"emailVisibility": false,
"name": "John Doe"
"id": "tu4Z9CkLW36mPJz",
"collectionId": "BHKW36mJl3ZPt6z",
"collectionName": "comments",
"created": "2022-01-01 01:10:00.123Z",
"updated": "2022-01-01 02:39:00.456Z",
"post": "WyAw4bDrvws6gGl",
"user": "FtHAW9feB5rze7D",
"message": "hello...",
"expand": {
"user": {
"id": "FtHAW9feB5rze7D",
"collectionId": "srmAo0hLxEqYF7F",
"collectionName": "users",
"created": "2022-01-01 00:00:00.000Z",
"updated": "2022-01-01 00:00:00.000Z",
"username": "users54126",
"verified": false,
"emailVisibility": false,
"name": "John Doe"
###### Back-relation caveats
-By default the back-relation reference is resolved as a dynamic
multiple relation field, even when the back-relation field itself is marked as
single.
This is because the main record could have more than one single
back-relation reference (see in the above example that the comments_via_post
expand is returned as array, although the original comments.post field is a
single relation).
The only case where the back-relation will be treated as a single
relation field is when there is
UNIQUE index constraint defined on the relation field.
-Back-relation expand is limited to max 1000 records per relation field. If you
need to fetch larger number of back-related records a better approach could be to send a
separate paginated getList() request to the back-related collection to avoid transferring
large JSON payloads and to reduce the memory usage.

## 8.Use as framework
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

## 9.Web APIs reference - API Records
collectionIdOrName(String):ID or name of the records' collection.
page(Number):The page (aka. offset) of the paginated list (default to 1).
perPage(Number):The max returned records per page (default to 30).
sort(String):Specify the ORDER BY fields.
Add - / + (default) in front of the attribute for DESC /
ASC order, eg.:
// DESC by created and ASC by id
?sort=-created,id
Supported record sort fields:
@random, id, created, updated,
and any other field from the collection schema.
filter(String):Filter expression to filter/search the returned records list (in addition to the
collection's listRule), eg.:
?filter=(title~'abc' && created>'2022-01-01')
Supported record filter fields:
id, created, updated,
+ any field from the collection schema.
The syntax basically follows the format
OPERAND
OPERATOR
OPERAND, where:
OPERAND - could be any of the above field literal, string (single or double
quoted), number, null, true, false
OPERATOR - is one of:
Equal
!=
NOT equal
Greater than
>=
Greater than or equal
Less than
<=
Less than or equal
Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for wildcard
match)
!~
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for
wildcard match)
?=
Any/At least one of
Equal
?!=
Any/At least one of
NOT equal
?>
Any/At least one of
Greater than
?>=
Any/At least one of
Greater than or equal
?<
Any/At least one of
Less than
?<=
Any/At least one of
Less than or equal
?~
Any/At least one of
Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for wildcard
match)
?!~
Any/At least one of
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for
wildcard match)
To group and combine several expressions you could use parenthesis
(...), && (AND) and || (OR) tokens.
Single line comments are also supported: // Example comment.
expand(String):Auto expand record relations. Ex.:
?expand=relField1,relField2.subRelField
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
expand property (eg. "expand": {"relField1": {...}, ...}).
Only the relations to which the request user has permissions to view will
be expanded.
fields(String):Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
skipTotal(Boolean):If it is set the total counts query will be skipped and the response fields
totalItems and totalPages will have -1 value.
This could drastically speed up the search queries when the total counters are not needed or cursor based
pagination is used.
For optimization purposes, it is set by default for the
getFirstListItem()
and
getFullList() SDKs methods.
collectionIdOrName(String):ID or name of the record's collection.
recordId(String):ID of the record to view.
expand(String):Auto expand record relations. Ex.:
?expand=relField1,relField2.subRelField
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
expand property (eg. "expand": {"relField1": {...}, ...}).
Only the relations to which the request user has permissions to view will
be expanded.
fields(String):Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
collectionIdOrName(String):ID or name of the record's collection.
Optional
id(String):15 characters string to store as record ID.
If not set, it will be auto generated.
Optional
username(String):The username of the auth record.
If not set, it will be auto generated.
Optional
email(String):Auth record email address.
Optional
emailVisibility(Boolean):Whether to show/hide the auth record email when fetching the record data.
Required
password(String):Auth record password.
Required
passwordConfirm(String):Auth record password confirmation.
Optional
verified(Boolean):Indicates whether the auth record is verified or not.
This field can be set only by admins or auth records with "Manage" access.
expand(String):Auto expand record relations. Ex.:
?expand=relField1,relField2.subRelField
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
expand property (eg. "expand": {"relField1": {...}, ...}).
Only the relations to which the request user has permissions to view will
be expanded.
fields(String):Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
collectionIdOrName(String):ID or name of the record's collection.
recordId(String):ID of the record to update.
Optional
username(String):The username of the auth record.
Optional
email(String):The auth record email address.
This field can be updated only by admins or auth records with "Manage" access.
Regular accounts can update their email by calling "Request email change".
Optional
emailVisibility(Boolean):Whether to show/hide the auth record email when fetching the record data.
Optional
oldPassword*(String):Old auth record password.
This field is required only when changing the record password. Admins and auth records with
"Manage" access can skip this field.
Optional
password(String):New auth record password.
Optional
passwordConfirm(String):New auth record password confirmation.
Optional
verified(Boolean):Indicates whether the auth record is verified or not.
This field can be set only by admins or auth records with "Manage" access.
expand(String):Auto expand record relations. Ex.:
?expand=relField1,relField2.subRelField
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
expand property (eg. "expand": {"relField1": {...}, ...}).
Only the relations to which the request user has permissions to view will
be expanded.
fields(String):Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
collectionIdOrName(String):ID or name of the record's collection.
recordId(String):ID of the record to delete.
collectionIdOrName(String):ID or name of the auth collection.
fields(String):Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
collectionIdOrName(String):ID or name of the auth collection.
Required
identity(String):Auth record username or email address.
Required
password(String):Auth record password.
expand(String):Auto expand record relations. Ex.:
?expand=relField1,relField2.subRelField
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
expand property (eg. "expand": {"relField1": {...}, ...}).
Only the relations to which the request user has permissions to view will
be expanded.
fields(String):Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,record.expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,record.description:excerpt(200,true)
collectionIdOrName(String):ID or name of the auth collection.
Required
provider(String):The name of the OAuth2 client provider (eg. "google").
Required
code(String):The authorization code returned from the initial request.
Required
codeVerifier(String):The code verifier sent with the initial request as part of the code_challenge.
Required
redirectUrl(String):The redirect url sent with the initial request.
Optional
createData(Object):Optional data that will be used when creating the auth record on OAuth2 sign-up.
The created auth record must comply with the same requirements and validations in the
regular create action.
The data can only be in json, aka. multipart/form-data and
files upload currently are not supported during OAuth2 sign-ups.
expand(String):Auto expand record relations. Ex.:
?expand=relField1,relField2.subRelField
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
expand property (eg. "expand": {"relField1": {...}, ...}).
Only the relations to which the request user has permissions to view will
be expanded.
fields(String):Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,record.expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,record.description:excerpt(200,true)
collectionIdOrName(String):ID or name of the auth collection.
expand(String):Auto expand record relations. Ex.:
?expand=relField1,relField2.subRelField
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
expand property (eg. "expand": {"relField1": {...}, ...}).
Only the relations to which the request user has permissions to view will
be expanded.
fields(String):Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,record.expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,record.description:excerpt(200,true)
collectionIdOrName(String):ID or name of the auth collection.
Required
email(String):The email address to send the password reset request (if registered).
collectionIdOrName(String):ID or name of the auth collection.
Required
token(String):The token from the verification request email.
Required
email(String):The email address to send the password reset request (if registered).
collectionIdOrName(String):ID or name of the auth collection.
collectionIdOrName(String):ID or name of the auth collection.
Required
token(String):The token from the password reset request email.
Required
password(String):The new auth record password to set.
Required
passwordConfirm(String):New auth record password confirmation.
collectionIdOrName(String):ID or name of the auth collection.
Required
newEmail(String):The new email address to send the change email request.
collectionIdOrName(String):ID or name of the auth collection.
Required
token(String):The token from the change email request.
Required
password(String):The auth record password to confirm the email address change.
collectionIdOrName(String):ID or name of the auth collection.
id(String):ID of the auth record.
fields(String):Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
collectionIdOrName(String):ID or name of the auth collection.
id(String):ID of the auth record.
provider(String):The name of the auth provider to unlink, eg. google, twitter,
github, etc.
# Web APIs reference - API Records
- **collectionIdOrName** (String): ID or name of the records' collection.
- **page** (Number): The page (aka. offset) of the paginated list (default to 1).
- **perPage** (Number): The max returned records per page (default to 30).
- **sort** (String): Specify the ORDER BY fields.
Add - / + (default) in front of the attribute for DESC /
ASC order, eg.:
// DESC by created and ASC by id
?sort=-created,id
Supported record sort fields:
@random, id, created, updated,
and any other field from the collection schema.
- **filter** (String): Filter expression to filter/search the returned records list (in addition to the
collection's listRule), eg.:
?filter=(title~'abc' && created>'2022-01-01')
Supported record filter fields:
id, created, updated,
+ any field from the collection schema.
The syntax basically follows the format
OPERAND
OPERATOR
OPERAND, where:
OPERAND - could be any of the above field literal, string (single or double
quoted), number, null, true, false
OPERATOR - is one of:
Equal
NOT equal
Greater than
Greater than or equal
Less than
Less than or equal
Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for wildcard
match)
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for
wildcard match)
Any/At least one of
Equal
Any/At least one of
NOT equal
Any/At least one of
Greater than
Any/At least one of
Greater than or equal
Any/At least one of
Less than
Any/At least one of
Less than or equal
Any/At least one of
Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for wildcard
match)
?!~
Any/At least one of
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for
wildcard match)
To group and combine several expressions you could use parenthesis
(...), && (AND) and || (OR) tokens.
Single line comments are also supported: // Example comment.
- **expand** (String): Auto expand record relations. Ex.:
?expand=relField1,relField2.subRelField
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
expand property (eg. "expand": {"relField1": {...}, ...}).
Only the relations to which the request user has permissions to view will
be expanded.
- **fields** (String): Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
- **skipTotal** (Boolean): If it is set the total counts query will be skipped and the response fields
totalItems and totalPages will have -1 value.
This could drastically speed up the search queries when the total counters are not needed or cursor based
pagination is used.
For optimization purposes, it is set by default for the
getFirstListItem()
and
getFullList() SDKs methods.
- **collectionIdOrName** (String): ID or name of the record's collection.
- **recordId** (String): ID of the record to view.
- **expand** (String): Auto expand record relations. Ex.:
?expand=relField1,relField2.subRelField
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
expand property (eg. "expand": {"relField1": {...}, ...}).
Only the relations to which the request user has permissions to view will
be expanded.
- **fields** (String): Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
- **collectionIdOrName** (String): ID or name of the record's collection.
- **Optional
id** (String): 15 characters string to store as record ID.
If not set, it will be auto generated.
- **Optional
username** (String): The username of the auth record.
If not set, it will be auto generated.
- **Optional
email** (String): Auth record email address.
- **Optional
emailVisibility** (Boolean): Whether to show/hide the auth record email when fetching the record data.
- **Required
password** (String): Auth record password.
- **Required
passwordConfirm** (String): Auth record password confirmation.
- **Optional
verified** (Boolean): Indicates whether the auth record is verified or not.
This field can be set only by admins or auth records with "Manage" access.
- **expand** (String): Auto expand record relations. Ex.:
?expand=relField1,relField2.subRelField
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
expand property (eg. "expand": {"relField1": {...}, ...}).
Only the relations to which the request user has permissions to view will
be expanded.
- **fields** (String): Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
- **collectionIdOrName** (String): ID or name of the record's collection.
- **recordId** (String): ID of the record to update.
- **Optional
username** (String): The username of the auth record.
- **Optional
email** (String): The auth record email address.
This field can be updated only by admins or auth records with "Manage" access.
Regular accounts can update their email by calling "Request email change".
- **Optional
emailVisibility** (Boolean): Whether to show/hide the auth record email when fetching the record data.
- **Optional
oldPassword** (String) (required): Old auth record password.
This field is required only when changing the record password. Admins and auth records with
"Manage" access can skip this field.
- **Optional
password** (String): New auth record password.
- **Optional
passwordConfirm** (String): New auth record password confirmation.
- **Optional
verified** (Boolean): Indicates whether the auth record is verified or not.
This field can be set only by admins or auth records with "Manage" access.
- **expand** (String): Auto expand record relations. Ex.:
?expand=relField1,relField2.subRelField
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
expand property (eg. "expand": {"relField1": {...}, ...}).
Only the relations to which the request user has permissions to view will
be expanded.
- **fields** (String): Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
- **collectionIdOrName** (String): ID or name of the record's collection.
- **recordId** (String): ID of the record to delete.
- **collectionIdOrName** (String): ID or name of the auth collection.
- **fields** (String): Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
- **collectionIdOrName** (String): ID or name of the auth collection.
- **Required
identity** (String): Auth record username or email address.
- **Required
password** (String): Auth record password.
- **expand** (String): Auto expand record relations. Ex.:
?expand=relField1,relField2.subRelField
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
expand property (eg. "expand": {"relField1": {...}, ...}).
Only the relations to which the request user has permissions to view will
be expanded.
- **fields** (String): Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,record.expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,record.description:excerpt(200,true)
- **collectionIdOrName** (String): ID or name of the auth collection.
- **Required
provider** (String): The name of the OAuth2 client provider (eg. "google").
- **Required
code** (String): The authorization code returned from the initial request.
- **Required
codeVerifier** (String): The code verifier sent with the initial request as part of the code_challenge.
- **Required
redirectUrl** (String): The redirect url sent with the initial request.
- **Optional
createData** (Object): Optional data that will be used when creating the auth record on OAuth2 sign-up.
The created auth record must comply with the same requirements and validations in the
regular create action.
The data can only be in json, aka. multipart/form-data and
files upload currently are not supported during OAuth2 sign-ups.
- **expand** (String): Auto expand record relations. Ex.:
?expand=relField1,relField2.subRelField
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
expand property (eg. "expand": {"relField1": {...}, ...}).
Only the relations to which the request user has permissions to view will
be expanded.
- **fields** (String): Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,record.expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,record.description:excerpt(200,true)
- **collectionIdOrName** (String): ID or name of the auth collection.
- **expand** (String): Auto expand record relations. Ex.:
?expand=relField1,relField2.subRelField
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
expand property (eg. "expand": {"relField1": {...}, ...}).
Only the relations to which the request user has permissions to view will
be expanded.
- **fields** (String): Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,record.expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,record.description:excerpt(200,true)
- **collectionIdOrName** (String): ID or name of the auth collection.
- **Required
email** (String): The email address to send the password reset request (if registered).
- **collectionIdOrName** (String): ID or name of the auth collection.
- **Required
token** (String): The token from the verification request email.
- **Required
email** (String): The email address to send the password reset request (if registered).
- **collectionIdOrName** (String): ID or name of the auth collection.
- **collectionIdOrName** (String): ID or name of the auth collection.
- **Required
token** (String): The token from the password reset request email.
- **Required
password** (String): The new auth record password to set.
- **Required
passwordConfirm** (String): New auth record password confirmation.
- **collectionIdOrName** (String): ID or name of the auth collection.
- **Required
newEmail** (String): The new email address to send the change email request.
- **collectionIdOrName** (String): ID or name of the auth collection.
- **Required
token** (String): The token from the change email request.
- **Required
password** (String): The auth record password to confirm the email address change.
- **collectionIdOrName** (String): ID or name of the auth collection.
- **id** (String): ID of the auth record.
- **fields** (String): Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
- **collectionIdOrName** (String): ID or name of the auth collection.
- **id** (String): ID of the auth record.
- **provider** (String): The name of the auth provider to unlink, eg. google, twitter,
github, etc.
### CRUD actions
List/Search records
Returns a paginated records list, supporting sorting and filtering.
Depending on the collection&#39;s listRule value, the access to this action may or may not
have been restricted.
You could find individual generated records API documentation in the &quot;Admin UI &gt; Collections &gt;
API Preview&quot;.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
// fetch a paginated records list
const resultList = await pb.collection('posts').getList(1, 50, {
filter: 'created >= "2022-01-01 00:00:00" &amp;&amp; someField1 != someField2',
// you can also fetch all records at once via getFullList
const records = await pb.collection('posts').getFullList({
sort: '-created',
// or fetch only the first record that matches the specified filter
const record = await pb.collection('posts').getFirstListItem('someField="test"', {
expand: 'relField1,relField2.subRelField',
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
// fetch a paginated records list
final resultList = await pb.collection('posts').getList(
page: 1,
perPage: 50,
filter: 'created >= "2022-01-01 00:00:00" &amp;&amp; someField1 != someField2',
// you can also fetch all records at once via getFullList
final records = await pb.collection('posts').getFullList(sort: '-created');
// or fetch only the first record that matches the specified filter
final record = await pb.collection('posts').getFirstListItem(
'someField="test"',
expand: 'relField1,relField2.subRelField',
GET
/api/collections/`collectionIdOrName`/records
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the records&#39; collection.
Query parameters
Param
Type
Description
page
Number
The page (aka. offset) of the paginated list (default to 1).
perPage
Number
The max returned records per page (default to 30).
sort
String
Specify the ORDER BY fields.
Add - / + (default) in front of the attribute for DESC /
ASC order, eg.:
// DESC by created and ASC by id
?sort=-created,id
Supported record sort fields:
@random, id, created, updated,
and any other field from the collection schema.
filter
String
Filter expression to filter/search the returned records list (in addition to the
collection&#39;s listRule), eg.:
`?filter=(title~'abc' &amp;&amp; created>'2022-01-01')`
Supported record filter fields:
id, created, updated,
+ any field from the collection schema.
The syntax basically follows the format
OPERAND
OPERATOR
OPERAND, where:
-OPERAND - could be any of the above field literal, string (single or double
quoted), number, null, true, false
-OPERATOR - is one of:
Equal
NOT equal
Greater than
Greater than or equal
-&lt;
Less than
-&lt;=
Less than or equal
Like/Contains (if not specified auto wraps the right string OPERAND in a &quot;%&quot; for wildcard
match)
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a &quot;%&quot; for
wildcard match)
Any/At least one of
Equal
Any/At least one of
NOT equal
Any/At least one of
Greater than
Any/At least one of
Greater than or equal
-?&lt;
Any/At least one of
Less than
-?&lt;=
Any/At least one of
Less than or equal
Any/At least one of
Like/Contains (if not specified auto wraps the right string OPERAND in a &quot;%&quot; for wildcard
match)
-?!~
Any/At least one of
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a &quot;%&quot; for
wildcard match)
To group and combine several expressions you could use parenthesis
(...), &amp;&amp; (AND) and || (OR) tokens.
Single line comments are also supported: // Example comment.
expand
String
Auto expand record relations. Ex.:
`?expand=relField1,relField2.subRelField`
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
`expand` property (eg. `"expand": {"relField1": {...}, ...}`).
Only the relations to which the request user has permissions to view will
be expanded.
fields
String
Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
-:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
skipTotal
Boolean
If it is set the total counts query will be skipped and the response fields
`totalItems` and `totalPages` will have `-1` value.
This could drastically speed up the search queries when the total counters are not needed or cursor based
pagination is used.
For optimization purposes, it is set by default for the
`getFirstListItem()`
and
`getFullList()` SDKs methods.
Responses
"page": 1,
"perPage": 100,
"totalItems": 2,
"totalPages": 1,
"items": [
"id": "ae40239d2bc4477",
"collectionId": "a98f514eb05f454",
"collectionName": "posts",
"updated": "2022-06-25 11:03:50.052",
"created": "2022-06-25 11:03:35.163",
"title": "test1"
"id": "d08dfc4f4d84419",
"collectionId": "a98f514eb05f454",
"collectionName": "posts",
"updated": "2022-06-25 11:03:45.876",
"created": "2022-06-25 11:03:45.876",
"title": "test2"
"code": 400,
"message": "Something went wrong while processing your request. Invalid filter.",
"data": {}
"code": 403,
"message": "Only admins can filter by '@collection.*'",
"data": {}
View record
Returns a single collection record by its ID.
Depending on the collection&#39;s viewRule value, the access to this action may or may not
have been restricted.
You could find individual generated records API documentation in the &quot;Admin UI &gt; Collections &gt;
API Preview&quot;.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const record1 = await pb.collection('posts').getOne('RECORD_ID', {
expand: 'relField1,relField2.subRelField',
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
final record1 = await pb.collection('posts').getOne('RECORD_ID',
expand: 'relField1,relField2.subRelField',
GET
/api/collections/`collectionIdOrName`/records/`recordId`
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the record&#39;s collection.
recordId
String
ID of the record to view.
Query parameters
Param
Type
Description
expand
String
Auto expand record relations. Ex.:
`?expand=relField1,relField2.subRelField`
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
`expand` property (eg. `"expand": {"relField1": {...}, ...}`).
Only the relations to which the request user has permissions to view will
be expanded.
fields
String
Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
-:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
Responses
"id": "ae40239d2bc4477",
"collectionId": "a98f514eb05f454",
"collectionName": "posts",
"updated": "2022-06-25 11:03:50.052",
"created": "2022-06-25 11:03:35.163",
"title": "test1"
"code": 403,
"message": "Only admins can perform this action.",
"data": {}
"code": 404,
"message": "The requested resource wasn't found.",
"data": {}
Create record
Creates a new collection Record.
Depending on the collection&#39;s createRule value, the access to this action may or may not
have been restricted.
You could find individual generated records API documentation from the admin UI.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const record = await pb.collection('demo').create({
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
final record = await pb.collection('demo').create(body: {
POST
/api/collections/`collectionIdOrName`/records
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the record&#39;s collection.
Body Parameters
Param
Type
Description
Optional
String
15 characters string to store as record ID.
If not set, it will be auto generated.
Schema fields
Any field from the collection&#39;s schema.
Auth record fields
Optional
username
String
The username of the auth record.
If not set, it will be auto generated.
Optional
email
String
Auth record email address.
Optional
emailVisibility
Boolean
Whether to show/hide the auth record email when fetching the record data.
Required
password
String
Auth record password.
Required
passwordConfirm
String
Auth record password confirmation.
Optional
verified
Boolean
Indicates whether the auth record is verified or not.
This field can be set only by admins or auth records with &quot;Manage&quot; access.
Body parameters could be sent as JSON or
multipart/form-data.
File upload is supported only through multipart/form-data.
Query parameters
Param
Type
Description
expand
String
Auto expand record relations. Ex.:
`?expand=relField1,relField2.subRelField`
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
`expand` property (eg. `"expand": {"relField1": {...}, ...}`).
Only the relations to which the request user has permissions to view will
be expanded.
fields
String
Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
-:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
Responses
"@collectionId": "a98f514eb05f454",
"@collectionName": "demo",
"id": "ae40239d2bc4477",
"updated": "2022-06-25 11:03:50.052",
"created": "2022-06-25 11:03:35.163",
"code": 400,
"message": "Failed to create record.",
"data": {
"title": {
"code": "validation_required",
"message": "Missing required value."
"code": 403,
"message": "Only admins can perform this action.",
"data": {}
"code": 404,
"message": "The requested resource wasn't found. Missing collection context.",
"data": {}
Update record
Updates an existing collection Record.
Depending on the collection&#39;s updateRule value, the access to this action may or may not
have been restricted.
You could find individual generated records API documentation from the admin UI.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const record = await pb.collection('demo').update('YOUR_RECORD_ID', {
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
final record = await pb.collection('demo').update('YOUR_RECORD_ID', body: {
PATCH
/api/collections/`collectionIdOrName`/records/`recordId`
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the record&#39;s collection.
recordId
String
ID of the record to update.
Body Parameters
Param
Type
Description
Schema fields
Any field from the collection&#39;s schema.
Auth record fields
Optional
username
String
The username of the auth record.
Optional
email
String
The auth record email address.
This field can be updated only by admins or auth records with &quot;Manage&quot; access.
Regular accounts can update their email by calling &quot;Request email change&quot;.
Optional
emailVisibility
Boolean
Whether to show/hide the auth record email when fetching the record data.
Optional
oldPassword
String
Old auth record password.
This field is required only when changing the record password. Admins and auth records with
&quot;Manage&quot; access can skip this field.
Optional
password
String
New auth record password.
Optional
passwordConfirm
String
New auth record password confirmation.
Optional
verified
Boolean
Indicates whether the auth record is verified or not.
This field can be set only by admins or auth records with &quot;Manage&quot; access.
Body parameters could be sent as JSON or
multipart/form-data.
File upload is supported only through multipart/form-data.
Query parameters
Param
Type
Description
expand
String
Auto expand record relations. Ex.:
`?expand=relField1,relField2.subRelField`
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
`expand` property (eg. `"expand": {"relField1": {...}, ...}`).
Only the relations to which the request user has permissions to view will
be expanded.
fields
String
Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
-:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
Responses
"@collectionId": "a98f514eb05f454",
"@collectionName": "demo",
"id": "ae40239d2bc4477",
"updated": "2022-06-25 11:03:50.052",
"created": "2022-06-25 11:03:35.163",
"code": 400,
"message": "Failed to create record.",
"data": {
"title": {
"code": "validation_required",
"message": "Missing required value."
"code": 403,
"message": "Only admins can perform this action.",
"data": {}
"code": 404,
"message": "The requested resource wasn't found. Missing collection context.",
"data": {}
Delete record
Deletes a single collection Record by its ID.
Depending on the collection&#39;s deleteRule value, the access to this action may or may not
have been restricted.
You could find individual generated records API documentation from the admin UI.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
await pb.collection('demo').delete('YOUR_RECORD_ID');
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
await pb.collection('demo').delete('YOUR_RECORD_ID');
DELETE
/api/collections/`collectionIdOrName`/records/`recordId`
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the record&#39;s collection.
recordId
String
ID of the record to delete.
Responses
`null`
"code": 400,
"message": "Failed to delete record. Make sure that the record is not part of a required relation reference.",
"data": {}
"code": 403,
"message": "Only admins can perform this action.",
"data": {}
"code": 404,
"message": "The requested resource wasn't found.",
"data": {}
### Auth record actions
List auth methods
Returns a public list with the allowed collection authentication methods.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const result = await pb.collection('users').listAuthMethods();
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
final result = await pb.collection('users').listAuthMethods();
GET
/api/collections/`collectionIdOrName`/auth-methods
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the auth collection.
Query parameters
Param
Type
Description
fields
String
Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
-:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
Responses
"usernamePassword": false,
"emailPassword": true,
"authProviders": [
"name": "github",
"state": "3Yd8jNkK_6PJG6hPWwBjLqKwse6Ejd",
"codeVerifier": "KxFDWz1B3fxscCDJ_9gHQhLuh__ie7",
"codeChallenge": "NM1oVexB6Q6QH8uPtOUfK7tq4pmu4Jz6lNDIwoxHZNE=",
"codeChallengeMethod": "S256",
"authUrl": "https://github.com/login/oauth/authorize?client_id=demo&amp;code_challenge=NM1oVexB6Q6QH8uPtOUfK7tq4pmu4Jz6lNDIwoxHZNE%3D&amp;code_challenge_method=S256&amp;response_type=code&amp;scope=user&amp;state=3Yd8jNkK_6PJG6hPWwBjLqKwse6Ejd&amp;redirect_uri="
"name": "gitlab",
"state": "NeQSbtO5cShr_mk5__3CUukiMnymeb",
"codeVerifier": "ahTFHOgua8mkvPAlIBGwCUJbWKR_xi",
"codeChallenge": "O-GATkTj4eXDCnfonsqGLCd6njvTixlpCMvy5kjgOOg=",
"codeChallengeMethod": "S256",
"authUrl": "https://gitlab.com/oauth/authorize?client_id=demo&amp;code_challenge=O-GATkTj4eXDCnfonsqGLCd6njvTixlpCMvy5kjgOOg%3D&amp;code_challenge_method=S256&amp;response_type=code&amp;scope=read_user&amp;state=NeQSbtO5cShr_mk5__3CUukiMnymeb&amp;redirect_uri="
"name": "google",
"state": "zB3ZPifV1TW2GMuvuFkamSXfSNkHPQ",
"codeVerifier": "t3CmO5VObGzdXqieakvR_fpjiW0zdO",
"codeChallenge": "KChwoQPKYlz2anAdqtgsSTdIo8hdwtc1fh2wHMwW2Yk=",
"codeChallengeMethod": "S256",
"authUrl": "https://accounts.google.com/o/oauth2/auth?client_id=demo&amp;code_challenge=KChwoQPKYlz2anAdqtgsSTdIo8hdwtc1fh2wHMwW2Yk%3D&amp;code_challenge_method=S256&amp;response_type=code&amp;scope=https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.profile+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.email&amp;state=zB3ZPifV1TW2GMuvuFkamSXfSNkHPQ&amp;redirect_uri="
Auth with password
Authenticate a single auth record by their username/email and password.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const authData = await pb.collection('users').authWithPassword(
'YOUR_USERNAME_OR_EMAIL',
'YOUR_PASSWORD',
// after the above you can also access the auth data from the authStore
console.log(pb.authStore.isValid);
console.log(pb.authStore.token);
console.log(pb.authStore.model.id);
// "logout" the last authenticated account
pb.authStore.clear();
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
final authData = await pb.collection('users').authWithPassword(
'YOUR_USERNAME_OR_EMAIL',
'YOUR_PASSWORD',
// after the above you can also access the auth data from the authStore
print(pb.authStore.isValid);
print(pb.authStore.token);
print(pb.authStore.model.id);
// "logout" the last authenticated account
pb.authStore.clear();
POST
/api/collections/`collectionIdOrName`/auth-with-password
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the auth collection.
Body Parameters
Param
Type
Description
Required
identity
String
Auth record username or email address.
Required
password
String
Auth record password.
Body parameters could be sent as JSON or
multipart/form-data.
Query parameters
Param
Type
Description
expand
String
Auto expand record relations. Ex.:
`?expand=relField1,relField2.subRelField`
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
`expand` property (eg. `"expand": {"relField1": {...}, ...}`).
Only the relations to which the request user has permissions to view will
be expanded.
fields
String
Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,record.expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
-:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,record.description:excerpt(200,true)
Responses
"token": "eyJhbGciOiJIUzI1NiJ9.eyJpZCI6IjRxMXhsY2xtZmxva3UzMyIsInR5cGUiOiJhdXRoUmVjb3JkIiwiY29sbGVjdGlvbklkIjoiX3BiX3VzZXJzX2F1dGhfIiwiZXhwIjoyMjA4OTg1MjYxfQ.UwD8JvkbQtXpymT09d7J6fdA0aP9g4FJ1GPh_ggEkzc",
"record": {
"id": "8171022dc95a4ed",
"collectionId": "d2972397d45614e",
"collectionName": "users",
"created": "2022-06-24 06:24:18.434Z",
"updated": "2022-06-24 06:24:18.889Z",
"verified": false,
"emailVisibility": true,
"someCustomField": "example 123"
"code": 400,
"message": "An error occurred while submitting the form.",
"data": {
"password": {
"code": "validation_required",
"message": "Missing required value."
Auth with OAuth2
Authenticate with an OAuth2 provider and returns a new auth token and record data.
This action usually should be called right after the provider login page redirect.
You could also check the
OAuth2 web integration example.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const authData = await pb.collection('users').authWithOAuth2Code(
'google',
'CODE',
'VERIFIER',
'REDIRECT_URL',
// optional data that will be used for the new account on OAuth2 sign-up
'name': 'test',
// after the above you can also access the auth data from the authStore
console.log(pb.authStore.isValid);
console.log(pb.authStore.token);
console.log(pb.authStore.model.id);
// "logout" the last authenticated account
pb.authStore.clear();
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
final authData = await pb.collection('users').authWithOAuth2Code(
'google',
'CODE',
'VERIFIER',
'REDIRECT_URL',
// optional data that will be used for the new account on OAuth2 sign-up
createData: {
'name': 'test',
// after the above you can also access the auth data from the authStore
print(pb.authStore.isValid);
print(pb.authStore.token);
print(pb.authStore.model.id);
// "logout" the last authenticated account
pb.authStore.clear();
POST
/api/collections/`collectionIdOrName`/auth-with-oauth2
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the auth collection.
Body Parameters
Param
Type
Description
Required
provider
String
The name of the OAuth2 client provider (eg. &quot;google&quot;).
Required
code
String
The authorization code returned from the initial request.
Required
codeVerifier
String
The code verifier sent with the initial request as part of the code_challenge.
Required
redirectUrl
String
The redirect url sent with the initial request.
Optional
createData
Object
Optional data that will be used when creating the auth record on OAuth2 sign-up.
The created auth record must comply with the same requirements and validations in the
regular create action.
The data can only be in json, aka. multipart/form-data and
files upload currently are not supported during OAuth2 sign-ups.
Body parameters could be sent as JSON or
multipart/form-data.
Query parameters
Param
Type
Description
expand
String
Auto expand record relations. Ex.:
`?expand=relField1,relField2.subRelField`
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
`expand` property (eg. `"expand": {"relField1": {...}, ...}`).
Only the relations to which the request user has permissions to view will
be expanded.
fields
String
Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,record.expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
-:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,record.description:excerpt(200,true)
Responses
"token": "eyJhbGciOiJIUzI1NiJ9.eyJpZCI6IjRxMXhsY2xtZmxva3UzMyIsInR5cGUiOiJhdXRoUmVjb3JkIiwiY29sbGVjdGlvbklkIjoiX3BiX3VzZXJzX2F1dGhfIiwiZXhwIjoyMjA4OTg1MjYxfQ.UwD8JvkbQtXpymT09d7J6fdA0aP9g4FJ1GPh_ggEkzc",
"record": {
"id": "8171022dc95a4ed",
"collectionId": "d2972397d45614e",
"collectionName": "users",
"created": "2022-06-24 06:24:18.434Z",
"updated": "2022-06-24 06:24:18.889Z",
"verified": true,
"emailVisibility": false,
"someCustomField": "example 123"
"meta": {
"id": "abc123",
"name": "John Doe",
"username": "john.doe",
"isNew": false,
"accessToken": "...",
"refreshToken": "...",
"expiry": "..."
"code": 400,
"message": "An error occurred while submitting the form.",
"data": {
"provider": {
"code": "validation_required",
"message": "Missing required value."
Auth refresh
Returns a new auth response (token and user data) for already authenticated auth record.
stored data in pb.authStore is still valid and up-to-date.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const authData = await pb.collection('users').authRefresh();
// after the above you can also access the refreshed auth data from the authStore
console.log(pb.authStore.isValid);
console.log(pb.authStore.token);
console.log(pb.authStore.model.id);
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
final authData = await pb.collection('users').authRefresh();
// after the above you can also access the refreshed auth data from the authStore
print(pb.authStore.isValid);
print(pb.authStore.token);
print(pb.authStore.model.id);
POST
/api/collections/`collectionIdOrName`/auth-refresh
Requires `Authorization: TOKEN`
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the auth collection.
Query parameters
Param
Type
Description
expand
String
Auto expand record relations. Ex.:
`?expand=relField1,relField2.subRelField`
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
`expand` property (eg. `"expand": {"relField1": {...}, ...}`).
Only the relations to which the request user has permissions to view will
be expanded.
fields
String
Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,record.expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
-:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,record.description:excerpt(200,true)
Responses
"token": "eyJhbGciOiJIUzI1NiJ9.eyJpZCI6IjRxMXhsY2xtZmxva3UzMyIsInR5cGUiOiJhdXRoUmVjb3JkIiwiY29sbGVjdGlvbklkIjoiX3BiX3VzZXJzX2F1dGhfIiwiZXhwIjoyMjA4OTg1MjYxfQ.UwD8JvkbQtXpymT09d7J6fdA0aP9g4FJ1GPh_ggEkzc",
"record": {
"id": "8171022dc95a4ed",
"collectionId": "d2972397d45614e",
"collectionName": "users",
"created": "2022-06-24 06:24:18.434Z",
"updated": "2022-06-24 06:24:18.889Z",
"verified": false,
"emailVisibility": true,
"someCustomField": "example 123"
"code": 401,
"message": "The request requires valid record authorization token to be set.",
"data": {}
"code": 403,
"message": "The authorized record model is not allowed to perform this action.",
"data": {}
"code": 404,
"message": "Missing auth record context.",
"data": {}
Request verification
Sends auth record verification email request.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
POST
/api/collections/`collectionIdOrName`/request-verification
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the auth collection.
Body Parameters
Param
Type
Description
Required
email
String
The email address to send the password reset request (if registered).
Body parameters could be sent as JSON or
multipart/form-data.
Responses
`null`
"code": 400,
"message": "An error occurred while submitting the form.",
"data": {
"email": {
"code": "validation_required",
"message": "Missing required value."
Confirm verification
Confirms an email address verification request.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
await pb.collection('users').confirmVerification('TOKEN');
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
await pb.collection('users').confirmVerification('TOKEN');
POST
/api/collections/`collectionIdOrName`/confirm-verification
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the auth collection.
Body Parameters
Param
Type
Description
Required
token
String
The token from the verification request email.
Body parameters could be sent as JSON or
multipart/form-data.
Responses
`null`
"code": 400,
"message": "An error occurred while submitting the form.",
"data": {
"token": {
"code": "validation_required",
"message": "Missing required value."
Request password reset
Sends a password reset email to a specified auth record email.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
POST
/api/collections/`collectionIdOrName`/request-password-reset
Body Parameters
Param
Type
Description
Required
email
String
The email address to send the password reset request (if registered).
Body parameters could be sent as JSON or
multipart/form-data.
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the auth collection.
Responses
`null`
"code": 400,
"message": "An error occurred while submitting the form.",
"data": {
"email": {
"code": "validation_required",
"message": "Missing required value."
Confirm password reset
Confirms a password reset request and sets a new auth record password.
invalidated.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
await pb.collection('users').confirmPasswordReset('TOKEN', '1234567890', '1234567890');
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
await pb.collection('users').confirmPasswordReset('TOKEN', '1234567890', '1234567890');
POST
/api/collections/`collectionIdOrName`/confirm-password-reset
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the auth collection.
Body Parameters
Param
Type
Description
Required
token
String
The token from the password reset request email.
Required
password
String
The new auth record password to set.
Required
passwordConfirm
String
New auth record password confirmation.
Body parameters could be sent as JSON or
multipart/form-data.
Responses
`null`
"code": 400,
"message": "An error occurred while submitting the form.",
"data": {
"password": {
"code": "validation_required",
"message": "Missing required value."
Request email change
Sends an email change request for an authenticated record.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
POST
/api/collections/`collectionIdOrName`/request-email-change
Requires record `Authorization:TOKEN` header
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the auth collection.
Body Parameters
Param
Type
Description
Required
newEmail
String
The new email address to send the change email request.
Body parameters could be sent as JSON or
multipart/form-data.
Responses
`null`
"code": 400,
"message": "An error occurred while submitting the form.",
"data": {
"newEmail": {
"code": "validation_required",
"message": "Missing required value."
"code": 401,
"message": "The request requires valid record authorization token to be set.",
"data": {}
"code": 403,
"message": "You are not allowed to perform this request.",
"data": {}
Confirm email change
Confirms email address change.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
await pb.collection('users').confirmEmailChange('TOKEN', 'YOUR_PASSWORD');
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
await pb.collection('users').confirmEmailChange('TOKEN', 'YOUR_PASSWORD');
POST
/api/collections/`collectionIdOrName`/confirm-email-change
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the auth collection.
Body Parameters
Param
Type
Description
Required
token
String
The token from the change email request.
Required
password
String
The auth record password to confirm the email address change.
Body parameters could be sent as JSON or
multipart/form-data.
Responses
`null`
"code": 400,
"message": "An error occurred while submitting the form.",
"data": {
"token": {
"code": "validation_required",
"message": "Missing required value."
List linked external auth providers
Return a list with all external auth providers linked to a single record.
Only admins and the account owner can access this action.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const result = await pb.collection('users').listExternalAuths(pb.authStore.model.id);
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
final result = await pb.collection('users').listExternalAuths(pb.authStore.model.id);
GET
/api/collections/`collectionIdOrName`/records/`id`/external-auths
Requires `Authorization: TOKEN`
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the auth collection.
String
ID of the auth record.
Query parameters
Param
Type
Description
fields
String
Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
-:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
Responses
"id": "8171022dc95a4e8",
"created": "2022-09-01 10:24:18.434",
"updated": "2022-09-01 10:24:18.889",
"recordId": "e22581b6f1d44ea",
"collectionId": "systemprofiles0",
"provider": "google",
"providerId": "2da15468800514p"
"id": "171022dc895a4e8",
"created": "2022-09-01 10:24:18.434",
"updated": "2022-09-01 10:24:18.889",
"recordId": "e22581b6f1d44ea",
"collectionId": "systemprofiles0",
"provider": "twitter",
"providerId": "720688005140514"
"code": 401,
"message": "The request requires admin or record authorization token to be set.",
"data": {}
"code": 403,
"message": "You are not allowed to perform this request.",
"data": {}
"code": 404,
"message": "The requested resource wasn't found.",
"data": {}
Unlink external auth provider
Unlink a single external OAuth2 provider from an auth record.
Only admins and the account owner can access this action.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
await pb.collection('users').unlinkExternalAuth(pb.authStore.model.id, 'google');
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
await pb.collection('users').unlinkExternalAuth(pb.authStore.model.id, 'google');
DELETE
/api/collections/`collectionIdOrName`/records/`id`/external-auths/`provider`
*`Authorization: TOKEN`
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the auth collection.
String
ID of the auth record.
provider
String
The name of the auth provider to unlink, eg. `google`, `twitter`,
`github`, etc.
Responses
`null`
"code": 401,
"message": "The request requires valid record authorization token to be set.",
"data": {}
"code": 403,
"message": "You are not allowed to perform this request.",
"data": {}
"code": 404,
"message": "Missing external auth provider relation.",
"data": {}

## 10.Going to production
# Going to production
### Deployment strategies
##### Minimal setup
One of the best PocketBase features is that it&#39;s completely portable. This mean that it doesn&#39;t require
any external dependency and
could be deployed by just uploading the executable on your server.
Here is an example for starting a production HTTPS server (auto managed TLS with Let&#39;s Encrypt) on clean
Ubuntu 22.04 installation.
-Consider the following app directory structure:
myapp/
pb_migrations/
pb_hooks/
pocketbase
-Upload the binary and anything else related to your remote server, for example using
rsync:
rsync -avz -e ssh /local/path/to/myapp root@YOUR_SERVER_IP:/root/pb
-Start a SSH session with your server:
ssh root@YOUR_SERVER_IP
-Start the executable (specifying a domain name will issue a Let&#39;s encrypt certificate for it)
[root@dev ~]$ /root/pb/pocketbase serve yourdomain.com
Notice that in the above example we are logged in as root which allow us to
bind to the
privileged 80 and 443 ports.
For non-root users usually you&#39;ll need special privileges to be able to do
that. You have several options depending on your OS - authbind,
setcap,
iptables, sysctl, etc. Here is an example using setcap:
[myuser@dev ~]$ sudo setcap 'cap_net_bind_service=+ep' /root/pb/pocketbase
-(Optional) systemd service
You can skip step 3 and create a Systemd service
to allow your application to start/restart on its own.
Here is an example service file (usually created in
/lib/systemd/system/pocketbase.service):
[Unit]
Description = pocketbase
[Service]
Type           = simple
User           = root
Group          = root
LimitNOFILE    = 4096
Restart        = always
RestartSec     = 5s
StandardOutput = append:/root/pb/errors.log
StandardError  = append:/root/pb/errors.log
ExecStart      = /root/pb/pocketbase serve yourdomain.com
[Install]
WantedBy = multi-user.target
After that we just have to enable it and start the service using systemctl:
[root@dev ~]$ systemctl enable pocketbase.service
[root@dev ~]$ systemctl start pocketbase
##### Using reverse proxy
If you plan hosting multiple applications on a single server or need finer network controls (rate limiter,
IPs whitelisting, etc.), you could always put PocketBase behind a reverse proxy such as
NGINX, Apache, Caddy, etc.
Here is a minimal NGINX example configuration:
server {
listen 80;
client_max_body_size 10M;
location / {
# check http://nginx.org/en/docs/http/ngx_http_upstream_module.html#keepalive
proxy_set_header Connection '';
proxy_http_version 1.1;
proxy_read_timeout 360s;
proxy_set_header Host $host;
proxy_set_header X-Real-IP $remote_addr;
proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
proxy_set_header X-Forwarded-Proto $scheme;
# enable if you are serving under a subpath location
# rewrite /yourSubpath/(.*) /$1  break;
proxy_pass http://127.0.0.1:8090;
Corresponding Caddy configuration is:
request_body {
max_size 10MB
reverse_proxy 127.0.0.1:8090 {
transport http {
read_timeout 360s
##### Using Docker
Some hosts (eg. fly.io) use Docker
for deployments. PocketBase doesn&#39;t have an official Docker image yet, but you could use the below
Dockerfile as an example:
FROM alpine:latest
ARG PB_VERSION=0.22.38
RUN apk add --no-cache \
unzip \
ca-certificates
# uncomment to copy the local pb_migrations dir into the image
# COPY ./pb_migrations /pb/pb_migrations
# uncomment to copy the local pb_hooks dir into the image
# COPY ./pb_hooks /pb/pb_hooks
EXPOSE 8080
# start PocketBase
CMD ["/old/pb/pocketbase", "serve", "--http=0.0.0.0:8080"]
To persist your data you need to mount a volume at /pb/pb_data.
For a full example you could check the
&quot;Host for free on Fly.io&quot;
guide.
### Backup and Restore
PocketBase v0.16+ comes with built-in backups and restore APIs that could be accessed from the Admin UI (Settings
&gt; Backups):
Note that the application will be temporary set in read-only mode during the backup&#39;s ZIP generation.
Backups can be stored locally (default) or in an external S3 storage.
Alternatively, you can always manually copy your pb_data directory
(for transactional safety make sure that the application is not running).
### Recommendations
By default, PocketBase uses the internal Unix sendmail command for sending emails.
While it&#39;s OK for development, it&#39;s not very useful for production, because your emails most likely will get
marked as spam or even fail to deliver.
To avoid deliverability issues, consider using a local SMTP server or an external mail service like
MailerSend,
Brevo,
SendGrid,
Mailgun,
AWS SES, etc.
Once you&#39;ve decided on a mail service, you could configure the PocketBase SMTP settings from the Admin UI
(Settings &gt; Mail settings):
Unix uses &quot;file descriptors&quot; also for network connections and most systems have a default limit
of ~ 1024.
If your application has a lot of concurrent realtime connections, it is possible that at some point you would
get an error such as: Too many open files.
One way to mitigate this is to check your current account resource limits by running
ulimit -a and find the parameter you want to change. For example, if you want to increase the
open files limit (-n), you could run
ulimit -n 4096 before starting PocketBase.
It is fine to ignore the below if you are not sure whether you need it.
By default, PocketBase stores the applications settings in the database as plain JSON text, including the
secret keys for the OAuth2 clients and the SMTP password.
While this is not a security issue on its own (PocketBase applications live entirely on a single server
and its expected only authorized users to have access to your server and application data), in some
situations it may be a good idea to store the settings encrypted in case someone get their hands on your
database file (eg. from an external stored backup).
To store your PocketBase settings encrypted:
-Create a new environment variable and set a random 32 characters string as its value.
eg. add
export PB_ENCRYPTION_KEY=&quot;LSidmCcAa6gwn777BwRZJVLdDkSWfvVL&quot;
in your shell profile file
-Start the application with --encryptionEnv=YOUR_ENV_VAR flag.
eg. pocketbase serve --encryptionEnv=PB_ENCRYPTION_KEY

## 11.Web APIs reference - API Records
collectionIdOrName(String):ID or name of the records' collection.
page(Number):The page (aka. offset) of the paginated list (default to 1).
perPage(Number):The max returned records per page (default to 30).
sort(String):Specify the ORDER BY fields.
Add - / + (default) in front of the attribute for DESC /
ASC order, eg.:
// DESC by created and ASC by id
?sort=-created,id
Supported record sort fields:
@random, id, created, updated,
and any other field from the collection schema.
filter(String):Filter expression to filter/search the returned records list (in addition to the
collection's listRule), eg.:
?filter=(title~'abc' && created>'2022-01-01')
Supported record filter fields:
id, created, updated,
+ any field from the collection schema.
The syntax basically follows the format
OPERAND
OPERATOR
OPERAND, where:
OPERAND - could be any of the above field literal, string (single or double
quoted), number, null, true, false
OPERATOR - is one of:
Equal
!=
NOT equal
Greater than
>=
Greater than or equal
Less than
<=
Less than or equal
Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for wildcard
match)
!~
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for
wildcard match)
?=
Any/At least one of
Equal
?!=
Any/At least one of
NOT equal
?>
Any/At least one of
Greater than
?>=
Any/At least one of
Greater than or equal
?<
Any/At least one of
Less than
?<=
Any/At least one of
Less than or equal
?~
Any/At least one of
Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for wildcard
match)
?!~
Any/At least one of
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for
wildcard match)
To group and combine several expressions you could use parenthesis
(...), && (AND) and || (OR) tokens.
Single line comments are also supported: // Example comment.
expand(String):Auto expand record relations. Ex.:
?expand=relField1,relField2.subRelField
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
expand property (eg. "expand": {"relField1": {...}, ...}).
Only the relations to which the request user has permissions to view will
be expanded.
fields(String):Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
skipTotal(Boolean):If it is set the total counts query will be skipped and the response fields
totalItems and totalPages will have -1 value.
This could drastically speed up the search queries when the total counters are not needed or cursor based
pagination is used.
For optimization purposes, it is set by default for the
getFirstListItem()
and
getFullList() SDKs methods.
collectionIdOrName(String):ID or name of the record's collection.
recordId(String):ID of the record to view.
expand(String):Auto expand record relations. Ex.:
?expand=relField1,relField2.subRelField
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
expand property (eg. "expand": {"relField1": {...}, ...}).
Only the relations to which the request user has permissions to view will
be expanded.
fields(String):Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
collectionIdOrName(String):ID or name of the record's collection.
Optional
id(String):15 characters string to store as record ID.
If not set, it will be auto generated.
Optional
username(String):The username of the auth record.
If not set, it will be auto generated.
Optional
email(String):Auth record email address.
Optional
emailVisibility(Boolean):Whether to show/hide the auth record email when fetching the record data.
Required
password(String):Auth record password.
Required
passwordConfirm(String):Auth record password confirmation.
Optional
verified(Boolean):Indicates whether the auth record is verified or not.
This field can be set only by admins or auth records with "Manage" access.
expand(String):Auto expand record relations. Ex.:
?expand=relField1,relField2.subRelField
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
expand property (eg. "expand": {"relField1": {...}, ...}).
Only the relations to which the request user has permissions to view will
be expanded.
fields(String):Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
collectionIdOrName(String):ID or name of the record's collection.
recordId(String):ID of the record to update.
Optional
username(String):The username of the auth record.
Optional
email(String):The auth record email address.
This field can be updated only by admins or auth records with "Manage" access.
Regular accounts can update their email by calling "Request email change".
Optional
emailVisibility(Boolean):Whether to show/hide the auth record email when fetching the record data.
Optional
oldPassword*(String):Old auth record password.
This field is required only when changing the record password. Admins and auth records with
"Manage" access can skip this field.
Optional
password(String):New auth record password.
Optional
passwordConfirm(String):New auth record password confirmation.
Optional
verified(Boolean):Indicates whether the auth record is verified or not.
This field can be set only by admins or auth records with "Manage" access.
expand(String):Auto expand record relations. Ex.:
?expand=relField1,relField2.subRelField
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
expand property (eg. "expand": {"relField1": {...}, ...}).
Only the relations to which the request user has permissions to view will
be expanded.
fields(String):Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
collectionIdOrName(String):ID or name of the record's collection.
recordId(String):ID of the record to delete.
collectionIdOrName(String):ID or name of the auth collection.
fields(String):Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
collectionIdOrName(String):ID or name of the auth collection.
Required
identity(String):Auth record username or email address.
Required
password(String):Auth record password.
expand(String):Auto expand record relations. Ex.:
?expand=relField1,relField2.subRelField
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
expand property (eg. "expand": {"relField1": {...}, ...}).
Only the relations to which the request user has permissions to view will
be expanded.
fields(String):Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,record.expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,record.description:excerpt(200,true)
collectionIdOrName(String):ID or name of the auth collection.
Required
provider(String):The name of the OAuth2 client provider (eg. "google").
Required
code(String):The authorization code returned from the initial request.
Required
codeVerifier(String):The code verifier sent with the initial request as part of the code_challenge.
Required
redirectUrl(String):The redirect url sent with the initial request.
Optional
createData(Object):Optional data that will be used when creating the auth record on OAuth2 sign-up.
The created auth record must comply with the same requirements and validations in the
regular create action.
The data can only be in json, aka. multipart/form-data and
files upload currently are not supported during OAuth2 sign-ups.
expand(String):Auto expand record relations. Ex.:
?expand=relField1,relField2.subRelField
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
expand property (eg. "expand": {"relField1": {...}, ...}).
Only the relations to which the request user has permissions to view will
be expanded.
fields(String):Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,record.expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,record.description:excerpt(200,true)
collectionIdOrName(String):ID or name of the auth collection.
expand(String):Auto expand record relations. Ex.:
?expand=relField1,relField2.subRelField
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
expand property (eg. "expand": {"relField1": {...}, ...}).
Only the relations to which the request user has permissions to view will
be expanded.
fields(String):Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,record.expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,record.description:excerpt(200,true)
collectionIdOrName(String):ID or name of the auth collection.
Required
email(String):The email address to send the password reset request (if registered).
collectionIdOrName(String):ID or name of the auth collection.
Required
token(String):The token from the verification request email.
Required
email(String):The email address to send the password reset request (if registered).
collectionIdOrName(String):ID or name of the auth collection.
collectionIdOrName(String):ID or name of the auth collection.
Required
token(String):The token from the password reset request email.
Required
password(String):The new auth record password to set.
Required
passwordConfirm(String):New auth record password confirmation.
collectionIdOrName(String):ID or name of the auth collection.
Required
newEmail(String):The new email address to send the change email request.
collectionIdOrName(String):ID or name of the auth collection.
Required
token(String):The token from the change email request.
Required
password(String):The auth record password to confirm the email address change.
collectionIdOrName(String):ID or name of the auth collection.
id(String):ID of the auth record.
fields(String):Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
collectionIdOrName(String):ID or name of the auth collection.
id(String):ID of the auth record.
provider(String):The name of the auth provider to unlink, eg. google, twitter,
github, etc.
# Web APIs reference - API Records
- **collectionIdOrName** (String): ID or name of the records' collection.
- **page** (Number): The page (aka. offset) of the paginated list (default to 1).
- **perPage** (Number): The max returned records per page (default to 30).
- **sort** (String): Specify the ORDER BY fields.
Add - / + (default) in front of the attribute for DESC /
ASC order, eg.:
// DESC by created and ASC by id
?sort=-created,id
Supported record sort fields:
@random, id, created, updated,
and any other field from the collection schema.
- **filter** (String): Filter expression to filter/search the returned records list (in addition to the
collection's listRule), eg.:
?filter=(title~'abc' && created>'2022-01-01')
Supported record filter fields:
id, created, updated,
+ any field from the collection schema.
The syntax basically follows the format
OPERAND
OPERATOR
OPERAND, where:
OPERAND - could be any of the above field literal, string (single or double
quoted), number, null, true, false
OPERATOR - is one of:
Equal
NOT equal
Greater than
Greater than or equal
Less than
Less than or equal
Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for wildcard
match)
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for
wildcard match)
Any/At least one of
Equal
Any/At least one of
NOT equal
Any/At least one of
Greater than
Any/At least one of
Greater than or equal
Any/At least one of
Less than
Any/At least one of
Less than or equal
Any/At least one of
Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for wildcard
match)
?!~
Any/At least one of
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for
wildcard match)
To group and combine several expressions you could use parenthesis
(...), && (AND) and || (OR) tokens.
Single line comments are also supported: // Example comment.
- **expand** (String): Auto expand record relations. Ex.:
?expand=relField1,relField2.subRelField
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
expand property (eg. "expand": {"relField1": {...}, ...}).
Only the relations to which the request user has permissions to view will
be expanded.
- **fields** (String): Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
- **skipTotal** (Boolean): If it is set the total counts query will be skipped and the response fields
totalItems and totalPages will have -1 value.
This could drastically speed up the search queries when the total counters are not needed or cursor based
pagination is used.
For optimization purposes, it is set by default for the
getFirstListItem()
and
getFullList() SDKs methods.
- **collectionIdOrName** (String): ID or name of the record's collection.
- **recordId** (String): ID of the record to view.
- **expand** (String): Auto expand record relations. Ex.:
?expand=relField1,relField2.subRelField
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
expand property (eg. "expand": {"relField1": {...}, ...}).
Only the relations to which the request user has permissions to view will
be expanded.
- **fields** (String): Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
- **collectionIdOrName** (String): ID or name of the record's collection.
- **Optional
id** (String): 15 characters string to store as record ID.
If not set, it will be auto generated.
- **Optional
username** (String): The username of the auth record.
If not set, it will be auto generated.
- **Optional
email** (String): Auth record email address.
- **Optional
emailVisibility** (Boolean): Whether to show/hide the auth record email when fetching the record data.
- **Required
password** (String): Auth record password.
- **Required
passwordConfirm** (String): Auth record password confirmation.
- **Optional
verified** (Boolean): Indicates whether the auth record is verified or not.
This field can be set only by admins or auth records with "Manage" access.
- **expand** (String): Auto expand record relations. Ex.:
?expand=relField1,relField2.subRelField
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
expand property (eg. "expand": {"relField1": {...}, ...}).
Only the relations to which the request user has permissions to view will
be expanded.
- **fields** (String): Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
- **collectionIdOrName** (String): ID or name of the record's collection.
- **recordId** (String): ID of the record to update.
- **Optional
username** (String): The username of the auth record.
- **Optional
email** (String): The auth record email address.
This field can be updated only by admins or auth records with "Manage" access.
Regular accounts can update their email by calling "Request email change".
- **Optional
emailVisibility** (Boolean): Whether to show/hide the auth record email when fetching the record data.
- **Optional
oldPassword** (String) (required): Old auth record password.
This field is required only when changing the record password. Admins and auth records with
"Manage" access can skip this field.
- **Optional
password** (String): New auth record password.
- **Optional
passwordConfirm** (String): New auth record password confirmation.
- **Optional
verified** (Boolean): Indicates whether the auth record is verified or not.
This field can be set only by admins or auth records with "Manage" access.
- **expand** (String): Auto expand record relations. Ex.:
?expand=relField1,relField2.subRelField
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
expand property (eg. "expand": {"relField1": {...}, ...}).
Only the relations to which the request user has permissions to view will
be expanded.
- **fields** (String): Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
- **collectionIdOrName** (String): ID or name of the record's collection.
- **recordId** (String): ID of the record to delete.
- **collectionIdOrName** (String): ID or name of the auth collection.
- **fields** (String): Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
- **collectionIdOrName** (String): ID or name of the auth collection.
- **Required
identity** (String): Auth record username or email address.
- **Required
password** (String): Auth record password.
- **expand** (String): Auto expand record relations. Ex.:
?expand=relField1,relField2.subRelField
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
expand property (eg. "expand": {"relField1": {...}, ...}).
Only the relations to which the request user has permissions to view will
be expanded.
- **fields** (String): Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,record.expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,record.description:excerpt(200,true)
- **collectionIdOrName** (String): ID or name of the auth collection.
- **Required
provider** (String): The name of the OAuth2 client provider (eg. "google").
- **Required
code** (String): The authorization code returned from the initial request.
- **Required
codeVerifier** (String): The code verifier sent with the initial request as part of the code_challenge.
- **Required
redirectUrl** (String): The redirect url sent with the initial request.
- **Optional
createData** (Object): Optional data that will be used when creating the auth record on OAuth2 sign-up.
The created auth record must comply with the same requirements and validations in the
regular create action.
The data can only be in json, aka. multipart/form-data and
files upload currently are not supported during OAuth2 sign-ups.
- **expand** (String): Auto expand record relations. Ex.:
?expand=relField1,relField2.subRelField
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
expand property (eg. "expand": {"relField1": {...}, ...}).
Only the relations to which the request user has permissions to view will
be expanded.
- **fields** (String): Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,record.expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,record.description:excerpt(200,true)
- **collectionIdOrName** (String): ID or name of the auth collection.
- **expand** (String): Auto expand record relations. Ex.:
?expand=relField1,relField2.subRelField
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
expand property (eg. "expand": {"relField1": {...}, ...}).
Only the relations to which the request user has permissions to view will
be expanded.
- **fields** (String): Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,record.expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,record.description:excerpt(200,true)
- **collectionIdOrName** (String): ID or name of the auth collection.
- **Required
email** (String): The email address to send the password reset request (if registered).
- **collectionIdOrName** (String): ID or name of the auth collection.
- **Required
token** (String): The token from the verification request email.
- **Required
email** (String): The email address to send the password reset request (if registered).
- **collectionIdOrName** (String): ID or name of the auth collection.
- **collectionIdOrName** (String): ID or name of the auth collection.
- **Required
token** (String): The token from the password reset request email.
- **Required
password** (String): The new auth record password to set.
- **Required
passwordConfirm** (String): New auth record password confirmation.
- **collectionIdOrName** (String): ID or name of the auth collection.
- **Required
newEmail** (String): The new email address to send the change email request.
- **collectionIdOrName** (String): ID or name of the auth collection.
- **Required
token** (String): The token from the change email request.
- **Required
password** (String): The auth record password to confirm the email address change.
- **collectionIdOrName** (String): ID or name of the auth collection.
- **id** (String): ID of the auth record.
- **fields** (String): Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
- **collectionIdOrName** (String): ID or name of the auth collection.
- **id** (String): ID of the auth record.
- **provider** (String): The name of the auth provider to unlink, eg. google, twitter,
github, etc.
### CRUD actions
List/Search records
Returns a paginated records list, supporting sorting and filtering.
Depending on the collection&#39;s listRule value, the access to this action may or may not
have been restricted.
You could find individual generated records API documentation in the &quot;Admin UI &gt; Collections &gt;
API Preview&quot;.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
// fetch a paginated records list
const resultList = await pb.collection('posts').getList(1, 50, {
filter: 'created >= "2022-01-01 00:00:00" &amp;&amp; someField1 != someField2',
// you can also fetch all records at once via getFullList
const records = await pb.collection('posts').getFullList({
sort: '-created',
// or fetch only the first record that matches the specified filter
const record = await pb.collection('posts').getFirstListItem('someField="test"', {
expand: 'relField1,relField2.subRelField',
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
// fetch a paginated records list
final resultList = await pb.collection('posts').getList(
page: 1,
perPage: 50,
filter: 'created >= "2022-01-01 00:00:00" &amp;&amp; someField1 != someField2',
// you can also fetch all records at once via getFullList
final records = await pb.collection('posts').getFullList(sort: '-created');
// or fetch only the first record that matches the specified filter
final record = await pb.collection('posts').getFirstListItem(
'someField="test"',
expand: 'relField1,relField2.subRelField',
GET
/api/collections/`collectionIdOrName`/records
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the records&#39; collection.
Query parameters
Param
Type
Description
page
Number
The page (aka. offset) of the paginated list (default to 1).
perPage
Number
The max returned records per page (default to 30).
sort
String
Specify the ORDER BY fields.
Add - / + (default) in front of the attribute for DESC /
ASC order, eg.:
// DESC by created and ASC by id
?sort=-created,id
Supported record sort fields:
@random, id, created, updated,
and any other field from the collection schema.
filter
String
Filter expression to filter/search the returned records list (in addition to the
collection&#39;s listRule), eg.:
`?filter=(title~'abc' &amp;&amp; created>'2022-01-01')`
Supported record filter fields:
id, created, updated,
+ any field from the collection schema.
The syntax basically follows the format
OPERAND
OPERATOR
OPERAND, where:
-OPERAND - could be any of the above field literal, string (single or double
quoted), number, null, true, false
-OPERATOR - is one of:
Equal
NOT equal
Greater than
Greater than or equal
-&lt;
Less than
-&lt;=
Less than or equal
Like/Contains (if not specified auto wraps the right string OPERAND in a &quot;%&quot; for wildcard
match)
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a &quot;%&quot; for
wildcard match)
Any/At least one of
Equal
Any/At least one of
NOT equal
Any/At least one of
Greater than
Any/At least one of
Greater than or equal
-?&lt;
Any/At least one of
Less than
-?&lt;=
Any/At least one of
Less than or equal
Any/At least one of
Like/Contains (if not specified auto wraps the right string OPERAND in a &quot;%&quot; for wildcard
match)
-?!~
Any/At least one of
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a &quot;%&quot; for
wildcard match)
To group and combine several expressions you could use parenthesis
(...), &amp;&amp; (AND) and || (OR) tokens.
Single line comments are also supported: // Example comment.
expand
String
Auto expand record relations. Ex.:
`?expand=relField1,relField2.subRelField`
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
`expand` property (eg. `"expand": {"relField1": {...}, ...}`).
Only the relations to which the request user has permissions to view will
be expanded.
fields
String
Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
-:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
skipTotal
Boolean
If it is set the total counts query will be skipped and the response fields
`totalItems` and `totalPages` will have `-1` value.
This could drastically speed up the search queries when the total counters are not needed or cursor based
pagination is used.
For optimization purposes, it is set by default for the
`getFirstListItem()`
and
`getFullList()` SDKs methods.
Responses
"page": 1,
"perPage": 100,
"totalItems": 2,
"totalPages": 1,
"items": [
"id": "ae40239d2bc4477",
"collectionId": "a98f514eb05f454",
"collectionName": "posts",
"updated": "2022-06-25 11:03:50.052",
"created": "2022-06-25 11:03:35.163",
"title": "test1"
"id": "d08dfc4f4d84419",
"collectionId": "a98f514eb05f454",
"collectionName": "posts",
"updated": "2022-06-25 11:03:45.876",
"created": "2022-06-25 11:03:45.876",
"title": "test2"
"code": 400,
"message": "Something went wrong while processing your request. Invalid filter.",
"data": {}
"code": 403,
"message": "Only admins can filter by '@collection.*'",
"data": {}
View record
Returns a single collection record by its ID.
Depending on the collection&#39;s viewRule value, the access to this action may or may not
have been restricted.
You could find individual generated records API documentation in the &quot;Admin UI &gt; Collections &gt;
API Preview&quot;.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const record1 = await pb.collection('posts').getOne('RECORD_ID', {
expand: 'relField1,relField2.subRelField',
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
final record1 = await pb.collection('posts').getOne('RECORD_ID',
expand: 'relField1,relField2.subRelField',
GET
/api/collections/`collectionIdOrName`/records/`recordId`
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the record&#39;s collection.
recordId
String
ID of the record to view.
Query parameters
Param
Type
Description
expand
String
Auto expand record relations. Ex.:
`?expand=relField1,relField2.subRelField`
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
`expand` property (eg. `"expand": {"relField1": {...}, ...}`).
Only the relations to which the request user has permissions to view will
be expanded.
fields
String
Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
-:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
Responses
"id": "ae40239d2bc4477",
"collectionId": "a98f514eb05f454",
"collectionName": "posts",
"updated": "2022-06-25 11:03:50.052",
"created": "2022-06-25 11:03:35.163",
"title": "test1"
"code": 403,
"message": "Only admins can perform this action.",
"data": {}
"code": 404,
"message": "The requested resource wasn't found.",
"data": {}
Create record
Creates a new collection Record.
Depending on the collection&#39;s createRule value, the access to this action may or may not
have been restricted.
You could find individual generated records API documentation from the admin UI.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const record = await pb.collection('demo').create({
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
final record = await pb.collection('demo').create(body: {
POST
/api/collections/`collectionIdOrName`/records
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the record&#39;s collection.
Body Parameters
Param
Type
Description
Optional
String
15 characters string to store as record ID.
If not set, it will be auto generated.
Schema fields
Any field from the collection&#39;s schema.
Auth record fields
Optional
username
String
The username of the auth record.
If not set, it will be auto generated.
Optional
email
String
Auth record email address.
Optional
emailVisibility
Boolean
Whether to show/hide the auth record email when fetching the record data.
Required
password
String
Auth record password.
Required
passwordConfirm
String
Auth record password confirmation.
Optional
verified
Boolean
Indicates whether the auth record is verified or not.
This field can be set only by admins or auth records with &quot;Manage&quot; access.
Body parameters could be sent as JSON or
multipart/form-data.
File upload is supported only through multipart/form-data.
Query parameters
Param
Type
Description
expand
String
Auto expand record relations. Ex.:
`?expand=relField1,relField2.subRelField`
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
`expand` property (eg. `"expand": {"relField1": {...}, ...}`).
Only the relations to which the request user has permissions to view will
be expanded.
fields
String
Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
-:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
Responses
"@collectionId": "a98f514eb05f454",
"@collectionName": "demo",
"id": "ae40239d2bc4477",
"updated": "2022-06-25 11:03:50.052",
"created": "2022-06-25 11:03:35.163",
"code": 400,
"message": "Failed to create record.",
"data": {
"title": {
"code": "validation_required",
"message": "Missing required value."
"code": 403,
"message": "Only admins can perform this action.",
"data": {}
"code": 404,
"message": "The requested resource wasn't found. Missing collection context.",
"data": {}
Update record
Updates an existing collection Record.
Depending on the collection&#39;s updateRule value, the access to this action may or may not
have been restricted.
You could find individual generated records API documentation from the admin UI.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const record = await pb.collection('demo').update('YOUR_RECORD_ID', {
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
final record = await pb.collection('demo').update('YOUR_RECORD_ID', body: {
PATCH
/api/collections/`collectionIdOrName`/records/`recordId`
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the record&#39;s collection.
recordId
String
ID of the record to update.
Body Parameters
Param
Type
Description
Schema fields
Any field from the collection&#39;s schema.
Auth record fields
Optional
username
String
The username of the auth record.
Optional
email
String
The auth record email address.
This field can be updated only by admins or auth records with &quot;Manage&quot; access.
Regular accounts can update their email by calling &quot;Request email change&quot;.
Optional
emailVisibility
Boolean
Whether to show/hide the auth record email when fetching the record data.
Optional
oldPassword
String
Old auth record password.
This field is required only when changing the record password. Admins and auth records with
&quot;Manage&quot; access can skip this field.
Optional
password
String
New auth record password.
Optional
passwordConfirm
String
New auth record password confirmation.
Optional
verified
Boolean
Indicates whether the auth record is verified or not.
This field can be set only by admins or auth records with &quot;Manage&quot; access.
Body parameters could be sent as JSON or
multipart/form-data.
File upload is supported only through multipart/form-data.
Query parameters
Param
Type
Description
expand
String
Auto expand record relations. Ex.:
`?expand=relField1,relField2.subRelField`
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
`expand` property (eg. `"expand": {"relField1": {...}, ...}`).
Only the relations to which the request user has permissions to view will
be expanded.
fields
String
Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
-:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
Responses
"@collectionId": "a98f514eb05f454",
"@collectionName": "demo",
"id": "ae40239d2bc4477",
"updated": "2022-06-25 11:03:50.052",
"created": "2022-06-25 11:03:35.163",
"code": 400,
"message": "Failed to create record.",
"data": {
"title": {
"code": "validation_required",
"message": "Missing required value."
"code": 403,
"message": "Only admins can perform this action.",
"data": {}
"code": 404,
"message": "The requested resource wasn't found. Missing collection context.",
"data": {}
Delete record
Deletes a single collection Record by its ID.
Depending on the collection&#39;s deleteRule value, the access to this action may or may not
have been restricted.
You could find individual generated records API documentation from the admin UI.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
await pb.collection('demo').delete('YOUR_RECORD_ID');
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
await pb.collection('demo').delete('YOUR_RECORD_ID');
DELETE
/api/collections/`collectionIdOrName`/records/`recordId`
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the record&#39;s collection.
recordId
String
ID of the record to delete.
Responses
`null`
"code": 400,
"message": "Failed to delete record. Make sure that the record is not part of a required relation reference.",
"data": {}
"code": 403,
"message": "Only admins can perform this action.",
"data": {}
"code": 404,
"message": "The requested resource wasn't found.",
"data": {}
### Auth record actions
List auth methods
Returns a public list with the allowed collection authentication methods.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const result = await pb.collection('users').listAuthMethods();
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
final result = await pb.collection('users').listAuthMethods();
GET
/api/collections/`collectionIdOrName`/auth-methods
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the auth collection.
Query parameters
Param
Type
Description
fields
String
Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
-:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
Responses
"usernamePassword": false,
"emailPassword": true,
"authProviders": [
"name": "github",
"state": "3Yd8jNkK_6PJG6hPWwBjLqKwse6Ejd",
"codeVerifier": "KxFDWz1B3fxscCDJ_9gHQhLuh__ie7",
"codeChallenge": "NM1oVexB6Q6QH8uPtOUfK7tq4pmu4Jz6lNDIwoxHZNE=",
"codeChallengeMethod": "S256",
"authUrl": "https://github.com/login/oauth/authorize?client_id=demo&amp;code_challenge=NM1oVexB6Q6QH8uPtOUfK7tq4pmu4Jz6lNDIwoxHZNE%3D&amp;code_challenge_method=S256&amp;response_type=code&amp;scope=user&amp;state=3Yd8jNkK_6PJG6hPWwBjLqKwse6Ejd&amp;redirect_uri="
"name": "gitlab",
"state": "NeQSbtO5cShr_mk5__3CUukiMnymeb",
"codeVerifier": "ahTFHOgua8mkvPAlIBGwCUJbWKR_xi",
"codeChallenge": "O-GATkTj4eXDCnfonsqGLCd6njvTixlpCMvy5kjgOOg=",
"codeChallengeMethod": "S256",
"authUrl": "https://gitlab.com/oauth/authorize?client_id=demo&amp;code_challenge=O-GATkTj4eXDCnfonsqGLCd6njvTixlpCMvy5kjgOOg%3D&amp;code_challenge_method=S256&amp;response_type=code&amp;scope=read_user&amp;state=NeQSbtO5cShr_mk5__3CUukiMnymeb&amp;redirect_uri="
"name": "google",
"state": "zB3ZPifV1TW2GMuvuFkamSXfSNkHPQ",
"codeVerifier": "t3CmO5VObGzdXqieakvR_fpjiW0zdO",
"codeChallenge": "KChwoQPKYlz2anAdqtgsSTdIo8hdwtc1fh2wHMwW2Yk=",
"codeChallengeMethod": "S256",
"authUrl": "https://accounts.google.com/o/oauth2/auth?client_id=demo&amp;code_challenge=KChwoQPKYlz2anAdqtgsSTdIo8hdwtc1fh2wHMwW2Yk%3D&amp;code_challenge_method=S256&amp;response_type=code&amp;scope=https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.profile+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.email&amp;state=zB3ZPifV1TW2GMuvuFkamSXfSNkHPQ&amp;redirect_uri="
Auth with password
Authenticate a single auth record by their username/email and password.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const authData = await pb.collection('users').authWithPassword(
'YOUR_USERNAME_OR_EMAIL',
'YOUR_PASSWORD',
// after the above you can also access the auth data from the authStore
console.log(pb.authStore.isValid);
console.log(pb.authStore.token);
console.log(pb.authStore.model.id);
// "logout" the last authenticated account
pb.authStore.clear();
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
final authData = await pb.collection('users').authWithPassword(
'YOUR_USERNAME_OR_EMAIL',
'YOUR_PASSWORD',
// after the above you can also access the auth data from the authStore
print(pb.authStore.isValid);
print(pb.authStore.token);
print(pb.authStore.model.id);
// "logout" the last authenticated account
pb.authStore.clear();
POST
/api/collections/`collectionIdOrName`/auth-with-password
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the auth collection.
Body Parameters
Param
Type
Description
Required
identity
String
Auth record username or email address.
Required
password
String
Auth record password.
Body parameters could be sent as JSON or
multipart/form-data.
Query parameters
Param
Type
Description
expand
String
Auto expand record relations. Ex.:
`?expand=relField1,relField2.subRelField`
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
`expand` property (eg. `"expand": {"relField1": {...}, ...}`).
Only the relations to which the request user has permissions to view will
be expanded.
fields
String
Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,record.expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
-:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,record.description:excerpt(200,true)
Responses
"token": "eyJhbGciOiJIUzI1NiJ9.eyJpZCI6IjRxMXhsY2xtZmxva3UzMyIsInR5cGUiOiJhdXRoUmVjb3JkIiwiY29sbGVjdGlvbklkIjoiX3BiX3VzZXJzX2F1dGhfIiwiZXhwIjoyMjA4OTg1MjYxfQ.UwD8JvkbQtXpymT09d7J6fdA0aP9g4FJ1GPh_ggEkzc",
"record": {
"id": "8171022dc95a4ed",
"collectionId": "d2972397d45614e",
"collectionName": "users",
"created": "2022-06-24 06:24:18.434Z",
"updated": "2022-06-24 06:24:18.889Z",
"verified": false,
"emailVisibility": true,
"someCustomField": "example 123"
"code": 400,
"message": "An error occurred while submitting the form.",
"data": {
"password": {
"code": "validation_required",
"message": "Missing required value."
Auth with OAuth2
Authenticate with an OAuth2 provider and returns a new auth token and record data.
This action usually should be called right after the provider login page redirect.
You could also check the
OAuth2 web integration example.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const authData = await pb.collection('users').authWithOAuth2Code(
'google',
'CODE',
'VERIFIER',
'REDIRECT_URL',
// optional data that will be used for the new account on OAuth2 sign-up
'name': 'test',
// after the above you can also access the auth data from the authStore
console.log(pb.authStore.isValid);
console.log(pb.authStore.token);
console.log(pb.authStore.model.id);
// "logout" the last authenticated account
pb.authStore.clear();
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
final authData = await pb.collection('users').authWithOAuth2Code(
'google',
'CODE',
'VERIFIER',
'REDIRECT_URL',
// optional data that will be used for the new account on OAuth2 sign-up
createData: {
'name': 'test',
// after the above you can also access the auth data from the authStore
print(pb.authStore.isValid);
print(pb.authStore.token);
print(pb.authStore.model.id);
// "logout" the last authenticated account
pb.authStore.clear();
POST
/api/collections/`collectionIdOrName`/auth-with-oauth2
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the auth collection.
Body Parameters
Param
Type
Description
Required
provider
String
The name of the OAuth2 client provider (eg. &quot;google&quot;).
Required
code
String
The authorization code returned from the initial request.
Required
codeVerifier
String
The code verifier sent with the initial request as part of the code_challenge.
Required
redirectUrl
String
The redirect url sent with the initial request.
Optional
createData
Object
Optional data that will be used when creating the auth record on OAuth2 sign-up.
The created auth record must comply with the same requirements and validations in the
regular create action.
The data can only be in json, aka. multipart/form-data and
files upload currently are not supported during OAuth2 sign-ups.
Body parameters could be sent as JSON or
multipart/form-data.
Query parameters
Param
Type
Description
expand
String
Auto expand record relations. Ex.:
`?expand=relField1,relField2.subRelField`
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
`expand` property (eg. `"expand": {"relField1": {...}, ...}`).
Only the relations to which the request user has permissions to view will
be expanded.
fields
String
Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,record.expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
-:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,record.description:excerpt(200,true)
Responses
"token": "eyJhbGciOiJIUzI1NiJ9.eyJpZCI6IjRxMXhsY2xtZmxva3UzMyIsInR5cGUiOiJhdXRoUmVjb3JkIiwiY29sbGVjdGlvbklkIjoiX3BiX3VzZXJzX2F1dGhfIiwiZXhwIjoyMjA4OTg1MjYxfQ.UwD8JvkbQtXpymT09d7J6fdA0aP9g4FJ1GPh_ggEkzc",
"record": {
"id": "8171022dc95a4ed",
"collectionId": "d2972397d45614e",
"collectionName": "users",
"created": "2022-06-24 06:24:18.434Z",
"updated": "2022-06-24 06:24:18.889Z",
"verified": true,
"emailVisibility": false,
"someCustomField": "example 123"
"meta": {
"id": "abc123",
"name": "John Doe",
"username": "john.doe",
"isNew": false,
"accessToken": "...",
"refreshToken": "...",
"expiry": "..."
"code": 400,
"message": "An error occurred while submitting the form.",
"data": {
"provider": {
"code": "validation_required",
"message": "Missing required value."
Auth refresh
Returns a new auth response (token and user data) for already authenticated auth record.
stored data in pb.authStore is still valid and up-to-date.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const authData = await pb.collection('users').authRefresh();
// after the above you can also access the refreshed auth data from the authStore
console.log(pb.authStore.isValid);
console.log(pb.authStore.token);
console.log(pb.authStore.model.id);
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
final authData = await pb.collection('users').authRefresh();
// after the above you can also access the refreshed auth data from the authStore
print(pb.authStore.isValid);
print(pb.authStore.token);
print(pb.authStore.model.id);
POST
/api/collections/`collectionIdOrName`/auth-refresh
Requires `Authorization: TOKEN`
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the auth collection.
Query parameters
Param
Type
Description
expand
String
Auto expand record relations. Ex.:
`?expand=relField1,relField2.subRelField`
Supports up to 6-levels depth nested relations expansion.
The expanded relations will be appended to the record under the
`expand` property (eg. `"expand": {"relField1": {...}, ...}`).
Only the relations to which the request user has permissions to view will
be expanded.
fields
String
Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,record.expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
-:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,record.description:excerpt(200,true)
Responses
"token": "eyJhbGciOiJIUzI1NiJ9.eyJpZCI6IjRxMXhsY2xtZmxva3UzMyIsInR5cGUiOiJhdXRoUmVjb3JkIiwiY29sbGVjdGlvbklkIjoiX3BiX3VzZXJzX2F1dGhfIiwiZXhwIjoyMjA4OTg1MjYxfQ.UwD8JvkbQtXpymT09d7J6fdA0aP9g4FJ1GPh_ggEkzc",
"record": {
"id": "8171022dc95a4ed",
"collectionId": "d2972397d45614e",
"collectionName": "users",
"created": "2022-06-24 06:24:18.434Z",
"updated": "2022-06-24 06:24:18.889Z",
"verified": false,
"emailVisibility": true,
"someCustomField": "example 123"
"code": 401,
"message": "The request requires valid record authorization token to be set.",
"data": {}
"code": 403,
"message": "The authorized record model is not allowed to perform this action.",
"data": {}
"code": 404,
"message": "Missing auth record context.",
"data": {}
Request verification
Sends auth record verification email request.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
POST
/api/collections/`collectionIdOrName`/request-verification
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the auth collection.
Body Parameters
Param
Type
Description
Required
email
String
The email address to send the password reset request (if registered).
Body parameters could be sent as JSON or
multipart/form-data.
Responses
`null`
"code": 400,
"message": "An error occurred while submitting the form.",
"data": {
"email": {
"code": "validation_required",
"message": "Missing required value."
Confirm verification
Confirms an email address verification request.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
await pb.collection('users').confirmVerification('TOKEN');
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
await pb.collection('users').confirmVerification('TOKEN');
POST
/api/collections/`collectionIdOrName`/confirm-verification
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the auth collection.
Body Parameters
Param
Type
Description
Required
token
String
The token from the verification request email.
Body parameters could be sent as JSON or
multipart/form-data.
Responses
`null`
"code": 400,
"message": "An error occurred while submitting the form.",
"data": {
"token": {
"code": "validation_required",
"message": "Missing required value."
Request password reset
Sends a password reset email to a specified auth record email.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
POST
/api/collections/`collectionIdOrName`/request-password-reset
Body Parameters
Param
Type
Description
Required
email
String
The email address to send the password reset request (if registered).
Body parameters could be sent as JSON or
multipart/form-data.
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the auth collection.
Responses
`null`
"code": 400,
"message": "An error occurred while submitting the form.",
"data": {
"email": {
"code": "validation_required",
"message": "Missing required value."
Confirm password reset
Confirms a password reset request and sets a new auth record password.
invalidated.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
await pb.collection('users').confirmPasswordReset('TOKEN', '1234567890', '1234567890');
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
await pb.collection('users').confirmPasswordReset('TOKEN', '1234567890', '1234567890');
POST
/api/collections/`collectionIdOrName`/confirm-password-reset
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the auth collection.
Body Parameters
Param
Type
Description
Required
token
String
The token from the password reset request email.
Required
password
String
The new auth record password to set.
Required
passwordConfirm
String
New auth record password confirmation.
Body parameters could be sent as JSON or
multipart/form-data.
Responses
`null`
"code": 400,
"message": "An error occurred while submitting the form.",
"data": {
"password": {
"code": "validation_required",
"message": "Missing required value."
Request email change
Sends an email change request for an authenticated record.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
POST
/api/collections/`collectionIdOrName`/request-email-change
Requires record `Authorization:TOKEN` header
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the auth collection.
Body Parameters
Param
Type
Description
Required
newEmail
String
The new email address to send the change email request.
Body parameters could be sent as JSON or
multipart/form-data.
Responses
`null`
"code": 400,
"message": "An error occurred while submitting the form.",
"data": {
"newEmail": {
"code": "validation_required",
"message": "Missing required value."
"code": 401,
"message": "The request requires valid record authorization token to be set.",
"data": {}
"code": 403,
"message": "You are not allowed to perform this request.",
"data": {}
Confirm email change
Confirms email address change.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
await pb.collection('users').confirmEmailChange('TOKEN', 'YOUR_PASSWORD');
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
await pb.collection('users').confirmEmailChange('TOKEN', 'YOUR_PASSWORD');
POST
/api/collections/`collectionIdOrName`/confirm-email-change
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the auth collection.
Body Parameters
Param
Type
Description
Required
token
String
The token from the change email request.
Required
password
String
The auth record password to confirm the email address change.
Body parameters could be sent as JSON or
multipart/form-data.
Responses
`null`
"code": 400,
"message": "An error occurred while submitting the form.",
"data": {
"token": {
"code": "validation_required",
"message": "Missing required value."
List linked external auth providers
Return a list with all external auth providers linked to a single record.
Only admins and the account owner can access this action.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const result = await pb.collection('users').listExternalAuths(pb.authStore.model.id);
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
final result = await pb.collection('users').listExternalAuths(pb.authStore.model.id);
GET
/api/collections/`collectionIdOrName`/records/`id`/external-auths
Requires `Authorization: TOKEN`
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the auth collection.
String
ID of the auth record.
Query parameters
Param
Type
Description
fields
String
Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
-:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
Responses
"id": "8171022dc95a4e8",
"created": "2022-09-01 10:24:18.434",
"updated": "2022-09-01 10:24:18.889",
"recordId": "e22581b6f1d44ea",
"collectionId": "systemprofiles0",
"provider": "google",
"providerId": "2da15468800514p"
"id": "171022dc895a4e8",
"created": "2022-09-01 10:24:18.434",
"updated": "2022-09-01 10:24:18.889",
"recordId": "e22581b6f1d44ea",
"collectionId": "systemprofiles0",
"provider": "twitter",
"providerId": "720688005140514"
"code": 401,
"message": "The request requires admin or record authorization token to be set.",
"data": {}
"code": 403,
"message": "You are not allowed to perform this request.",
"data": {}
"code": 404,
"message": "The requested resource wasn't found.",
"data": {}
Unlink external auth provider
Unlink a single external OAuth2 provider from an auth record.
Only admins and the account owner can access this action.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
await pb.collection('users').unlinkExternalAuth(pb.authStore.model.id, 'google');
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
await pb.collection('users').unlinkExternalAuth(pb.authStore.model.id, 'google');
DELETE
/api/collections/`collectionIdOrName`/records/`id`/external-auths/`provider`
*`Authorization: TOKEN`
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the auth collection.
String
ID of the auth record.
provider
String
The name of the auth provider to unlink, eg. `google`, `twitter`,
`github`, etc.
Responses
`null`
"code": 401,
"message": "The request requires valid record authorization token to be set.",
"data": {}
"code": 403,
"message": "You are not allowed to perform this request.",
"data": {}
"code": 404,
"message": "Missing external auth provider relation.",
"data": {}

## 12.Web APIs reference - API Realtime
Required
clientId(String):ID of the SSE client connection.
Optional
subscriptions(Array<String>):The new client subscriptions to set in the format:
COLLECTION_ID_OR_NAME or
COLLECTION_ID_OR_NAME/RECORD_ID.
You can also attach optional query and header parameters as serialized json to a
single topic using the options
query parameter, eg.:
COLLECTION_ID_OR_NAME/RECORD_ID?options={"query": {"abc": "123"}, "headers": {"x-token": "..."}}
Leave empty to unsubscribe from everything.
# Web APIs reference - API Realtime
- **Required
clientId** (String): ID of the SSE client connection.
- **Optional
subscriptions** (Array<String>): The new client subscriptions to set in the format:
COLLECTION_ID_OR_NAME or
COLLECTION_ID_OR_NAME/RECORD_ID.
You can also attach optional query and header parameters as serialized json to a
single topic using the options
query parameter, eg.:
COLLECTION_ID_OR_NAME/RECORD_ID?options={"query": {"abc": "123"}, "headers": {"x-token": "..."}}
Leave empty to unsubscribe from everything.
The Realtime API is implemented via Server-Sent Events (SSE). Generally, it consists of 2 operations:
-establish SSE connection
-submit client subscriptions
SSE events are sent for create, update
and delete record operations.
You could subscribe to a single record or to an entire collection.
When you subscribe to a single record, the collection&#39;s
ViewRule will be used to determine whether the subscriber has access to receive the
event message.
When you subscribe to an entire collection, the collection&#39;s
ListRule will be used to determine whether the subscriber has access to receive the
event message.
Connect
GET
/api/realtime
Establishes a new SSE connection and immediately sends a PB_CONNECT SSE event with the
created client ID.
NB! The user/admin authorization happens during the first
Set subscriptions
call.
If the connected client doesn&#39;t receive any new messages for 5 minutes, the server will send a
disconnect signal (this is to prevent forgotten/leaked connections). The connection will be
automatically reestablished if the client is still active (eg. the browser tab is still open).
Set subscriptions
POST
/api/realtime
If Authorization header is set, will authorize the client SSE connection with the
associated user or admin.
Body Parameters
Param
Type
Description
Required
clientId
String
ID of the SSE client connection.
Optional
subscriptions
Array&lt;String>
The new client subscriptions to set in the format:
COLLECTION_ID_OR_NAME or
COLLECTION_ID_OR_NAME/RECORD_ID.
You can also attach optional query and header parameters as serialized json to a
single topic using the options
query parameter, eg.:
COLLECTION_ID_OR_NAME/RECORD_ID?options={"query": {"abc": "123"}, "headers": {"x-token": "..."}}
Leave empty to unsubscribe from everything.
Body parameters could be sent as JSON or
multipart/form-data.
Responses
`null`
"code": 400,
"message": "Something went wrong while processing your request.",
"data": {
"clientId": {
"code": "validation_required",
"message": "Missing required value."
"code": 403,
"data": {}
"code": 404,
"message": "Missing or invalid client id.",
"data": {}
All of this is seamlessly handled by the SDKs using just the subscribe and
unsubscribe methods:
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
// (Optionally) authenticate
// Subscribe to changes in any record in the collection
pb.collection('example').subscribe('*', function (e) {
console.log(e.action);
console.log(e.record);
}, { /* other options like expand, custom headers, etc. */ });
// Subscribe to changes only in the specified record
pb.collection('example').subscribe('RECORD_ID', function (e) {
console.log(e.action);
console.log(e.record);
}, { /* other options like expand, custom headers, etc. */ });
// Unsubscribe
pb.collection('example').unsubscribe('RECORD_ID'); // remove all 'RECORD_ID' subscriptions
pb.collection('example').unsubscribe('*'); // remove all '*' topic subscriptions
pb.collection('example').unsubscribe(); // remove all subscriptions in the collection
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
// (Optionally) authenticate
// Subscribe to changes in any record in the collection
pb.collection('example').subscribe('*', (e) {
print(e.action);
print(e.record);
}, /* other options like expand, custom headers, etc. */);
// Subscribe to changes only in the specified record
pb.collection('example').subscribe('RECORD_ID', (e) {
print(e.action);
print(e.record);
}, /* other options like expand, custom headers, etc. */);
// Unsubscribe
pb.collection('example').unsubscribe('RECORD_ID'); // remove all 'RECORD_ID' subscriptions
pb.collection('example').unsubscribe('*'); // remove all '*' topic subscriptions
pb.collection('example').unsubscribe(); // remove all subscriptions in the collection

## 13.Web APIs reference - API Files
collectionIdOrName(String):ID or name of the collection whose record model contains the file resource.
recordId(String):ID of the record model that contains the file resource.
filename(String):Name of the file resource.
thumb(String):Get the thumb of the requested file.
The following thumb formats are currently supported:
WxH
(eg. 100x300) - crop to WxH viewbox (from center)
WxHt
(eg. 100x300t) - crop to WxH viewbox (from top)
WxHb
(eg. 100x300b) - crop to WxH viewbox (from bottom)
WxHf
(eg. 100x300f) - fit inside a WxH viewbox (without cropping)
0xH
(eg. 0x300) - resize to H height preserving the aspect ratio
Wx0
(eg. 100x0) - resize to W width preserving the aspect ratio
If the thumb size is not defined in the file schema field options or the file resource is not
an image (jpg, png, gif), then the original file resource is returned unmodified.
token(String):Optional file token for granting access to
protected file(s).
For an example, you can check
"Files upload and handling".
download(Boolean):If it is set to a truthy value (1, t, true) the file will be
served with Content-Disposition: attachment header instructing the browser to
ignore the file preview for pdf, images, videos, etc. and to directly download the file.
# Web APIs reference - API Files
- **collectionIdOrName** (String): ID or name of the collection whose record model contains the file resource.
- **recordId** (String): ID of the record model that contains the file resource.
- **filename** (String): Name of the file resource.
- **thumb** (String): Get the thumb of the requested file.
The following thumb formats are currently supported:
WxH
(eg. 100x300) - crop to WxH viewbox (from center)
WxHt
(eg. 100x300t) - crop to WxH viewbox (from top)
WxHb
(eg. 100x300b) - crop to WxH viewbox (from bottom)
WxHf
(eg. 100x300f) - fit inside a WxH viewbox (without cropping)
0xH
(eg. 0x300) - resize to H height preserving the aspect ratio
Wx0
(eg. 100x0) - resize to W width preserving the aspect ratio
If the thumb size is not defined in the file schema field options or the file resource is not
an image (jpg, png, gif), then the original file resource is returned unmodified.
- **token** (String): Optional file token for granting access to
protected file(s).
For an example, you can check
"Files upload and handling".
served with Content-Disposition: attachment header instructing the browser to
Files are uploaded, updated or deleted via the
Records API.
manipulations, like generating thumbs).
GET
/api/files/`collectionIdOrName`/`recordId`/`filename`
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the collection whose record model contains the file resource.
recordId
String
ID of the record model that contains the file resource.
filename
String
Name of the file resource.
Query parameters
Param
Type
Description
thumb
String
Get the thumb of the requested file.
The following thumb formats are currently supported:
-WxH
(eg. 100x300) - crop to WxH viewbox (from center)
-WxHt
(eg. 100x300t) - crop to WxH viewbox (from top)
-WxHb
(eg. 100x300b) - crop to WxH viewbox (from bottom)
-WxHf
(eg. 100x300f) - fit inside a WxH viewbox (without cropping)
-0xH
(eg. 0x300) - resize to H height preserving the aspect ratio
-Wx0
(eg. 100x0) - resize to W width preserving the aspect ratio
If the thumb size is not defined in the file schema field options or the file resource is not
an image (jpg, png, gif), then the original file resource is returned unmodified.
token
String
Optional file token for granting access to
protected file(s).
For an example, you can check
&quot;Files upload and handling&quot;.
Boolean
If it is set to a truthy value (1, t, true) the file will be
served with `Content-Disposition: attachment` header instructing the browser to
Responses
`[file resource]`
"code": 400,
"message": "Filesystem initialization failure.",
"data": {}
"code": 404,
"message": "The requested resource wasn't found.",
"data": {}
Generate protected file token
Generates a short-lived file token for accessing
protected file(s).
The client must be admin or auth record authenticated (aka. have regular authorization token sent
with the request).
POST
/api/files/token
Requires `Authorization: TOKEN`
Responses
"token": "..."
"code": 400,
"message": "Failed to generate file token.",
"data": {}

## 14.Web APIs reference - API Admins
Required
identity(String):Admin account identifier (currently only the email address is supported).
Required
password(String):Admin password.
fields(String):Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
fields(String):Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
Required
email(String):The email address to send the password reset request (if existing).
Required
token(String):The token from the password reset request email.
Required
password(String):The new admin password to set.
Required
passwordConfirm(String):New admin password confirmation.
page(Number):The page (aka. offset) of the paginated list (default to 1).
perPage(Number):The max returned admins per page (default to 30).
sort(String):Specify the ORDER BY fields.
Add - / + (default) in front of the attribute for DESC /
ASC order, eg.:
// DESC by created and ASC by id
?sort=-created,id
Supported admin sort fields:
@random, id, created,
updated, email, name
filter(String):Filter expression to filter/search the returned admins list, eg.:
?filter=(id='abc' && created>'2022-01-01')
Supported admin filter fields:
id, created, updated, email,
name
The syntax basically follows the format
OPERAND
OPERATOR
OPERAND, where:
OPERAND - could be any of the above field literal, string (single or double
quoted), number, null, true, false
OPERATOR - is one of:
Equal
!=
NOT equal
Greater than
>=
Greater than or equal
Less than
<=
Less than or equal
Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for wildcard
match)
!~
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for
wildcard match)
?=
Any/At least one of
Equal
?!=
Any/At least one of
NOT equal
?>
Any/At least one of
Greater than
?>=
Any/At least one of
Greater than or equal
?<
Any/At least one of
Less than
?<=
Any/At least one of
Less than or equal
?~
Any/At least one of
Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for wildcard
match)
?!~
Any/At least one of
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for
wildcard match)
To group and combine several expressions you could use parenthesis
(...), && (AND) and || (OR) tokens.
Single line comments are also supported: // Example comment.
fields(String):Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
skipTotal(Boolean):If it is set the total counts query will be skipped and the response fields
totalItems and totalPages will have -1 value.
This could drastically speed up the search queries when the total counters are not needed or cursor based
pagination is used.
For optimization purposes, it is set by default for the
getFirstListItem()
and
getFullList() SDKs methods.
id(String):ID of the admin to view.
fields(String):Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
Optional
id(String):15 characters string to store as admin ID.
If not set, it will be auto generated.
Required
email(String):Admin email address.
Required
password(String):Admin password.
Required
passwordConfirm(String):Admin password confirmation.
Optional
avatar(Number):Admin avatar image key (0-9).
fields(String):Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
id(String):ID of the admin to update.
Optional
email(String):New admin email address.
Optional
password(String):New admin password.
Optional
passwordConfirm(String):New admin password confirmation.
Optional
avatar(Number):New admin avatar key (0-9).
fields(String):Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
id(String):ID of the admin to delete.
# Web APIs reference - API Admins
- **Required
identity** (String): Admin account identifier (currently only the email address is supported).
- **Required
password** (String): Admin password.
- **fields** (String): Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
- **fields** (String): Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
- **Required
email** (String): The email address to send the password reset request (if existing).
- **Required
token** (String): The token from the password reset request email.
- **Required
password** (String): The new admin password to set.
- **Required
passwordConfirm** (String): New admin password confirmation.
- **page** (Number): The page (aka. offset) of the paginated list (default to 1).
- **perPage** (Number): The max returned admins per page (default to 30).
- **sort** (String): Specify the ORDER BY fields.
Add - / + (default) in front of the attribute for DESC /
ASC order, eg.:
// DESC by created and ASC by id
?sort=-created,id
Supported admin sort fields:
@random, id, created,
updated, email, name
- **filter** (String): Filter expression to filter/search the returned admins list, eg.:
?filter=(id='abc' && created>'2022-01-01')
Supported admin filter fields:
id, created, updated, email,
name
The syntax basically follows the format
OPERAND
OPERATOR
OPERAND, where:
OPERAND - could be any of the above field literal, string (single or double
quoted), number, null, true, false
OPERATOR - is one of:
Equal
NOT equal
Greater than
Greater than or equal
Less than
Less than or equal
Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for wildcard
match)
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for
wildcard match)
Any/At least one of
Equal
Any/At least one of
NOT equal
Any/At least one of
Greater than
Any/At least one of
Greater than or equal
Any/At least one of
Less than
Any/At least one of
Less than or equal
Any/At least one of
Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for wildcard
match)
?!~
Any/At least one of
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for
wildcard match)
To group and combine several expressions you could use parenthesis
(...), && (AND) and || (OR) tokens.
Single line comments are also supported: // Example comment.
- **fields** (String): Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
- **skipTotal** (Boolean): If it is set the total counts query will be skipped and the response fields
totalItems and totalPages will have -1 value.
This could drastically speed up the search queries when the total counters are not needed or cursor based
pagination is used.
For optimization purposes, it is set by default for the
getFirstListItem()
and
getFullList() SDKs methods.
- **id** (String): ID of the admin to view.
- **fields** (String): Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
- **Optional
id** (String): 15 characters string to store as admin ID.
If not set, it will be auto generated.
- **Required
email** (String): Admin email address.
- **Required
password** (String): Admin password.
- **Required
passwordConfirm** (String): Admin password confirmation.
- **Optional
avatar** (Number): Admin avatar image key (0-9).
- **fields** (String): Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
- **id** (String): ID of the admin to update.
- **Optional
email** (String): New admin email address.
- **Optional
password** (String): New admin password.
- **Optional
passwordConfirm** (String): New admin password confirmation.
- **Optional
avatar** (Number): New admin avatar key (0-9).
- **fields** (String): Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
- **id** (String): ID of the admin to delete.
Auth with password
Authenticate an admin with their email and password.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
POST
/api/admins/auth-with-password
Body Parameters
Param
Type
Description
Required
identity
String
Admin account identifier (currently only the email address is supported).
Required
password
String
Admin password.
Body parameters could be sent as JSON or
multipart/form-data.
Query parameters
Param
Type
Description
fields
String
Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
-:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
Responses
"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6InN5d2JoZWNuaDQ2cmhtMCIsInR5cGUiOiJhZG1pbiIsImV4cCI6MjIwODk4MTYwMH0.han3_sG65zLddpcX2ic78qgy7FKecuPfOpFa8Dvi5Bg",
"admin": {
"id": "b6e4b08274f34e9",
"created": "2022-06-22 07:13:09.735Z",
"updated": "2022-06-22 07:13:09.735Z",
"avatar": 0
"code": 400,
"message": "An error occurred while submitting the form.",
"data": {
"password": {
"code": "validation_required",
"message": "Missing required value."
Auth refresh
Returns a new auth response (token and admin data) for already authenticated admin.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const newAuthData = await pb.admins.authRefresh();
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
final newAuthData = await pb.admins.authRefresh();
POST
/api/admins/auth-refresh
Requires `Authorization: TOKEN`
Query parameters
Param
Type
Description
fields
String
Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
-:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
Responses
"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6InN5d2JoZWNuaDQ2cmhtMCIsInR5cGUiOiJhZG1pbiIsImV4cCI6MjIwODk4MTYwMH0.han3_sG65zLddpcX2ic78qgy7FKecuPfOpFa8Dvi5Bg",
"admin": {
"id": "b6e4b08274f34e9",
"created": "2022-06-22 07:13:09.735Z",
"updated": "2022-06-22 07:13:09.735Z",
"avatar": 0
"code": 401,
"message": "The request requires admin authorization token to be set.",
"data": {}
"code": 403,
"message": "You are not allowed to perform this request.",
"data": {}
"code": 404,
"message": "Missing auth admin context.",
"data": {}
Request password reset
Sends a password reset email to an admin.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
POST
/api/admins/request-password-reset
Body Parameters
Param
Type
Description
Required
email
String
The email address to send the password reset request (if existing).
Body parameters could be sent as JSON or
multipart/form-data.
Responses
`null`
"code": 400,
"message": "An error occurred while submitting the form.",
"data": {
"email": {
"code": "validation_required",
"message": "Missing required value."
Confirm password reset
Confirms a password reset request and sets a new admin password.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
// change password
await pb.admins.confirmPasswordReset('TOKEN', '1234567890', '1234567890');
// reauthenticate
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
// change password
await pb.admins.confirmPasswordReset('TOKEN', '1234567890', '1234567890');
// reauthenticate
POST
/api/admins/confirm-password-reset
Body Parameters
Param
Type
Description
Required
token
String
The token from the password reset request email.
Required
password
String
The new admin password to set.
Required
passwordConfirm
String
New admin password confirmation.
Body parameters could be sent as JSON or
multipart/form-data.
Responses
`null`
"code": 400,
"message": "An error occurred while submitting the form.",
"data": {
"password": {
"code": "validation_required",
"message": "Missing required value."
List admins
Returns a paginated admins list.
Only admins can perform this action.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
// fetch a paginated admins list
const resultList = await pb.admins.getList(1, 100, {
filter: 'created >= '2022-01-01 00:00:00'',
// you can also fetch all admins at once via getFullList
const admins = await pb.admins.getFullList({ sort: '-created' });
// or fetch only the first admin that matches the specified filter
const admin = await pb.admins.getFirstListItem('email~"test"');
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
// fetch a paginated admins list
final resultList = await pb.admins.getList(
page: 1,
perPage: 100,
filter: 'created >= "2022-01-01 00:00:00"',
// alternatively you can also fetch all admins at once via getFullList
final admins = await pb.admins.getFullList(sort: '-created');
// or fetch only the first admin that matches the specified filter
final admin = await pb.admins.getFirstListItem('email~"test"');
GET
/api/admins
Requires `Authorization: TOKEN`
Query parameters
Param
Type
Description
page
Number
The page (aka. offset) of the paginated list (default to 1).
perPage
Number
The max returned admins per page (default to 30).
sort
String
Specify the ORDER BY fields.
Add - / + (default) in front of the attribute for DESC /
ASC order, eg.:
// DESC by created and ASC by id
?sort=-created,id
Supported admin sort fields:
@random, id, created,
updated, email, name
filter
String
Filter expression to filter/search the returned admins list, eg.:
`?filter=(id='abc' &amp;&amp; created>'2022-01-01')`
Supported admin filter fields:
id, created, updated, email,
name
The syntax basically follows the format
OPERAND
OPERATOR
OPERAND, where:
-OPERAND - could be any of the above field literal, string (single or double
quoted), number, null, true, false
-OPERATOR - is one of:
Equal
NOT equal
Greater than
Greater than or equal
-&lt;
Less than
-&lt;=
Less than or equal
Like/Contains (if not specified auto wraps the right string OPERAND in a &quot;%&quot; for wildcard
match)
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a &quot;%&quot; for
wildcard match)
Any/At least one of
Equal
Any/At least one of
NOT equal
Any/At least one of
Greater than
Any/At least one of
Greater than or equal
-?&lt;
Any/At least one of
Less than
-?&lt;=
Any/At least one of
Less than or equal
Any/At least one of
Like/Contains (if not specified auto wraps the right string OPERAND in a &quot;%&quot; for wildcard
match)
-?!~
Any/At least one of
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a &quot;%&quot; for
wildcard match)
To group and combine several expressions you could use parenthesis
(...), &amp;&amp; (AND) and || (OR) tokens.
Single line comments are also supported: // Example comment.
fields
String
Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
-:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
skipTotal
Boolean
If it is set the total counts query will be skipped and the response fields
`totalItems` and `totalPages` will have `-1` value.
This could drastically speed up the search queries when the total counters are not needed or cursor based
pagination is used.
For optimization purposes, it is set by default for the
`getFirstListItem()`
and
`getFullList()` SDKs methods.
Responses
"page": 1,
"perPage": 100,
"totalItems": 2,
"totalPages": 1,
"items": [
"id": "b6e4b08274f34e9",
"created": "2022-06-22 07:13:09.735Z",
"updated": "2022-06-22 07:15:09.735Z",
"avatar": 0
"id": "e99c3f2aff6d695",
"created": "2022-06-25 16:14:23.037Z",
"updated": "2022-06-25 16:14:27.495Z",
"avatar": 6
"code": 400,
"message": "Something went wrong while processing your request. Invalid filter.",
"data": {}
"code": 401,
"message": "The request requires admin authorization token to be set.",
"data": {}
"code": 403,
"message": "Only admins can perform this action.",
"data": {}
View admin
Return a single admin by its ID.
Only admins can perform this action.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const admin = await pb.admins.getOne('ADMIN_ID');
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
final admin = await pb.admins.getOne('ADMIN_ID');
GET
/api/admins/`id`
Requires `Authorization: TOKEN`
Path parameters
Param
Type
Description
String
ID of the admin to view.
Query parameters
Param
Type
Description
fields
String
Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
-:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
Responses
"id": "b6e4b08274f34e9",
"created": "2022-06-22 07:13:09.735Z",
"updated": "2022-06-22 07:15:09.735Z",
"avatar": 0
"code": 401,
"message": "The request requires admin authorization token to be set.",
"data": {}
"code": 403,
"message": "You are not allowed to perform this request.",
"data": {}
"code": 404,
"message": "The requested resource wasn't found.",
"data": {}
Create admin
Creates a new admin.
Only admins can perform this action.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const admin = await pb.admins.create({
password: '1234567890',
passwordConfirm: '1234567890',
avatar: 8,
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
final admin = await pb.admins.create(body: {
'password': '1234567890',
'passwordConfirm': '1234567890',
'avatar': 8,
POST
/api/admins
Requires `Authorization: TOKEN`
Body Parameters
Param
Type
Description
Optional
String
15 characters string to store as admin ID.
If not set, it will be auto generated.
Required
email
String
Admin email address.
Required
password
String
Admin password.
Required
passwordConfirm
String
Admin password confirmation.
Optional
avatar
Number
Admin avatar image key (0-9).
Body parameters could be sent as JSON or
multipart/form-data.
Query parameters
Param
Type
Description
fields
String
Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
-:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
Responses
"id": "b6e4b08274f34e9",
"created": "2022-06-22 07:13:09.735Z",
"updated": "2022-06-22 07:15:09.735Z",
"avatar": 8
"code": 400,
"message": "An error occurred while submitting the form.",
"data": {
"email": {
"code": "validation_required",
"message": "Missing required value."
"code": 401,
"message": "The request requires admin authorization token to be set.",
"data": {}
"code": 403,
"message": "You are not allowed to perform this request.",
"data": {}
Update admin
Update a single admin model by its ID.
Only admins can perform this action.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const admin = await pb.admins.update('ADMIN_ID', {
password: '0987654321',
passwordConfirm: '0987654321',
avatar: 4,
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
final admin = await pb.admins.update('ADMIN_ID', body: {
'password': '0987654321',
'passwordConfirm': '0987654321',
'avatar': 4,
PATCH
/api/admins/`id`
Requires `Authorization: TOKEN`
Path parameters
Param
Type
Description
String
ID of the admin to update.
Body Parameters
Param
Type
Description
Optional
email
String
New admin email address.
Optional
password
String
New admin password.
Optional
passwordConfirm
String
New admin password confirmation.
Optional
avatar
Number
New admin avatar key (0-9).
Body parameters could be sent as JSON or
multipart/form-data.
Query parameters
Param
Type
Description
fields
String
Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
-:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
Responses
"id": "b6e4b08274f34e9",
"created": "2022-06-22 07:13:09.735Z",
"updated": "2022-06-22 07:15:09.735Z",
"avatar": 4
"code": 400,
"message": "An error occurred while submitting the form.",
"data": {
"email": {
"code": "validation_invalid_email",
"message": "Invalid email value."
"code": 401,
"message": "The request requires admin authorization token to be set.",
"data": {}
"code": 403,
"message": "You are not allowed to perform this request.",
"data": {}
"code": 404,
"message": "The requested resource wasn't found.",
"data": {}
Delete admin
Deletes a single admin by its id.
Only admins can perform this action.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
await pb.admins.delete('ADMIN_ID');
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
await pb.admins.delete('ADMIN_ID');
DELETE
/api/admins/`id`
Requires `Authorization: User/Admin TOKEN`
Path parameters
Param
Type
Description
String
ID of the admin to delete.
Responses
`null`
"code": 400,
"message": "Failed to delete admin. You cannot delete the only existing admin.",
"data": {}
"code": 401,
"message": "The request requires admin authorization token to be set.",
"data": {}
"code": 403,
"message": "You are not allowed to perform this request.",
"data": {}
"code": 404,
"message": "The requested resource wasn't found.",
"data": {}

## 15.Web APIs reference - API Collections
page(Number):The page (aka. offset) of the paginated list (default to 1).
perPage(Number):The max returned collections per page (default to 30).
sort(String):Specify the ORDER BY fields.
Add - / + (default) in front of the attribute for DESC /
ASC order, eg.:
// DESC by created and ASC by id
?sort=-created,id
Supported collection sort fields:
@random, id, created,
updated, name, type,
system
filter(String):Filter expression to filter/search the returned collections list, eg.:
?filter=(name~'abc' && created>'2022-01-01')
Supported collection filter fields:
id, created, updated,
name, type, system
The syntax basically follows the format
OPERAND
OPERATOR
OPERAND, where:
OPERAND - could be any of the above field literal, string (single or double
quoted), number, null, true, false
OPERATOR - is one of:
Equal
!=
NOT equal
Greater than
>=
Greater than or equal
Less than
<=
Less than or equal
Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for wildcard
match)
!~
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for
wildcard match)
?=
Any/At least one of
Equal
?!=
Any/At least one of
NOT equal
?>
Any/At least one of
Greater than
?>=
Any/At least one of
Greater than or equal
?<
Any/At least one of
Less than
?<=
Any/At least one of
Less than or equal
?~
Any/At least one of
Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for wildcard
match)
?!~
Any/At least one of
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for
wildcard match)
To group and combine several expressions you could use parenthesis
(...), && (AND) and || (OR) tokens.
Single line comments are also supported: // Example comment.
fields(String):Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
skipTotal(Boolean):If it is set the total counts query will be skipped and the response fields
totalItems and totalPages will have -1 value.
This could drastically speed up the search queries when the total counters are not needed or cursor based
pagination is used.
For optimization purposes, it is set by default for the
getFirstListItem()
and
getFullList() SDKs methods.
collectionIdOrName(String):ID or name of the collection to view.
fields(String):Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
Optional
id(String):15 characters string to store as collection ID.
If not set, it will be auto generated.
Required
name(String):Unique collection name (used as a table name for the records table).
Required
type(String):The type of the collection - base (default), auth or
view.
Req|Opt
schema*(Array):List with the collection fields.
This field is required for base collections.
This field is optional for auth collections.
This field is optional and autopopulated for view
collections based on the
options.query.
For more info about the supported fields and their options, you could check the
pocketbase/models/schema
Go sub-package definitions.
Optional
system(Boolean):Marks the collection as "system", aka. cannot be renamed or deleted.
Optional
listRule(null|String):API List action rule.
Check
Rules/Filters syntax guide
for more details.
Optional
viewRule(null|String):API View action rule.
Check
Rules/Filters syntax guide
for more details.
Optional
createRule(null|String):API Create action rule.
Check
Rules/Filters syntax guide
for more details.
This rule must be null for view collections.
Optional
updateRule(null|String):API Update action rule.
Check
Rules/Filters syntax guide
for more details.
This rule must be null for view collections.
Optional
deleteRule(null|String):API Delete action rule.
Check
Rules/Filters syntax guide
for more details.
This rule must be null for view collections.
Optional
indexes(Array<String>):The collection indexes and unique constriants.
Note that view collections don't support indexes.
├─
Required
query(null|String):The SQL SELECT statement that will be used to create the underlying view of the
collection.
├─
Optional
manageRule(null|String):API rule that gives admin-like permissions to allow fully managing the auth record(s), eg.
changing the password without requiring to enter the old one, directly updating the
verified state or email, etc. This rule is executed in addition to the
createRule and updateRule.
├─
Optional
allowOAuth2Auth(Boolean):Whether to allow OAuth2 sign-in/sign-up for the auth collection.
├─
Optional
allowUsernameAuth(Boolean):Whether to allow username + password authentication for the auth collection.
├─
Optional
allowEmailAuth(Boolean):Whether to allow email + password authentication for the auth collection.
├─
Optional
requireEmail(Boolean):Whether to always require email address when creating or updating auth records.
├─
Optional
exceptEmailDomains(Array<String>):Whether to allow sign-ups only with the email domains not listed in the specified list.
├─
Optional
onlyEmailDomains(Array<String>):Whether to allow sign-ups only with the email domains listed in the specified list.
├─
Optional
onlyVerified(Boolean):If enabled, it will return 403 for any new auth request performed by unverified user.
Note that when authenticating with OAuth2 for the first time, the user would be created with
verified=true even if the provider doesn't return an email.
└─
Optional
minPasswordLength*(Boolean):The minimum required password length for new auth records.
fields(String):Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
collectionIdOrName(String):ID or name of the collection to view.
Required
name(String):Unique collection name (used as a table name for the records table).
Required
type(String):The type of the collection - base (default), auth.
Req|Opt
schema*(Array):List with the collection fields.
This field is required for base collections.
This field is optional for auth collections.
This field is optional and autopopulated for view
collections based on the
options.query.
For more info about the supported fields and their options, you could check the
pocketbase/models/schema
Go sub-package definitions.
Optional
system(Boolean):Marks the collection as "system", aka. cannot be renamed or deleted.
Optional
listRule(null|String):API List action rule.
Check
Rules/Filters syntax guide
for more details.
Optional
viewRule(null|String):API View action rule.
Check
Rules/Filters syntax guide
for more details.
Optional
createRule(null|String):API Create action rule.
Check
Rules/Filters syntax guide
for more details.
This rule must be null for view collections.
Optional
updateRule(null|String):API Update action rule.
Check
Rules/Filters syntax guide
for more details.
This rule must be null for view collections.
Optional
deleteRule(null|String):API Delete action rule.
Check
Rules/Filters syntax guide
for more details.
This rule must be null for view collections.
Optional
indexes(Array<String>):The collection indexes and unique constriants.
Note that view collections don't support indexes.
├─
Required
query(null|String):The SQL SELECT statement that will be used to create the underlying view of the
collection.
├─
Optional
manageRule(null|String):API rule that gives admin-like permissions to allow fully managing the auth record(s), eg.
changing the password without requiring to enter the old one, directly updating the
verified state or email, etc. This rule is executed in addition to the
createRule and updateRule.
├─
Optional
allowOAuth2Auth(Boolean):Whether to allow OAuth2 sign-in/sign-up for the auth collection.
├─
Optional
allowUsernameAuth(Boolean):Whether to allow username + password authentication for the auth collection.
├─
Optional
allowEmailAuth(Boolean):Whether to allow email + password authentication for the auth collection.
├─
Optional
requireEmail(Boolean):Whether to always require email address when creating or updating auth records.
├─
Optional
exceptEmailDomains(Array<String>):Whether to allow sign-ups only with the email domains not listed in the specified list.
├─
Optional
onlyEmailDomains(Array<String>):Whether to allow sign-ups only with the email domains listed in the specified list.
├─
Optional
onlyVerified(Boolean):If enabled, it will return 403 for any new auth request performed by unverified user.
└─
Optional
minPasswordLength*(Boolean):The minimum required password length for new auth records.
fields(String):Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
collectionIdOrName(String):ID or name of the collection to view.
Required
collections(Array<Collection>):List of collections to import (replace and create).
Optional
deleteMissing(Boolean):If true all existing collections and schema fields that are not present in the
imported configuration will be deleted, including their related records
data (default to
false).
# Web APIs reference - API Collections
- **page** (Number): The page (aka. offset) of the paginated list (default to 1).
- **perPage** (Number): The max returned collections per page (default to 30).
- **sort** (String): Specify the ORDER BY fields.
Add - / + (default) in front of the attribute for DESC /
ASC order, eg.:
// DESC by created and ASC by id
?sort=-created,id
Supported collection sort fields:
@random, id, created,
updated, name, type,
system
- **filter** (String): Filter expression to filter/search the returned collections list, eg.:
?filter=(name~'abc' && created>'2022-01-01')
Supported collection filter fields:
id, created, updated,
name, type, system
The syntax basically follows the format
OPERAND
OPERATOR
OPERAND, where:
OPERAND - could be any of the above field literal, string (single or double
quoted), number, null, true, false
OPERATOR - is one of:
Equal
NOT equal
Greater than
Greater than or equal
Less than
Less than or equal
Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for wildcard
match)
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for
wildcard match)
Any/At least one of
Equal
Any/At least one of
NOT equal
Any/At least one of
Greater than
Any/At least one of
Greater than or equal
Any/At least one of
Less than
Any/At least one of
Less than or equal
Any/At least one of
Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for wildcard
match)
?!~
Any/At least one of
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for
wildcard match)
To group and combine several expressions you could use parenthesis
(...), && (AND) and || (OR) tokens.
Single line comments are also supported: // Example comment.
- **fields** (String): Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
- **skipTotal** (Boolean): If it is set the total counts query will be skipped and the response fields
totalItems and totalPages will have -1 value.
This could drastically speed up the search queries when the total counters are not needed or cursor based
pagination is used.
For optimization purposes, it is set by default for the
getFirstListItem()
and
getFullList() SDKs methods.
- **collectionIdOrName** (String): ID or name of the collection to view.
- **fields** (String): Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
- **Optional
id** (String): 15 characters string to store as collection ID.
If not set, it will be auto generated.
- **Required
name** (String): Unique collection name (used as a table name for the records table).
- **Required
type** (String): The type of the collection - base (default), auth or
view.
- **Req|Opt
schema** (Array) (required): List with the collection fields.
This field is required for base collections.
This field is optional for auth collections.
This field is optional and autopopulated for view
collections based on the
options.query.
For more info about the supported fields and their options, you could check the
pocketbase/models/schema
Go sub-package definitions.
- **Optional
system** (Boolean): Marks the collection as "system", aka. cannot be renamed or deleted.
- **Optional
listRule** (null|String): API List action rule.
Check
Rules/Filters syntax guide
for more details.
- **Optional
viewRule** (null|String): API View action rule.
Check
Rules/Filters syntax guide
for more details.
- **Optional
createRule** (null|String): API Create action rule.
Check
Rules/Filters syntax guide
for more details.
This rule must be null for view collections.
- **Optional
updateRule** (null|String): API Update action rule.
Check
Rules/Filters syntax guide
for more details.
This rule must be null for view collections.
- **Optional
deleteRule** (null|String): API Delete action rule.
Check
Rules/Filters syntax guide
for more details.
This rule must be null for view collections.
- **Optional
indexes** (Array<String>): The collection indexes and unique constriants.
Note that view collections don't support indexes.
- **├─
Required
query** (null|String): The SQL SELECT statement that will be used to create the underlying view of the
collection.
- **├─
Optional
manageRule** (null|String): API rule that gives admin-like permissions to allow fully managing the auth record(s), eg.
changing the password without requiring to enter the old one, directly updating the
verified state or email, etc. This rule is executed in addition to the
createRule and updateRule.
- **├─
Optional
allowOAuth2Auth** (Boolean): Whether to allow OAuth2 sign-in/sign-up for the auth collection.
- **├─
Optional
allowUsernameAuth** (Boolean): Whether to allow username + password authentication for the auth collection.
- **├─
Optional
allowEmailAuth** (Boolean): Whether to allow email + password authentication for the auth collection.
- **├─
Optional
requireEmail** (Boolean): Whether to always require email address when creating or updating auth records.
- **├─
Optional
exceptEmailDomains** (Array<String>): Whether to allow sign-ups only with the email domains not listed in the specified list.
- **├─
Optional
onlyEmailDomains** (Array<String>): Whether to allow sign-ups only with the email domains listed in the specified list.
- **├─
Optional
onlyVerified** (Boolean): If enabled, it will return 403 for any new auth request performed by unverified user.
Note that when authenticating with OAuth2 for the first time, the user would be created with
verified=true even if the provider doesn't return an email.
- **└─
Optional
minPasswordLength** (Boolean) (required): The minimum required password length for new auth records.
- **fields** (String): Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
- **collectionIdOrName** (String): ID or name of the collection to view.
- **Required
name** (String): Unique collection name (used as a table name for the records table).
- **Required
type** (String): The type of the collection - base (default), auth.
- **Req|Opt
schema** (Array) (required): List with the collection fields.
This field is required for base collections.
This field is optional for auth collections.
This field is optional and autopopulated for view
collections based on the
options.query.
For more info about the supported fields and their options, you could check the
pocketbase/models/schema
Go sub-package definitions.
- **Optional
system** (Boolean): Marks the collection as "system", aka. cannot be renamed or deleted.
- **Optional
listRule** (null|String): API List action rule.
Check
Rules/Filters syntax guide
for more details.
- **Optional
viewRule** (null|String): API View action rule.
Check
Rules/Filters syntax guide
for more details.
- **Optional
createRule** (null|String): API Create action rule.
Check
Rules/Filters syntax guide
for more details.
This rule must be null for view collections.
- **Optional
updateRule** (null|String): API Update action rule.
Check
Rules/Filters syntax guide
for more details.
This rule must be null for view collections.
- **Optional
deleteRule** (null|String): API Delete action rule.
Check
Rules/Filters syntax guide
for more details.
This rule must be null for view collections.
- **Optional
indexes** (Array<String>): The collection indexes and unique constriants.
Note that view collections don't support indexes.
- **├─
Required
query** (null|String): The SQL SELECT statement that will be used to create the underlying view of the
collection.
- **├─
Optional
manageRule** (null|String): API rule that gives admin-like permissions to allow fully managing the auth record(s), eg.
changing the password without requiring to enter the old one, directly updating the
verified state or email, etc. This rule is executed in addition to the
createRule and updateRule.
- **├─
Optional
allowOAuth2Auth** (Boolean): Whether to allow OAuth2 sign-in/sign-up for the auth collection.
- **├─
Optional
allowUsernameAuth** (Boolean): Whether to allow username + password authentication for the auth collection.
- **├─
Optional
allowEmailAuth** (Boolean): Whether to allow email + password authentication for the auth collection.
- **├─
Optional
requireEmail** (Boolean): Whether to always require email address when creating or updating auth records.
- **├─
Optional
exceptEmailDomains** (Array<String>): Whether to allow sign-ups only with the email domains not listed in the specified list.
- **├─
Optional
onlyEmailDomains** (Array<String>): Whether to allow sign-ups only with the email domains listed in the specified list.
- **├─
Optional
onlyVerified** (Boolean): If enabled, it will return 403 for any new auth request performed by unverified user.
- **└─
Optional
minPasswordLength** (Boolean) (required): The minimum required password length for new auth records.
- **fields** (String): Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
- **collectionIdOrName** (String): ID or name of the collection to view.
- **Required
collections** (Array<Collection>): List of collections to import (replace and create).
- **Optional
deleteMissing** (Boolean): If true all existing collections and schema fields that are not present in the
imported configuration will be deleted, including their related records
data (default to
false).
List collections
Returns a paginated Collections list.
Only admins can perform this action.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
// fetch a paginated collections list
const pageResult = await pb.collections.getList(1, 100, {
filter: 'created >= "2022-01-01 00:00:00"',
// you can also fetch all collections at once via getFullList
const collections = await pb.collections.getFullList({ sort: '-created' });
// or fetch only the first collection that matches the specified filter
const collection = await pb.collections.getFirstListItem('type="auth"');
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
// fetch a paginated collections list
final pageResult = await pb.collections.getList(
page: 1,
perPage: 100,
filter: 'created >= "2022-01-01 00:00:00"',
// you can also fetch all collections at once via getFullList
final collections = await pb.collections.getFullList(sort: '-created');
// or fetch only the first collection that matches the specified filter
final collection = await pb.collections.getFirstListItem('type="auth"');
GET
/api/collections
Requires `Authorization: TOKEN`
Query parameters
Param
Type
Description
page
Number
The page (aka. offset) of the paginated list (default to 1).
perPage
Number
The max returned collections per page (default to 30).
sort
String
Specify the ORDER BY fields.
Add - / + (default) in front of the attribute for DESC /
ASC order, eg.:
// DESC by created and ASC by id
?sort=-created,id
Supported collection sort fields:
@random, id, created,
updated, name, type,
system
filter
String
Filter expression to filter/search the returned collections list, eg.:
`?filter=(name~'abc' &amp;&amp; created>'2022-01-01')`
Supported collection filter fields:
id, created, updated,
name, type, system
The syntax basically follows the format
OPERAND
OPERATOR
OPERAND, where:
-OPERAND - could be any of the above field literal, string (single or double
quoted), number, null, true, false
-OPERATOR - is one of:
Equal
NOT equal
Greater than
Greater than or equal
-&lt;
Less than
-&lt;=
Less than or equal
Like/Contains (if not specified auto wraps the right string OPERAND in a &quot;%&quot; for wildcard
match)
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a &quot;%&quot; for
wildcard match)
Any/At least one of
Equal
Any/At least one of
NOT equal
Any/At least one of
Greater than
Any/At least one of
Greater than or equal
-?&lt;
Any/At least one of
Less than
-?&lt;=
Any/At least one of
Less than or equal
Any/At least one of
Like/Contains (if not specified auto wraps the right string OPERAND in a &quot;%&quot; for wildcard
match)
-?!~
Any/At least one of
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a &quot;%&quot; for
wildcard match)
To group and combine several expressions you could use parenthesis
(...), &amp;&amp; (AND) and || (OR) tokens.
Single line comments are also supported: // Example comment.
fields
String
Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
-:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
skipTotal
Boolean
If it is set the total counts query will be skipped and the response fields
`totalItems` and `totalPages` will have `-1` value.
This could drastically speed up the search queries when the total counters are not needed or cursor based
pagination is used.
For optimization purposes, it is set by default for the
`getFirstListItem()`
and
`getFullList()` SDKs methods.
Responses
"page": 1,
"perPage": 100,
"totalItems": 3,
"totalPages": 1,
"items": [
"id": "d2972397d45614e",
"created": "2022-06-22 07:13:00.643Z",
"updated": "2022-06-22 07:13:00.643Z",
"name": "users",
"type": "base",
"system": true,
"schema": [
"system": false,
"id": "njnkhxa2",
"name": "title",
"type": "text",
"required": false,
"unique": false,
"options": {
"min": null,
"max": null,
"pattern": ""
"system": false,
"id": "9gvv0jkj",
"name": "avatar",
"type": "file",
"required": false,
"unique": false,
"options": {
"maxSelect": 1,
"maxSize": 5242880,
"mimeTypes": [
"image/jpg",
"image/jpeg",
"image/png",
"image/svg+xml",
"image/gif"
"thumbs": null
"listRule": "id = @request.user.id",
"viewRule": "id = @request.user.id",
"createRule": "id = @request.user.id",
"updateRule": "id = @request.user.id",
"deleteRule": null,
"options": {
"manageRule": null,
"allowOAuth2Auth": true,
"allowUsernameAuth": true,
"allowEmailAuth": true,
"requireEmail": true,
"exceptEmailDomains": [],
"onlyEmailDomains": [],
"minPasswordLength": 8
"indexes": ["create index title_idx on users (title)"]
"id": "a98f514eb05f454",
"created": "2022-06-23 10:46:16.462Z",
"updated": "2022-06-24 13:25:04.170Z",
"name": "posts",
"system": false,
"schema": [
"system": false,
"id": "b7olyhbx",
"name": "title",
"type": "text",
"required": false,
"unique": false,
"options": {
"min": null,
"max": null,
"pattern": ""
"listRule": "title ~ 'test'",
"viewRule": null,
"createRule": null,
"updateRule": null,
"deleteRule": null,
"options": {},
"indexes": []
"code": 400,
"message": "Something went wrong while processing your request. Invalid filter.",
"data": {}
"code": 401,
"message": "The request requires admin authorization token to be set.",
"data": {}
"code": 403,
"message": "Only admins can perform this action.",
"data": {}
View collection
Returns a single Collection by its ID or name.
Only admins can perform this action.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const collection = await pb.collections.getOne('demo');
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
final collection = await pb.collections.getOne('demo');
GET
/api/collections/`collectionIdOrName`
Requires `Authorization: TOKEN`
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the collection to view.
Query parameters
Param
Type
Description
fields
String
Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
-:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
Responses
"id": "d2972397d45614e",
"created": "2022-06-22 07:13:00.643Z",
"updated": "2022-06-22 07:13:00.643Z",
"name": "posts",
"type": "base",
"schema": [
"system": false,
"id": "njnkhxa2",
"name": "title",
"type": "text",
"required": false,
"unique": false,
"options": {
"min": null,
"max": null,
"pattern": ""
"system": false,
"id": "9gvv0jkj",
"name": "image",
"type": "file",
"required": false,
"unique": false,
"options": {
"maxSelect": 1,
"maxSize": 5242880,
"mimeTypes": [
"image/jpg",
"image/jpeg",
"image/png",
"image/svg+xml",
"image/gif"
"thumbs": null
"listRule": "id = @request.user.id",
"viewRule": "id = @request.user.id",
"createRule": "id = @request.user.id",
"updateRule": "id = @request.user.id",
"deleteRule": null,
"options": {},
"indexes": ["create index title_idx on posts (title)"]
"code": 401,
"message": "The request requires admin authorization token to be set.",
"data": {}
"code": 403,
"message": "You are not allowed to perform this request.",
"data": {}
"code": 404,
"message": "The requested resource wasn't found.",
"data": {}
Create collection
Creates a new Collection.
Only admins can perform this action.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
// create base collection
const base = await pb.collections.create({
name: 'exampleBase',
type: 'base',
schema: [
name: 'title',
type: 'text',
required: true,
options: {
min: 10,
name: 'status',
type: 'bool',
// create auth collection
const auth = await pb.collections.create({
name: 'exampleAuth',
type: 'auth',
createRule: 'id = @request.auth.id',
updateRule: 'id = @request.auth.id',
deleteRule: 'id = @request.auth.id',
// schema is optional for auth collections
schema: [
name: 'name',
type: 'text',
options: {
allowOAuth2Auth: true,
requireEmail: true,
// create view collection
const view = await pb.collections.create({
name: 'exampleView',
type: 'view',
listRule: '@request.auth.id != ""',
viewRule: null,
// the schema will be autogenerated from the below query
options: {
query: 'SELECT id, name from posts',
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
// create base collection
final base = await pb.collections.create(body: {
'name': 'exampleBase',
'type': 'base',
'schema': [
'name': 'title',
'type': 'text',
'required': true,
'options': {
'min': 10,
'name': 'status',
'type': 'bool',
// create auth collection
final auth = await pb.collections.create(body: {
'name': 'exampleAuth',
'type': 'auth',
'createRule': 'id = @request.auth.id',
'updateRule': 'id = @request.auth.id',
'deleteRule': 'id = @request.auth.id',
// schema is optional for auth collections
'schema': [
'name': 'name',
'type': 'text',
'options': {
'allowOAuth2Auth': true,
'requireEmail': true,
// create view collection
final view = await pb.collections.create(body: {
'name': 'exampleView',
'type': 'view',
'listRule': '@request.auth.id != ""',
'viewRule': null,
// the schema will be autogenerated from the below query
'options': {
'query': 'SELECT id, name from posts',
POST
/api/collections
Requires `Authorization: TOKEN`
Body Parameters
Param
Type
Description
Optional
String
15 characters string to store as collection ID.
If not set, it will be auto generated.
Required
name
String
Unique collection name (used as a table name for the records table).
Required
type
String
The type of the collection - `base` (default), `auth` or
`view`.
Req|Opt
schema
Array
List with the collection fields.
This field is required for base collections.
This field is optional for auth collections.
This field is optional and autopopulated for view
collections based on the
options.query.
For more info about the supported fields and their options, you could check the
pocketbase/models/schema
Go sub-package definitions.
Optional
system
Boolean
Marks the collection as &quot;system&quot;, aka. cannot be renamed or deleted.
Optional
listRule
null|String
API List action rule.
Check
Rules/Filters syntax guide
for more details.
Optional
viewRule
null|String
API View action rule.
Check
Rules/Filters syntax guide
for more details.
Optional
createRule
null|String
API Create action rule.
Check
Rules/Filters syntax guide
for more details.
This rule must be null for view collections.
Optional
updateRule
null|String
API Update action rule.
Check
Rules/Filters syntax guide
for more details.
This rule must be null for view collections.
Optional
deleteRule
null|String
API Delete action rule.
Check
Rules/Filters syntax guide
for more details.
This rule must be null for view collections.
Optional
indexes
Array&lt;String>
The collection indexes and unique constriants.
Note that view collections don&#39;t support indexes.
options (view)
├─
Required
query
null|String
The SQL `SELECT` statement that will be used to create the underlying view of the
collection.
options (auth)
├─
Optional
manageRule
null|String
API rule that gives admin-like permissions to allow fully managing the auth record(s), eg.
changing the password without requiring to enter the old one, directly updating the
verified state or email, etc. This rule is executed in addition to the
`createRule` and `updateRule`.
├─
Optional
allowOAuth2Auth
Boolean
Whether to allow OAuth2 sign-in/sign-up for the auth collection.
├─
Optional
allowUsernameAuth
Boolean
Whether to allow username + password authentication for the auth collection.
├─
Optional
allowEmailAuth
Boolean
Whether to allow email + password authentication for the auth collection.
├─
Optional
requireEmail
Boolean
Whether to always require email address when creating or updating auth records.
├─
Optional
exceptEmailDomains
Array&lt;String>
Whether to allow sign-ups only with the email domains not listed in the specified list.
├─
Optional
onlyEmailDomains
Array&lt;String>
Whether to allow sign-ups only with the email domains listed in the specified list.
├─
Optional
onlyVerified
Boolean
If enabled, it will return 403 for any new auth request performed by unverified user.
Note that when authenticating with OAuth2 for the first time, the user would be created with
`verified=true` even if the provider doesn&#39;t return an email.
└─
Optional
minPasswordLength
Boolean
The minimum required password length for new auth records.
Body parameters could be sent as JSON or
multipart/form-data.
Query parameters
Param
Type
Description
fields
String
Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
-:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
Responses
"id": "d2972397d45614e",
"created": "2022-06-22 07:13:00.643Z",
"updated": "2022-06-22 07:13:00.643Z",
"type": "base",
"name": "posts",
"system": true,
"schema": [
"system": false,
"id": "njnkhxa2",
"name": "name",
"type": "text",
"required": false,
"unique": false,
"options": {
"min": null,
"max": null,
"pattern": ""
"system": false,
"id": "9gvv0jkj",
"name": "avatar",
"type": "file",
"required": false,
"unique": false,
"options": {
"maxSelect": 1,
"maxSize": 5242880,
"mimeTypes": [
"image/jpg",
"image/jpeg",
"image/png",
"image/svg+xml",
"image/gif"
"thumbs": null
"listRule": "id = @request.user.id",
"viewRule": "id = @request.user.id",
"createRule": "id = @request.user.id",
"updateRule": "id = @request.user.id",
"deleteRule": null,
"indexes": ["create index name_idx on posts (name)"]
"code": 400,
"message": "An error occurred while submitting the form.",
"data": {
"email": {
"code": "validation_required",
"message": "Missing required value."
"code": 401,
"message": "The request requires admin authorization token to be set.",
"data": {}
"code": 403,
"message": "You are not allowed to perform this request.",
"data": {}
Update collection
Updates a single Collection by its ID or name.
Only admins can perform this action.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const collection = await pb.collections.update('demo', {
name: 'new_demo',
listRule: 'created > "2022-01-01 00:00:00"',
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
final collection = await pb.collections.update('demo', body: {
'name': 'new_demo',
'listRule': 'created > "2022-01-01 00:00:00"',
PATCH
/api/collections/`collectionIdOrName`
Requires `Authorization: TOKEN`
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the collection to view.
Body Parameters
Param
Type
Description
Required
name
String
Unique collection name (used as a table name for the records table).
Required
type
String
The type of the collection - `base` (default), `auth`.
Req|Opt
schema
Array
List with the collection fields.
This field is required for base collections.
This field is optional for auth collections.
This field is optional and autopopulated for view
collections based on the
options.query.
For more info about the supported fields and their options, you could check the
pocketbase/models/schema
Go sub-package definitions.
Optional
system
Boolean
Marks the collection as &quot;system&quot;, aka. cannot be renamed or deleted.
Optional
listRule
null|String
API List action rule.
Check
Rules/Filters syntax guide
for more details.
Optional
viewRule
null|String
API View action rule.
Check
Rules/Filters syntax guide
for more details.
Optional
createRule
null|String
API Create action rule.
Check
Rules/Filters syntax guide
for more details.
This rule must be null for view collections.
Optional
updateRule
null|String
API Update action rule.
Check
Rules/Filters syntax guide
for more details.
This rule must be null for view collections.
Optional
deleteRule
null|String
API Delete action rule.
Check
Rules/Filters syntax guide
for more details.
This rule must be null for view collections.
Optional
indexes
Array&lt;String>
The collection indexes and unique constriants.
Note that view collections don&#39;t support indexes.
options (view)
├─
Required
query
null|String
The SQL `SELECT` statement that will be used to create the underlying view of the
collection.
options (view)
├─
Optional
manageRule
null|String
API rule that gives admin-like permissions to allow fully managing the auth record(s), eg.
changing the password without requiring to enter the old one, directly updating the
verified state or email, etc. This rule is executed in addition to the
`createRule` and `updateRule`.
├─
Optional
allowOAuth2Auth
Boolean
Whether to allow OAuth2 sign-in/sign-up for the auth collection.
├─
Optional
allowUsernameAuth
Boolean
Whether to allow username + password authentication for the auth collection.
├─
Optional
allowEmailAuth
Boolean
Whether to allow email + password authentication for the auth collection.
├─
Optional
requireEmail
Boolean
Whether to always require email address when creating or updating auth records.
├─
Optional
exceptEmailDomains
Array&lt;String>
Whether to allow sign-ups only with the email domains not listed in the specified list.
├─
Optional
onlyEmailDomains
Array&lt;String>
Whether to allow sign-ups only with the email domains listed in the specified list.
├─
Optional
onlyVerified
Boolean
If enabled, it will return 403 for any new auth request performed by unverified user.
└─
Optional
minPasswordLength
Boolean
The minimum required password length for new auth records.
Body parameters could be sent as JSON or
multipart/form-data.
Query parameters
Param
Type
Description
fields
String
Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
-:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
Responses
"id": "d2972397d45614e",
"created": "2022-06-22 07:13:00.643Z",
"updated": "2022-06-22 08:00:00.341Z",
"type": "base",
"name": "posts",
"schema": [
"system": false,
"id": "njnkhxa2",
"name": "name",
"type": "text",
"required": false,
"unique": false,
"options": {
"min": null,
"max": null,
"pattern": ""
"system": false,
"id": "9gvv0jkj",
"name": "avatar",
"type": "file",
"required": false,
"unique": false,
"options": {
"maxSelect": 1,
"maxSize": 5242880,
"mimeTypes": [
"image/jpg",
"image/jpeg",
"image/png",
"image/svg+xml",
"image/gif"
"thumbs": null
"listRule": "id = @request.user.id",
"viewRule": "id = @request.user.id",
"createRule": "id = @request.user.id",
"updateRule": "id = @request.user.id",
"deleteRule": null,
"indexes": ["create index name_idx on posts (name)"]
"code": 400,
"message": "An error occurred while submitting the form.",
"data": {
"email": {
"code": "validation_required",
"message": "Missing required value."
"code": 401,
"message": "The request requires admin authorization token to be set.",
"data": {}
"code": 403,
"message": "You are not allowed to perform this request.",
"data": {}
Delete collection
Deletes a single Collection by its ID or name.
Only admins can perform this action.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
await pb.collections.delete('demo');
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
await pb.collections.delete('demo');
DELETE
/api/collections/`collectionIdOrName`
Requires `Authorization: TOKEN`
Path parameters
Param
Type
Description
collectionIdOrName
String
ID or name of the collection to view.
Responses
`null`
"code": 400,
"message": "Failed to delete collection. Make sure that the collection is not referenced by other collections.",
"data": {}
"code": 401,
"message": "The request requires admin authorization token to be set.",
"data": {}
"code": 403,
"message": "You are not allowed to perform this request.",
"data": {}
"code": 404,
"message": "The requested resource wasn't found.",
"data": {}
Import collections
Bulk imports the provided Collections configuration.
Only admins can perform this action.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const importData = [
name: 'collection1',
schema: [
name: 'status',
type: 'bool',
name: 'collection2',
schema: [
name: 'title',
type: 'text',
await pb.collections.import(importData, false);
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
final importData = [
CollectionModel(
name: "collection1",
schema: [
SchemaField(name: "status", type: "bool"),
CollectionModel(
name: "collection2",
schema: [
SchemaField(name: "title", type: "text"),
await pb.collections.import(importData, deleteMissing: false);
PUT
/api/collections/import
Requires `Authorization: TOKEN`
Body Parameters
Param
Type
Description
Required
collections
Array&lt;Collection>
List of collections to import (replace and create).
Optional
deleteMissing
Boolean
If true all existing collections and schema fields that are not present in the
imported configuration will be deleted, including their related records
data (default to
false).
Body parameters could be sent as JSON or
multipart/form-data.
Responses
`null`
"code": 400,
"message": "An error occurred while submitting the form.",
"data": {
"collections": {
"code": "collections_import_failure",
"message": "Failed to import the collections configuration."
"code": 401,
"message": "The request requires admin authorization token to be set.",
"data": {}
"code": 403,
"message": "You are not allowed to perform this request.",
"data": {}

## 16.Web APIs reference - API Settings
fields(String):Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
├─
Required
appName(String):The app name.
├─
Required
appUrl(String):The app public absolute url.
├─
Optional
hideControls(Boolean):Hides the collection create and update controls from the Admin UI.
Useful to prevent making accidental schema changes when in production environment.
├─
Required
senderName(String):Transactional mails sender name.
├─
Required
senderAddress(String):Transactional mails sender address.
├─
Required
verificationTemplate(Object):The default user verification email template.
├─
Required
resetPasswordTemplate(Object):The default user reset password email template.
└─
Required
confirmEmailChangeTemplate(Object):The default user email change confirmation email template.
└─
Optional
maxDays(Number):Max retention period. Set to 0 for no logs.
├─
Optional
cron(String):Cron expression to schedule auto backups, eg. 0 0 * * *.
├─
Optional
cronMaxKeep(Number):The max number of cron generated backups to keep before removing older entries.
└─
Optional
s3(Object):S3 configuration (the same fields as for the S3 file storage settings).
├─
Optional
enabled(Boolean):Enable the use of the SMTP mail server for sending emails.
├─
Required
host*(String):Mail server host (required if SMTP is enabled).
├─
Required
port*(Number):Mail server port (required if SMTP is enabled).
├─
Optional
username(String):Mail server username.
├─
Optional
password(String):Mail server password.
├─
Optional
tls(Boolean):Whether to enforce TLS connection encryption.
When false StartTLS command is send, leaving the server to decide whether
to upgrade the connection or not).
├─
Optional
authMethod(String):The SMTP AUTH method to use - PLAIN or LOGIN (used mainly by Microsoft).
Default to PLAIN if empty.
└─
Optional
localName(String):Optional domain name or (IP address) to use for the initial EHLO/HELO exchange.
If not explicitly set, localhost will be used.
Note that some SMTP providers, such as Gmail SMTP-relay, requires a proper domain name and
and will reject attempts to use localhost.
├─
Optional
enabled(Boolean):Enable the use of a S3 compatible storage.
├─
Required
bucket*(String):S3 storage bucket (required if enabled).
├─
Required
region*(String):S3 storage region (required if enabled).
├─
Required
endpoint*(String):S3 storage public endpoint (required if enabled).
├─
Required
accessKey*(String):S3 storage access key (required if enabled).
├─
Required
secret*(String):S3 storage secret (required if enabled).
└─
Optional
forcePathStyle(Boolean):Forces the S3 request to use path-style addressing, eg.
"https://s3.amazonaws.com/BUCKET/KEY" instead of the default
"https://BUCKET.s3.amazonaws.com/KEY".
├─
Required
secret(String):Token secret (random 30+ characters).
└─
Required
duration(Number):Token validity duration in seconds.
├─
Required
secret(String):Token secret (random 30+ characters).
└─
Required
duration(Number):Token validity duration in seconds.
├─
Required
secret(String):Token secret (random 30+ characters).
└─
Required
duration(Number):Token validity duration in seconds.
├─
Required
secret(String):Token secret (random 30+ characters).
└─
Required
duration(Number):Token validity duration in seconds.
├─
Required
secret(String):Token secret (random 30+ characters).
└─
Required
duration(Number):Token validity duration in seconds.
├─
Required
secret(String):Token secret (random 30+ characters).
└─
Required
duration(Number):Token validity duration in seconds.
├─
Optional
enabled(Boolean):Enable the OAuth2 provider.
├─
Required
clientId*(String):The provider's app client id (required if enabled).
├─
Required
clientSecret*(String):The provider's app client secret (required if enabled).
├─
Optional
authUrl(String):The provider's authorization endpoint URL.
Default to https://accounts.google.com/o/oauth2/auth.
├─
Optional
tokenUrl(String):The provider's token endpoint URL.
Default to https://accounts.google.com/o/oauth2/token.
└─
Optional
userApiUrl(String):The provider's user profile endpoint URL.
Default to https://www.googleapis.com/oauth2/v1/userinfo.
├─
Optional
enabled(Boolean):Enable the OAuth2 provider.
├─
Required
clientId*(String):The provider's app client id (required if enabled).
├─
Required
clientSecret*(String):The provider's app client secret (required if enabled).
├─
Optional
authUrl(String):The provider's authorization endpoint URL.
Default to https://www.facebook.com/dialog/oauth.
├─
Optional
tokenUrl(String):The provider's token endpoint URL.
Default to https://graph.facebook.com/oauth/access_token.
└─
Optional
userApiUrl(String):The provider's user profile endpoint URL.
Default to https://graph.facebook.com/me?fields=name,email,picture.type(large).
├─
Optional
enabled(Boolean):Enable the OAuth2 provider.
├─
Required
clientId*(String):The provider's app client id (required if enabled).
├─
Required
clientSecret*(String):The provider's app client secret (required if enabled).
├─
Optional
authUrl(String):The provider's authorization endpoint URL.
Default to https://github.com/login/oauth/authorize.
├─
Optional
tokenUrl(String):The provider's token endpoint URL.
Default to https://github.com/login/oauth/access_token.
└─
Optional
userApiUrl(String):The provider's user profile endpoint URL.
Default to https://api.github.com/user.
├─
Optional
enabled(Boolean):Enable the OAuth2 provider.
├─
Required
clientId*(String):The provider's app client id (required if enabled).
├─
Required
clientSecret*(String):The provider's app client secret (required if enabled).
├─
Optional
authUrl(String):The provider's authorization endpoint URL.
Default to https://gitlab.com/oauth/authorize.
├─
Optional
tokenUrl(String):The provider's token endpoint URL.
Default to https://gitlab.com/oauth/token.
└─
Optional
userApiUrl(String):The provider's user profile endpoint URL.
Default to https://gitlab.com/api/v4/user.
├─
Optional
enabled(Boolean):Enable the OAuth2 provider.
├─
Required
clientId*(String):The provider's app client id (required if enabled).
└─
Required
clientSecret*(String):The provider's app client secret (required if enabled).
├─
Optional
enabled(Boolean):Enable the OAuth2 provider.
├─
Required
clientId*(String):The provider's app client id (required if enabled).
└─
Required
clientSecret*(String):The provider's app client secret (required if enabled).
├─
Optional
enabled(Boolean):Enable the OAuth2 provider.
├─
Required
clientId*(String):The provider's app client id (required if enabled).
├─
Required
clientSecret*(String):The provider's app client secret (required if enabled).
├─
Optional
authUrl(String):The provider's authorization endpoint URL.
Default to https://gitlab.com/oauth/authorize.
└─
Optional
tokenUrl(String):The provider's token endpoint URL.
Default to https://gitlab.com/oauth/token.
├─
Optional
enabled(Boolean):Enable the OAuth2 provider.
├─
Required
clientId*(String):The provider's app client id (required if enabled).
└─
Required
clientSecret*(String):The provider's app client secret (required if enabled).
fields(String):Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
Required
filesystem(String):The storage filesystem to test (storage or backups).
Required
email(String):The receiver of the test email.
Required
template(String):The test email template to send:
verification,
password-reset or
email-change.
Required
clientId(String):The identifier of your app (aka. Service ID).
Required
teamId(String):10-character string associated with your developer account (usually could be found next to
your name in the Apple Developer site).
Required
keyId(String):10-character key identifier generated for the "Sign in with Apple" private key associated
with your developer account.
Required
privateKey(String):PrivateKey is the private key associated to your app.
Required
duration(Number):Duration specifies how long the generated JWT token should be considered valid.
The specified value must be in seconds and max 15777000 (~6months).
# Web APIs reference - API Settings
- **fields** (String): Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
- **├─
Required
appName** (String): The app name.
- **├─
Required
appUrl** (String): The app public absolute url.
- **├─
Optional
hideControls** (Boolean): Hides the collection create and update controls from the Admin UI.
Useful to prevent making accidental schema changes when in production environment.
- **├─
Required
senderName** (String): Transactional mails sender name.
- **├─
Required
senderAddress** (String): Transactional mails sender address.
- **├─
Required
verificationTemplate** (Object): The default user verification email template.
- **├─
Required
resetPasswordTemplate** (Object): The default user reset password email template.
- **└─
Required
confirmEmailChangeTemplate** (Object): The default user email change confirmation email template.
- **└─
Optional
maxDays** (Number): Max retention period. Set to 0 for no logs.
- **├─
Optional
cron** (String): Cron expression to schedule auto backups, eg. 0 0 * * *.
- **├─
Optional
cronMaxKeep** (Number): The max number of cron generated backups to keep before removing older entries.
- **└─
Optional
s3** (Object): S3 configuration (the same fields as for the S3 file storage settings).
- **├─
Optional
enabled** (Boolean): Enable the use of the SMTP mail server for sending emails.
- **├─
Required
host** (String) (required): Mail server host (required if SMTP is enabled).
- **├─
Required
port** (Number) (required): Mail server port (required if SMTP is enabled).
- **├─
Optional
username** (String): Mail server username.
- **├─
Optional
password** (String): Mail server password.
- **├─
Optional
tls** (Boolean): Whether to enforce TLS connection encryption.
When false StartTLS command is send, leaving the server to decide whether
to upgrade the connection or not).
- **├─
Optional
authMethod** (String): The SMTP AUTH method to use - PLAIN or LOGIN (used mainly by Microsoft).
Default to PLAIN if empty.
- **└─
Optional
localName** (String): Optional domain name or (IP address) to use for the initial EHLO/HELO exchange.
If not explicitly set, localhost will be used.
Note that some SMTP providers, such as Gmail SMTP-relay, requires a proper domain name and
and will reject attempts to use localhost.
- **├─
Optional
enabled** (Boolean): Enable the use of a S3 compatible storage.
- **├─
Required
bucket** (String) (required): S3 storage bucket (required if enabled).
- **├─
Required
region** (String) (required): S3 storage region (required if enabled).
- **├─
Required
endpoint** (String) (required): S3 storage public endpoint (required if enabled).
- **├─
Required
accessKey** (String) (required): S3 storage access key (required if enabled).
- **├─
Required
secret** (String) (required): S3 storage secret (required if enabled).
- **└─
Optional
forcePathStyle** (Boolean): Forces the S3 request to use path-style addressing, eg.
"https://s3.amazonaws.com/BUCKET/KEY" instead of the default
"https://BUCKET.s3.amazonaws.com/KEY".
- **├─
Required
secret** (String): Token secret (random 30+ characters).
- **└─
Required
duration** (Number): Token validity duration in seconds.
- **├─
Required
secret** (String): Token secret (random 30+ characters).
- **└─
Required
duration** (Number): Token validity duration in seconds.
- **├─
Required
secret** (String): Token secret (random 30+ characters).
- **└─
Required
duration** (Number): Token validity duration in seconds.
- **├─
Required
secret** (String): Token secret (random 30+ characters).
- **└─
Required
duration** (Number): Token validity duration in seconds.
- **├─
Required
secret** (String): Token secret (random 30+ characters).
- **└─
Required
duration** (Number): Token validity duration in seconds.
- **├─
Required
secret** (String): Token secret (random 30+ characters).
- **└─
Required
duration** (Number): Token validity duration in seconds.
- **├─
Optional
enabled** (Boolean): Enable the OAuth2 provider.
- **├─
Required
clientId** (String) (required): The provider's app client id (required if enabled).
- **├─
Required
clientSecret** (String) (required): The provider's app client secret (required if enabled).
- **├─
Optional
authUrl** (String): The provider's authorization endpoint URL.
Default to https://accounts.google.com/o/oauth2/auth.
- **├─
Optional
tokenUrl** (String): The provider's token endpoint URL.
Default to https://accounts.google.com/o/oauth2/token.
- **└─
Optional
userApiUrl** (String): The provider's user profile endpoint URL.
Default to https://www.googleapis.com/oauth2/v1/userinfo.
- **├─
Optional
enabled** (Boolean): Enable the OAuth2 provider.
- **├─
Required
clientId** (String) (required): The provider's app client id (required if enabled).
- **├─
Required
clientSecret** (String) (required): The provider's app client secret (required if enabled).
- **├─
Optional
authUrl** (String): The provider's authorization endpoint URL.
Default to https://www.facebook.com/dialog/oauth.
- **├─
Optional
tokenUrl** (String): The provider's token endpoint URL.
Default to https://graph.facebook.com/oauth/access_token.
- **└─
Optional
userApiUrl** (String): The provider's user profile endpoint URL.
Default to https://graph.facebook.com/me?fields=name,email,picture.type(large).
- **├─
Optional
enabled** (Boolean): Enable the OAuth2 provider.
- **├─
Required
clientId** (String) (required): The provider's app client id (required if enabled).
- **├─
Required
clientSecret** (String) (required): The provider's app client secret (required if enabled).
- **├─
Optional
authUrl** (String): The provider's authorization endpoint URL.
Default to https://github.com/login/oauth/authorize.
- **├─
Optional
tokenUrl** (String): The provider's token endpoint URL.
Default to https://github.com/login/oauth/access_token.
- **└─
Optional
userApiUrl** (String): The provider's user profile endpoint URL.
Default to https://api.github.com/user.
- **├─
Optional
enabled** (Boolean): Enable the OAuth2 provider.
- **├─
Required
clientId** (String) (required): The provider's app client id (required if enabled).
- **├─
Required
clientSecret** (String) (required): The provider's app client secret (required if enabled).
- **├─
Optional
authUrl** (String): The provider's authorization endpoint URL.
Default to https://gitlab.com/oauth/authorize.
- **├─
Optional
tokenUrl** (String): The provider's token endpoint URL.
Default to https://gitlab.com/oauth/token.
- **└─
Optional
userApiUrl** (String): The provider's user profile endpoint URL.
Default to https://gitlab.com/api/v4/user.
- **├─
Optional
enabled** (Boolean): Enable the OAuth2 provider.
- **├─
Required
clientId** (String) (required): The provider's app client id (required if enabled).
- **└─
Required
clientSecret** (String) (required): The provider's app client secret (required if enabled).
- **├─
Optional
enabled** (Boolean): Enable the OAuth2 provider.
- **├─
Required
clientId** (String) (required): The provider's app client id (required if enabled).
- **└─
Required
clientSecret** (String) (required): The provider's app client secret (required if enabled).
- **├─
Optional
enabled** (Boolean): Enable the OAuth2 provider.
- **├─
Required
clientId** (String) (required): The provider's app client id (required if enabled).
- **├─
Required
clientSecret** (String) (required): The provider's app client secret (required if enabled).
- **├─
Optional
authUrl** (String): The provider's authorization endpoint URL.
Default to https://gitlab.com/oauth/authorize.
- **└─
Optional
tokenUrl** (String): The provider's token endpoint URL.
Default to https://gitlab.com/oauth/token.
- **├─
Optional
enabled** (Boolean): Enable the OAuth2 provider.
- **├─
Required
clientId** (String) (required): The provider's app client id (required if enabled).
- **└─
Required
clientSecret** (String) (required): The provider's app client secret (required if enabled).
- **fields** (String): Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
- **Required
filesystem** (String): The storage filesystem to test (storage or backups).
- **Required
email** (String): The receiver of the test email.
- **Required
template** (String): The test email template to send:
verification,
password-reset or
email-change.
- **Required
clientId** (String): The identifier of your app (aka. Service ID).
- **Required
your name in the Apple Developer site).
- **Required
keyId** (String): 10-character key identifier generated for the "Sign in with Apple" private key associated
with your developer account.
- **Required
privateKey** (String): PrivateKey is the private key associated to your app.
- **Required
duration** (Number): Duration specifies how long the generated JWT token should be considered valid.
The specified value must be in seconds and max 15777000 (~6months).
List settings
Returns a list with all available application settings.
Secret/password fields are automatically redacted with ****** characters.
Only admins can perform this action.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const settings = await pb.settings.getAll();
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
final settings = await pb.settings.getAll();
GET
/api/settings
Requires `Authorization: TOKEN`
Query parameters
Param
Type
Description
fields
String
Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
-:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
Responses
"meta": {
"appName": "Acme",
"appUrl": "http://127.0.0.1:8090",
"hideControls": false,
"senderName": "Support",
"verificationTemplate": { ... },
"resetPasswordTemplate": { ... },
"confirmEmailChangeTemplate": { ... }
"logs": {
"maxDays": 7
"backups": {
"cron": "0 0 * * *",
"cronMaxKeep": 1,
"s3": {
"enabled": false,
"bucket": "",
"region": "",
"endpoint": "",
"accessKey": "",
"secret": "",
"forcePathStyle": false
"smtp": {
"enabled": false,
"port": 587,
"username": "",
"password": "",
"tls": true,
"authMethod": "",
"localName": ""
"s3": {
"enabled": false,
"bucket": "",
"region": "",
"endpoint": "",
"accessKey": "",
"secret": "",
"forcePathStyle": false
"adminAuthToken": {
"secret": "******",
"duration": 1209600
"adminPasswordResetToken": {
"secret": "******",
"duration": 1800
"recordAuthToken": {
"secret": "******",
"duration": 1209600
"recordPasswordResetToken": {
"secret": "******",
"duration": 1800
"recordEmailChangeToken": {
"secret": "******",
"duration": 1800
"recordVerificationToken": {
"secret": "******",
"duration": 604800
"googleAuth": {
"enabled": true,
"clientId": "demo",
"clientSecret": "******"
"facebookAuth": {
"enabled": false,
"allowRegistrations": false
"githubAuth": {
"enabled": true,
"clientId": "demo",
"clientSecret": "******"
"gitlabAuth": {
"enabled": true,
"clientId": "demo",
"clientSecret": "******"
"discordAuth": {
"enabled": true,
"clientId": "demo",
"clientSecret": "******"
"twitterAuth": {
"enabled": true,
"clientId": "demo",
"clientSecret": "******"
"microsoftAuth": {
"enabled": true,
"clientId": "demo",
"clientSecret": "******"
"spotifyAuth": {
"enabled": true,
"clientId": "demo",
"clientSecret": "******"
"code": 401,
"message": "The request requires admin authorization token to be set.",
"data": {}
"code": 403,
"message": "You are not allowed to perform this request.",
"data": {}
Update settings
Bulk updates application settings and returns the updated settings list.
Only admins can perform this action.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const settings = await pb.settings.update({
meta: {
appName: 'YOUR_APP',
appUrl: 'http://127.0.0.1:8090',
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
final settings = await pb.settings.update(body: {
'meta': {
'appName': 'YOUR_APP',
'appUrl': 'http://127.0.0.1:8090',
PATCH
/api/settings
Requires `Authorization: TOKEN`
Body Parameters
Param
Type
Description
meta
Application meta data (name, url, support email, etc.).
├─
Required
appName
String
The app name.
├─
Required
appUrl
String
The app public absolute url.
├─
Optional
hideControls
Boolean
Hides the collection create and update controls from the Admin UI.
Useful to prevent making accidental schema changes when in production environment.
├─
Required
senderName
String
Transactional mails sender name.
├─
Required
senderAddress
String
Transactional mails sender address.
├─
Required
verificationTemplate
Object
The default user verification email template.
├─
Required
resetPasswordTemplate
Object
The default user reset password email template.
└─
Required
confirmEmailChangeTemplate
Object
The default user email change confirmation email template.
logs
Request logs settings.
└─
Optional
maxDays
Number
Max retention period. Set to 0 for no logs.
backups
App data backups settings.
├─
Optional
cron
String
Cron expression to schedule auto backups, eg. `0 0 * * *`.
├─
Optional
cronMaxKeep
Number
The max number of cron generated backups to keep before removing older entries.
└─
Optional
Object
S3 configuration (the same fields as for the S3 file storage settings).
smtp
SMTP mail server settings.
├─
Optional
enabled
Boolean
Enable the use of the SMTP mail server for sending emails.
├─
Required
host
String
Mail server host (required if SMTP is enabled).
├─
Required
port
Number
Mail server port (required if SMTP is enabled).
├─
Optional
username
String
Mail server username.
├─
Optional
password
String
Mail server password.
├─
Optional
tls
Boolean
Whether to enforce TLS connection encryption.
When false StartTLS command is send, leaving the server to decide whether
to upgrade the connection or not).
├─
Optional
authMethod
String
The SMTP AUTH method to use - PLAIN or LOGIN (used mainly by Microsoft).
Default to PLAIN if empty.
└─
Optional
localName
String
Optional domain name or (IP address) to use for the initial EHLO/HELO exchange.
If not explicitly set, `localhost` will be used.
Note that some SMTP providers, such as Gmail SMTP-relay, requires a proper domain name and
and will reject attempts to use localhost.
S3 compatible file storage settings.
├─
Optional
enabled
Boolean
Enable the use of a S3 compatible storage.
├─
Required
bucket
String
S3 storage bucket (required if enabled).
├─
Required
region
String
S3 storage region (required if enabled).
├─
Required
endpoint
String
S3 storage public endpoint (required if enabled).
├─
Required
accessKey
String
S3 storage access key (required if enabled).
├─
Required
secret
String
S3 storage secret (required if enabled).
└─
Optional
forcePathStyle
Boolean
Forces the S3 request to use path-style addressing, eg.
&quot;https://s3.amazonaws.com/BUCKET/KEY&quot; instead of the default
&quot;https://BUCKET.s3.amazonaws.com/KEY&quot;.
adminAuthToken
Admin authentication token options.
├─
Required
secret
String
Token secret (random 30+ characters).
└─
Required
duration
Number
Token validity duration in seconds.
adminPasswordResetToken
Admin password reset token options.
├─
Required
secret
String
Token secret (random 30+ characters).
└─
Required
duration
Number
Token validity duration in seconds.
recordAuthToken
Record authentication token options.
├─
Required
secret
String
Token secret (random 30+ characters).
└─
Required
duration
Number
Token validity duration in seconds.
recordPasswordResetToken
Record password reset token options.
├─
Required
secret
String
Token secret (random 30+ characters).
└─
Required
duration
Number
Token validity duration in seconds.
recordEmailChangeToken
Record email change token options.
├─
Required
secret
String
Token secret (random 30+ characters).
└─
Required
duration
Number
Token validity duration in seconds.
recordVerificationToken
Record verification token options.
├─
Required
secret
String
Token secret (random 30+ characters).
└─
Required
duration
Number
Token validity duration in seconds.
googleAuth
Google OAuth2 provider settings.
├─
Optional
enabled
Boolean
Enable the OAuth2 provider.
├─
Required
clientId
String
The provider&#39;s app client id (required if enabled).
├─
Required
clientSecret
String
The provider&#39;s app client secret (required if enabled).
├─
Optional
authUrl
String
The provider&#39;s authorization endpoint URL.
Default to https://accounts.google.com/o/oauth2/auth.
├─
Optional
tokenUrl
String
The provider&#39;s token endpoint URL.
Default to https://accounts.google.com/o/oauth2/token.
└─
Optional
userApiUrl
String
The provider&#39;s user profile endpoint URL.
Default to https://www.googleapis.com/oauth2/v1/userinfo.
facebookAuth
Facebook OAuth2 provider settings.
├─
Optional
enabled
Boolean
Enable the OAuth2 provider.
├─
Required
clientId
String
The provider&#39;s app client id (required if enabled).
├─
Required
clientSecret
String
The provider&#39;s app client secret (required if enabled).
├─
Optional
authUrl
String
The provider&#39;s authorization endpoint URL.
Default to https://www.facebook.com/dialog/oauth.
├─
Optional
tokenUrl
String
The provider&#39;s token endpoint URL.
Default to https://graph.facebook.com/oauth/access_token.
└─
Optional
userApiUrl
String
The provider&#39;s user profile endpoint URL.
Default to https://graph.facebook.com/me?fields=name,email,picture.type(large).
githubAuth
GitHub OAuth2 provider settings.
├─
Optional
enabled
Boolean
Enable the OAuth2 provider.
├─
Required
clientId
String
The provider&#39;s app client id (required if enabled).
├─
Required
clientSecret
String
The provider&#39;s app client secret (required if enabled).
├─
Optional
authUrl
String
The provider&#39;s authorization endpoint URL.
Default to https://github.com/login/oauth/authorize.
├─
Optional
tokenUrl
String
The provider&#39;s token endpoint URL.
Default to https://github.com/login/oauth/access_token.
└─
Optional
userApiUrl
String
The provider&#39;s user profile endpoint URL.
Default to https://api.github.com/user.
gitlabAuth
GitLab OAuth2 provider settings.
├─
Optional
enabled
Boolean
Enable the OAuth2 provider.
├─
Required
clientId
String
The provider&#39;s app client id (required if enabled).
├─
Required
clientSecret
String
The provider&#39;s app client secret (required if enabled).
├─
Optional
authUrl
String
The provider&#39;s authorization endpoint URL.
Default to https://gitlab.com/oauth/authorize.
├─
Optional
tokenUrl
String
The provider&#39;s token endpoint URL.
Default to https://gitlab.com/oauth/token.
└─
Optional
userApiUrl
String
The provider&#39;s user profile endpoint URL.
Default to https://gitlab.com/api/v4/user.
discordAuth
Discord OAuth2 provider settings.
├─
Optional
enabled
Boolean
Enable the OAuth2 provider.
├─
Required
clientId
String
The provider&#39;s app client id (required if enabled).
└─
Required
clientSecret
String
The provider&#39;s app client secret (required if enabled).
twitterAuth
Twitter OAuth2 provider settings.
├─
Optional
enabled
Boolean
Enable the OAuth2 provider.
├─
Required
clientId
String
The provider&#39;s app client id (required if enabled).
└─
Required
clientSecret
String
The provider&#39;s app client secret (required if enabled).
microsoftAuth
Microsoft Azure AD OAuth2 provider settings.
├─
Optional
enabled
Boolean
Enable the OAuth2 provider.
├─
Required
clientId
String
The provider&#39;s app client id (required if enabled).
├─
Required
clientSecret
String
The provider&#39;s app client secret (required if enabled).
├─
Optional
authUrl
String
The provider&#39;s authorization endpoint URL.
Default to https://gitlab.com/oauth/authorize.
└─
Optional
tokenUrl
String
The provider&#39;s token endpoint URL.
Default to https://gitlab.com/oauth/token.
spotifyAuth
Spotify OAuth2 provider settings.
├─
Optional
enabled
Boolean
Enable the OAuth2 provider.
├─
Required
clientId
String
The provider&#39;s app client id (required if enabled).
└─
Required
clientSecret
String
The provider&#39;s app client secret (required if enabled).
Body parameters could be sent as JSON or
multipart/form-data.
Query parameters
Param
Type
Description
fields
String
Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
-:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
Responses
"meta": {
"appName": "Acme",
"appUrl": "http://127.0.0.1:8090",
"hideControls": false,
"senderName": "Support",
"verificationTemplate": { ... },
"resetPasswordTemplate": { ... },
"confirmEmailChangeTemplate": { ...}"
"logs": {
"maxDays": 7
"backups": {
"cron": "0 0 * * *",
"cronMaxKeep": 1,
"s3": {
"enabled": false,
"bucket": "",
"region": "",
"endpoint": "",
"accessKey": "",
"secret": "",
"forcePathStyle": false
"smtp": {
"enabled": false,
"port": 587,
"username": "",
"password": "",
"tls": true,
"authMethod": "",
"localName": ""
"s3": {
"enabled": false,
"bucket": "",
"region": "",
"endpoint": "",
"accessKey": "",
"secret": "",
"forcePathStyle": false
"adminAuthToken": {
"secret": "******",
"duration": 1209600
"adminPasswordResetToken": {
"secret": "******",
"duration": 1800
"recordAuthToken": {
"secret": "******",
"duration": 1209600
"recordPasswordResetToken": {
"secret": "******",
"duration": 1800
"recordEmailChangeToken": {
"secret": "******",
"duration": 1800
"recordVerificationToken": {
"secret": "******",
"duration": 604800
"googleAuth": {
"enabled": true,
"clientId": "demo",
"clientSecret": "******"
"facebookAuth": {
"enabled": false,
"githubAuth": {
"enabled": true,
"clientId": "demo",
"clientSecret": "******"
"gitlabAuth": {
"enabled": true,
"clientId": "demo",
"clientSecret": "******"
"discordAuth": {
"enabled": true,
"clientId": "demo",
"clientSecret": "******"
"twitterAuth": {
"enabled": true,
"clientId": "demo",
"clientSecret": "******"
"microsoftAuth": {
"enabled": true,
"clientId": "demo",
"clientSecret": "******"
"spotifyAuth": {
"enabled": true,
"clientId": "demo",
"clientSecret": "******"
"code": 400,
"message": "An error occurred while submitting the form.",
"data": {
"meta": {
"appName": {
"code": "validation_required",
"message": "Missing required value."
"code": 401,
"message": "The request requires admin authorization token to be set.",
"data": {}
"code": 403,
"message": "You are not allowed to perform this request.",
"data": {}
Test S3 storage connection
Performs a S3 storage connection test.
Only admins can perform this action.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
await pb.settings.testS3("backups");
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
await pb.settings.testS3("backups");
POST
/api/settings/test/s3
Requires `Authorization: TOKEN`
Body Parameters
Param
Type
Description
Required
filesystem
String
The storage filesystem to test (`storage` or `backups`).
Body parameters could be sent as JSON or
multipart/form-data.
Responses
`null`
"code": 400,
"data": {}
"code": 401,
"message": "The request requires admin authorization token to be set.",
"data": {}
Send test email
Sends a test user email.
Only admins can perform this action.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
POST
/api/settings/test/email
Requires `Authorization: TOKEN`
Body Parameters
Param
Type
Description
Required
email
String
The receiver of the test email.
Required
template
String
The test email template to send:
`verification`,
`password-reset` or
`email-change`.
Body parameters could be sent as JSON or
multipart/form-data.
Responses
`null`
"code": 400,
"message": "Failed to send the test email.",
"data": {
"email": {
"code": "validation_required",
"message": "Missing required value."
"code": 401,
"message": "The request requires admin authorization token to be set.",
"data": {}
Generate Apple client secret
Generates a new Apple OAuth2 client secret key.
Only admins can perform this action.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
await pb.settings.generateAppleClientSecret(clientId, teamId, keyId, privateKey, duration)
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
await pb.settings.generateAppleClientSecret(clientId, teamId, keyId, privateKey, duration)
POST
/api/settings/apple/generate-client-secret
Requires `Authorization: TOKEN`
Body Parameters
Param
Type
Description
Required
clientId
String
The identifier of your app (aka. Service ID).
Required
teamId
String
your name in the Apple Developer site).
Required
keyId
String
10-character key identifier generated for the &quot;Sign in with Apple&quot; private key associated
with your developer account.
Required
privateKey
String
PrivateKey is the private key associated to your app.
Required
duration
Number
Duration specifies how long the generated JWT token should be considered valid.
The specified value must be in seconds and max 15777000 (~6months).
Body parameters could be sent as JSON or
multipart/form-data.
Responses
"secret": "..."
"code": 400,
"data": {}
"code": 401,
"message": "The request requires admin authorization token to be set.",
"data": {}

## 17.Web APIs reference - API Logs
POST /api/collections/users/auth-with-password"
page(Number):The page (aka. offset) of the paginated list (default to 1).
perPage(Number):The max returned logs per page (default to 30).
sort(String):Specify the ORDER BY fields.
Add - / + (default) in front of the attribute for DESC /
ASC order, eg.:
// DESC by the insertion rowid and ASC by level
?sort=-rowid,level
Supported log sort fields:
@random, rowid, id, created,
updated, level, message and any
data.* attribute.
filter(String):Filter expression to filter/search the returned logs list, eg.:
?filter=(data.url~'test.com' && level>0)
Supported log filter fields:
id, created, updated,
level, message and any data.* attribute.
The syntax basically follows the format
OPERAND
OPERATOR
OPERAND, where:
OPERAND - could be any of the above field literal, string (single or double
quoted), number, null, true, false
OPERATOR - is one of:
Equal
!=
NOT equal
Greater than
>=
Greater than or equal
Less than
<=
Less than or equal
Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for wildcard
match)
!~
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for
wildcard match)
?=
Any/At least one of
Equal
?!=
Any/At least one of
NOT equal
?>
Any/At least one of
Greater than
?>=
Any/At least one of
Greater than or equal
?<
Any/At least one of
Less than
?<=
Any/At least one of
Less than or equal
?~
Any/At least one of
Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for wildcard
match)
?!~
Any/At least one of
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for
wildcard match)
To group and combine several expressions you could use parenthesis
(...), && (AND) and || (OR) tokens.
Single line comments are also supported: // Example comment.
fields(String):Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
id(String):ID of the log to view.
fields(String):Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
filter(String):Filter expression to filter/search the logs, eg.:
?filter=(data.url~'test.com' && level>0)
Supported log filter fields:
rowid, id, created,
updated, level, message and any
data.* attribute.
The syntax basically follows the format
OPERAND
OPERATOR
OPERAND, where:
OPERAND - could be any of the above field literal, string (single or double
quoted), number, null, true, false
OPERATOR - is one of:
Equal
!=
NOT equal
Greater than
>=
Greater than or equal
Less than
<=
Less than or equal
Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for wildcard
match)
!~
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for
wildcard match)
?=
Any/At least one of
Equal
?!=
Any/At least one of
NOT equal
?>
Any/At least one of
Greater than
?>=
Any/At least one of
Greater than or equal
?<
Any/At least one of
Less than
?<=
Any/At least one of
Less than or equal
?~
Any/At least one of
Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for wildcard
match)
?!~
Any/At least one of
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for
wildcard match)
To group and combine several expressions you could use parenthesis
(...), && (AND) and || (OR) tokens.
Single line comments are also supported: // Example comment.
fields(String):Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
# Web APIs reference - API Logs
`POST /api/collections/users/auth-with-password"`
- **page** (Number): The page (aka. offset) of the paginated list (default to 1).
- **perPage** (Number): The max returned logs per page (default to 30).
- **sort** (String): Specify the ORDER BY fields.
Add - / + (default) in front of the attribute for DESC /
ASC order, eg.:
// DESC by the insertion rowid and ASC by level
?sort=-rowid,level
Supported log sort fields:
@random, rowid, id, created,
updated, level, message and any
data.* attribute.
- **filter** (String): Filter expression to filter/search the returned logs list, eg.:
?filter=(data.url~'test.com' && level>0)
Supported log filter fields:
id, created, updated,
level, message and any data.* attribute.
The syntax basically follows the format
OPERAND
OPERATOR
OPERAND, where:
OPERAND - could be any of the above field literal, string (single or double
quoted), number, null, true, false
OPERATOR - is one of:
Equal
NOT equal
Greater than
Greater than or equal
Less than
Less than or equal
Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for wildcard
match)
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for
wildcard match)
Any/At least one of
Equal
Any/At least one of
NOT equal
Any/At least one of
Greater than
Any/At least one of
Greater than or equal
Any/At least one of
Less than
Any/At least one of
Less than or equal
Any/At least one of
Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for wildcard
match)
?!~
Any/At least one of
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for
wildcard match)
To group and combine several expressions you could use parenthesis
(...), && (AND) and || (OR) tokens.
Single line comments are also supported: // Example comment.
- **fields** (String): Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
- **id** (String): ID of the log to view.
- **fields** (String): Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
- **filter** (String): Filter expression to filter/search the logs, eg.:
?filter=(data.url~'test.com' && level>0)
Supported log filter fields:
rowid, id, created,
updated, level, message and any
data.* attribute.
The syntax basically follows the format
OPERAND
OPERATOR
OPERAND, where:
OPERAND - could be any of the above field literal, string (single or double
quoted), number, null, true, false
OPERATOR - is one of:
Equal
NOT equal
Greater than
Greater than or equal
Less than
Less than or equal
Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for wildcard
match)
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for
wildcard match)
Any/At least one of
Equal
Any/At least one of
NOT equal
Any/At least one of
Greater than
Any/At least one of
Greater than or equal
Any/At least one of
Less than
Any/At least one of
Less than or equal
Any/At least one of
Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for wildcard
match)
?!~
Any/At least one of
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for
wildcard match)
To group and combine several expressions you could use parenthesis
(...), && (AND) and || (OR) tokens.
Single line comments are also supported: // Example comment.
- **fields** (String): Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
List logs
Returns a paginated logs list.
Only admins can perform this action.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const pageResult = await pb.logs.getList(1, 20, {
filter: 'data.status >= 400'
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
final pageResult = await pb.logs.getList(
page: 1,
perPage: 20,
filter: 'data.status >= 400',
GET
/api/logs
Requires `Authorization: TOKEN`
Query parameters
Param
Type
Description
page
Number
The page (aka. offset) of the paginated list (default to 1).
perPage
Number
The max returned logs per page (default to 30).
sort
String
Specify the ORDER BY fields.
Add - / + (default) in front of the attribute for DESC /
ASC order, eg.:
// DESC by the insertion rowid and ASC by level
?sort=-rowid,level
Supported log sort fields:
@random, rowid, id, created,
updated, level, message and any
data.* attribute.
filter
String
Filter expression to filter/search the returned logs list, eg.:
`?filter=(data.url~'test.com' &amp;&amp; level>0)`
Supported log filter fields:
id, created, updated,
level, message and any data.* attribute.
The syntax basically follows the format
OPERAND
OPERATOR
OPERAND, where:
-OPERAND - could be any of the above field literal, string (single or double
quoted), number, null, true, false
-OPERATOR - is one of:
Equal
NOT equal
Greater than
Greater than or equal
-&lt;
Less than
-&lt;=
Less than or equal
Like/Contains (if not specified auto wraps the right string OPERAND in a &quot;%&quot; for wildcard
match)
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a &quot;%&quot; for
wildcard match)
Any/At least one of
Equal
Any/At least one of
NOT equal
Any/At least one of
Greater than
Any/At least one of
Greater than or equal
-?&lt;
Any/At least one of
Less than
-?&lt;=
Any/At least one of
Less than or equal
Any/At least one of
Like/Contains (if not specified auto wraps the right string OPERAND in a &quot;%&quot; for wildcard
match)
-?!~
Any/At least one of
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a &quot;%&quot; for
wildcard match)
To group and combine several expressions you could use parenthesis
(...), &amp;&amp; (AND) and || (OR) tokens.
Single line comments are also supported: // Example comment.
fields
String
Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
-:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
Responses
"page": 1,
"perPage": 20,
"totalItems": 2,
"items": [
"id": "9ajmzgd99r039k9",
"created": "2023-12-12 04:41:59.973Z",
"updated": "2023-12-12 04:41:59.973Z",
"data": {
"auth": "authRecord",
"execTime": 364.961387,
"method": "POST",
"remoteIp": "127.0.0.1",
"status": 200,
"type": "request",
"url": "/old/api/collections/users/auth-with-password",
"userAgent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36",
"userIp": "127.0.0.1"
"message": "POST /api/collections/users/auth-with-password",
"level": 0
"id": "26apis4s3sm9yqm",
"created": "2023-12-12 04:27:21.583Z",
"updated": "2023-12-12 04:27:21.583Z",
"data": {
"auth": "authRecord",
"execTime": 403.664712,
"method": "POST",
"remoteIp": "127.0.0.1",
"status": 200,
"type": "request",
"url": "/old/api/collections/users/auth-with-password?expand=rel&amp;fields=*%2Crecord.*%2Crecord.expand.rel.id",
"userAgent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36",
"userIp": "127.0.0.1"
"message": "POST /api/collections/users/auth-with-password?expand=rel&amp;fields=*%2Crecord.*%2Crecord.expand.rel.id",
"level": 0
"code": 400,
"message": "Something went wrong while processing your request. Invalid filter.",
"data": {}
"code": 401,
"message": "The request requires admin authorization token to be set.",
"data": {}
"code": 403,
"message": "You are not allowed to perform this request.",
"data": {}
View log
Returns a single log by its ID.
Only admins can perform this action.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const log = await pb.logs.getOne('LOG_ID');
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
final log = await pb.logs.getOne('LOG_ID');
GET
/api/logs/`id`
Requires `Authorization: TOKEN`
Path parameters
Param
Type
Description
String
ID of the log to view.
Query parameters
Param
Type
Description
fields
String
Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
-:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
Responses
"id": "twjnabervu5log8",
"created": "2023-12-10 19:56:29.556Z",
"updated": "2023-12-10 19:56:29.556Z",
"data": {
"auth": "guest",
"execTime": 0.66452,
"method": "GET",
"remoteIp": "127.0.0.1",
"status": 200,
"type": "request",
"url": "/old/api/collections/users/auth-methods",
"userAgent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36",
"userIp": "127.0.0.1"
"message": "GET /api/collections/users/auth-methods",
"level": 0
"code": 401,
"message": "The request requires admin authorization token to be set.",
"data": {}
"code": 403,
"message": "You are not allowed to perform this request.",
"data": {}
"code": 404,
"message": "The requested resource wasn't found.",
"data": {}
Logs statistics
Returns hourly aggregated logs statistics.
Only admins can perform this action.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const stats = await pb.logs.getStats({
filter: 'data.status >= 400'
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
final stats = await pb.logs.getStats(
filter: 'data.status >= 400'
GET
/api/logs/stats
Requires `Authorization: TOKEN`
Query parameters
Param
Type
Description
filter
String
Filter expression to filter/search the logs, eg.:
`?filter=(data.url~'test.com' &amp;&amp; level>0)`
Supported log filter fields:
rowid, id, created,
updated, level, message and any
data.* attribute.
The syntax basically follows the format
OPERAND
OPERATOR
OPERAND, where:
-OPERAND - could be any of the above field literal, string (single or double
quoted), number, null, true, false
-OPERATOR - is one of:
Equal
NOT equal
Greater than
Greater than or equal
-&lt;
Less than
-&lt;=
Less than or equal
Like/Contains (if not specified auto wraps the right string OPERAND in a &quot;%&quot; for wildcard
match)
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a &quot;%&quot; for
wildcard match)
Any/At least one of
Equal
Any/At least one of
NOT equal
Any/At least one of
Greater than
Any/At least one of
Greater than or equal
-?&lt;
Any/At least one of
Less than
-?&lt;=
Any/At least one of
Less than or equal
Any/At least one of
Like/Contains (if not specified auto wraps the right string OPERAND in a &quot;%&quot; for wildcard
match)
-?!~
Any/At least one of
NOT Like/Contains (if not specified auto wraps the right string OPERAND in a &quot;%&quot; for
wildcard match)
To group and combine several expressions you could use parenthesis
(...), &amp;&amp; (AND) and || (OR) tokens.
Single line comments are also supported: // Example comment.
fields
String
Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
-:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
Responses
"total": 4,
"date": "2022-06-01 19:00:00.000"
"total": 1,
"date": "2022-06-02 12:00:00.000"
"total": 8,
"date": "2022-06-02 13:00:00.000"
"code": 400,
"message": "Something went wrong while processing your request. Invalid filter.",
"data": {}
"code": 401,
"message": "The request requires admin authorization token to be set.",
"data": {}
"code": 403,
"message": "You are not allowed to perform this request.",
"data": {}

## 18.Web APIs reference - API Backups
fields(String):Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
Optional
name(String):The base name of the backup file to create.
Must be in the format [a-z0-9_-].zip
If not set, it will be auto generated.
Required
file(File):The zip archive to upload.
key(String):The key of the backup file to delete.
key(String):The key of the backup file to restore.
key(String):The key of the backup file to download.
token(String):Admin file token for granting access to the
backup file.
# Web APIs reference - API Backups
- **fields** (String): Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
- **Optional
name** (String): The base name of the backup file to create.
If not set, it will be auto generated.
- **Required
file** (File): The zip archive to upload.
- **key** (String): The key of the backup file to delete.
- **key** (String): The key of the backup file to restore.
- **token** (String): Admin file token for granting access to the
backup file.
List backups
Returns list with all available backup files.
Only admins can perform this action.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const backups = await pb.backups.getFullList();
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
final backups = await pb.backups.getFullList();
GET
/api/backups
Query parameters
Param
Type
Description
fields
String
Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
-:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
Responses
"modified": "2023-05-19 16:25:57.542Z",
"size": 251316185
"modified": "2023-05-18 16:25:57.542Z",
"size": 251314010
"code": 400,
"message": "Failed to load backups filesystem.",
"data": {}
"code": 401,
"message": "The request requires admin authorization token to be set.",
"data": {}
"code": 403,
"message": "Only admins can perform this action.",
"data": {}
Create backup
Creates a new app data backup.
This action will return an error if there is another backup/restore operation already in progress.
Only admins can perform this action.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
POST
/api/backups
Requires `Authorization: TOKEN`
Body Parameters
Param
Type
Description
Optional
name
String
The base name of the backup file to create.
If not set, it will be auto generated.
Body parameters could be sent as JSON or
multipart/form-data.
Responses
`null`
"code": 400,
"message": "Try again later - another backup/restore process has already been started.",
"data": {}
"code": 401,
"message": "The request requires admin authorization token to be set.",
"data": {}
"code": 403,
"message": "You are not allowed to perform this request.",
"data": {}
Upload backup
Uploads an existing backup zip file.
Only admins can perform this action.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
await pb.backups.upload({ file: new Blob([...]) });
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
await pb.backups.upload(http.MultipartFile.fromBytes('file', ...));
POST
/api/backups/upload
Requires `Authorization: TOKEN`
Body Parameters
Param
Type
Description
Required
file
File
The zip archive to upload.
Uploading files is supported only via multipart/form-data.
Responses
`null`
"code": 400,
"message": "Something went wrong while processing your request.",
"data": {
"file": {
"code": "validation_invalid_mime_type",
"message": "\"test_backup.txt\" mime type must be one of: application/zip."
"code": 401,
"message": "The request requires admin authorization token to be set.",
"data": {}
"code": 403,
"message": "You are not allowed to perform this request.",
"data": {}
Delete backup
Deletes a single backup by its name.
This action will return an error if the backup to delete is still being generated or part of a
restore operation.
Only admins can perform this action.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
DELETE
/api/backups/`key`
Requires `Authorization: TOKEN`
Path parameters
Param
Type
Description
key
String
The key of the backup file to delete.
Responses
`null`
"code": 400,
"message": "Try again later - another backup/restore process has already been started.",
"data": {}
"code": 401,
"message": "The request requires admin authorization token to be set.",
"data": {}
"code": 403,
"message": "You are not allowed to perform this request.",
"data": {}
Restore backup
Restore a single backup by its name and restarts the current running PocketBase process.
This action will return an error if there is another backup/restore operation already in progress.
Only admins can perform this action.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
POST
/api/backups/`key`/restore
Requires `Authorization: TOKEN`
Path parameters
Param
Type
Description
key
String
The key of the backup file to restore.
Responses
`null`
"code": 400,
"message": "Try again later - another backup/restore process has already been started.",
"data": {}
"code": 401,
"message": "The request requires admin authorization token to be set.",
"data": {}
"code": 403,
"message": "You are not allowed to perform this request.",
"data": {}
Only admins can perform this action.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const token = await pb.files.getToken();
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
final token = await pb.files.getToken();
GET
/api/backups/`key`
Path parameters
Param
Type
Description
key
String
Query parameters
Param
Type
Description
token
String
Admin file token for granting access to the
backup file.
Responses
`[file resource]`
"code": 400,
"message": "Filesystem initialization failure.",
"data": {}
"code": 404,
"message": "The requested resource wasn't found.",
"data": {}

## 19.Web APIs reference - API Health
fields(String):Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
# Web APIs reference - API Health
- **fields** (String): Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
Health check
Returns the health status of the server.
GET/HEAD
/api/health
Query parameters
Param
Type
Description
fields
String
Comma separated string of the fields to return in the JSON response
(by default returns all fields). Ex.:
?fields=*,expand.relField.name
* targets all keys from the specific depth level.
In addition, the following field modifiers are also supported:
-:excerpt(maxLength, withEllipsis?)
Returns a short plain text version of the field string value.
Ex.:
?fields=*,description:excerpt(200,true)
Responses
"code": 200,
"message": "API is healthy.",
"data": {
"canBackup": false

## 20.Extend with JavaScript - Overview
# Extend with JavaScript - Overview
### JavaScript engine
The prebuilt PocketBase v0.17+ executable comes with embedded ES5 JavaScript engine (goja) which enables you to write custom server-side code using plain JavaScript.
You can start by creating *.pb.js file(s) inside a pb_hooks
// pb_hooks/main.pb.js
routerAdd("GET", "/hello/:name", (c) => {
let name = c.pathParam("name")
return c.json(200, { "message": "Hello " + name })
onModelAfterUpdate((e) => {
console.log("user updated...", e.model.get("email"))
}, "users")
For convenience, when making changes to the files inside pb_hooks, the process will
automatically restart/reload itself (currently supported only on UNIX based platforms). The
*.pb.js files are loaded per their filename sort order.
For most parts, the JavaScript APIs are derived from Go with 2 main differences:
-Go exported method and field names are converted to camelCase, for example:
app.Dao().FindRecordById(&quot;example&quot;, &quot;RECORD_ID&quot;) becomes
$app.dao().findRecordById(&quot;example&quot;, &quot;RECORD_ID&quot;).
-Errors are thrown as regular JavaScript exceptions and not returned as Go values.
##### Global objects
Below is a list with some of the commonly used global objects that are accessible from everywhere:
-__hooks
- The absolute path to the app pb_hooks directory.
-$app - The current running PocketBase application instance.
-$apis.* - API routing helpers and middlewares.
-$os.* - OS level primitives (deleting directories, executing shell commands, etc.).
-$security.* - Low level helpers for creating and parsing JWTs, random string generation, AES encryption, etc.
-And many more - for all exposed APIs, please refer to the
JSVM reference docs.
### TypeScript declarations and code completion
While you can&#39;t use directly TypeScript (without transpiling it to JS on your own), PocketBase
comes with builtin ambient TypeScript declarations that can help providing information
and documentation about the available global variables, methods and arguments, code completion, etc. as
long as your editor has TypeScript LSP support
(most editors either have it builtin or available as plugin).
The types declarations are stored in
pb_data/types.d.ts file. You can point to those declarations using the
reference tripple-slash directive
at the top of your JS file:
/// &lt;reference path="../pb_data/types.d.ts" />
onAfterBootstrap((e) => {
console.log("App initialized!")
If after referencing the types your editor still doesn&#39;t perform linting, then you can try to rename your
file to have .pb.ts extension.
### Caveats and limitations
##### Handlers scope
Each handler function (hook, route, middleware, etc.) is
serialized and executed in its own isolated context as a separate &quot;program&quot;. This means
that you don&#39;t have access to custom variables and functions declared outside of the handler scope. For
example, the below code will fail:
const name = "test"
onAfterBootstrap((e) => {
console.log(name) // &lt;-- name will be undefined inside the handler
The above serialization and isolation context is also the reason why error stack trace line numbers may
not be accurate.
One possible workaround for sharing/reusing code across different handlers could be to move and export the
reusable code portion as local module and load it with require() inside the handler but keep in
mind that the loaded modules use a shared registry and mutations should be avoided when possible to prevent
concurrency issues:
onAfterBootstrap((e) => {
const config = require(`${__hooks}/config.js`)
console.log(config.name)
##### Relative paths
Relative file paths are relative to the current working directory (CWD) and not to the
pb_hooks.
To get an absolute path to the pb_hooks directory you can use the global
__hooks variable.
##### Loading modules
Please note that the embedded JavaScript engine is not a Node.js or browser environment, meaning
that modules that relies on APIs like window, fs,
fetch, buffer or any other runtime specific API not part of the ES5 spec may not
work!
You can load modules either by specifying their local filesystem path or by using their name, which will
automatically search in:
-the current working directory (affects also relative paths)
-any node_modules directory
-any parent node_modules directory
Currently only CommonJS (CJS) modules are supported and can be loaded with
const x = require(...).
ECMAScript modules (ESM) can be loaded by first precompiling and transforming your dependencies with a bundler
like
rollup,
webpack,
browserify, etc.
A common usage of local modules is for loading shared helpers or configuration parameters, for example:
// pb_hooks/utils.js
module.exports = {
hello: (name) => {
console.log("Hello " + name)
// pb_hooks/main.pb.js
onAfterBootstrap((e) => {
const utils = require(`${__hooks}/utils.js`)
utils.hello("world")
Loaded modules use a shared registry and mutations should be avoided when possible to prevent
concurrency issues.
##### Performance
The prebuilt executable comes with a prewarmed pool of 25 JS runtimes, which helps
maintaining the handlers execution times on par with the Go equivalent code (see
benchmarks). You can adjust the pool size manually with the --hooksPool=100 flag (increasing the pool size may improve the performance in high concurrent scenarios but also will
increase the memory usage).
Note that the handlers performance may degrade if you have heavy computational tasks in pure JavaScript
(eg. encryption, random generators, etc.). For such cases prefer using the exposed Go bindings
(eg. $security.randomString(10)).
##### Engine limitations
We inherit some of the limitations and caveats of the embedded JavaScript engine
(goja):
-Has most of ES6 functionality already implemented but it is not fully spec compliant yet.
-No concurrent execution inside a single handler (aka. no setTimeout/setInterval).
-Wrapped Go structural types (such as maps, slices) comes with some peculiarities and do not behave the
exact same way as native ECMAScript values (for more details see
goja ToValue).
-In relation to the above, DB json field values require the use of get() and
set() helpers (this may change in the future).

## 21.Extend with JavaScript - Event hooks
# Extend with JavaScript - Event hooks
You can extend the default PocketBase behavior with custom server-side code using the exposed JavaScript
app level hooks.
Throwing an error or returning false inside a hook handler function stops the hook execution chain.
### App hooks
onBeforeBootstrap
`onBeforeBootstrap` hook is triggered before initializing the main
application resources (eg. before db open and initial settings load).
onBeforeBootstrap((e) => {
console.log(e.app)
onAfterBootstrap
`onAfterBootstrap` hook is triggered after initializing the main
application resources (eg. after db open and initial settings load).
onAfterBootstrap((e) => {
console.log(e.app)
onBeforeApiError
`onBeforeApiError` hook is triggered right before sending an error API
response to the client, allowing you to further modify the error data
or to return a completely different API response.
onBeforeApiError((e) => {
console.log(e.httpContext)
console.log(e.error)
onAfterApiError
`onAfterApiError` hook is triggered right after sending an error API
response to the client.
It could be used for example to log the final API error in external services.
onAfterApiError((e) => {
console.log(e.httpContext)
console.log(e.error)
onTerminate
`onTerminate` hook is triggered when the app is in the process
of being terminated (eg. on `SIGTERM` signal).
Note that the app could be terminated abruptly without awaiting the hook completion.
onTerminate((e) => {
console.log("terminating...")
### DB hooks
onModelBeforeCreate
onModelBeforeCreate hook is triggered before inserting a new
model in the DB, allowing you to modify or validate the stored data.
If the optional "tags" list (table names and/or the Collection id for Record models)
is specified, then all event handlers registered via the created hook
will be triggered and called only if their event data origin matches the tags.
// fires for every db model
onModelBeforeCreate((e) => {
console.log(e.model.tableName())
console.log(e.model.id)
// fires only for "users" and "members"
onModelBeforeCreate((e) => {
console.log(e.model.tableName())
console.log(e.model.id)
}, "users", "members")
onModelAfterCreate
onModelAfterCreate hook is triggered after successfully
inserting a new model in the DB.
If the optional "tags" list (table names and/or the Collection id for Record models)
is specified, then all event handlers registered via the created hook
will be triggered and called only if their event data origin matches the tags.
// fires for every db model
onModelAfterCreate((e) => {
console.log(e.model.tableName())
console.log(e.model.id)
// fires only for "users" and "members"
onModelAfterCreate((e) => {
console.log(e.model.tableName())
console.log(e.model.id)
}, "users", "members")
onModelBeforeUpdate
onModelBeforeUpdate hook is triggered before updating existing
model in the DB, allowing you to modify or validate the stored data.
If the optional "tags" list (table names and/or the Collection id for Record models)
is specified, then all event handlers registered via the created hook
will be triggered and called only if their event data origin matches the tags.
// fires for every db model
onModelBeforeUpdate((e) => {
console.log(e.model.tableName())
console.log(e.model.id)
// fires only for "users" and "members"
onModelBeforeUpdate((e) => {
console.log(e.model.tableName())
console.log(e.model.id)
}, "users", "members")
onModelAfterUpdate
onModelAfterUpdate hook is triggered after successfully updating
existing model in the DB.
If the optional "tags" list (table names and/or the Collection id for Record models)
is specified, then all event handlers registered via the created hook
will be triggered and called only if their event data origin matches the tags.
// fires for every db model
onModelAfterUpdate((e) => {
console.log(e.model.tableName())
console.log(e.model.id)
// fires only for "users" and "members"
onModelAfterUpdate((e) => {
console.log(e.model.tableName())
console.log(e.model.id)
}, "users", "members")
onModelBeforeDelete
onModelBeforeDelete hook is triggered before deleting an
existing model from the DB.
If the optional "tags" list (table names and/or the Collection id for Record models)
is specified, then all event handlers registered via the created hook
will be triggered and called only if their event data origin matches the tags.
// fires for every db model
onModelBeforeDelete((e) => {
console.log(e.model.tableName())
console.log(e.model.id)
// fires only for "users" and "members"
onModelBeforeDelete((e) => {
console.log(e.model.tableName())
console.log(e.model.id)
}, "users", "members")
onModelAfterDelete
onModelAfterDelete hook is triggered after successfully
deleting an existing model from the DB.
If the optional "tags" list (table names and/or the Collection id for Record models)
is specified, then all event handlers registered via the created hook
will be triggered and called only if their event data origin matches the tags.
// fires for every db model
onModelAfterDelete((e) => {
console.log(e.model.tableName())
console.log(e.model.id)
// fires only for "users" and "members"
onModelAfterDelete((e) => {
console.log(e.model.tableName())
console.log(e.model.id)
}, "users", "members")
### Mailer hooks
onMailerBeforeAdminResetPasswordSend
`onMailerBeforeAdminResetPasswordSend` hook is triggered right
before sending a password reset email to an admin, allowing you
to inspect and customize the email message that is being sent.
onMailerBeforeAdminResetPasswordSend((e) => {
console.log(e.mailClient)
console.log(e.message)
console.log(e.admin)
console.log(e.meta)
// change the mail subject
e.message.subject = "new subject"
onMailerAfterAdminResetPasswordSend
`onMailerAfterAdminResetPasswordSend` hook is triggered after
admin password reset email was successfully sent.
onMailerAfterAdminResetPasswordSend((e) => {
console.log(e.mailClient)
console.log(e.message)
console.log(e.admin)
console.log(e.meta)
onMailerBeforeRecordResetPasswordSend
onMailerBeforeRecordResetPasswordSend hook is triggered right
before sending a password reset email to an auth record, allowing
you to inspect and customize the email message that is being sent.
If the optional "tags" list (Collection ids or names) is specified,
then all event handlers registered via the created hook will be
triggered and called only if their event data origin matches the tags.
onMailerBeforeRecordResetPasswordSend((e) => {
console.log(e.mailClient)
console.log(e.message)
console.log(e.record)
console.log(e.meta)
// change the mail subject
e.message.subject = "new subject"
onMailerAfterRecordResetPasswordSend
onMailerAfterRecordResetPasswordSend hook is triggered
after an auth record password reset email was successfully sent.
If the optional "tags" list (Collection ids or names) is specified,
then all event handlers registered via the created hook will be
triggered and called only if their event data origin matches the tags.
onMailerAfterRecordResetPasswordSend((e) => {
console.log(e.mailClient)
console.log(e.message)
console.log(e.record)
console.log(e.meta)
onMailerBeforeRecordVerificationSend
onMailerBeforeRecordVerificationSend hook is triggered right
before sending a verification email to an auth record, allowing
you to inspect and customize the email message that is being sent.
If the optional "tags" list (Collection ids or names) is specified,
then all event handlers registered via the created hook will be
triggered and called only if their event data origin matches the tags.
onMailerBeforeRecordVerificationSend((e) => {
console.log(e.mailClient)
console.log(e.message)
console.log(e.record)
console.log(e.meta)
// change the mail subject
e.message.subject = "new subject"
onMailerAfterRecordVerificationSend
onMailerAfterRecordVerificationSend hook is triggered
after a verification email was successfully sent to an auth record.
If the optional "tags" list (Collection ids or names) is specified,
then all event handlers registered via the created hook will be
triggered and called only if their event data origin matches the tags.
onMailerAfterRecordVerificationSend((e) => {
console.log(e.mailClient)
console.log(e.message)
console.log(e.record)
console.log(e.meta)
onMailerBeforeRecordChangeEmailSend
onMailerBeforeRecordChangeEmailSend hook is triggered right before
sending a confirmation new address email to an auth record, allowing
you to inspect and customize the email message that is being sent.
If the optional "tags" list (Collection ids or names) is specified,
then all event handlers registered via the created hook will be
triggered and called only if their event data origin matches the tags.
onMailerBeforeRecordChangeEmailSend((e) => {
console.log(e.mailClient)
console.log(e.message)
console.log(e.record)
console.log(e.meta)
// change the mail subject
e.message.subject = "new subject"
onMailerAfterRecordChangeEmailSend
onMailerAfterRecordChangeEmailSend hook is triggered
after a verification email was successfully sent to an auth record.
If the optional "tags" list (Collection ids or names) is specified,
then all event handlers registered via the created hook will be
triggered and called only if their event data origin matches the tags.
onMailerAfterRecordChangeEmailSend((e) => {
console.log(e.mailClient)
console.log(e.message)
console.log(e.record)
console.log(e.meta)
### Record CRUD API hooks
onRecordsListRequest
onRecordsListRequest hook is triggered on each API Records list request.
Could be used to validate or modify the response before returning it to the client.
If the optional "tags" list (Collection ids or names) is specified,
then all event handlers registered via the created hook will be
triggered and called only if their event data origin matches the tags.
// fires for every collection
onRecordsListRequest((e) => {
console.log(e.httpContext)
console.log(e.result)
// fires only for "users" and "articles" collections
onRecordsListRequest((e) => {
console.log(e.httpContext)
console.log(e.result)
}, "users", "articles")
onRecordViewRequest
onRecordViewRequest hook is triggered on each API Record view request.
Could be used to validate or modify the response before returning it to the client.
If the optional "tags" list (Collection ids or names) is specified,
then all event handlers registered via the created hook will be
triggered and called only if their event data origin matches the tags.
// fires for every collection
onRecordViewRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
// fires only for "users" and "articles" collections
onRecordViewRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
}, "users", "articles")
onRecordBeforeCreateRequest
onRecordBeforeCreateRequest hook is triggered before each API Record
create request (after request data load and before model persistence).
Could be used to additionally validate the request data or implement
completely different persistence behavior.
If the optional "tags" list (Collection ids or names) is specified,
then all event handlers registered via the created hook will be
triggered and called only if their event data origin matches the tags.
// fires for every collection
onRecordBeforeCreateRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
console.log(e.uploadedFiles)
// fires only for "users" and "articles" collections
onRecordBeforeCreateRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
console.log(e.uploadedFiles)
}, "users", "articles")
onRecordAfterCreateRequest
onRecordAfterCreateRequest hook is triggered after each
successful API Record create request.
Could be used to additionally validate the request data or implement
completely different persistence behavior.
If the optional "tags" list (Collection ids or names) is specified,
then all event handlers registered via the created hook will be
triggered and called only if their event data origin matches the tags.
// fires for every collection
onRecordAfterCreateRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
console.log(e.uploadedFiles)
// fires only for "users" and "articles" collections
onRecordAfterCreateRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
console.log(e.uploadedFiles)
}, "users", "articles")
onRecordBeforeUpdateRequest
onRecordBeforeUpdateRequest hook is triggered before each API Record
update request (after request data load and before model persistence).
Could be used to additionally validate the request data or implement
completely different persistence behavior.
If the optional "tags" list (Collection ids or names) is specified,
then all event handlers registered via the created hook will be
triggered and called only if their event data origin matches the tags.
// fires for every collection
onRecordBeforeUpdateRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
console.log(e.uploadedFiles)
// fires only for "users" and "articles" collections
onRecordBeforeUpdateRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
console.log(e.uploadedFiles)
}, "users", "articles")
onRecordAfterUpdateRequest
onRecordAfterUpdateRequest hook is triggered
after each successful API Record update request.
If the optional "tags" list (Collection ids or names) is specified,
then all event handlers registered via the created hook will be
triggered and called only if their event data origin matches the tags.
// fires for every collection
onRecordAfterUpdateRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
console.log(e.uploadedFiles)
// fires only for "users" and "articles" collections
onRecordAfterUpdateRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
console.log(e.uploadedFiles)
}, "users", "articles")
onRecordBeforeDeleteRequest
onRecordBeforeDeleteRequest hook is triggered before each API Record
delete request (after model load and before actual deletion).
Could be used to additionally validate the request data or implement
completely different delete behavior.
If the optional "tags" list (Collection ids or names) is specified,
then all event handlers registered via the created hook will be
triggered and called only if their event data origin matches the tags.
// fires for every collection
onRecordBeforeDeleteRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
// fires only for "users" and "articles" collections
onRecordBeforeDeleteRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
}, "users", "articles")
onRecordAfterDeleteRequest
onRecordAfterDeleteRequest hook is triggered
after each successful API Record delete request.
If the optional "tags" list (Collection ids or names) is specified,
then all event handlers registered via the created hook will be
triggered and called only if their event data origin matches the tags.
// fires for every collection
onRecordAfterDeleteRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
// fires only for "users" and "articles" collections
onRecordAfterDeleteRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
}, "users", "articles")
### Record Auth API hooks
onRecordAuthRequest
onRecordAuthRequest hook is triggered on each successful API
record authentication request (sign-in, token refresh, etc.).
Could be used to additionally validate or modify the authenticated
record data and token.
If the optional "tags" list (Collection ids or names) is specified,
then all event handlers registered via the created hook will be
triggered and called only if their event data origin matches the tags.
// fires for every auth collection
onRecordAuthRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
console.log(e.token)
console.log(e.meta)
// fires only for "users" and "managers" auth collections
onRecordAuthRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
console.log(e.token)
console.log(e.meta)
}, "users", "managers")
onRecordBeforeAuthWithPasswordRequest
onRecordBeforeAuthWithPasswordRequest hook is triggered before each Record
auth with password API request (after request data load and before password validation).
Could be used to implement for example a custom password validation
or to locate a different Record model (by reassigning RecordAuthWithPasswordEvent.Record).
If the optional "tags" list (Collection ids or names) is specified,
then all event handlers registered via the created hook will be
triggered and called only if their event data origin matches the tags.
// fires for every auth collection
onRecordBeforeAuthWithPasswordRequest((e) => {
console.log(e.httpContext)
console.log(e.record) // could be null
console.log(e.identity)
console.log(e.password)
// fires only for "users" and "managers" auth collections
onRecordBeforeAuthWithPasswordRequest((e) => {
console.log(e.httpContext)
console.log(e.record) // could be null
console.log(e.identity)
console.log(e.password)
}, "users", "managers")
onRecordAfterAuthWithPasswordRequest
onRecordAfterAuthWithPasswordRequest hook is triggered after each
successful Record auth with password API request.
If the optional "tags" list (Collection ids or names) is specified,
then all event handlers registered via the created hook will be
triggered and called only if their event data origin matches the tags.
// fires for every auth collection
onRecordAfterAuthWithPasswordRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
console.log(e.identity)
console.log(e.password)
// fires only for "users" and "managers" auth collections
onRecordAfterAuthWithPasswordRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
console.log(e.identity)
console.log(e.password)
}, "users", "managers")
onRecordBeforeAuthWithOAuth2Request
onRecordBeforeAuthWithOAuth2Request hook is triggered before each Record
OAuth2 sign-in/sign-up API request (after token exchange and before external provider linking).
If the RecordAuthWithOAuth2Event.Record is not set,
then the OAuth2 request will try to create a new auth Record.
To assign or link a different existing record model you can
change the RecordAuthWithOAuth2Event.Record field.
If the optional "tags" list (Collection ids or names) is specified,
then all event handlers registered via the created hook will be
triggered and called only if their event data origin matches the tags.
// fires for every auth collection
onRecordBeforeAuthWithOAuth2Request((e) => {
console.log(e.httpContext)
console.log(e.providerName)
console.log(e.providerClient)
console.log(e.record) // could be null
console.log(e.oAuth2User)
console.log(e.isNewRecord)
// fires only for "users" and "managers" auth collections
onRecordBeforeAuthWithOAuth2Request((e) => {
console.log(e.httpContext)
console.log(e.providerName)
console.log(e.providerClient)
console.log(e.record) // could be null
console.log(e.oAuth2User)
console.log(e.isNewRecord)
}, "users", "managers")
onRecordAfterAuthWithOAuth2Request
onRecordAfterAuthWithOAuth2Request hook is triggered
after each successful Record OAuth2 API request.
If the optional "tags" list (Collection ids or names) is specified,
then all event handlers registered via the created hook will be
triggered and called only if their event data origin matches the tags.
// fires for every auth collection
onRecordAfterAuthWithOAuth2Request((e) => {
console.log(e.httpContext)
console.log(e.providerName)
console.log(e.providerClient)
console.log(e.record)
console.log(e.oAuth2User)
console.log(e.isNewRecord)
// fires only for "users" and "managers" auth collections
onRecordAfterAuthWithOAuth2Request((e) => {
console.log(e.httpContext)
console.log(e.providerName)
console.log(e.providerClient)
console.log(e.record)
console.log(e.oAuth2User)
console.log(e.isNewRecord)
}, "users", "managers")
onRecordBeforeAuthRefreshRequest
onRecordBeforeAuthRefreshRequest hook is triggered before each Record
auth refresh API request (right before generating a new auth token).
Could be used to additionally validate the request data or implement
completely different auth refresh behavior.
If the optional "tags" list (Collection ids or names) is specified,
then all event handlers registered via the created hook will be
triggered and called only if their event data origin matches the tags.
// fires for every auth collection
onRecordBeforeAuthRefreshRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
// fires only for "users" and "managers" auth collections
onRecordBeforeAuthRefreshRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
}, "users", "managers")
onRecordAfterAuthRefreshRequest
onRecordAfterAuthRefreshRequest hook is triggered after each
successful auth refresh API request (right after generating a new auth token).
If the optional "tags" list (Collection ids or names) is specified,
then all event handlers registered via the created hook will be
triggered and called only if their event data origin matches the tags.
// fires for every auth collection
onRecordAfterAuthRefreshRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
// fires only for "users" and "managers" auth collections
onRecordAfterAuthRefreshRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
}, "users", "managers")
onRecordListExternalAuthsRequest
onRecordListExternalAuthsRequest hook is triggered on each API record external auths list request.
Could be used to validate or modify the response before returning it to the client.
If the optional "tags" list (Collection ids or names) is specified,
then all event handlers registered via the created hook will be
triggered and called only if their event data origin matches the tags.
// fires for every auth collection
onRecordListExternalAuthsRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
console.log(e.externalAuths)
// fires only for "users" and "managers" auth collections
onRecordListExternalAuthsRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
console.log(e.externalAuths)
}, "users", "managers")
onRecordBeforeUnlinkExternalAuthRequest
onRecordBeforeUnlinkExternalAuthRequest hook is triggered before each API record
external auth unlink request (after models load and before the actual relation deletion).
Could be used to additionally validate the request data or implement
completely different delete behavior.
If the optional "tags" list (Collection ids or names) is specified,
then all event handlers registered via the created hook will be
triggered and called only if their event data origin matches the tags.
// fires for every auth collection
onRecordAfterUnlinkExternalAuthRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
console.log(e.externalAuth)
// fires only for "users" and "managers" auth collections
onRecordBeforeUnlinkExternalAuthRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
console.log(e.externalAuth)
}, "users", "managers")
onRecordAfterUnlinkExternalAuthRequest
onRecordAfterUnlinkExternalAuthRequest hook is triggered
after each successful API record external auth unlink request.
If the optional "tags" list (Collection ids or names) is specified,
then all event handlers registered via the created hook will be
triggered and called only if their event data origin matches the tags.
// fires for every auth collection
onRecordAfterUnlinkExternalAuthRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
console.log(e.externalAuth)
// fires only for "users" and "managers" auth collections
onRecordAfterUnlinkExternalAuthRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
console.log(e.externalAuth)
}, "users", "managers")
onRecordBeforeRequestPasswordResetRequest
onRecordBeforeRequestPasswordResetRequest hook is triggered before each Record
request password reset API request (after request data load and before sending the reset email).
Could be used to additionally validate the request data or implement
completely different password reset behavior.
If the optional "tags" list (Collection ids or names) is specified,
then all event handlers registered via the created hook will be
triggered and called only if their event data origin matches the tags.
// fires for every auth collection
onRecordBeforeRequestPasswordResetRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
// fires only for "users" and "managers" auth collections
onRecordBeforeRequestPasswordResetRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
}, "users", "managers")
onRecordAfterRequestPasswordResetRequest
onRecordAfterRequestPasswordResetRequest hook is triggered
after each successful request password reset API request.
If the optional "tags" list (Collection ids or names) is specified,
then all event handlers registered via the created hook will be
triggered and called only if their event data origin matches the tags.
// fires for every auth collection
onRecordAfterRequestPasswordResetRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
// fires only for "users" and "managers" auth collections
onRecordAfterRequestPasswordResetRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
}, "users", "managers")
onRecordBeforeConfirmPasswordResetRequest
onRecordBeforeConfirmPasswordResetRequest hook is triggered before each Record
confirm password reset API request (after request data load and before persistence).
Could be used to additionally validate the request data or implement
completely different persistence behavior.
If the optional "tags" list (Collection ids or names) is specified,
then all event handlers registered via the created hook will be
triggered and called only if their event data origin matches the tags.
// fires for every auth collection
onRecordBeforeConfirmPasswordResetRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
// fires only for "users" and "managers" auth collections
onRecordBeforeConfirmPasswordResetRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
}, "users", "managers")
onRecordAfterConfirmPasswordResetRequest
onRecordAfterConfirmPasswordResetRequest hook is triggered
after each successful confirm password reset API request.
If the optional "tags" list (Collection ids or names) is specified,
then all event handlers registered via the created hook will be
triggered and called only if their event data origin matches the tags.
// fires for every auth collection
onRecordAfterConfirmPasswordResetRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
// fires only for "users" and "managers" auth collections
onRecordAfterConfirmPasswordResetRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
}, "users", "managers")
onRecordBeforeRequestVerificationRequest
onRecordBeforeRequestVerificationRequest hook is triggered before each Record
request verification API request (after request data load and before sending the verification email).
Could be used to additionally validate the loaded request data or implement
completely different verification behavior.
If the optional "tags" list (Collection ids or names) is specified,
then all event handlers registered via the created hook will be
triggered and called only if their event data origin matches the tags.
// fires for every auth collection
onRecordBeforeRequestVerificationRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
// fires only for "users" and "managers" auth collections
onRecordBeforeRequestVerificationRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
}, "users", "managers")
onRecordAfterRequestVerificationRequest
onRecordAfterRequestVerificationRequest  hook is triggered
after each successful request verification API request.
If the optional "tags" list (Collection ids or names) is specified,
then all event handlers registered via the created hook will be
triggered and called only if their event data origin matches the tags.
// fires for every auth collection
onRecordAfterRequestVerificationRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
// fires only for "users" and "managers" auth collections
onRecordAfterRequestVerificationRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
}, "users", "managers")
onRecordBeforeConfirmVerificationRequest
onRecordBeforeConfirmVerificationRequest hook is triggered before each Record
confirm verification API request (after request data load and before persistence).
Could be used to additionally validate the request data or implement
completely different persistence behavior.
If the optional "tags" list (Collection ids or names) is specified,
then all event handlers registered via the created hook will be
triggered and called only if their event data origin matches the tags.
// fires for every auth collection
onRecordBeforeConfirmVerificationRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
// fires only for "users" and "managers" auth collections
onRecordBeforeConfirmVerificationRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
}, "users", "managers")
onRecordAfterConfirmVerificationRequest
onRecordAfterConfirmVerificationRequest hook is triggered after each
successful confirm verification API request.
If the optional "tags" list (Collection ids or names) is specified,
then all event handlers registered via the created hook will be
triggered and called only if their event data origin matches the tags.
// fires for every auth collection
onRecordAfterConfirmVerificationRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
// fires only for "users" and "managers" auth collections
onRecordAfterConfirmVerificationRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
}, "users", "managers")
onRecordBeforeRequestEmailChangeRequest
onRecordBeforeRequestEmailChangeRequest hook is triggered before each Record request email change API request
(after request data load and before sending the email link to confirm the change).
Could be used to additionally validate the request data or implement
completely different request email change behavior.
If the optional "tags" list (Collection ids or names) is specified,
then all event handlers registered via the created hook will be
triggered and called only if their event data origin matches the tags.
// fires for every auth collection
onRecordBeforeRequestEmailChangeRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
// fires only for "users" and "managers" auth collections
onRecordBeforeRequestEmailChangeRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
}, "users", "managers")
onRecordAfterRequestEmailChangeRequest
onRecordAfterRequestEmailChangeRequest hook is triggered
after each successful request email change API request.
If the optional "tags" list (Collection ids or names) is specified,
then all event handlers registered via the created hook will be
triggered and called only if their event data origin matches the tags.
// fires for every auth collection
onRecordAfterRequestEmailChangeRequest(e) => {
console.log(e.httpContext)
console.log(e.record)
// fires only for "users" and "managers" auth collections
onRecordAfterRequestEmailChangeRequest((e) => {
console.log(e.httpContext)
console.log(e.record)
}, "users", "managers")
### Realtime API hooks
onRealtimeConnectRequest
`onRealtimeConnectRequest` hook is triggered right before establishing
the SSE client connection.
onRealtimeConnectRequest((e) => {
console.log(e.httpContext)
console.log(e.client.id())
console.log(e.idleTimeout) // in nanosec
onRealtimeDisconnectRequest
`onRealtimeDisconnectRequest` hook is triggered on disconnected/interrupted
SSE client connection.
onRealtimeDisconnectRequest((e) => {
console.log(e.httpContext)
console.log(e.client.id())
onRealtimeBeforeMessageSend
onRealtimeBeforeMessageSend hook is triggered right before sending
an SSE message to a client.
Returning false will prevent sending the message.
Returning any other error will close the realtime connection.
onRealtimeBeforeMessageSend((e) => {
console.log(e.httpContext)
console.log(e.client.id())
console.log(e.message)
onRealtimeAfterMessageSend
`onRealtimeAfterMessageSend` hook is triggered right after sending
an SSE message to a client.
onRealtimeAfterMessageSend((e) => {
console.log(e.httpContext)
console.log(e.client.id())
console.log(e.message)
onRealtimeBeforeSubscribeRequest
`onRealtimeBeforeSubscribeRequest` hook is triggered before changing
the client subscriptions, allowing you to further validate and
modify the submitted change.
onRealtimeBeforeSubscribeRequest((e) => {
console.log(e.httpContext)
console.log(e.client.id())
console.log(e.subscriptions)
onRealtimeAfterSubscribeRequest
`onRealtimeAfterSubscribeRequest` hook is triggered after the client
subscriptions were successfully changed.
onRealtimeAfterSubscribeRequest((e) => {
console.log(e.httpContext)
console.log(e.client.id())
console.log(e.subscriptions)
### File API hooks
Could be used to validate or modify the file response before returning it to the client.
console.log(e.httpContext)
console.log(e.record)
console.log(e.fileField)
console.log(e.servedPath)
console.log(e.servedName)
onFileBeforeTokenRequest
onFileBeforeTokenRequest hook is triggered before each file
token API request.
If no token or model was submitted, e.Model and e.Token will be empty,
allowing you to implement your own custom model file auth implementation.
If the optional "tags" list (Collection ids or names) is specified,
then all event handlers registered via the created hook will be
triggered and called only if their event data origin matches the tags.
// fires for every auth model
onFileBeforeTokenRequest((e) => {
console.log(e.httpContext)
console.log(e.token)
// fires only for "users"
onFileBeforeTokenRequest((e) => {
console.log(e.httpContext)
console.log(e.token)
}, "users")
onFileAfterTokenRequest
onFileAfterTokenRequest hook is triggered after each
successful file token API request.
If the optional "tags" list (Collection ids or names) is specified,
then all event handlers registered via the created hook will be
triggered and called only if their event data origin matches the tags.
// fires for every auth model
onFileAfterTokenRequest((e) => {
console.log(e.httpContext)
console.log(e.token)
// fires only for "users"
onFileAfterTokenRequest((e) => {
console.log(e.httpContext)
console.log(e.token)
}, "users")
### Collection API hooks
onCollectionsListRequest
`onCollectionsListRequest` hook is triggered on each API Collections list request.
Could be used to validate or modify the response before returning it to the client.
onCollectionsListRequest((e) => {
console.log(e.httpContext)
console.log(e.collections)
console.log(e.result)
onCollectionViewRequest
`onCollectionViewRequest` hook is triggered on each API Collection view request.
Could be used to validate or modify the response before returning it to the client.
onCollectionViewRequest((e) => {
console.log(e.httpContext)
console.log(e.collection)
onCollectionBeforeCreateRequest
`onCollectionBeforeCreateRequest` hook is triggered before each API Collection
create request (after request data load and before model persistence).
Could be used to additionally validate the request data or implement
completely different persistence behavior.
onCollectionBeforeCreateRequest((e) => {
console.log(e.httpContext)
console.log(e.collection)
onCollectionAfterCreateRequest
`onCollectionAfterCreateRequest` hook is triggered after each
successful API Collection create request.
onCollectionAfterCreateRequest((e) => {
console.log(e.httpContext)
console.log(e.collection)
onCollectionBeforeUpdateRequest
`onCollectionBeforeUpdateRequest` hook is triggered before each API Collection
update request (after request data load and before model persistence).
Could be used to additionally validate the request data or implement
completely different persistence behavior.
onCollectionBeforeUpdateRequest((e) => {
console.log(e.httpContext)
console.log(e.collection)
onCollectionAfterUpdateRequest
`onCollectionAfterUpdateRequest` hook is triggered after each
successful API Collection update request.
onCollectionAfterUpdateRequest((e) => {
console.log(e.httpContext)
console.log(e.collection)
onCollectionBeforeDeleteRequest
`onCollectionBeforeDeleteRequest` hook is triggered before each API
Collection delete request (after model load and before actual deletion).
Could be used to additionally validate the request data or implement
completely different delete behavior.
onCollectionBeforeDeleteRequest((e) => {
console.log(e.httpContext)
console.log(e.collection)
onCollectionAfterDeleteRequest
`onCollectionAfterDeleteRequest` hook is triggered after each
successful API Collection delete request.
onCollectionAfterDeleteRequest((e) => {
console.log(e.httpContext)
console.log(e.collection)
onCollectionsBeforeImportRequest
`onCollectionsBeforeImportRequest` hook is triggered before each API
collections import request (after request data load and before the actual import).
Could be used to additionally validate the imported collections or
to implement completely different import behavior.
onCollectionsBeforeImportRequest((e) => {
console.log(e.httpContext)
console.log(e.collections)
onCollectionsAfterImportRequest
`onCollectionsAfterImportRequest` hook is triggered after each
successful API collections import request.
onCollectionsAfterImportRequest((e) => {
console.log(e.httpContext)
console.log(e.collections)
### Settings API hooks
onSettingsListRequest
`onSettingsListRequest` hook is triggered on each successful
API Settings list request.
Could be used to validate or modify the response before
returning it to the client.
onSettingsListRequest((e) => {
console.log(e.httpContext)
console.log(e.redactedSettings)
onSettingsBeforeUpdateRequest
`onSettingsBeforeUpdateRequest` hook is triggered on each successful
API Settings list request.
Could be used to validate or modify the response before
returning it to the client.
onSettingsBeforeUpdateRequest((e) => {
console.log(e.httpContext)
console.log(e.oldSettings)
console.log(e.newSettings)
onSettingsAfterUpdateRequest
`onSettingsAfterUpdateRequest` hook is triggered after each
successful API Settings update request.
onSettingsAfterUpdateRequest((e) => {
console.log(e.httpContext)
console.log(e.oldSettings)
console.log(e.newSettings)
### Admin CRUD API hooks
onAdminsListRequest
`onAdminsListRequest` hook is triggered on each API Admins list request.
Could be used to validate or modify the response before returning it to the client.
onAdminsListRequest((e) => {
console.log(e.httpContext)
console.log(e.admins)
console.log(e.result)
onAdminViewRequest
`onAdminViewRequest` hook is triggered on each API Admin view request.
Could be used to validate or modify the response before returning it to the client.
onAdminViewRequest((e) => {
console.log(e.httpContext)
console.log(e.admin)
onAdminBeforeCreateRequest
`onAdminBeforeCreateRequest` hook is triggered before each API
Admin create request (after request data load and before model persistence).
Could be used to additionally validate the request data or implement
completely different persistence behavior.
onAdminBeforeCreateRequest((e) => {
console.log(e.httpContext)
console.log(e.admin)
onAdminAfterCreateRequest
`onAdminAfterCreateRequest` hook is triggered after each
successful API Admin create request.
onAdminAfterCreateRequest((e) => {
console.log(e.httpContext)
console.log(e.admin)
onAdminBeforeUpdateRequest
`onAdminBeforeUpdateRequest` hook is triggered before each API
Admin update request (after request data load and before model persistence).
Could be used to additionally validate the request data or implement
completely different persistence behavior.
onAdminBeforeUpdateRequest((e) => {
console.log(e.httpContext)
console.log(e.admin)
onAdminAfterUpdateRequest
`onAdminAfterUpdateRequest` hook is triggered after each
successful API Admin update request.
onAdminAfterUpdateRequest((e) => {
console.log(e.httpContext)
console.log(e.admin)
onAdminBeforeDeleteRequest
`onAdminBeforeDeleteRequest` hook is triggered before each API
Admin delete request (after model load and before actual deletion).
Could be used to additionally validate the request data or implement
completely different delete behavior.
onAdminBeforeDeleteRequest((e) => {
console.log(e.httpContext)
console.log(e.admin)
onAdminAfterDeleteRequest
`onAdminAfterDeleteRequest` hook is triggered after each
successful API Admin delete request.
onAdminAfterDeleteRequest((e) => {
console.log(e.httpContext)
console.log(e.admin)
### Admin Auth API hooks
onAdminAuthRequest
`onAdminAuthRequest` hook is triggered on each successful API Admin
authentication request (sign-in, token refresh, etc.).
Could be used to additionally validate or modify the
authenticated admin data and token.
onAdminAuthRequest((e) => {
console.log(e.httpContext)
console.log(e.admin)
console.log(e.token)
onAdminBeforeAuthWithPasswordRequest
`onAdminBeforeAuthWithPasswordRequest` hook is triggered before each Admin
auth with password API request (after request data load and before password validation).
Could be used to implement for example a custom password validation
or to locate a different Admin identity (by assigning `AdminAuthWithPasswordEvent.Admin`).
onAdminBeforeAuthWithPasswordRequest((e) => {
console.log(e.httpContext)
console.log(e.admin)
console.log(e.identity)
console.log(e.password)
onAdminAfterAuthWithPasswordRequest
`onAdminAfterAuthWithPasswordRequest` hook is triggered after each
successful Admin auth with password API request.
onAdminAfterAuthWithPasswordRequest((e) => {
console.log(e.httpContext)
console.log(e.admin)
console.log(e.identity)
console.log(e.password)
onAdminBeforeAuthRefreshRequest
`onAdminBeforeAuthRefreshRequest` hook is triggered before each Admin
auth refresh API request (right before generating a new auth token).
Could be used to additionally validate the request data or implement
completely different auth refresh behavior.
onAdminBeforeAuthRefreshRequest((e) => {
console.log(e.httpContext)
console.log(e.admin)
onAdminAfterAuthRefreshRequest
`onAdminAfterAuthRefreshRequest` hook is triggered after each
successful auth refresh API request (right after generating a new auth token).
onAdminAfterAuthRefreshRequest((e) => {
console.log(e.httpContext)
console.log(e.admin)
onAdminBeforeRequestPasswordResetRequest
`onAdminBeforeRequestPasswordResetRequest` hook is triggered before each Admin
request password reset API request (after request data load and before sending the reset email).
Could be used to additionally validate the request data or implement
completely different password reset behavior.
onAdminBeforeRequestPasswordResetRequest((e) => {
console.log(e.httpContext)
console.log(e.admin)
onAdminAfterRequestPasswordResetRequest
`onAdminAfterRequestPasswordResetRequest` hook is triggered after each
successful request password reset API request.
onAdminBeforeRequestPasswordResetRequest((e) => {
console.log(e.httpContext)
console.log(e.admin)
onAdminBeforeConfirmPasswordResetRequest
`onAdminBeforeConfirmPasswordResetRequest` hook is triggered before each Admin
confirm password reset API request (after request data load and before persistence).
Could be used to additionally validate the request data or implement
completely different persistence behavior.
onAdminBeforeConfirmPasswordResetRequest((e) => {
console.log(e.httpContext)
console.log(e.admin)
onAdminAfterConfirmPasswordResetRequest
`onAdminAfterConfirmPasswordResetRequest` hook is triggered after each
successful confirm password reset API request.
onAdminAfterConfirmPasswordResetRequest((e) => {
console.log(e.httpContext)
console.log(e.admin)

## 22.Extend with JavaScript - Routing
GET /hello/:name
# Extend with JavaScript - Routing
`GET /hello/:name`
You can register custom routes and middlewares by using the top-level routerAdd()
and routerUse() functions.
### Routes
##### Registering new routes
Each route consists of at least a path and a handler function. For example, the below code registers
GET /hello/:name route that responds with a json body:
routerAdd("GET", "/hello/:name", (c) => {
let name = c.pathParam("name")
return c.json(200, { "message": "Hello " + name })
}, /* optional middlewares */)
To avoid collisions with future internal routes you should avoid using the /api/...
base path or consider combining it with a unique prefix like /api/myapp/....
Each handler function receives a request context argument (usually named c).
The request context is also accessible in the event request hooks under the httpContext key.
Below you can find common request context operations.
##### Request context store
The request context comes with a local store that you can use to share data related only to the current
request between routes and middlewares.
// store for the duration of the request
c.set("someKey", 123)
// retrieve later
const val = c.get("someKey") // 123
##### Retrieving the current auth state
We also use the store to manage the current auth state with the admin and
authRecord special keys.
const admin  = c.get("admin")      // empty if not authenticated as admin
const record = c.get("authRecord") // empty if not authenticated as regular auth record
// alternatively, you can also read the auth state from the cached request info
const info   = $apis.requestInfo(c);
const admin  = info.admin;      // empty if not authenticated as admin
const record = info.authRecord; // empty if not authenticated as regular auth record
const isGuest = !admin &amp;&amp; !record
##### Reading path parameters
c.pathParam(&quot;paramName&quot;).
`const id = c.pathParam("id")`
##### Reading query parameters
const search = c.queryParam("search")
// or via the cached request object
const search = $apis.requestInfo(c).query.search
##### Reading request headers
const token = c.request().header.get("Some-Header")
// or via the cached request object (the header value is always normalized)
const token = $apis.requestInfo(c).headers["some_header"]
##### Writing response headers
`c.response().header().set("Some-Header", "123")`
##### Reading request body
// read the body via the cached request object
// (this method is commonly used in hook handlers because it allows reading the body more than once)
const data = $apis.requestInfo(c).data
console.log(data.title)
// read/scan the request body fields into a typed object
// (note that a body cannot be read twice with "bind" because it is a stream)
const data = new DynamicModel({
// describe the fields to read (used also as initial values)
someTextField:   "",
someNumberField: 0,
someBoolField:   false,
someArrayField:  [],
someObjectField: {}, // object props are accessible via .get(key)
c.bind(data)
console.log(data.sometextField)
// read single multipart/form-data field
const title = c.formValue("title")
// read single multipart/form-data file
const doc = c.formFile("document")
##### Writing response body
// send response with json body
c.json(200, {"name": "John"})
// send response with string body
// send response with html body
// (check also the "Rendering templates" section)
c.html(200, "&lt;h1>Hello!&lt;/h1>")
// redirect
// send response with no body
c.noContent(204)
### Middlewares
Middlewares could be used to apply a shared behavior or to intercept and modify route requests.
Middlewares can be registered both to a single route (by passing them after the handler) and globally usually
by using routerUse(someMiddlereFunc).
// attach a middleware globally to all routes
routerUse(someMiddlereFunc)
// attach multiple middlewares to a single route
// each route will execute their own middlewares + the global ones
routerAdd("GET", "/hello", (c) => {
return c.string(200, "Hello world!")
}, $apis.activityLogger($app), $apis.requireAdminAuth())
##### Builtin middlewares
// logs the request in the Admin UI > Logs
$apis.activityLogger($app)
// requires the request client to be unauthenticated, aka. guest
$apis.requireGuestOnly()
// requires the request client to be authenticated as an auth record
$apis.requireRecordAuth(optCollectionNames...)
// require the request client to be authenticated as admin
$apis.requireAdminAuth()
// require the request client to be authenticated as admin OR auth record
$apis.requireAdminOrRecordAuth(optCollectionNames...)
// require the request client to be authenticated as admin OR auth record
// that matches the ownerIdParam path parameter
$apis.requireAdminOrOwnerAuth(ownerIdParam = "id")
// compresses HTTP response using gzip
$apis.gzip()
// sets the maximum allowed size (in bytes) for a request body
$apis.bodyLimit(bytes)
##### Custom middlewares
return (c) => {
// eg. inspect some header value before processing the request
const header = c.request().header.get("Some-Header")
if (!header) {
// throw or return an error
throw new BadRequestError("Invalid request")
routerUse(myCustomMiddleware)
### Error response
PocketBase has a global error handler and every returned or thrown Error from a route or
middleware will be safely converted by default to a generic HTTP 400 error to avoid accidentally leaking
sensitive information (the original error will be visible only in the Admin UI &gt; Logs or when in
--dev mode).
To make it easier returning formatted json error responses, PocketBase provides
ApiError constructor that can be instantiated directly or using the builtin factories.
ApiError.data will be returned in the response only if it is a map of
ValidationError items.
// construct ApiError with custom status code and validation data error
throw new ApiError(500, "something went wrong", {
"title": new ValidationError("invalid_title", "Invalid or missing title"),
// if message is empty string, a default one will be set
throw new BadRequestError(optMessage, optData)   // 400 ApiError
throw new UnauthorizedError(optMessage, optData) // 401 ApiError
throw new ForbiddenError(optMessage, optData)    // 403 ApiError
throw new NotFoundError(optMessage, optData)     // 404 ApiError
### Helpers
The global $apis namespace expose several helpers you can use as part of your route hooks.
##### Auth response
$apis.recordAuthResponse() writes standardised json record auth response (aka. token + record
data) into the specified request context. Could be used as a return result from a custom auth route.
routerAdd("GET", "/phone-login", (c) => {
const data = new DynamicModel({
phone:    "",
password: "",
c.bind(data)
const record = $app.dao().findFirstRecordByData("users", "phone", data.phone)
if (!record.validatePassword(data.password)) {
throw new BadRequestError("invalid credentials")
return $apis.recordAuthResponse($app, c, record)
}, $apis.activityLogger($app))
##### Enrich record(s)
$apis.enrichRecord() and $apis.enrichRecords() helpers parses the request context
and enrich the provided record(s) by:
-expands relations (if defaultExpands and/or ?expand query parameter is set)
-ensures that the emails of the auth record and its expanded auth relations are visible only for the
current logged admin, record owner or record with manage access
routerAdd("GET", "/custom-article", (c) => {
const records = $app.dao().findRecordsByFilter("article", "status = 'active'", '-created', 40)
// enrich the records with the "categories" relation as default expand
$apis.enrichRecords(c, $app.dao(), records, "categories")
return c.json(200, records)
}, $apis.activityLogger($app))
##### Serving static files
`routerAdd("GET", "/*", $apis.staticDirectoryHandler("/path/to/public", false))`
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

## 23.Extend with JavaScript - Database
# Extend with JavaScript - Database
The main interface to interact with your application database is via $app.dao().
$app.dao()
provides read and write helpers (see Collection operations
and Record operations) and it is responsible for triggering the
onModel* event hooks.
It also exposes $app.dao().db() builder that allows executing various SQL statements (including
### Executing queries
To execute DB queries you can start with the newQuery(&quot;...&quot;) statement and then call one of:
-execute()
- for any query statement that is not meant to retrieve data:
$app.dao().db()
.newQuery("CREATE INDEX name_idx ON users (name)")
.execute() // throw an error on db failure
-one()
- to populate a single row into DynamicModel object:
const result = new DynamicModel({
// describe the shape of the data (used also as initial values)
"id":     "",
"status": false,
"age":    0,
"roles":  [], // serialized json db arrays are decoded as plain arrays
$app.dao().db()
.newQuery("SELECT id, status, age, roles FROM users WHERE id=1")
.one(result) // throw an error on db failure or missing row
console.log(result.id)
-all()
- to populate multiple rows into an array of objects (note that the array must be created with
arrayOf):
const result = arrayOf(new DynamicModel({
// describe the shape of the data (used also as initial values)
"id":     "",
"status": false,
"age":    0,
"roles":  [], // serialized json db arrays are decoded as plain arrays
$app.dao().db()
.newQuery("SELECT id, status, age, roles FROM users LIMIT 100")
.all(result) // throw an error on db failure
if (result.length > 0) {
console.log(result[0].id)
### Binding parameters
To prevent SQL injection attacks, you should use named parameters for any expression value that comes from
user input. This could be done using the named {:paramName}
bind(params). For example:
const result = arrayOf(new DynamicModel({
"name":    "",
"created": "",
$app.dao().db()
.newQuery("SELECT name, created FROM posts WHERE created >= {:from} and created &lt;= {:to}")
.bind({
"from": "2023-06-25 00:00:00.000Z",
"to":   "2023-06-28 23:59:59.999Z",
.all(result)
console.log(result.length)
### Query builder
Instead of writting plain SQLs, you can also compose SQL statements programmatically using the db query
builder.
Every SQL keyword has a corresponding query building method. For example, SELECT corresponds
to select(), FROM corresponds to from(),
WHERE corresponds to where(), and so on.
const result = arrayOf(new DynamicModel({
"id":    "",
"email": "",
$app.dao().db()
.select("id", "email")
.from("users")
.limit(100)
.orderBy("created ASC")
.all(result)
##### select(), andSelect(), distinct()
The select(...cols) method initializes a SELECT query builder. It accepts a list
of the column names to be selected.
To add additional columns to an existing select query, you can call andSelect().
To select distinct rows, you can call distinct().
$app.dao().db()
.select("id", "avatar as image")
.andSelect("(firstName || ' ' || lastName) as fullName")
.distinct()
##### from()
The from(...tables) method specifies which tables to select from (plain table names are automatically
quoted).
$app.dao().db()
.select("table1.id", "table2.name")
.from("table1", "table2")
##### join()
The join(type, table, on) method specifies a JOIN clause. It takes 3 parameters:
-type - join type string like INNER JOIN, LEFT JOIN, etc.
-table - the name of the table to be joined
-on - optional dbx.Expression as an ON clause
For convenience, you can also use the shortcuts innerJoin(table, on),
leftJoin(table, on),
rightJoin(table, on) to specify INNER JOIN, LEFT JOIN and
RIGHT JOIN, respectively.
$app.dao().db()
.select("users.*")
.from("users")
.innerJoin("profiles", $dbx.exp("profiles.user_id = users.id"))
.join("FULL OUTER JOIN", "department", $dbx.exp("department.id = {:id}", {id: "someId"}))
##### where(), andWhere(), orWhere()
The where(exp) method specifies the WHERE condition of the query.
You can also use andWhere(exp) or orWhere(exp) to append additional one or more
conditions to an existing WHERE clause.
Each where condition accepts a single dbx.Expression (see below for full list).
SELECT users.*
FROM users
WHERE id = "someId" AND
status = "public" AND
name like "%john%" OR
role = "manager" AND
fullTime IS TRUE AND
experience > 10
$app.dao().db()
.select("users.*")
.from("users")
.where($dbx.exp("id = {:id}", { id: "someId" }))
.andWhere($dbx.hashExp({ status: "public" }))
.andWhere($dbx.like("name", "john"))
.orWhere($dbx.and(
$dbx.hashExp({
role:     "manager",
fullTime: true,
$dbx.exp("experience > {:exp}", { exp: 10 })
The following dbx.Expression methods are available:
parameters to the expression.
$dbx.exp("status = 'public'")
$dbx.exp("total > {:min} AND total &lt; {:max}", { min: 10, max: 30 })
-$dbx.hashExp(pairs)
Generates a hash expression from a map whose keys are DB column names which need to be filtered according
to the corresponding values.
// slug = "example" AND active IS TRUE AND tags in ("tag1", "tag2", "tag3") AND parent IS NULL
$dbx.hashExp({
slug:   "example",
active: true,
tags:   ["tag1", "tag2", "tag3"],
parent: null,
-$dbx.not(exp)
Negates a single expression by wrapping it with NOT().
// NOT(status = 1)
$dbx.not($dbx.exp("status = 1"))
-$dbx.and(...exps)
Creates a new expression by concatenating the specified ones with AND.
// (status = 1 AND username like "%john%")
$dbx.and($dbx.exp("status = 1"), $dbx.like("username", "john"))
-$dbx.or(...exps)
Creates a new expression by concatenating the specified ones with OR.
// (status = 1 OR username like "%john%")
$dbx.or($dbx.exp("status = 1"), $dbx.like("username", "john"))
-$dbx.in(col, ...values)
Generates an IN expression for the specified column and the list of allowed values.
// status IN ("public", "reviewed")
$dbx.in("status", "public", "reviewed")
-$dbx.notIn(col, ...values)
Generates an NOT IN expression for the specified column and the list of allowed values.
// status NOT IN ("public", "reviewed")
$dbx.notIn("status", "public", "reviewed")
-$dbx.like(col, ...values)
Generates a LIKE expression for the specified column and the possible strings that the
column should be like. If multiple values are present, the column should be like
all of them.
By default, each value will be surrounded by &quot;%&quot; to enable partial matching. Special
characters like &quot;%&quot;, &quot;\&quot;, &quot;_&quot; will also be properly escaped. You may call
escape(...pairs) and/or match(left, right) to change the default behavior.
// name LIKE "%test1%" AND name LIKE "%test2%"
$dbx.like("name", "test1", "test2")
// name LIKE "test1%"
$dbx.like("name", "test1").match(false, true)
-$dbx.notLike(col, ...values)
Generates a NOT LIKE expression in similar manner as like().
// name NOT LIKE "%test1%" AND name NOT LIKE "%test2%"
$dbx.notLike("name", "test1", "test2")
// name NOT LIKE "test1%"
$dbx.notLike("name", "test1").match(false, true)
-$dbx.orLike(col, ...values)
This is similar to like() except that the column must be one of the provided values, aka.
multiple values are concatenated with OR instead of AND.
// name LIKE "%test1%" OR name LIKE "%test2%"
$dbx.orLike("name", "test1", "test2")
// name LIKE "test1%" OR name LIKE "test2%"
$dbx.orLike("name", "test1", "test2").match(false, true)
-$dbx.orNotLike(col, ...values)
This is similar to notLike() except that the column must not be one of the provided
values, aka. multiple values are concatenated with OR instead of AND.
// name NOT LIKE "%test1%" OR name NOT LIKE "%test2%"
$dbx.orNotLike("name", "test1", "test2")
// name NOT LIKE "test1%" OR name NOT LIKE "test2%"
$dbx.orNotLike("name", "test1", "test2").match(false, true)
-$dbx.exists(exp)
Prefix with EXISTS the specified expression (usually a subquery).
// EXISTS (SELECT 1 FROM users WHERE status = 'active')
$dbx.exists(dbx.exp("SELECT 1 FROM users WHERE status = 'active'"))
-$dbx.notExists(exp)
Prefix with NOT EXISTS the specified expression (usually a subquery).
// NOT EXISTS (SELECT 1 FROM users WHERE status = 'active')
$dbx.notExists(dbx.exp("SELECT 1 FROM users WHERE status = 'active'"))
-$dbx.between(col, from, to)
Generates a BETWEEN expression with the specified range.
// age BETWEEN 3 and 99
$dbx.between("age", 3, 99)
-$dbx.notBetween(col, from, to)
Generates a NOT BETWEEN expression with the specified range.
// age NOT BETWEEN 3 and 99
$dbx.notBetween("age", 3, 99)
##### orderBy(), andOrderBy()
The orderBy(...cols) specifies the ORDER BY clause of the query.
A column name can contain &quot;ASC&quot; or &quot;DESC&quot; to indicate its ordering direction.
You can also use andOrderBy(...cols) to append additional columns to an existing
ORDER BY clause.
$app.dao().db()
.select("users.*")
.from("users")
.orderBy("created ASC", "updated DESC")
.andOrderBy("title ASC")
##### groupBy(), andGroupBy()
The groupBy(...cols) specifies the GROUP BY clause of the query.
You can also use andGroupBy(...cols) to append additional columns to an existing
GROUP BY clause.
$app.dao().db()
.select("users.*")
.from("users")
.groupBy("department", "level")
##### having(), andHaving(), orHaving()
The having(exp) specifies the HAVING clause of the query.
Similarly to
where(exp), it accept a single dbx.Expression (see all available expressions
listed above).
You can also use andHaving(exp) or orHaving(exp) to append additional one or
more conditions to an existing HAVING clause.
$app.dao().db()
.select("users.*")
.from("users")
.groupBy("department", "level")
.having($dbx.exp("sum(level) > {:sum}", { sum: 10 }))
##### limit()
The limit(number) method specifies the LIMIT clause of the query.
$app.dao().db()
.select("users.*")
.from("users")
.limit(30)
##### offset()
The offset(number) method specifies the OFFSET clause of the query. Usually used
together with limit(number).
$app.dao().db()
.select("users.*")
.from("users")
.offset(5)
.limit(30)
### Transaction
To execute multiple queries in a transaction you can use $app.dao().runInTransaction()
You can nest Dao.runInTransaction() as many times as you want.
The transaction will be committed only if there are no errors.
$app.dao().runInTransaction((txDao) => {
// update a record
const record = txDao.findRecordById("articles", "RECORD_ID")
record.set("status", "active")
txDao.saveRecord(record)
txDao.db().newQuery("DELETE FROM articles WHERE status = 'pending'").execute()
### Dao without event hooks
By default all Dao write operations (create, update, delete) trigger the onModel* event
hooks.
If you don&#39;t want this behavior, you can create a new Dao without hooks from an existing one by calling
Dao.withoutHooks()
or instantiate a new one with new Dao(db, [nonconcurrentDB]):
const record = $app.dao().findRecordById("articles", "RECORD_ID")
// the below WILL fire the onModelBeforeUpdate and onModelAfterUpdate hooks
$app.dao().saveRecord(record)
// the below WILL NOT fire the onModelBeforeUpdate and onModelAfterUpdate hooks
const dao = $app.dao().withoutHooks() // or new Dao($app.dao().db())
dao.saveRecord(record)

## 24.Extend with JavaScript - Record operations
# Extend with JavaScript - Record operations
The most common task when extending PocketBase probably would be querying and operating with your
collection records.
### Get/Set record fields
// export the public safe record fields as map[string]any
record.publicExport()
// returns a new model copy populated with the original/intial record data
// (could be useful if you want to compare old and new field values)
record.originalCopy()
// returns a copy of the current record model populated only
// with its latest data state and everything else reset to the defaults
record.cleanCopy()
// set the value of a single record field
record.set("someField", 123)
// bulk set fields from a map
record.load(data)
// retrieve a single record field value
record.get("someField")            // -> as any
record.getBool("someField")        // -> as bool
record.getString("someField")      // -> as string
record.getInt("someField")         // -> as int
record.getFloat("someField")       // -> as float64
record.getTime("someField")        // -> as time.Time
record.getDateTime("someField")    // -> as types.DateTime
record.getStringSlice("someField") // -> as []string
// unmarshal a single json field value into the provided result
const result = new DynamicModel({ ... })
record.unmarshalJSONField("someJsonField", result)
// retrieve a single or multiple expanded data
record.expandedOne("author")     // -> as null|Record
record.expandedAll("categories") // -> as []Record
// auth records only
record.setPassword("123456")
record.validatePassword("123456")
record.passwordHash()
record.username()
record.setUsername("john.doe")
record.email()
record.emailVisibility()
record.setEmailVisibility(false)
record.verified()
record.setVerified(false)
record.tokenKey()
record.setTokenKey("ABCD123")
record.refreshTokenKey() // sets autogenerated TokenKey
record.lastResetSentAt()
record.setLastResetSentAt(new DateTime())
record.lastVerificationSentAt()
record.setLastVerificationSentAt(new DateTime())
### Fetch records
##### Fetch single record
// retrieve a single "articles" collection record by its id
const record = $app.dao().findRecordById("articles", "RECORD_ID")
// retrieve a single "articles" collection record by a single key-value pair
const record = $app.dao().findFirstRecordByData("articles", "slug", "test")
// retrieve a single "articles" collection record by a string filter expression
const record = $app.dao().findFirstRecordByFilter(
"articles", "status = 'public' &amp;&amp; category = {:category}",
{ category: "news" },
##### Fetch multiple records
// retrieve multiple "articles" collection records by their ids
const records = $app.dao().findRecordsByIds("articles", ["RECORD_ID1", "RECORD_ID2"])
// retrieve multiple "articles" collection records by a custom dbx expression(s)
// (for all avalaible expressions, please check the Database guide)
const records = $app.dao().findRecordsByExpr("articles",
$dbx.exp("LOWER(username) = {:username}", { "username": "John.Doe" }),
$dbx.hashExp({ status: "pending" })
// retrieve multiple "articles" collection records by a string filter expression
const records = $app.dao().findRecordsByFilter(
"articles",                                    // collection
"status = 'public' &amp;&amp; category = {:category}", // filter
"-publised",                                   // sort
10,                                            // limit
0,                                             // offset
{ category: "news" },                          // optional filter params
##### Fetch auth records
// retrieve a single auth collection record by its email
// retrieve a single auth collection record by its username (case insensitive)
const user = $app.dao().findAuthRecordByUsername("users", "John.Doe")
// retrieve a single auth collection record by its JWT (auth, password reset, etc.)
const user = $app.dao().findAuthRecordByToken("YOUR_TOKEN", $app.settings().recordAuthToken.secret)
##### Custom record query
In addition to the above read and write helpers, you can also create custom Record model queries using
Dao.recordQuery(collection)
method. It returns a DB builder that can be used with the same methods described in the
Database guide.
For retrieving a single Record model with the one() executor, you can use a
blank new Record() model to populate the result in.
function findTopArticle() {
const record = new Record();
$app.dao().recordQuery("articles")
.andWhere($dbx.hashExp({ "status": "active" }))
.orderBy("rank ASC")
.limit(1)
.one(record)
return record
const article = findTopArticle()
For retrieving multiple Record models with the all() executor, you can use
arrayOf(new Record)
// the below is identical to
// dao.findRecordsByFilter("articles", "status = 'active'", '-published', 10)
// but allows more advanced use cases and filtering (aggregations, subqueries, etc.)
function findLatestArticles() {
const records = arrayOf(new Record);
$app.dao().recordQuery("articles")
.andWhere($dbx.hashExp({ "status": "active" }))
.orderBy("published DESC")
.limit(10)
.all(records)
return records
const articles = findLatestArticles()
### Create new record
##### Create new record WITHOUT data validations
const collection = $app.dao().findCollectionByNameOrId("articles")
const record = new Record(collection, {
// bulk load the record data during initialization
"active": true
// or load individual fields separately
record.set("someOtherField", 123)
$app.dao().saveRecord(record)
##### Create new record WITH data validations
const collection = $app.dao().findCollectionByNameOrId("articles")
const record = new Record(collection)
const form = new RecordUpsertForm($app, record)
// or form.loadRequest(request, "")
form.loadData({
"active":         true,
"someOtherField": 123,
// manually upload file(s)
const f1 = $filesystem.fileFromPath("/path/to/file1")
const f2 = $filesystem.fileFromPath("/path/to/file2")
form.addFiles("yourFileField1", f1, f2)
// or mark file(s) for deletion
form.removeFiles("yourFileField2", "demo_xzihx0w.png")
// validate and submit (internally it calls $app.dao().saveRecord(record) in a transaction)
form.submit()
##### Intercept record before create API hook
onRecordBeforeCreateRequest((e) => {
if (e.httpContext.get("admin")) {
return null // ignore for admins
// overwrite the submitted "active" field value to false
e.record.set("active", false)
// or you can also prevent the create event by returning an error, eg.:
if (e.record.get("status") != "pending") {
throw new BadRequestError("status must be pending")
}, "articles")
### Update existing record
##### Update record WITHOUT data validations
const record = $app.dao().findRecordById("articles", "RECORD_ID")
// set individual fields
// or bulk load with record.load({...})
record.set("active", true)
record.set("someOtherField", 123)
$app.dao().saveRecord(record)
##### Update record WITH data validations
const record = $app.dao().findRecordById("articles", "RECORD_ID")
const form = new RecordUpsertForm($app, record)
// or form.loadRequest(request, "")
form.loadData({
"active":         true,
"someOtherField": 123,
// validate and submit (internally it calls $app.dao().saveRecord(record) in a transaction)
form.submit();
##### Intercept record before update API hook
onRecordBeforeUpdateRequest((e) => {
if (e.httpContext.get("admin")) {
return null // ignore for admins
// overwrite the submitted "active" field value to false
e.record.set("active", false)
// or you can also prevent the create event by returning an error, eg.:
if (e.record.get("status") != "pending") {
throw new BadRequestError("status must be pending")
}, "articles")
### Delete record
const record = $app.dao().findRecordById("articles", "RECORD_ID")
$app.dao().deleteRecord(record)
### Transaction
const titles = ["title1", "title2", "title3"]
const collection = $app.dao().findCollectionByNameOrId("articles")
$app.dao().runInTransaction((txDao) => {
// create new record for each title
for (let title of titles) {
const record = new Record(collection)
record.set("title", title)
txDao.saveRecord(record)
### Programmatically expanding relations
To expand record relations programmatically you can use the
$app.dao().expandRecord(record, expands, customFetchFunc) or
$app.dao().expandRecords(records, expands, customFetchFunc)
methods.
Once loaded, you can access the expanded relations via
record.expandedOne(relName) or
record.expandedAll(relName) methods.
For example:
const record = $app.dao().findFirstRecordByData("articles", "slug", "lorem-ipsum")
// expand the "author" and "categories" relations
$app.dao().expandRecord(record, ["author", "categories"], null)
// print the expanded records
console.log(record.expandedOne("author"))
console.log(record.expandedAll("categories"))
### Check if record can be accessed
To check whether a custom client request or user can access a single record, you can use the
$app.dao().canAccessRecord(record, requestInfo, rule) method.
For example:
// allow access to the article with the specified slug
// only if the current client request satisfy the articles view rule
routerAdd("get", "/articles/:slug", (c) => {
const info = $apis.requestInfo(c)
const slug = c.pathParam("slug")
const record = $app.dao().findFirstRecordByData("articles", "slug", slug)
const canAccess = $app.dao().canAccessRecord(record, info, record.collection().viewRule)
if (!canAccess) {
throw new ForbiddenError()
return c.json(200, record)

## 25.Extend with JavaScript - Collection operations
# Extend with JavaScript - Collection operations
Collections are usually managed via the Admin UI, but there are some situations where you may want to
create or edit a collection programmatically (usually as part of a
DB migration). PocketBase exposes several helpers to simplify the
Collection model operations.
### Fetch collections
##### Fetch collection by name or id
`const collection = $app.dao().findCollectionByNameOrId("example")`
##### Fetch collections by type
const baseCollections = $app.dao().findCollectionsByType("base")
const authCollections = $app.dao().findCollectionsByType("auth")
const viewCollections = $app.dao().findCollectionsByType("view")
### Create new collection
##### Create new collection WITHOUT data validations
const collection = new Collection({
// the id is autogenerated, but you can set a specific one if you want to
// id:      "...",
name:       "example",
type:       "base",
listRule:   null,
viewRule:   "@request.auth.id != ''",
createRule: "",
updateRule: "@request.auth.id != ''",
deleteRule: null,
schema:     [
name:     "title",
type:     "text",
required: true,
options:  {
max: 10,
name:     "user",
type:     "relation",
required: true,
options:  {
maxSelect:     1,
collectionId:  "ae40239d2bc4477",
cascadeDelete: true,
indexes: [
"CREATE UNIQUE INDEX idx_user ON example (user)"
options: {}
$app.dao().saveCollection(collection)
##### Create new collection WITH data validations
const collection = new Collection()
const form = new CollectionUpsertForm($app, collection)
form.name = "example"
form.type = "base"
form.listRule = null
form.viewRule = "@request.auth.id != ''"
form.createRule = ""
form.updateRule = "@request.auth.id != ''"
form.deleteRule = null
form.schema.addField(new SchemaField({
name:     "title",
type:     "text",
required: true,
options: {
max: 10,
form.schema.addField(new SchemaField({
name:     "user",
type:     "relation",
options: {
maxSelect:     1,
collectionId:  "ae40239d2bc4477",
cascadeDelete: true,
// validate and submit (internally it calls $app.dao().saveCollection(collection) in a transaction)
form.submit()
### Update existing collection
##### Update collection WITHOUT data validations
const collection = $app.dao().findCollectionByNameOrId("example")
// change the collection name
collection.name = "example_update"
// add new field
collection.schema.addField(new SchemaField({
name: "description",
type: "text",
$app.dao().saveCollection(collection)
##### Update collection WITH data validations
const collection = $app.dao().findCollectionByNameOrId("example")
const form = new CollectionUpsertForm(app, collection)
// change the collection name
form.name = "example_update"
// add new field
form.schema.addField(new SchemaField{
name: "description",
type: "text",
// validate and submit (internally it calls $app.dao().saveCollection(collection) in a transaction)
form.submit()
### Delete collection
const collection = $app.dao().findCollectionByNameOrId("example")
$app.dao().deleteCollection(collection)

## 26.Extend with JavaScript - Migrations
# Extend with JavaScript - Migrations
PocketBase comes with a builtin DB and data migration utility, allowing you to version your DB structure,
create collections programmatically, initialize default settings and/or run anything that needs to be
executed only once.
The user defined migrations are located in pb_migrations directory (it can be changed using
the
--migrationsDir flag) and each unapplied migration inside it will be executed automatically
in a transaction on serve (or on migrate up).
The generated migrations are safe to be commited to version control and can be shared with your other team
members.
### Automigrate
The prebuilt executable has the --automigrate flag enabled by default, meaning that every collection
configuration change from the Admin UI will generate the related migration file automatically for you.
### Creating migrations
To create a new blank migration you can run migrate create.
`[root@dev app]$ ./pocketbase migrate create "your_new_migration"`
// pb_migrations/1687801097_your_new_migration.js
migrate((db) => {
// add up queries...
}, (db) => {
// add down queries...
New migrations are applied automatically on serve.
Optionally, you could apply new migrations manually by running migrate up.
To revert the last applied migration(s), you could run migrate down [number].
##### Migration file
Each migration file should have a single migrate(upFunc, downFunc) call.
In the migration file, you are expected to write your &quot;upgrade&quot; code in the upFunc callback.
The downFunc is optional and it should contains the &quot;downgrade&quot; operations to revert the
changes made by the
upFunc.
Both callbacks accept a single db argument (dbx.Builder) that you can use
directly or create a Dao instance and use its available helpers. You can explore the
Database guide
for more details how to operate with the db object and its available methods.
### Collections snapshot
PocketBase comes also with a migrate collections command that will generate a full snapshot of
your current Collections configuration without having to type it manually:
`[root@dev app]$ ./pocketbase migrate collections`
Similar to the migrate create command, this will generate a new migration file in the
pb_migrations directory.
It is safe to run the command multiple times and generate multiple snapshot migration files.
### Migrations history
All applied migration filenames are stored in the internal _migrations table.
During local development often you might end up making various collection changes to test different approaches.
When --automigrate is enabled (which is the default) this could lead in a migration
history with unnecessary intermediate steps that may not be wanted in the final migration history.
To avoid the clutter and to prevent applying the intermediate steps in production, you can remove (or
squash) the unnecessary migration files manually and then update the local migrations history by running:
`[root@dev app]$ ./pocketbase migrate history-sync`
The above command will remove any entry from the _migrations table that doesn&#39;t have a related
migration file associated with it.
// pb_migrations/1687801090_set_pending_status.js
// set a default "pending" status to all empty status articles
migrate((db) => {
db.newQuery("UPDATE articles SET status = 'pending' WHERE status = ''")
.execute()
##### Initialize default application settings
// pb_migrations/1687801090_initial_settings.js
migrate((db) => {
const dao = new Dao(db);
const settings = dao.findSettings()
settings.meta.appName = "test"
settings.logs.maxDays = 2
dao.saveSettings(settings)
##### Creating new admin
// pb_migrations/1687801090_initial_admin.js
migrate((db) => {
const dao = new Dao(db);
const admin = new Admin();
admin.setPassword("1234567890")
dao.saveAdmin(admin)
}, (db) => { // optional revert
const dao = new Dao(db);
try {
dao.deleteAdmin(admin)
} catch (_) { /* most likely already deleted */ }
##### Creating new auth record
// pb_migrations/1687801090_new_users_record.js
migrate((db) => {
const dao = new Dao(db);
const collection = dao.findCollectionByNameOrId("users")
const record = new Record(collection)
record.setUsername("u_" + $security.randomStringWithAlphabet(5, "123456789"))
record.setPassword("1234567890")
record.set("name", "John Doe")
dao.saveRecord(record)
}, (db) => { // optional revert
const dao = new Dao(db);
try {
dao.deleteRecord(record)
} catch (_) { /* most likely already deleted */ }

## 27.Extend with JavaScript - Jobs scheduling
# Extend with JavaScript - Jobs scheduling
If you have tasks that need to be performed periodically, you could setup crontab-like jobs with
cronAdd(name, expr, handler).
Each scheduled job runs in its own goroutine as part of the serve command process and must have:
-name - identifier for the scheduled job; could be used to replace or remove an existing job
-cron expression like 0 0 * * * (
supports numeric list, steps, ranges or
macros
-handler - the function that will be executed everytime when the job runs
Here is an example:
// prints "Hello!" every 2 minutes
cronAdd("hello", "*/2 * * * *", () => {
console.log("Hello!")
To remove a single registered cron job you can call cronRemove(name).

## 28.Extend with JavaScript - Console commands
# Extend with JavaScript - Console commands
You can register custom console commands using
app.rootCmd.addCommand(cmd), where cmd is a
Command instance.
Here is an example:
$app.rootCmd.addCommand(new Command({
use: "hello",
run: (cmd, args) => {
console.log("Hello world!")
To run the command you can execute:
`./pocketbase hello`
Keep in mind that the console commands execute in their own separate app process and run
independently from the main serve command (aka. hook events between different processes
are not shared with one another).

## 29.Extend with JavaScript - Sending emails
# Extend with JavaScript - Sending emails
PocketBase provides a simple abstraction for sending emails via the
$app.newMailClient() helper.
Depending on your configured mail settings (Admin UI &gt; Settings &gt; Mail settings) it will use the
sendmail command or a SMTP client.
### Send custom email
You can send your own custom emails from everywhere within your app (hooks, middlewares, routes, etc.) by
using $app.newMailClient().send(message). Here is an example of sending a custom email after
user registration:
onRecordAfterCreateRequest((e) => {
const message = new MailerMessage({
from: {
address: $app.settings().meta.senderAddress,
name:    $app.settings().meta.senderName,
to:      [{address: e.record.email()}],
subject: "YOUR_SUBJECT...",
html:    "YOUR_HTML_BODY...",
// bcc, cc and custom headers are also supported...
$app.newMailClient().send(message)
}, "users")
### Intercept system emails
If you want to change the default system emails for forgotten password, verification, etc., you can adjust
the default templates from the Admin UI &gt; Settings &gt; Mail settings.
Alternatively, you can also apply individual changes by binding to one of the
mailer hooks. Here is an example of appending a Record
field value to the subject using the onMailerBeforeRecordResetPasswordSend hook:
onMailerBeforeRecordResetPasswordSend().add((e) => {
// modify the subject
e.message.subject += (" " + e.record.get("name"))

## 30.Extend with JavaScript - Sending HTTP requests
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

## 31.Extend with JavaScript - Rendering templates
Response 200:
{{template "placeholderName" .}}
Response 200:
{{block "placeholderName" .}}default...{{end}}
Response 200:
{{define "placeholderName"}}custom...{{end}}
# Extend with JavaScript - Rendering templates
### Overview
A common task when creating custom routes or emails is the need of generating HTML output. To assist with
this, PocketBase provides the global $template helper for parsing and rendering HTML templates.
const html = $template.loadFiles(
`${__hooks}/views/base.html`,
`${__hooks}/views/partial1.html`,
`${__hooks}/views/partial2.html`,
).render(data)
The general flow when working with composed and nested templates is that you create &quot;base&quot; template(s)
The dot object (.) in the above represents the data passed to the templates
via the render(data) method.
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
pb_hooks/
views/
layout.html
hello.html
main.pb.js
pocketbase
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
routerAdd("get", "/hello/:name", (c) => {
const name = c.pathParam("name")
const html = $template.loadFiles(
`${__hooks}/views/layout.html`,
`${__hooks}/views/hello.html`,
).render({
"name": name,
return c.html(200, html)

## 32.Extend with JavaScript - Logging
# Extend with JavaScript - Logging
$app.logger() could be used to writes any logs into the database so that they can be later
explored from the PocketBase Admin UI &gt; Logs section.
For better performance and to minimize blocking on hot paths, note that logs are written with
debounce and on batches:
-3 seconds after the last debounced log write
-when the batch threshold is reached (currently 200)
-right before app termination to attempt saving everything from the existing logs queue
### Logger methods
All standard
slog.Logger
methods are available but below is a list with some of the most notable ones. Note that attributes are represented
as key-value pair arguments.
##### debug(message, attrs...)
$app.logger().debug("Debug message!")
$app.logger().debug(
"Debug message with attributes!",
"name", "John Doe",
"id", 123,
##### info(message, attrs...)
$app.logger().info("Info message!")
$app.logger().info(
"Info message with attributes!",
"name", "John Doe",
"id", 123,
##### warn(message, attrs...)
$app.logger().warn("Warning message!")
$app.logger().warn(
"Warning message with attributes!",
"name", "John Doe",
"id", 123,
##### error(message, attrs...)
$app.logger().error("Error message!")
$app.logger().error(
"Error message with attributes!",
"id", 123,
"error", err,
##### with(attrs...)
with(atrs...) creates a new local logger that will &quot;inject&quot; the specified attributes with each
following log.
const l = $app.logger().with("total", 123)
// results in log with data {"total": 123}
l.info("message A")
// results in log with data {"total": 123, "name": "john"}
l.info("message B", "name", "john")
##### withGroup(name)
withGroup(name) creates a new local logger that wraps all logs attributes under the specified
group name.
const l = $app.logger().withGroup("sub")
// results in log with data {"sub": { "total": 123 }}
l.info("message A", "total", 123)
### Logs settings
You can control various log settings like logs retention period, minimal log level, request IP logging,
etc. from the logs settings panel:

## 33.pb-ext - Enhanced PocketBase Server
# pb-ext
Enhanced PocketBase server with monitoring, logging & API docs.
<img width="3840" height="2160" alt="pb-ext" src="https://github.com/user-attachments/assets/af360704-c3d6-4d1f-9b49-80229d6570d2" />
<img width="1920" height="2153" alt="Screenshot_2026-02-10_14-42-37" src="https://github.com/user-attachments/assets/d74cf16e-7b5a-4bd0-9f73-1ea81b8c175c" />
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/magooney-loon/pb-ext)
## Core Features
- **API Schema**: Auto-generates OpenAPI docs UI for your endpoints
- **Cron Tracking**: Logs and monitors scheduled cron jobs
- **System Monitoring**: Real-time CPU, memory, disk, network, and runtime metrics
- **Structured Logging**: Complete logging with error tracking and request tracing
- **Visitor Analytics**: Track visitor stats, page views, device types, and browsers
- **PocketBase Integration**: Uses PocketBase's auth system and styling
## Access
- Admin panel:
```bash
127.0.0.1:8090/_
```
- pb-ext dashboard:
```bash
127.0.0.1:8090/_/_
```
## Quick Start
> 🆕 New to Golang and/or PocketBase? [Read this beginner tutorial](TUTORIAL.md).
```go
package main
import (
"flag"
"log"
app "github.com/magooney-loon/pb-ext/core"
"github.com/pocketbase/pocketbase/core"
func main() {
devMode := flag.Bool("dev", false, "Run in developer mode")
flag.Parse()
initApp(*devMode)
func initApp(devMode bool) {
var opts []app.Option
if devMode {
opts = append(opts, app.InDeveloperMode())
} else {
opts = append(opts, app.InNormalMode())
srv := app.New(opts...)
app.SetupLogging(srv)
registerCollections(srv.App())
registerRoutes(srv.App())
registerJobs(srv.App())
srv.App().OnServe().BindFunc(func(e *core.ServeEvent) error {
app.SetupRecovery(srv.App(), e)
if err := srv.Start(); err != nil {
srv.App().Logger().Error("Fatal application error",
"error", err,
"uptime", srv.Stats().StartTime,
"total_requests", srv.Stats().TotalRequests.Load(),
"active_connections", srv.Stats().ActiveConnections.Load(),
"last_request_time", srv.Stats().LastRequestTime.Load(),
log.Fatal(err)
// Example models in cmd/server/collections.go
// Example routes in cmd/server/routes.go
// Example handlers in cmd/server/handlers.go
// Example cron jobs in cmd/server/jobs.go
// You can restructure Your project as You wish,
// just keep this main.go in cmd/server/main.go
// Need a pre-built Svelte5Kit starter template?
// https://github.com/magooney-loon/svelte-gui
// Ready for a production build deployment?
// https://github.com/magooney-loon/pb-deployer
```
```bash
go mod tidy
go run cmd/scripts/main.go --run-only
```
See `**/*/README.md` for detailed docs.
Having issues with Your API Docs?
```bash
127.0.0.1:8090/api/docs/debug/ast
```

## 34.pb-ext - Scripts Documentation
## [COMMAND SEQUENCES]
### > Standard Development Mode
```
$ go run cmd/scripts/main.go
```
*Builds frontend + starts development server*
### > Full System Installation
```
$ go run cmd/scripts/main.go --install
```
### > Frontend Compilation Only
```
$ go run cmd/scripts/main.go --build-only
```
*Compiles frontend assets without server daemon*
### > Development Server Only
```
$ go run cmd/scripts/main.go --run-only
```
*Starts server daemon, skips build sequence*
### > Production Deployment Build
```
$ go run cmd/scripts/main.go --production
```
*Creates optimized production binary + assets*
### > Test Suite Execution
```
$ go run cmd/scripts/main.go --test-only
```
*Runs comprehensive test suite with coverage reports*
### > Custom Output Directory
```
$ go run cmd/scripts/main.go --production --dist release
```
*Production build with custom target directory*
### > System Help Terminal
```
$ go run cmd/scripts/main.go --help
```
*Displays all available command flags and options*
## [DEPLOYMENT INTEGRATION]
### Automated VPS Deployment via pb-deployer:
```
$ git clone https://github.com/magooney-loon/pb-deployer
$ cd pb-deployer && go run cmd/scripts/main.go --install
```
### pb-deployer Features:
[✓] Automated server provisioning + security hardening
[✓] Zero-downtime deployment cycles with rollback
[✓] Production systemd service management
[✓] Full PocketBase v0.20+ compatibility
## [SYSTEM REQUIREMENTS]
[REQUIRED]
├── Go 1.19+        (backend compilation)
├── Node.js 16+     (frontend build system)
├── npm 8+          (dependency management)
└── Git             (version control)
[OPTIONAL]
└── pb-deployer     (production deployment automation)
## [BUILD PROCESS]
[DEVELOPMENT MODE]
1. System validation    → Check Go/Node/npm availability
2. Dependency install   → npm install + go mod tidy
3. Frontend build       → npm run build
4. Asset deployment     → Copy to pb_public/
5. Server startup       → go run ./cmd/server --dev serve
[PRODUCTION MODE]
1. Environment prep     → Clean dist/ directory
2. Dependency install   → Full dependency resolution
3. Frontend build       → Optimized production build
4. Server compilation   → go build -ldflags="-s -w"
5. Asset packaging      → Create deployment archive
6. Metadata generation  → Build info + package metadata
[TEST MODE]
1. System validation    → Verify test environment
2. Test execution       → Run all test suites
3. Coverage analysis    → Generate coverage reports
4. Report generation    → HTML/JSON/TXT outputs
## [TROUBLESHOOTING]
[ERROR: Command not found]
→ Ensure Go/Node/npm are installed and in system PATH
[ERROR: Frontend build failed]
→ Check package.json and run 'npm install' manually
→ Verify frontend/ directory exists with valid source
[ERROR: Server compilation failed]
→ Run 'go mod tidy' to resolve dependencies
→ Check cmd/server/main.go exists
[ERROR: Permission denied]
→ Ensure write permissions for pb_public/ and dist/

## 35.pb-ext - Collections Implementation
```go
package main
// Collection example
import (
"github.com/pocketbase/pocketbase/core"
// registerCollections sets up all database collections for the application
func registerCollections(app core.App) {
app.OnServe().BindFunc(func(e *core.ServeEvent) error {
if existingCollection != nil {
return nil
// Find users collection for optional relation (v2 auth)
usersCollection, err := app.FindCollectionByNameOrId("users")
if err != nil {
return err
// Add optional user relation
collection.Fields.Add(&core.RelationField{
Name:          "user",
Required:      false, // Optional - v1 routes won't use this
CollectionId:  usersCollection.Id,
CascadeDelete: true,
// Add title field (required)
collection.Fields.Add(&core.TextField{
Name:     "title",
Required: true,
Max:      200,
// Add description field (optional)
collection.Fields.Add(&core.TextField{
Name:     "description",
Required: false,
Max:      1000,
// Add completed field (boolean, default false)
collection.Fields.Add(&core.BoolField{
Name: "completed",
// Add priority field (select)
collection.Fields.Add(&core.SelectField{
Name:   "priority",
Values: []string{"low", "medium", "high"},
// Add auto-date fields
collection.Fields.Add(&core.AutodateField{
Name:     "created",
OnCreate: true,
collection.Fields.Add(&core.AutodateField{
Name:     "updated",
OnCreate: true,
OnUpdate: true,
// Set collection rules - public access for v1
collection.ViewRule = nil   // Public read
collection.CreateRule = nil // Public create
collection.UpdateRule = nil // Public update
collection.DeleteRule = nil // Public delete
// Add indexes
// Save the collection
if err := app.Save(collection); err != nil {
return err
return nil
```

## 36.pb-ext - Handlers Implementation
```go
package main
// API_SOURCE
import (
"encoding/json"
"net/http"
"strconv"
"time"
"github.com/pocketbase/pocketbase/core"
// Request types
Title       string `json:"title"`
Description string `json:"description,omitempty"`
Priority    string `json:"priority,omitempty"` // low, medium, high
Completed   bool   `json:"completed"`
Title       *string `json:"title,omitempty"`
Description *string `json:"description,omitempty"`
Priority    *string `json:"priority,omitempty"`
Completed   *bool   `json:"completed,omitempty"`
// API_DESC Get current server time in multiple formats
// API_TAGS Utility
func timeHandler(c *core.RequestEvent) error {
now := time.Now()
return c.JSON(http.StatusOK, map[string]any{
"time": map[string]string{
"iso":       now.Format(time.RFC3339),
"unix":      strconv.FormatInt(now.Unix(), 10),
"unix_nano": strconv.FormatInt(now.UnixNano(), 10),
"utc":       now.UTC().Format(time.RFC3339),
"server":  "pb-ext",
"version": "1.0.0",
// Check authentication - required for creation
if c.Auth == nil {
return c.JSON(http.StatusUnauthorized, map[string]any{"error": "Authentication required"})
if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
return c.JSON(http.StatusBadRequest, map[string]any{"error": "Invalid JSON payload"})
if req.Title == "" {
return c.JSON(http.StatusBadRequest, map[string]any{"error": "Title is required"})
// Validate priority if provided
if req.Priority != "" && req.Priority != "low" && req.Priority != "medium" && req.Priority != "high" {
return c.JSON(http.StatusBadRequest, map[string]any{"error": "Priority must be 'low', 'medium', or 'high'"})
// Default priority to medium if not provided
if req.Priority == "" {
req.Priority = "medium"
"title":       req.Title,
"description": req.Description,
"priority":    req.Priority,
"completed":   req.Completed,
// Only set user field if authenticated user is from users collection
// Superusers/admins don't have records in users collection
if c.Auth.Collection().Name == "users" {
// If authenticated as superuser, leave user field empty or handle differently
if err != nil {
return c.JSON(http.StatusInternalServerError, map[string]any{"error": "Collection not found"})
record := core.NewRecord(collection)
if err := c.App.Save(record); err != nil {
return c.JSON(http.StatusCreated, map[string]any{
"id":          record.Id,
"title":       record.GetString("title"),
"description": record.GetString("description"),
"priority":    record.GetString("priority"),
"completed":   record.GetBool("completed"),
"created_at":  record.GetDateTime("created"),
"user_id":     record.GetString("user"),
"created_by":  c.Auth.Collection().Name, // Show if created by user or admin
if err != nil {
return c.JSON(http.StatusInternalServerError, map[string]any{"error": "Collection not found"})
// Build query with optional filters
filter := ""
filterParams := make(map[string]any)
// Filter by completion status if provided
if completed := c.Request.URL.Query().Get("completed"); completed != "" {
if completed == "true" || completed == "1" {
filter = "completed = true"
} else if completed == "false" || completed == "0" {
filter = "completed = false"
// Filter by priority if provided
if priority := c.Request.URL.Query().Get("priority"); priority != "" {
if filter != "" {
filter += " && "
filter += "priority = {:priority}"
filterParams["priority"] = priority
// For authenticated requests, filter by user (only if user is from users collection)
if c.Auth != nil && c.Auth.Collection().Name == "users" {
if filter != "" {
filter += " && "
filter += "user = {:userId}"
filterParams["userId"] = c.Auth.Id
records, err := c.App.FindRecordsByFilter(collection, filter, "-created", 100, 0, filterParams)
if err != nil {
for i, record := range records {
"id":          record.Id,
"title":       record.GetString("title"),
"description": record.GetString("description"),
"priority":    record.GetString("priority"),
"completed":   record.GetBool("completed"),
"created_at":  record.GetDateTime("created"),
"updated_at":  record.GetDateTime("updated"),
// Include user info if available
if userId := record.GetString("user"); userId != "" {
return c.JSON(http.StatusOK, map[string]any{
"filters": map[string]any{
"completed": c.Request.URL.Query().Get("completed"),
"priority":  c.Request.URL.Query().Get("priority"),
if err != nil {
return c.JSON(http.StatusInternalServerError, map[string]any{"error": "Collection not found"})
if err != nil {
// For authenticated requests, check ownership (only enforce for regular users)
if c.Auth != nil && c.Auth.Collection().Name == "users" {
if userID := record.GetString("user"); userID != "" && userID != c.Auth.Id {
return c.JSON(http.StatusForbidden, map[string]any{"error": "Access denied"})
return c.JSON(http.StatusOK, map[string]any{
"id":          record.Id,
"title":       record.GetString("title"),
"description": record.GetString("description"),
"priority":    record.GetString("priority"),
"completed":   record.GetBool("completed"),
"created_at":  record.GetDateTime("created"),
"updated_at":  record.GetDateTime("updated"),
"user_id":     record.GetString("user"),
// Check authentication - required for updates
if c.Auth == nil {
return c.JSON(http.StatusUnauthorized, map[string]any{"error": "Authentication required"})
if err != nil {
return c.JSON(http.StatusInternalServerError, map[string]any{"error": "Collection not found"})
if err != nil {
if c.Auth.Collection().Name == "users" {
if userID := record.GetString("user"); userID != c.Auth.Id {
if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
return c.JSON(http.StatusBadRequest, map[string]any{"error": "Invalid JSON payload"})
// Apply updates
updates := make(map[string]any)
if req.Title != nil {
if *req.Title == "" {
return c.JSON(http.StatusBadRequest, map[string]any{"error": "Title cannot be empty"})
record.Set("title", *req.Title)
updates["title"] = *req.Title
if req.Description != nil {
record.Set("description", *req.Description)
updates["description"] = *req.Description
if req.Priority != nil {
if *req.Priority != "low" && *req.Priority != "medium" && *req.Priority != "high" {
return c.JSON(http.StatusBadRequest, map[string]any{"error": "Priority must be 'low', 'medium', or 'high'"})
record.Set("priority", *req.Priority)
updates["priority"] = *req.Priority
if req.Completed != nil {
record.Set("completed", *req.Completed)
updates["completed"] = *req.Completed
if err := c.App.Save(record); err != nil {
return c.JSON(http.StatusOK, map[string]any{
"id":          record.Id,
"title":       record.GetString("title"),
"description": record.GetString("description"),
"priority":    record.GetString("priority"),
"completed":   record.GetBool("completed"),
"updated_at":  record.GetDateTime("updated"),
"updates": updates,
// Check authentication - required for deletion
if c.Auth == nil {
return c.JSON(http.StatusUnauthorized, map[string]any{"error": "Authentication required"})
if err != nil {
return c.JSON(http.StatusInternalServerError, map[string]any{"error": "Collection not found"})
if err != nil {
if c.Auth.Collection().Name == "users" {
if userID := record.GetString("user"); userID != c.Auth.Id {
if err := c.App.Delete(record); err != nil {
return c.JSON(http.StatusOK, map[string]any{
"title": record.GetString("title"),
"deleted_at": time.Now().Format(time.RFC3339),
```

## 37.pb-ext - Jobs Implementation
```go
package main
import (
"fmt"
"time"
"github.com/magooney-loon/pb-ext/core/server"
"github.com/pocketbase/pocketbase/core"
func registerJobs(app core.App) {
app.OnServe().BindFunc(func(e *core.ServeEvent) error {
// Register example cron jobs
if err := helloJob(app); err != nil {
app.Logger().Error("Failed to register hello job", "error", err)
return err
if err := dailyCleanupJob(app); err != nil {
app.Logger().Error("Failed to register daily cleanup job", "error", err)
return err
if err := weeklyStatsJob(app); err != nil {
app.Logger().Error("Failed to register weekly stats job", "error", err)
return err
app.Logger().Info("All cron jobs registered successfully")
func helloJob(app core.App) error {
jobManager := server.GetJobManager()
if jobManager == nil {
return fmt.Errorf("job manager not initialized")
return jobManager.RegisterJob("helloWorld", "Hello World Job",
"A simple demonstration job that runs every 5 minutes, outputs timestamped hello messages and simulates basic task processing",
"*/5 * * * *", func(jobLogger *server.JobExecutionLogger) {
jobLogger.Start("Hello World Job")
jobLogger.Info("Current time: %s", time.Now().Format("2006-01-02 15:04:05"))
jobLogger.Progress("Processing hello world task...")
// Simulate some work
time.Sleep(100 * time.Millisecond)
jobLogger.Success("Hello from cron job! Task completed successfully.")
jobLogger.Complete(fmt.Sprintf("Job finished at: %s", time.Now().Format("2006-01-02 15:04:05")))
func dailyCleanupJob(app core.App) error {
jobManager := server.GetJobManager()
if jobManager == nil {
return fmt.Errorf("job manager not initialized")
return jobManager.RegisterJob("dailyCleanup", "Daily Cleanup Job",
"0 2 * * *", func(jobLogger *server.JobExecutionLogger) {
jobLogger.Start("Daily Cleanup Job")
jobLogger.Info("Cleanup job started at: %s", time.Now().Format("2006-01-02 15:04:05"))
app.Logger().Info("Running daily cleanup job", "time", time.Now())
if err != nil {
jobLogger.Fail(err)
return
cutoffDate := time.Now().AddDate(0, 0, -30)
filter := "completed = true && created < {:cutoff}"
records, err := app.FindRecordsByFilter(collection, filter, "", 100, 0, map[string]any{
"cutoff": cutoffDate.Format("2006-01-02 15:04:05.000Z"),
if err != nil {
jobLogger.Fail(err)
return
deletedCount := 0
for _, record := range records {
if err := app.Delete(record); err != nil {
} else {
deletedCount++
jobLogger.Statistics(map[string]interface{}{
"total_found": len(records),
"deleted":     deletedCount,
"failed":      len(records) - deletedCount,
jobLogger.Complete(fmt.Sprintf("Deleted %d/%d records", deletedCount, len(records)))
app.Logger().Info("Daily cleanup completed", "deleted_records", deletedCount)
func weeklyStatsJob(app core.App) error {
jobManager := server.GetJobManager()
if jobManager == nil {
return fmt.Errorf("job manager not initialized")
return jobManager.RegisterJob("weeklyStats", "Weekly Statistics Job",
"0 0 * * 0", func(jobLogger *server.JobExecutionLogger) {
jobLogger.Start("Weekly Statistics Job")
jobLogger.Info("Generating weekly report for week ending: %s", time.Now().Format("2006-01-02"))
app.Logger().Info("Generating weekly statistics", "time", time.Now())
if err != nil {
jobLogger.Fail(err)
return
weekAgo := time.Now().AddDate(0, 0, -7)
filter := "created >= {:week_ago}"
records, err := app.FindRecordsByFilter(collection, filter, "", 1000, 0, map[string]any{
"week_ago": weekAgo.Format("2006-01-02 15:04:05.000Z"),
if err != nil {
jobLogger.Fail(err)
return
completed := 0
pending := 0
for _, record := range records {
if record.GetBool("completed") {
completed++
} else {
pending++
completionRate := float64(0)
if len(records) > 0 {
completionRate = float64(completed) / float64(len(records)) * 100
// Log statistics using the structured method
stats := map[string]interface{}{
"Completion rate":     fmt.Sprintf("%.1f%%", completionRate),
jobLogger.Info("WEEKLY STATISTICS REPORT")
jobLogger.Statistics(stats)
jobLogger.Complete("Weekly statistics report generated successfully")
app.Logger().Info("Weekly statistics generated",
"completion_rate", completionRate,
```

## 38.pb-ext - Routes Implementation
```go
package main
// API_SOURCE
import (
"github.com/magooney-loon/pb-ext/core/server/api"
"github.com/pocketbase/pocketbase/apis"
"github.com/pocketbase/pocketbase/core"
func registerRoutes(app core.App) {
// Initialize version manager with configs
versionManager := api.InitializeVersionedSystem(createAPIVersions(), "v1")
app.OnServe().BindFunc(func(e *core.ServeEvent) error {
// Get version-specific routers
v1Router, err := versionManager.GetVersionRouter("v1", e)
if err != nil {
return err
v2Router, err := versionManager.GetVersionRouter("v2", e)
if err != nil {
return err
// Register v1 routes
registerV1Routes(v1Router)
// Register v2 routes
registerV2Routes(v2Router)
// Register version management endpoints
versionManager.RegisterWithServer(app)
// createAPIVersions creates version configurations with reduced duplication
func createAPIVersions() map[string]*api.APIDocsConfig {
baseConfig := &api.APIDocsConfig{
Title:       "pb-ext demo api",
Description: "Hello world",
BaseURL:     "http://127.0.0.1:8090/",
Enabled:     true,
ContactName:  "pb-ext Team",
ContactEmail: "contact@magooney.org",
ContactURL:   "https://github.com/magooney-loon/pb-ext",
LicenseName: "MIT",
LicenseURL:  "https://opensource.org/licenses/MIT",
ExternalDocsURL:  "https://github.com/magooney-loon/pb-ext",
ExternalDocsDesc: "pb-ext documentation",
PublicSwagger: true,
// Create v1 config
v1Config := *baseConfig
v1Config.Version = "1.0.0"
v1Config.Status = "stable"
// Create v2 config
v2Config := *baseConfig
v2Config.Version = "2.0.0"
v2Config.Status = "testing"
return map[string]*api.APIDocsConfig{
"v1": &v1Config,
"v2": &v2Config,
// registerV1Routes registers all v1 API routes
func registerV1Routes(router *api.VersionedAPIRouter) {
// Option 1: Manual route registration (explicit control)
prefix := "/api/v1"
// Option 2: CRUD convenience method (less boilerplate)
// Uncomment to use instead of manual registration above:
// v1 := router.SetPrefix("/api/v1")
// }, apis.RequireAuth()) // Auth applied to Create, Update, Patch, Delete
// registerV2Routes registers all v2 API routes
func registerV2Routes(router *api.VersionedAPIRouter) {
// Using prefixed router for cleaner code
v2 := router.SetPrefix("/api/v2")
// Utility routes (no auth required)
v2.GET("/time", timeHandler)
```

## 39.pb-deployer - PocketBase Production Deployment
<div align="center">
<img src="frontend/static/favicon.svg" alt="Logo" width="200">
<h1 align="center">pb-deployer</h1>
<h3 align="center">Automates the lifecycle of deploying PocketBase apps to production</h3>
<a href="https://github.com/magooney-loon/pb-deployer/stargazers"><img src="https://img.shields.io/github/stars/magooney-loon/pb-deployer?style=for-the-badge&color=blue" alt="Stargazers"></a>
<a href="https://github.com/magooney-loon/pb-deployer/graphs/contributors"><img src="https://img.shields.io/github/contributors/magooney-loon/pb-deployer?style=for-the-badge&color=blue" alt="Contributors"></a>
<a href="https://github.com/magooney-loon/pb-deployer/blob/main/LICENSE"><img src="https://img.shields.io/github/license/magooney-loon/pb-deployer?style=for-the-badge&color=blue" alt="AGPL-3.0"></a>
<br>
<img src="frontend/static/deployer.png" alt="Screenshot" width="100%">
<h5 align="center">**WARNING**HOBBY PROJECT**</h5>
</div>
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/magooney-loon/pb-deployer)
## 🚀 Quick Start
```bash
git clone https://github.com/magooney-loon/pb-deployer
cd pb-deployer
go run cmd/scripts/main.go --install
```
## Core Workflow
1. **Server Registration**: Add remote host connection details
2. **Server Setup**: Automated user creation and directory structure
3. **Security Lockdown**: Firewall, fail2ban, disable root SSH (Optional)
4. **App Deployment**: Upload prod dist, systemd service creation
5. **Version Management**: Rollback support with file storage
## Directory Structure
```
/opt/pocketbase/
├── apps/           # Application deployments (per app directory)
├── backups/        # Deployment backups (timestamped)
├── logs/           # Application logs
└── staging/        # Temporary staging during deployments
```
## Deployment Steps
2. **Checking service status**
3. **Stopping existing service**
4. **Creating backup of current deployment**
5. **Preparing deployment directory**
6. **Installing new version**
7. **Creating/updating systemd service**
8. **Creating superuser (if initial deployment)**
9. **Starting service**
10. **Verifying & finalizing deployment**
<div align="center">
<img src="frontend/static/deployer2.png" alt="Logo" width="100%">
</div>
See `**/*/README.md` for detailed docs.
Make sure you loaded your SSH keys, check with `ssh-add -l`
## Contribution
