# RepoMetadataExtraData


## Supported Types

### 

```go
repoMetadataExtraData := shared.CreateRepoMetadataExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
repoMetadataExtraData := shared.CreateRepoMetadataExtraDataStr(string{/* values here */})
```

### 

```go
repoMetadataExtraData := shared.CreateRepoMetadataExtraDataNumber(float64{/* values here */})
```

### 

```go
repoMetadataExtraData := shared.CreateRepoMetadataExtraDataBoolean(bool{/* values here */})
```

### 

```go
repoMetadataExtraData := shared.CreateRepoMetadataExtraDataArrayOfRepoMetadata5([]shared.RepoMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch repoMetadataExtraData.Type {
	case shared.RepoMetadataExtraDataTypeMapOfAny:
		// repoMetadataExtraData.MapOfAny is populated
	case shared.RepoMetadataExtraDataTypeStr:
		// repoMetadataExtraData.Str is populated
	case shared.RepoMetadataExtraDataTypeNumber:
		// repoMetadataExtraData.Number is populated
	case shared.RepoMetadataExtraDataTypeBoolean:
		// repoMetadataExtraData.Boolean is populated
	case shared.RepoMetadataExtraDataTypeArrayOfRepoMetadata5:
		// repoMetadataExtraData.ArrayOfRepoMetadata5 is populated
}
```
