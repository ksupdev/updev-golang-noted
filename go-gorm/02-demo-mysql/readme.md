> go mod init updev/go-gorm/02-demo-mysql
1. Install GIN ``go get github.com/gin-gonic/gin``
go get gorm.io/driver/sqlite
go get gorm.io/gorm

```golang
type Product struct {
	gorm.Model // Use for genereate standare column ID,CreatedAt,UpdatedAt,DeletedAt
	Code       string
	Price      uint
}

```
