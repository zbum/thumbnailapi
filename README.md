# thumbnailapi
* Restful thumbnail maker built with golang and ImageMagick
## support format
* heic
* jpeg
* png

## build and run
* build
```shell
$ cd docker && sh build.sh
```
* run
```shell
$ sh run.sh
```
## apis
### convert
* POST /v1/thumbnail/convert
#### request
* Multipart params
```
sizes=300-200 /* cols-rows */
file=@aaa.jpg /* file upload */
```
#### response
* header
```
X-Content-Lengths : 300-200:12345 /* cols-rows:size(byte) */
```
* body
```
/* thumbnail binary */
```