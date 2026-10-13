# LegalPartyRole

the party role on the matter

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.LegalPartyRoleClient

// Open enum: custom values can be created with a direct type cast
custom := shared.LegalPartyRole("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `LegalPartyRoleClient`          | CLIENT                          |
| `LegalPartyRoleOpposingParty`   | OPPOSING_PARTY                  |
| `LegalPartyRoleOpposingCounsel` | OPPOSING_COUNSEL                |
| `LegalPartyRoleCoCounsel`       | CO_COUNSEL                      |
| `LegalPartyRoleWitness`         | WITNESS                         |
| `LegalPartyRoleExpert`          | EXPERT                          |
| `LegalPartyRoleJudge`           | JUDGE                           |
| `LegalPartyRoleReferrer`        | REFERRER                        |
| `LegalPartyRoleOther`           | OTHER                           |