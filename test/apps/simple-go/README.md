blobstache
==========
BlobStache is a warehousing service that logs to stdout and stores pbjects in pluggable backends as well as a plugable metadata database

Usage of blobstache
  -backend="local": Backend data storage
  -backendcredentials="/tmp/data": Backend data credentials
  -dbCredentials="dbname=live sslmode=disable": Connection string for database connection
  -port="8080": Port to listen on


example start
blobstache -backend=local -backendcredentials=/path/to/glustermount -dbCredentials="dbname=live sslmode=disable username=whatever password=whatever host=whatever port=whatever" -port=1234