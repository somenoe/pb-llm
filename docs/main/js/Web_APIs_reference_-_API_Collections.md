# POCKETBASE DOCS|2026-02-25|1 sections

## 1.Web APIs reference - API Collections
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
