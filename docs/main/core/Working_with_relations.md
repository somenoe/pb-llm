# POCKETBASE DOCS|2026-02-25|1 sections

## 1.Working with relations
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
