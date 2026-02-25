# POCKETBASE DOCS|2026-02-25|1 sections

## 1.Extend with JavaScript - Database
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
