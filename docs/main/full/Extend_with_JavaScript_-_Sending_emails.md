# POCKETBASE DOCS|2026-02-25|1 sections

## 1.Extend with JavaScript - Sending emails
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
