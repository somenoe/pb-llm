# POCKETBASE DOCS|2026-02-25|1 sections

## 1.Collections
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
