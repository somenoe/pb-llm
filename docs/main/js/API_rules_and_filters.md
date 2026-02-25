# POCKETBASE DOCS|2026-02-25|1 sections

## 1.API rules and filters
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
