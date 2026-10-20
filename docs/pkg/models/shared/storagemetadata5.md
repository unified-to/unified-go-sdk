# StorageMetadata5


## Supported Types

### StorageMetadata1

```go
storageMetadata5 := shared.CreateStorageMetadata5StorageMetadata1(shared.StorageMetadata1{/* values here */})
```

### 

```go
storageMetadata5 := shared.CreateStorageMetadata5Str(string{/* values here */})
```

### 

```go
storageMetadata5 := shared.CreateStorageMetadata5Number(float64{/* values here */})
```

### 

```go
storageMetadata5 := shared.CreateStorageMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch storageMetadata5.Type {
	case shared.StorageMetadata5TypeStorageMetadata1:
		// storageMetadata5.StorageMetadata1 is populated
	case shared.StorageMetadata5TypeStr:
		// storageMetadata5.Str is populated
	case shared.StorageMetadata5TypeNumber:
		// storageMetadata5.Number is populated
	case shared.StorageMetadata5TypeBoolean:
		// storageMetadata5.Boolean is populated
}
```
