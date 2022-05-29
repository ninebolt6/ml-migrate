# ml-migrate
ml-migrate is a CLI-based migration tool for MySQL.
- Create SQL files in specific directory. (default: `./migrations`)
- ml-migrate executes SQL files and stores its' filenames in `_migration` table.
- SQL files will be executed in ascending order. Files stored in `_migration` table will be ignored.
- ⚠️When execution fails, ml-migrate will call `ROLLBACK;`. It is recommended to add `BEGIN;` and `COMMIT;` to your SQL file.

## Usage
```bash
$ migrate info -p dbPassword -D path/to/migrations/dir dbName
$ migrate run -p dbPassword -D path/to/migrations/dir dbName
```
