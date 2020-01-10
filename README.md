# crm2go

## Development

Start `postgres` and import `sql/crm.sql`.

```
docker run -e POSTGRES_DB=crm -d -p 5432:5432 postgres
psql -d crm -h localhost -U postgres < sql/crm.sql
```
