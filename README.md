# 🔧 Subscription Feature Flag Platform

This project implements a backend service for managing **feature flags per client** in a subscription platform. It provides a RESTful API to create, update, and query feature flags, handle flag dependencies, and manage client-specific flag configurations.

## 🚀 Features

- Create and manage **clients**
- Create and manage **feature flags**
- Set feature flag **ON/OFF** per client
- Define **dependencies** between feature flags
- Enforce dependency rules (e.g., B depends on A)
- **Cascade toggling**: disabling a flag disables all dependent flags
- Query all feature flags or only enabled flags per client


## 🧪 Test-Driven Development (TDD)

This project is built with **TDD** in mind:

- All core features include test coverage
- Follows table-driven test format in Go
- To run tests and generate coverage:
  ```bash
  go test ./... -v -coverprofile=coverage.out
  go tool cover -func=coverage.out
  ```