# StorageMetadataValue


## Supported Types

### 

```go
storageMetadataValue := shared.CreateStorageMetadataValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
storageMetadataValue := shared.CreateStorageMetadataValueStr(string{/* values here */})
```

### 

```go
storageMetadataValue := shared.CreateStorageMetadataValueNumber(float64{/* values here */})
```

### 

```go
storageMetadataValue := shared.CreateStorageMetadataValueBoolean(bool{/* values here */})
```

### 

```go
storageMetadataValue := shared.CreateStorageMetadataValueArrayOfStorageMetadataSchemas5([]shared.StorageMetadataSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch storageMetadataValue.Type {
	case shared.StorageMetadataValueTypeMapOfAny:
		// storageMetadataValue.MapOfAny is populated
	case shared.StorageMetadataValueTypeStr:
		// storageMetadataValue.Str is populated
	case shared.StorageMetadataValueTypeNumber:
		// storageMetadataValue.Number is populated
	case shared.StorageMetadataValueTypeBoolean:
		// storageMetadataValue.Boolean is populated
	case shared.StorageMetadataValueTypeArrayOfStorageMetadataSchemas5:
		// storageMetadataValue.ArrayOfStorageMetadataSchemas5 is populated
}
```
