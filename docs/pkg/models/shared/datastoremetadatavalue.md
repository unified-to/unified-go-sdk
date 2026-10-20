# DatastoreMetadataValue


## Supported Types

### 

```go
datastoreMetadataValue := shared.CreateDatastoreMetadataValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
datastoreMetadataValue := shared.CreateDatastoreMetadataValueStr(string{/* values here */})
```

### 

```go
datastoreMetadataValue := shared.CreateDatastoreMetadataValueNumber(float64{/* values here */})
```

### 

```go
datastoreMetadataValue := shared.CreateDatastoreMetadataValueBoolean(bool{/* values here */})
```

### 

```go
datastoreMetadataValue := shared.CreateDatastoreMetadataValueArrayOfDatastoreMetadataSchemas5([]shared.DatastoreMetadataSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch datastoreMetadataValue.Type {
	case shared.DatastoreMetadataValueTypeMapOfAny:
		// datastoreMetadataValue.MapOfAny is populated
	case shared.DatastoreMetadataValueTypeStr:
		// datastoreMetadataValue.Str is populated
	case shared.DatastoreMetadataValueTypeNumber:
		// datastoreMetadataValue.Number is populated
	case shared.DatastoreMetadataValueTypeBoolean:
		// datastoreMetadataValue.Boolean is populated
	case shared.DatastoreMetadataValueTypeArrayOfDatastoreMetadataSchemas5:
		// datastoreMetadataValue.ArrayOfDatastoreMetadataSchemas5 is populated
}
```
