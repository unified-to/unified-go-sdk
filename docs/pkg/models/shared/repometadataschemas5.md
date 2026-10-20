# RepoMetadataSchemas5


## Supported Types

### RepoMetadataSchemas1

```go
repoMetadataSchemas5 := shared.CreateRepoMetadataSchemas5RepoMetadataSchemas1(shared.RepoMetadataSchemas1{/* values here */})
```

### 

```go
repoMetadataSchemas5 := shared.CreateRepoMetadataSchemas5Str(string{/* values here */})
```

### 

```go
repoMetadataSchemas5 := shared.CreateRepoMetadataSchemas5Number(float64{/* values here */})
```

### 

```go
repoMetadataSchemas5 := shared.CreateRepoMetadataSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch repoMetadataSchemas5.Type {
	case shared.RepoMetadataSchemas5TypeRepoMetadataSchemas1:
		// repoMetadataSchemas5.RepoMetadataSchemas1 is populated
	case shared.RepoMetadataSchemas5TypeStr:
		// repoMetadataSchemas5.Str is populated
	case shared.RepoMetadataSchemas5TypeNumber:
		// repoMetadataSchemas5.Number is populated
	case shared.RepoMetadataSchemas5TypeBoolean:
		// repoMetadataSchemas5.Boolean is populated
}
```
