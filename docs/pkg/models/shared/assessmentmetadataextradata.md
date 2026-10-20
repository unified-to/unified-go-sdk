# AssessmentMetadataExtraData


## Supported Types

### 

```go
assessmentMetadataExtraData := shared.CreateAssessmentMetadataExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
assessmentMetadataExtraData := shared.CreateAssessmentMetadataExtraDataStr(string{/* values here */})
```

### 

```go
assessmentMetadataExtraData := shared.CreateAssessmentMetadataExtraDataNumber(float64{/* values here */})
```

### 

```go
assessmentMetadataExtraData := shared.CreateAssessmentMetadataExtraDataBoolean(bool{/* values here */})
```

### 

```go
assessmentMetadataExtraData := shared.CreateAssessmentMetadataExtraDataArrayOfAssessmentMetadata5([]shared.AssessmentMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch assessmentMetadataExtraData.Type {
	case shared.AssessmentMetadataExtraDataTypeMapOfAny:
		// assessmentMetadataExtraData.MapOfAny is populated
	case shared.AssessmentMetadataExtraDataTypeStr:
		// assessmentMetadataExtraData.Str is populated
	case shared.AssessmentMetadataExtraDataTypeNumber:
		// assessmentMetadataExtraData.Number is populated
	case shared.AssessmentMetadataExtraDataTypeBoolean:
		// assessmentMetadataExtraData.Boolean is populated
	case shared.AssessmentMetadataExtraDataTypeArrayOfAssessmentMetadata5:
		// assessmentMetadataExtraData.ArrayOfAssessmentMetadata5 is populated
}
```
