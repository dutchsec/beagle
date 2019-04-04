# Beagle

## Usage

### Beagle DB

Generate code for database usage.

```
//go:generate beagle db --table users --key user_id user.go
```

```
go generate user.go
```

### Beagle Schema

Generate code for json schema's.

```
//go:generate beagle schema *.json
package schema
```

```
go generate schema.go
```

### Beagle License

Update license headers for all go files.

```
beagle license .
```

https://stackoverflow.com/questions/6135376/mysql-select-where-field-in-subquery-extremely-slow-why
