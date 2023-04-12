# tbox
    
    Model generation tool, which mainly generates mysql table 
    as golang type, avoids the inconvenience caused by manual
    operation, and facilitates project maintenance and expansion.

# usage
    
    $ go get -v github.com/daheige/tbox/cmd/tbox
    
    go 1.16.x version
    $ go install github.com/daheige/tbox/cmd/tbox@latest
    
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
            table,eg:-t=user;order
        -tag string
            tag key,eg:-tag=db (default "db")
        -u    whether uc first only,eg:-u=true (default false)
        -v    whether output cmd,eg:-v=true
        -n    whether all field no null eg:-n=true

    take tbox_demo as an example:
    $ cd ~
    $ mkdir tbox_demo
    $ cd tbox_demo
    $ tbox -dsn="root:root1234@tcp(127.0.0.1:3306)/test?charset=utf8mb4"
    2021/12/25 19:36:18 generating code success

# the style of code generation

```go
// Package model of db entity
// Code generated by tbox. DO NOT EDIT!!!

package model

import (
	"time"
)

// NewsTable for news
const NewsTable = "news"

// NewsEntity for news table entity struct.
type NewsEntity struct {
	ID        int64      `json:"id" db:"id"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
	Title     *string    `json:"title" db:"title"`
	Slug      *string    `json:"slug" db:"slug"`
	Content   *string    `json:"content" db:"content"`
	Status    *string    `json:"status" db:"status"`
}

// TableName for news
func (NewsEntity) TableName() string {
	return NewsTable
}

```
