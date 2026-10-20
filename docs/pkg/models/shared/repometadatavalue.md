# RepoMetadataValue


## Supported Types

### 

```go
repoMetadataValue := shared.CreateRepoMetadataValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
repoMetadataValue := shared.CreateRepoMetadataValueStr(string{/* values here */})
```

### 

```go
repoMetadataValue := shared.CreateRepoMetadataValueNumber(float64{/* values here */})
```

### 

```go
repoMetadataValue := shared.CreateRepoMetadataValueBoolean(bool{/* values here */})
```

### 

```go
repoMetadataValue := shared.CreateRepoMetadataValueArrayOfRepoMetadataSchemas5([]shared.RepoMetadataSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch repoMetadataValue.Type {
	case shared.RepoMetadataValueTypeMapOfAny:
		// repoMetadataValue.MapOfAny is populated
	case shared.RepoMetadataValueTypeStr:
		// repoMetadataValue.Str is populated
	case shared.RepoMetadataValueTypeNumber:
		// repoMetadataValue.Number is populated
	case shared.RepoMetadataValueTypeBoolean:
		// repoMetadataValue.Boolean is populated
	case shared.RepoMetadataValueTypeArrayOfRepoMetadataSchemas5:
		// repoMetadataValue.ArrayOfRepoMetadataSchemas5 is populated
}
```
