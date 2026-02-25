# POCKETBASE DOCS|2026-02-25|1 sections

## 1.Files upload and handling
### Uploading files
To upload files, you must first add a file field to your collection:
Once added, you could create/update a Record and upload &quot;documents&quot; files by sending a
multipart/form-data request using the Records create/update APIs.
The client SDK makes things a little easier and auto detect the request content type based on the
parameters that you provide. Here is an example how to create a new Record and upload multiple files to
the example file field &quot;documents&quot; using the SDKs:
Dart
// Example HTML:
// &lt;input type="file" id="fileInput" />
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const formData = new FormData();
const fileInput = document.getElementById('fileInput');
// listen to file input changes and add the selected files to the form data
fileInput.addEventListener('change', function () {
for (let file of fileInput.files) {
formData.append('documents', file);
// set some other regular text field value
formData.append('title', 'Hello world!');
// upload and create new record
const createdRecord = await pb.collection('example').create(formData);
import 'package:pocketbase/pocketbase.dart';
import 'package:http/http.dart' as http;
final pb = PocketBase('http://127.0.0.1:8090');
// create a new record and upload multiple files
final record = await pb.collection('example').create(
body: {
'title': 'Hello world!', // some regular text field
files: [
http.MultipartFile.fromString(
'document',
'example content 1...',
filename: 'file1.txt',
http.MultipartFile.fromString(
'document',
'example content 2...',
filename: 'file2.txt',
Each uploaded file will be stored with the original filename (sanitized) and suffixed with a
random (10 chars) part (eg. test_52iWbGinWd.png).
When you upload a new file to a single file upload field (aka.
file (if any) and upload the new one in its place.
When you upload a new file to a multiple file upload field (aka.
Max Files option is &gt; 1) the new file will be appended to the existing field
values (as long as the Max Files limit is not reached, otherwise an error will be
thrown).
### Deleting files
To delete uploaded file(s), you could either edit the Record from the admin UI, or use the API and set the
file field to a zero-value  (null, [], empty string, etc.).
If you want to delete individual file(s) from a multiple file upload field, you could
suffix the field name with - and specify the filename(s) you want to delete. Here are some examples
using the SDKs:
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
// delete all "documents" files
await pb.collection('example').update('RECORD_ID', {
'documents': null,
// delete individual files
await pb.collection('example').update('RECORD_ID', {
'documents-': ["file1.pdf", "file2.txt"],
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
// delete all "documents" files
await pb.collection('example').update('RECORD_ID', body: {
'documents': null,
// delete individual files
await pb.collection('example').update('RECORD_ID', body: {
'documents-': ["file1.pdf", "file2.txt"],
The above examples use the JSON object data format, but you could also use FormData instance
for multipart/form-data requests. If using
FormData set the file field to an empty string.
### File URL
Each uploaded file could be accessed by requesting its file url:
http://127.0.0.1:8090/api/files/COLLECTION_ID_OR_NAME/RECORD_ID/FILENAME
If your file field has the Thumb sizes option, you can get a thumb of the image file
(currently limited to jpg, png, and partially gif – its first frame) by adding the thumb
query parameter to the url like this:
http://127.0.0.1:8090/api/files/COLLECTION_ID_OR_NAME/RECORD_ID/FILENAME?thumb=100x300
The following thumb formats are currently supported:
-WxH
(eg. 100x300) - crop to WxH viewbox (from center)
-WxHt
(eg. 100x300t) - crop to WxH viewbox (from top)
-WxHb
(eg. 100x300b) - crop to WxH viewbox (from bottom)
-WxHf
(eg. 100x300f) - fit inside a WxH viewbox (without cropping)
-0xH
(eg. 0x300) - resize to H height preserving the aspect ratio
-Wx0
(eg. 100x0) - resize to W width preserving the aspect ratio
The original file would be returned, if the requested thumb size is not found or the file is not an image!
If you already have a Record model instance, the SDKs provide a convenient method to generate a file url
by its name.
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
const record = await pb.collection('example').getOne('RECORD_ID');
// get only the first filename from "documents"
// note:
// "documents" is an array of filenames because
// the "documents" field was created with "Max Files" option > 1;
// if "Max Files" was 1, then the result property would be just a string
const firstFilename = record.documents[0];
// returns something like:
// http://127.0.0.1:8090/api/files/example/kfzjt5oy8r34hvn/test_52iWbGinWd.png?thumb=100x250
const url = pb.files.getUrl(record, firstFilename, {'thumb': '100x250'});
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
final record = await pb.collection('example').getOne('RECORD_ID');
// get only the first filename from "documents"
// note:
// "documents" is an array of filenames because
// the "documents" field was created with "Max Files" option > 1;
// if "Max Files" was 1, then the result property would be just a string
final firstFilename = record.getListValue&lt;String>('documents')[0];
// returns something like:
// http://127.0.0.1:8090/api/files/example/kfzjt5oy8r34hvn/test_52iWbGinWd.png?thumb=100x250
final url = pb.files.getUrl(record, firstFilename, thumb: '100x250');
### Protected files
By default all files are public accessible if you know their full url.
For most applications this is fine since all files have a random part, but in some cases you may want an
extra security to prevent unauthorized access to sensitive files like ID card or Passport copies,
contracts, etc.
To do this you can mark the file field as Protected and then request the file with a
special short-lived file token.
Only requests that satisfy the View API rule of the record collection will be able
Dart
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
// authenticate
// generate a file token
const fileToken = await pb.files.getToken();
// retrieve an example protected file url (will be valid ~2min)
const record = await pb.collection('example').getOne('RECORD_ID');
const url = pb.files.getUrl(record, record.myPrivateFile, {'token': fileToken});
import 'package:pocketbase/pocketbase.dart';
final pb = PocketBase('http://127.0.0.1:8090');
// authenticate
// generate a file token
final fileToken = await pb.files.getToken();
// retrieve an example protected file url (will be valid ~2min)
final record = await pb.collection('example').getOne('RECORD_ID');
final url = pb.files.getUrl(record, record.getStringValue('myPrivateFile'), token: fileToken);
### Storage options
By default PocketBase uses the local file system to store uploaded files (in the
pb_data/storage directory).
If you have limited disk space, you could use a S3 compatible storage (AWS S3, MinIO, Wasabi, DigitalOcean
Spaces, Vultr Object Storage, etc.). The easiest way to setup the connection settings is from the admin UI
(Settings &gt; Files storage):
