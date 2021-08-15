module github.com/aristletl/go_advanced

go 1.16

require (
	github.com/gin-gonic/gin v1.7.2
	github.com/go-sql-driver/mysql v1.6.0
	github.com/pkg/errors v0.9.1
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
)

replace golang.org/x/sync v0.0.0-20210220032951-036812b2e83c => github.com/golang/sync v0.0.0-20210220032951-036812b2e83c
