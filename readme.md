# tbox
    
    Model generation tool, which mainly generates mysql table 
    as golang type, avoids the inconvenience caused by manual
    operation, and facilitates project maintenance and expansion.

# usage
    
    $ go get -v github.com/daheige/tbox/cmd/tbox
    
    go 1.16.x version
    $ go install github.com/daheige/tbox/cmd/tbox
    
    $ tbox -h
    Usage of tbox:
        -d string
            pkg dir path,eg:-d=./model (default "./model")
        -dsn string
            mysql dsn,eg:-dsn="root:root1234@tcp(127.0.0.1:3306)/test?charset=utf8mb4"
        -j    whether add json tag eg:-j=true
        -m    whether add TableName func eg:-m=true
        -p string
            pkg name,eg:-p=model (default "model")
        -t string
            table,eg:-u=user;order
        -tag string
            tag key,eg:-tag=db (default "db")
        -u    whether uc first only,eg:-u=true (default true)
        -v    whether output cmd,eg:-v=true

    take tbox_demo as an example:
    $ cd ~
    $ mkdir tbox_demo
    $ cd tbox_demo
    $ tbox -dsn="root:root1234@tcp(127.0.0.1:3306)/test?charset=utf8mb4"
    2021/12/25 19:36:18 generating code success
