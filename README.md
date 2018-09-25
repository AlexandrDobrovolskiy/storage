## Golang storage server
### Standart API
#### Images
* https://hostname/images/news Method `POST` multipart form data field `news`
* https://hostname/static/images/news/image-name.ext Method `GET`
#### Files
* https://hostname/files/news Method `POST` multipart form data field `news`
* https://hostname/static/files/news/file-name.ext Method `GET`
### FilePond API
#### Process 
Store file into temprary folder, recomended async requests with one file.
* URL `https://hostname.com/filepond`
* Content type `multipart/form-data`
* Method `POST`
* Server returns unique location id `12345` in `text/plain `response
#### Revert
There’s one way the client can deviate from the previous path and that is by reverting the upload.
* URL `https://hostname.com/filepond`
* Method `DELETE`
* Body: location id (`12345`)
* Server removes temporary folder matching the supplied id and returns an empty response
#### Submit
Submit the form and move files from temprary folder
* URL `https://hostname.com/filepond/submit`
* Method `POST`
* Body: array of unique id's `["12345"]`
___
Restore, Load and Fetch endpoins are in development.