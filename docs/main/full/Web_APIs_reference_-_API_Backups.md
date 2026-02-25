# POCKETBASE DOCS|2026-02-25|1 sections

## 1.Web APIs reference - API Backups
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
