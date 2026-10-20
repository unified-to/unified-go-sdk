# ClubsMetadataValue


## Supported Types

### 

```go
clubsMetadataValue := shared.CreateClubsMetadataValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
clubsMetadataValue := shared.CreateClubsMetadataValueStr(string{/* values here */})
```

### 

```go
clubsMetadataValue := shared.CreateClubsMetadataValueNumber(float64{/* values here */})
```

### 

```go
clubsMetadataValue := shared.CreateClubsMetadataValueBoolean(bool{/* values here */})
```

### 

```go
clubsMetadataValue := shared.CreateClubsMetadataValueArrayOfClubsMetadataSchemas5([]shared.ClubsMetadataSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch clubsMetadataValue.Type {
	case shared.ClubsMetadataValueTypeMapOfAny:
		// clubsMetadataValue.MapOfAny is populated
	case shared.ClubsMetadataValueTypeStr:
		// clubsMetadataValue.Str is populated
	case shared.ClubsMetadataValueTypeNumber:
		// clubsMetadataValue.Number is populated
	case shared.ClubsMetadataValueTypeBoolean:
		// clubsMetadataValue.Boolean is populated
	case shared.ClubsMetadataValueTypeArrayOfClubsMetadataSchemas5:
		// clubsMetadataValue.ArrayOfClubsMetadataSchemas5 is populated
}
```
