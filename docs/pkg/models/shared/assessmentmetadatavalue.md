# AssessmentMetadataValue


## Supported Types

### 

```go
assessmentMetadataValue := shared.CreateAssessmentMetadataValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
assessmentMetadataValue := shared.CreateAssessmentMetadataValueStr(string{/* values here */})
```

### 

```go
assessmentMetadataValue := shared.CreateAssessmentMetadataValueNumber(float64{/* values here */})
```

### 

```go
assessmentMetadataValue := shared.CreateAssessmentMetadataValueBoolean(bool{/* values here */})
```

### 

```go
assessmentMetadataValue := shared.CreateAssessmentMetadataValueArrayOfAssessmentMetadataSchemas5([]shared.AssessmentMetadataSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch assessmentMetadataValue.Type {
	case shared.AssessmentMetadataValueTypeMapOfAny:
		// assessmentMetadataValue.MapOfAny is populated
	case shared.AssessmentMetadataValueTypeStr:
		// assessmentMetadataValue.Str is populated
	case shared.AssessmentMetadataValueTypeNumber:
		// assessmentMetadataValue.Number is populated
	case shared.AssessmentMetadataValueTypeBoolean:
		// assessmentMetadataValue.Boolean is populated
	case shared.AssessmentMetadataValueTypeArrayOfAssessmentMetadataSchemas5:
		// assessmentMetadataValue.ArrayOfAssessmentMetadataSchemas5 is populated
}
```
