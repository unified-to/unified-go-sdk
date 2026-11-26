# Resource

A resource of the searched type


## Supported Types

### 

```go
resource := shared.CreateResourceMapOfAny(map[string]any{/* values here */})
```

### 

```go
resource := shared.CreateResourceStr(string{/* values here */})
```

### 

```go
resource := shared.CreateResourceNumber(float64{/* values here */})
```

### 

```go
resource := shared.CreateResourceBoolean(bool{/* values here */})
```

### 

```go
resource := shared.CreateResourceArrayOfFhirBundleEntry5([]shared.FhirBundleEntry5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch resource.Type {
	case shared.ResourceUnionTypeMapOfAny:
		// resource.MapOfAny is populated
	case shared.ResourceUnionTypeStr:
		// resource.Str is populated
	case shared.ResourceUnionTypeNumber:
		// resource.Number is populated
	case shared.ResourceUnionTypeBoolean:
		// resource.Boolean is populated
	case shared.ResourceUnionTypeArrayOfFhirBundleEntry5:
		// resource.ArrayOfFhirBundleEntry5 is populated
}
```
