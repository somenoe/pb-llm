# POCKETBASE DOCS|2026-02-25|1 sections

## 1.Authentication
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
