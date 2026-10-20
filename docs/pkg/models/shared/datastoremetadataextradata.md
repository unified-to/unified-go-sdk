# DatastoreMetadataExtraData


## Supported Types

### 

```go
datastoreMetadataExtraData := shared.CreateDatastoreMetadataExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
datastoreMetadataExtraData := shared.CreateDatastoreMetadataExtraDataStr(string{/* values here */})
```

### 

```go
datastoreMetadataExtraData := shared.CreateDatastoreMetadataExtraDataNumber(float64{/* values here */})
```

### 

```go
datastoreMetadataExtraData := shared.CreateDatastoreMetadataExtraDataBoolean(bool{/* values here */})
```

### 

```go
datastoreMetadataExtraData := shared.CreateDatastoreMetadataExtraDataArrayOfDatastoreMetadata5([]shared.DatastoreMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch datastoreMetadataExtraData.Type {
	case shared.DatastoreMetadataExtraDataTypeMapOfAny:
		// datastoreMetadataExtraData.MapOfAny is populated
	case shared.DatastoreMetadataExtraDataTypeStr:
		// datastoreMetadataExtraData.Str is populated
	case shared.DatastoreMetadataExtraDataTypeNumber:
		// datastoreMetadataExtraData.Number is populated
	case shared.DatastoreMetadataExtraDataTypeBoolean:
		// datastoreMetadataExtraData.Boolean is populated
	case shared.DatastoreMetadataExtraDataTypeArrayOfDatastoreMetadata5:
		// datastoreMetadataExtraData.ArrayOfDatastoreMetadata5 is populated
}
```
