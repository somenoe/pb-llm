# POCKETBASE DOCS|2026-02-25|1 sections

## 1.Web APIs reference - API Admins
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
