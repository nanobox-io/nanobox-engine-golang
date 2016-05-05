package main

import (
	// "net/http"
	"flag"
	"fmt"
)

func main() {
	if len(flag.Args()) < 1 {
		fmt.Println("Thanks for using The Client")
		help()
		return
	}

	switch flag.Arg(0) {
	case "list-repos", "list_repos", "listrepos", "list-users", "list_users", "listusers":
		listUsers()
	case "create-repo", "create_repo", "createrepo", "create-user", "create_user", "createuser":
		createUser(flag.Arg(1))
	case "delete-repo", "delete_repo", "deleterepo", "delete-user", "delete_user", "deleteuser":
		deleteUser(flag.Arg(1))
	case "list-buckets", "list_buckets", "listbuckets":
		listBuckets()
	case "show-bucket", "show_bucket", "showbucket":
		showBucket(flag.Arg(1))
	case "create-bucket", "create_bucket", "createbucket":
		createBucket(flag.Arg(1))
	case "delete-bucket", "delete_bucket", "deletebucket":
		deleteBucket(flag.Arg(1))
	case "list-objects", "list_objects", "listobjects":
		listObjects()
	case "show-object-info", "show_object_info", "showobjectinfo":
		getObjectInfo(flag.Arg(1))
	case "set-object-public", "set_object_public", "setobjectpublic":
		setObjectPublic(flag.Arg(1))
	case "object-size", "object_size", "objectsize":
		showObjectSize(flag.Arg(1))
	case "get-object", "get_object", "getobject":
		getObject(flag.Arg(1))
	case "create-object", "create_object", "createobject":
		createObject(flag.Arg(1))
	case "delete-object", "delete_object", "deleteobject":
		deleteObject(flag.Arg(1))

	case "help", "h":
		help()
	default:
		fmt.Printf("I dont know what to do with %s\n", flag.Arg(0))
		help()
	}
}

func help() {
	fmt.Printf(`
Usage of client:
  Option:
    When using the client two arguments must be present:
      userkey and userid
    All Options can be set using flags or environment variables
    Flag version:
`)
	flag.PrintDefaults()
	fmt.Printf(`
    Environment Variables version:
    (these options set the same as the flag version)
HOST        = host
LOCATION    = host
SERVER      = host
USERKEY     = key
KEY         = key
USERID      = id
ID          = id
BUCKETID    = bucketid
BUCKETNAME  = bucketname
OBJECTID    = objectid
OBJECTALIAS = objectalias

  Tasks:
    client [options] list-users             : List all users (admin only)
    client [options] create-user            : Create New User (admin only)
    client [options] delete-user ID         : Delete User (admin only)
    client [options] list-buckets           : List all buckets you own
    client [options] show-bucket (ID|NAME)  : Show a bucket and its size
    client [options] create-bucket NAME     : Create a new Bucket
    client [options] delete-bucket          : Delete a bucket
    client [options] list-objects           : List all objects in a bucket (requires bucketid)
    client [options] object-size (ID|ALIAS) : Get the objects size (requires bucketid)
    client [options] get-object (ID|ALIAS)  : Get a object form the system. This outputs all of the object to stdout (requires bucketid)
    client [options] create-object ALIAS    : Create a New object and takes stdin as the object data (requires bucketid)
    client [options] delete-object          : Delete an object from the data store. (requires bucketid)

  Examples:
    Create a new Object:
      Environment Variable way:
        ID=<user uuid> KEY=<secret key> BUCKETID=<bucket uuid> cat something.txt | client create-object something.txt

      Flag way:
        cat something.txt | client -id <user uuid> --key=<secret key> --bucketid <bucket uuid> create-object something.txt

    Getting a Objects Contents:
      client -id <user uuid> --key=<secret key> --bucketid <bucket uuid> get-object something.txt > something.txt

`)
}
