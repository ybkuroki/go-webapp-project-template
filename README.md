# go-webapp-project-template

## Preface
This repository is a template for implementing web application in golang.
The sample of implementation based on this template is [here](https://github.com/ybkuroki/go-webapp-sample).
To use this repository, please click the 'Use this template' button.

## Install
Perform the following steps:
1. Download and install [MinGW(gcc)](https://sourceforge.net/projects/mingw-w64/files/?source=navbar).
1. Download and install [Visual Studio Code(VS Code)](https://code.visualstudio.com/).
1. Download and install [Golang](https://golang.org/).
1. Get the source code of forked this repository.

## Starting Server
Perform the following steps:
1. Starting this web application by the following command.
    ```bash
    go run main.go
    ```
1. When startup is complete, the console shows the following message:
    ```
    http server started on [::]:8080
    ```
1. Access [http://localhost:8080/api/health](http://localhost:8080/api/health) in your browser and confirm that this application has started.
    ```
    healthy
    ```
1. Login with the following username and password.
    - username : ``test``
    - password : ``test``

## Build executable file
Build this source code by the following command.
```bash
go build main.go
```

## Project Structure
This project uses MVC architecture. The summary of each packages is the following.

```
- go-webapp-sample
  + config                  … Define configurations of this system.
  + logger                  … Provide loggers.
  + migration               … Provide database migration service for development.
  + router                  … Define routing.
  + controller              … Define controllers.
  + model                   … Define models.
  + repository              … Provide a service of database access.
  + service                 … Provide a service of book management.
  + session                 … Provide session management.
  + test                    … Provide utilities for unit test.
  - main.go                 … Entry point of this application.
```

I will explain about the specification of some packages in the following.

### Configuration
The config package is the package for defining the configuration settings of this application.
When this application have started up, this package will load `application.yml` defined the settings.
The constructre of `application.yml` defines as struct of golang.
For providing this mechanism, this sample uses the library is called configor.

We can switch many configuration files depending on execution environments.
For switching environments, we have to change the execute parameter of this application.
The defalut value is set `DEV`, so when this application have started up, it will be loaded `application.dev.yml`.

```bash
$ go run main.go -env=prod
```

### Logging
The logger package provides 3 logging functions: Http request logger, SQL logger, Action logger.

Http request logger outputs http request by Echo framework and the log format is defined in `application.yml`.
The following is the sample of log and the configuration settings.

**Request Log Sample**
```
2020-08-09T16:23:14+09:00 [INFO] ::1 GET /api/health 200
```

**application.yml**
```yml
log:
  format: ${time_rfc3339} [${level}] ${remote_ip} ${method} ${uri} ${status}
  level: 1
  file_path: 
```

SQL logger outputs SQL statements created by Gorm. This repository uses the logger that customized default Gorm logger.
The following is the sample of log.

**SQL Log Sample**
```
2020-08-09T16:22:54+09:00 [DEBUG]     [gorm] : CREATE TABLE "account_master" ("id" integer primary key autoincrement,"name" varchar(255),"password" varchar(255),"authority_id" integer )
2020-08-09T16:22:54+09:00 [DEBUG]     [gorm] : CREATE TABLE "authority_master" ("id" integer primary key autoincrement,"name" varchar(255) )
2020-08-09T16:22:54+09:00 [DEBUG]     [gorm] : INSERT INTO "authority_master" ("name") VALUES ('Admin')
2020-08-09T16:22:54+09:00 [DEBUG]     [gorm] : UPDATE "authority_master" SET "name" = 'Admin'  WHERE "authority_master"."id" = 1
2020-08-09T16:22:54+09:00 [DEBUG]     [gorm] : INSERT INTO "account_master" ("name","password","authority_id") VALUES ('test','$2a$10$2NQi6QUVYv/16DB9kQ/fru9sRp.cDVSBpuaMlKOayxGRsmsrmNq5C',1)
```

Action logger outputs the start and end of the controller method that executed by Http request.
This is implemented as echo middleware and I assume that is used it only when debug.
The following is the sample of log.

**Action Log Sample**
```
2020-08-09T16:23:14+09:00 [DEBUG]     /api/health Action Start
2020-08-09T16:23:14+09:00 [DEBUG]     /api/health Action End
```


### Data Migration
The migration package provides 2 functions: the initialize of database, the creation of master data.
Those functions exists for debugging and don't use in the production environment.

The initialize of database is written in `/migraiton/db_generator.go`.
It will delete the existing tables and will create new tables that is corresponded to models.

The creation of master data is written in `/migraiton/master_generator.go`.
It will create the master data needed to run this application such as the account data and the authority data.

### Session
The session package provides the session management function.
This package will authorize a user, and we can get the session of login user or set a data on the user session by using this package.

In this session package, the authoritzation is implemented by using the middleware of echo framework.
This package will judge whether a user have role performs a accessed request.
If a user have role that performs a accessed request, this package will perform a method of a controller.
Otherwise, this package will return a authoritzation error.

## Implemented Services
We have already implemented a account management service in this template repository.
There are 4 functions of the following in the account management service:

### Account Management
|Service Name|HTTP Method|URL|Parameter|Summary|
|:---|:---:|:---|:---|:---|
|Login Service|Form POST|``/api/account/login``|Session ID, User Name, Password|Session authentication with username and password.|
|Logout Service|POST|``/api/account/logout``|Session ID|Logout a user.|
|Login Status Check Service|GET|``/api/account/loginStatus``|Session ID|Check if the user is logged in.|
|Login Username Service|GET|``/api/account/loginAccount``|Session ID|Get the login user's username.|

Those functions is implemented in the account controller and the account service.

The login function performs the session authentication by using a given user name and a password. 
This function will compare a given password by a user and a encrypted password stored in the database.
If it have equaled, this function will return authentication success and a session id managed by this application.
Otherwise, it will return authentication failure.

The logout function disables the current session id having a login user.
A client cannot use the disable session.

The login status function returns whether a user is login or not.
If the user have logged in, this function will return true. Otherwise, this function will return false.

## Libraries
This sample uses the following libraries.

|Library Name|Version|
|:---|:---:|
|Echo|4.1.16|
|Gorm|1.9.12|
|go-playground/validator.v9|9.31.0|

## License
The License of this sample is *MIT License*.
