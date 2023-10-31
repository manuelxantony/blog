# blog
Simple blog server written in golang and mysql

# Steps to run the server
* make sure mysql is running.
* In mysql run "CREATE DATABASE blog' for creating the database.
* Then run the query.sql file in the repo, command: mysql -u root -p < query.sql this will create table 
* For running the server type: go run *.go
* I have attached a postman file to test the apis, available in postman folder.
* By default server is running on port 8000 and it can be changed in config.yml file.


# APIS
* /blog/create  -> POST request to create blogs
* /blog/showall  -> GET request to read all blog posts
* /blog/showbyid?id=<number> -> GET request to read a blog post by id
* /blog/update -> PUT request to update a blog post
* /blog/delete -> DELETE request to delete a blog post

# Config file
* config.yml file can be use for setting the server and database.

# Things to do/pending..
*  No documentation is done, because it is a simple curd app and function are straight forward to understand.
*  No unit tests are written. TODO.
*  No session  implemented.
*  No user profile added.
*  No admin added.

