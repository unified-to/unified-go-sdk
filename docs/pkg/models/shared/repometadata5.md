# RepoMetadata5


## Supported Types

### RepoMetadata1

```go
repoMetadata5 := shared.CreateRepoMetadata5RepoMetadata1(shared.RepoMetadata1{/* values here */})
```

### 

```go
repoMetadata5 := shared.CreateRepoMetadata5Str(string{/* values here */})
```

### 

```go
repoMetadata5 := shared.CreateRepoMetadata5Number(float64{/* values here */})
```

### 

```go
repoMetadata5 := shared.CreateRepoMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch repoMetadata5.Type {
	case shared.RepoMetadata5TypeRepoMetadata1:
		// repoMetadata5.RepoMetadata1 is populated
	case shared.RepoMetadata5TypeStr:
		// repoMetadata5.Str is populated
	case shared.RepoMetadata5TypeNumber:
		// repoMetadata5.Number is populated
	case shared.RepoMetadata5TypeBoolean:
		// repoMetadata5.Boolean is populated
}
```
