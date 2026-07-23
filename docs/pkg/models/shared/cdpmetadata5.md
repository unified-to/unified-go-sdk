# CdpMetadata5


## Supported Types

### CdpMetadata1

```go
cdpMetadata5 := shared.CreateCdpMetadata5CdpMetadata1(shared.CdpMetadata1{/* values here */})
```

### 

```go
cdpMetadata5 := shared.CreateCdpMetadata5Str(string{/* values here */})
```

### 

```go
cdpMetadata5 := shared.CreateCdpMetadata5Number(float64{/* values here */})
```

### 

```go
cdpMetadata5 := shared.CreateCdpMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch cdpMetadata5.Type {
	case shared.CdpMetadata5TypeCdpMetadata1:
		// cdpMetadata5.CdpMetadata1 is populated
	case shared.CdpMetadata5TypeStr:
		// cdpMetadata5.Str is populated
	case shared.CdpMetadata5TypeNumber:
		// cdpMetadata5.Number is populated
	case shared.CdpMetadata5TypeBoolean:
		// cdpMetadata5.Boolean is populated
}
```
