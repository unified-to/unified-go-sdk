# DbType

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.DbTypeMongodb

// Open enum: custom values can be created with a direct type cast
custom := shared.DbType("custom_value")
```


## Values

| Name             | Value            |
| ---------------- | ---------------- |
| `DbTypeMongodb`  | mongodb          |
| `DbTypeMysql`    | mysql            |
| `DbTypePostgres` | postgres         |
| `DbTypeMssql`    | mssql            |
| `DbTypeMariadb`  | mariadb          |
| `DbTypeSupabase` | supabase         |