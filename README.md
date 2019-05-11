# mini-db
A very simple in-memory key-value store like Redis.

I need to save data in the memory
And also in the disk

so, for disk..
I can create files as the key and content as value
and I can read from the filesystem for when application starts

All the methods,
GET
SET
DELETE
FLUSH
MGET
MSET

Supported data types,
String
Binary
Numbers
<!-- Null -->
Arrays
Dictionaries


# Flow...

> CMD = "GET" or Client.get()
> Read from Map
> Return Data from Server

> CMD = "SET"
> Write to Map
> Write to FS
> Return Stored Data