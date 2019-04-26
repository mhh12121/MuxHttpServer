### This is a Simple Http Server


#### Environment
1. Ubuntu 16.04
2. Golang 1.11
3. MySQL driver
```"github.com/go-sql-driver/mysql"```
4. Gorilla Mux
```"github.com/gorilla/mux"```

#### Step
1. Clone down this repo, then go into direcotry ```/server/```
2. run ```go run httpServer.go```

#### File Structure
```
    |--conf //save os,filepath,host address configuration
    |--dao //save mysql DAO
    |--data //save basic data structure
    |--server //main.go
    |--static
        |-- img
        |-- view //save templates
    |--test //test 
    |--util // common constant like errorcode and so on

    authenticationTable.sql, simpleUserTable.sql, userTable.sql 
    contain creating test Table and inserting test data

```
#### API Doc
1.  Route: ```/v{version}/{action}```
    Methods:GET
    Request:{}
    Header:
        Content-type: application/x-www-form-urlencoded
    
2.  Route: ```/{namespace}/{resource}/{action}```
    Methods:GET

3.  Route: ```/user/self/{option} ```


