# ClubsMetadata5


## Supported Types

### ClubsMetadata1

```go
clubsMetadata5 := shared.CreateClubsMetadata5ClubsMetadata1(shared.ClubsMetadata1{/* values here */})
```

### 

```go
clubsMetadata5 := shared.CreateClubsMetadata5Str(string{/* values here */})
```

### 

```go
clubsMetadata5 := shared.CreateClubsMetadata5Number(float64{/* values here */})
```

### 

```go
clubsMetadata5 := shared.CreateClubsMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch clubsMetadata5.Type {
	case shared.ClubsMetadata5TypeClubsMetadata1:
		// clubsMetadata5.ClubsMetadata1 is populated
	case shared.ClubsMetadata5TypeStr:
		// clubsMetadata5.Str is populated
	case shared.ClubsMetadata5TypeNumber:
		// clubsMetadata5.Number is populated
	case shared.ClubsMetadata5TypeBoolean:
		// clubsMetadata5.Boolean is populated
}
```
