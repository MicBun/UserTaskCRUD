# UserSimpleCRUD
a Simple CRUD for User and User Login

PreInstall
1. Docker https://docs.docker.com/desktop/install/windows-install/

How to use
1. Clone this Repository using `git clone`
2. Run Docker Desktop
3. use CMD Command on this project root 
   1. `docker-compose build`
   2. `docker-compose up`
      1. The first `docker-compose up` might be failed so stop it by `Ctrl + C`
      * run `docker-compose up` again and you are good to go
4. access http://localhost:8080/swagger/index.html#/
5. The app is ready to use.
---
# Database
Database included and can be accessed at http://localhost:5003

---
# Screenshot
Screenshot can be seen at screenshot folder

---
# Structure Folder
```
root
-config // Contain Db Configuration file
-controllers // Contains Controllers
-docker // Contain app.dockerfile for dockerise
-docs // auto-generated file by Swagger
-middlewares // middlewares for auth
-models // Contains Models for database structure
-routes // Contains Routes
-screenshot // Contains some Screenshot
-utils // Conatins utility file for multiple uses
.env // an env file
.docker-compose.yml // docker related file
go.mod // go mod
main.go // the main.go
README.md // this Readme file
```
---
# Credential
1. phpmyadmin
```
username: root
password: root
```
2. application's admin 
```
username: admin
password: admin
```