# Subscription Feature Flag Platform

A backend service for managing **feature flags for each client** of a subscription platform: create flags, turn them on or off per client, and define dependencies between flags (if B depends on A, turning A off also turns B off).

Built with **Go, Gin, GORM and PostgreSQL** as an assignment for FinBox.

> **Status: work in progress.** The data model for every feature is in place, but only **creating a feature flag** has an API so far. The rest is listed under [Roadmap](#roadmap).

## Data model

| Table | Purpose |
|---|---|
| `feature_flags` | Flag name (unique), description, default value |
| `dependencies` | `parent_id → child_id` links between flags |
| `clients` | Subscribing clients |
| `client_feature_flags` | Whether each flag is on or off for each client |

Tables are created automatically at startup with GORM's `AutoMigrate`.

## API

### `POST /api/feature-flags`

```bash
curl -X POST http://localhost:8080/api/feature-flags \
  -H "Content-Type: application/json" \
  -d '{"name": "dark_mode", "description": "New UI theme", "defaultValue": false}'
```

```json
{ "id": 1, "name": "dark_mode", "description": "New UI theme", "defaultValue": false }
```

## Running locally

1. Start PostgreSQL. The bundled compose file runs Postgres 13 on port 5432:

   ```bash
   docker-compose up -d postgres
   ```

2. Create a `.env` file in the project root:

   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=subarna
   DB_PASSWORD=passw0rd
   DB_NAME=localtest_db
   ```

3. Run the server:

   ```bash
   go run main.go     # http://localhost:8080
   ```

> The `app` service in `docker-compose.yml` expects a `Dockerfile`, which hasn't been added yet. Run the Go service directly for now.

## Roadmap

- [x] Create feature flag
- [ ] List, update and delete feature flags
- [ ] Client CRUD
- [ ] Turn a flag on or off for a client
- [ ] Add dependencies between flags, rejecting cycles
- [ ] Cascade: turning off a parent turns off its dependents
- [ ] Query all flags, or only enabled flags, for a client
- [ ] Table-driven tests (`go test ./... -cover`)
- [ ] Dockerfile for the app service

## Project structure

```
main.go                         DB connection, migrations, router
routers/router.go               Routes under /api
controllers/feature_flag_controller.go
models/feature_flag.go          FeatureFlag, Dependency, Client, ClientFeatureFlag
database/database.go            Singleton GORM connection (reads .env)
utils/migration.go              AutoMigrate
```

## License

MIT. See [LICENSE](LICENSE).
