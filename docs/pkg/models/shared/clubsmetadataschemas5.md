# ClubsMetadataSchemas5


## Supported Types

### ClubsMetadataSchemas1

```go
clubsMetadataSchemas5 := shared.CreateClubsMetadataSchemas5ClubsMetadataSchemas1(shared.ClubsMetadataSchemas1{/* values here */})
```

### 

```go
clubsMetadataSchemas5 := shared.CreateClubsMetadataSchemas5Str(string{/* values here */})
```

### 

```go
clubsMetadataSchemas5 := shared.CreateClubsMetadataSchemas5Number(float64{/* values here */})
```

### 

```go
clubsMetadataSchemas5 := shared.CreateClubsMetadataSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch clubsMetadataSchemas5.Type {
	case shared.ClubsMetadataSchemas5TypeClubsMetadataSchemas1:
		// clubsMetadataSchemas5.ClubsMetadataSchemas1 is populated
	case shared.ClubsMetadataSchemas5TypeStr:
		// clubsMetadataSchemas5.Str is populated
	case shared.ClubsMetadataSchemas5TypeNumber:
		// clubsMetadataSchemas5.Number is populated
	case shared.ClubsMetadataSchemas5TypeBoolean:
		// clubsMetadataSchemas5.Boolean is populated
}
```
