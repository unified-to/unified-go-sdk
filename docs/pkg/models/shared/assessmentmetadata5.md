# AssessmentMetadata5


## Supported Types

### AssessmentMetadata1

```go
assessmentMetadata5 := shared.CreateAssessmentMetadata5AssessmentMetadata1(shared.AssessmentMetadata1{/* values here */})
```

### 

```go
assessmentMetadata5 := shared.CreateAssessmentMetadata5Str(string{/* values here */})
```

### 

```go
assessmentMetadata5 := shared.CreateAssessmentMetadata5Number(float64{/* values here */})
```

### 

```go
assessmentMetadata5 := shared.CreateAssessmentMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch assessmentMetadata5.Type {
	case shared.AssessmentMetadata5TypeAssessmentMetadata1:
		// assessmentMetadata5.AssessmentMetadata1 is populated
	case shared.AssessmentMetadata5TypeStr:
		// assessmentMetadata5.Str is populated
	case shared.AssessmentMetadata5TypeNumber:
		// assessmentMetadata5.Number is populated
	case shared.AssessmentMetadata5TypeBoolean:
		// assessmentMetadata5.Boolean is populated
}
```
