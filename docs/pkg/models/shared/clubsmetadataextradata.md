# ClubsMetadataExtraData


## Supported Types

### 

```go
clubsMetadataExtraData := shared.CreateClubsMetadataExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
clubsMetadataExtraData := shared.CreateClubsMetadataExtraDataStr(string{/* values here */})
```

### 

```go
clubsMetadataExtraData := shared.CreateClubsMetadataExtraDataNumber(float64{/* values here */})
```

### 

```go
clubsMetadataExtraData := shared.CreateClubsMetadataExtraDataBoolean(bool{/* values here */})
```

### 

```go
clubsMetadataExtraData := shared.CreateClubsMetadataExtraDataArrayOfClubsMetadata5([]shared.ClubsMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch clubsMetadataExtraData.Type {
	case shared.ClubsMetadataExtraDataTypeMapOfAny:
		// clubsMetadataExtraData.MapOfAny is populated
	case shared.ClubsMetadataExtraDataTypeStr:
		// clubsMetadataExtraData.Str is populated
	case shared.ClubsMetadataExtraDataTypeNumber:
		// clubsMetadataExtraData.Number is populated
	case shared.ClubsMetadataExtraDataTypeBoolean:
		// clubsMetadataExtraData.Boolean is populated
	case shared.ClubsMetadataExtraDataTypeArrayOfClubsMetadata5:
		// clubsMetadataExtraData.ArrayOfClubsMetadata5 is populated
}
```
