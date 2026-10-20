# DatastoreMetadataSchemas5


## Supported Types

### DatastoreMetadataSchemas1

```go
datastoreMetadataSchemas5 := shared.CreateDatastoreMetadataSchemas5DatastoreMetadataSchemas1(shared.DatastoreMetadataSchemas1{/* values here */})
```

### 

```go
datastoreMetadataSchemas5 := shared.CreateDatastoreMetadataSchemas5Str(string{/* values here */})
```

### 

```go
datastoreMetadataSchemas5 := shared.CreateDatastoreMetadataSchemas5Number(float64{/* values here */})
```

### 

```go
datastoreMetadataSchemas5 := shared.CreateDatastoreMetadataSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch datastoreMetadataSchemas5.Type {
	case shared.DatastoreMetadataSchemas5TypeDatastoreMetadataSchemas1:
		// datastoreMetadataSchemas5.DatastoreMetadataSchemas1 is populated
	case shared.DatastoreMetadataSchemas5TypeStr:
		// datastoreMetadataSchemas5.Str is populated
	case shared.DatastoreMetadataSchemas5TypeNumber:
		// datastoreMetadataSchemas5.Number is populated
	case shared.DatastoreMetadataSchemas5TypeBoolean:
		// datastoreMetadataSchemas5.Boolean is populated
}
```
