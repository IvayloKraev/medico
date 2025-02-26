To run you need to install 
- GoLang
- MySQL (Windows) / MariaDB (Linux)
- Redis (for windows only through WSL2)

Run
```bash
go mod download
```

Create separate database for medico.
Update the files in ./config to match the compute's data.
Run
```bash
go run .
```
