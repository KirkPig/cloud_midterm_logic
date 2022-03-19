# 2110524 CLOUD COMP TECH Midterm Project

## Running the app
1. Rename `config/config.example.yml` to `config/config.yml` and fill in missing configurations.
2. ```shell
   $ go run main.go
   ```
   
## Migration
### Create migration
```shell
$  migrate create -ext sql -dir migrations -seq <migration_name>
```
### Apply migration (up)
Linux:
```shell
$  ./migrations/migrate.sh
```

Windows:
```shell
$  ./migrations/migrate.ps1
```
