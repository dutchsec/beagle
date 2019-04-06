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

### Beagle Run

Run the application and restart on file modifications.

```
beagle run go run main.go
```

