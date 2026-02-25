# POCKETBASE DOCS|2026-02-25|1 sections

## 1.Web APIs reference - API Settings
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
