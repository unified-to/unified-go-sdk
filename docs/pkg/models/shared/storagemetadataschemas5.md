# StorageMetadataSchemas5


## Supported Types

### StorageMetadataSchemas1

```go
storageMetadataSchemas5 := shared.CreateStorageMetadataSchemas5StorageMetadataSchemas1(shared.StorageMetadataSchemas1{/* values here */})
```

### 

```go
storageMetadataSchemas5 := shared.CreateStorageMetadataSchemas5Str(string{/* values here */})
```

### 

```go
storageMetadataSchemas5 := shared.CreateStorageMetadataSchemas5Number(float64{/* values here */})
```

### 

```go
storageMetadataSchemas5 := shared.CreateStorageMetadataSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch storageMetadataSchemas5.Type {
	case shared.StorageMetadataSchemas5TypeStorageMetadataSchemas1:
		// storageMetadataSchemas5.StorageMetadataSchemas1 is populated
	case shared.StorageMetadataSchemas5TypeStr:
		// storageMetadataSchemas5.Str is populated
	case shared.StorageMetadataSchemas5TypeNumber:
		// storageMetadataSchemas5.Number is populated
	case shared.StorageMetadataSchemas5TypeBoolean:
		// storageMetadataSchemas5.Boolean is populated
}
```
