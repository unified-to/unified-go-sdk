# AssessmentMetadataSchemas5


## Supported Types

### AssessmentMetadataSchemas1

```go
assessmentMetadataSchemas5 := shared.CreateAssessmentMetadataSchemas5AssessmentMetadataSchemas1(shared.AssessmentMetadataSchemas1{/* values here */})
```

### 

```go
assessmentMetadataSchemas5 := shared.CreateAssessmentMetadataSchemas5Str(string{/* values here */})
```

### 

```go
assessmentMetadataSchemas5 := shared.CreateAssessmentMetadataSchemas5Number(float64{/* values here */})
```

### 

```go
assessmentMetadataSchemas5 := shared.CreateAssessmentMetadataSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch assessmentMetadataSchemas5.Type {
	case shared.AssessmentMetadataSchemas5TypeAssessmentMetadataSchemas1:
		// assessmentMetadataSchemas5.AssessmentMetadataSchemas1 is populated
	case shared.AssessmentMetadataSchemas5TypeStr:
		// assessmentMetadataSchemas5.Str is populated
	case shared.AssessmentMetadataSchemas5TypeNumber:
		// assessmentMetadataSchemas5.Number is populated
	case shared.AssessmentMetadataSchemas5TypeBoolean:
		// assessmentMetadataSchemas5.Boolean is populated
}
```
