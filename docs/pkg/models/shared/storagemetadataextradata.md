# StorageMetadataExtraData


## Supported Types

### 

```go
storageMetadataExtraData := shared.CreateStorageMetadataExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
storageMetadataExtraData := shared.CreateStorageMetadataExtraDataStr(string{/* values here */})
```

### 

```go
storageMetadataExtraData := shared.CreateStorageMetadataExtraDataNumber(float64{/* values here */})
```

### 

```go
storageMetadataExtraData := shared.CreateStorageMetadataExtraDataBoolean(bool{/* values here */})
```

### 

```go
storageMetadataExtraData := shared.CreateStorageMetadataExtraDataArrayOfStorageMetadata5([]shared.StorageMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch storageMetadataExtraData.Type {
	case shared.StorageMetadataExtraDataTypeMapOfAny:
		// storageMetadataExtraData.MapOfAny is populated
	case shared.StorageMetadataExtraDataTypeStr:
		// storageMetadataExtraData.Str is populated
	case shared.StorageMetadataExtraDataTypeNumber:
		// storageMetadataExtraData.Number is populated
	case shared.StorageMetadataExtraDataTypeBoolean:
		// storageMetadataExtraData.Boolean is populated
	case shared.StorageMetadataExtraDataTypeArrayOfStorageMetadata5:
		// storageMetadataExtraData.ArrayOfStorageMetadata5 is populated
}
```
