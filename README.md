# Go DB Application – Setup, Git Workflow, and API Behavior

## 1. Initializing the Go Module

```bash
go mod init go-k8s
```

Initializes a new Go module named `go-k8s` and creates:
- `go.mod` – module metadata and dependency list
- `go.sum` – checksums to ensure dependency integrity

---

## 2. Adding HTTP Router Dependency (Gorilla Mux)

```bash
go get -u github.com/gorilla/mux
```

Downloads and adds the Gorilla Mux router.
- `-u` updates to the latest compatible version
- Marked as `// indirect` until explicitly imported

---

## 3. Adding Database ORM (GORM + PostgreSQL)

```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```

Adds:
- GORM ORM
- PostgreSQL driver
- Required transitive dependencies

---

## 4. go.mod vs go.sum

- `go.mod`: Declares module name, Go version, dependencies
- `go.sum`: Ensures reproducible builds and dependency integrity  
Both **must be committed**.

---

## 5. Git Status Before Pull

```bash
git status
```

- `go.mod` → tracked and modified
- `go-k8s` → untracked Go binary

---

## 6. Git Stash Behavior

```bash
git stash push -m "server local changes"
```

- Stashes **tracked files only**
- Untracked files are ignored

To include untracked files:
```bash
git stash push -u -m "include untracked"
```

---

## 7. Fetching Remote Changes

```bash
git fetch origin
git log --oneline HEAD..origin/main
git diff HEAD..origin/main
```

- Fetch downloads updates without modifying working tree

---

## 8. PostgreSQL DSN Fix

Incorrect:
```go
password=%s dbname=%s
```

Correct:
```go
password=%s dbname=%s
```

Correct ordering is required for authentication to succeed.

---

## 9. Merging Remote Branch

```bash
git merge origin/main
```

Fast-forward merge applied successfully.

---

## 10. Applying Stash

```bash
git stash apply stash@{0}
```

Restores tracked files only.

---

## 11. Go Binary (go-k8s)

- Compiled executable
- Should **not** be committed

Recommended `.gitignore`:
```gitignore
go-k8s
```

---

## 12. Gorilla Mux Routing Fix

Incorrect:
```go
/user{id}
```

Correct:
```go
/user/{id}
```

Curly braces are mandatory for path variables.

---

## 13. API Testing

Create user:
```bash
curl -X POST http://localhost:8080/user \
  -H "Content-Type: application/json" \
  -d '{"name":"Anuroop","email":"anu@example.com"}'
```

List users:
```bash
curl http://localhost:8080/users
```

Update user:
```bash
curl -X PUT http://localhost:8080/user/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Anuroop PS"}'
```

Delete user:
```bash
curl -X DELETE http://localhost:8080/user/1
```

---

## 14. Why `/users/1` Returns 404

No route defined for:
```go
/users/{id}
```

Only `/users` and `/user/{id}` exist.

---

## 15. Security Warning

Never commit GitHub tokens or secrets.

Use environment variables:
```bash
export GITHUB_TOKEN=xxxx
```

---

## 16. Best Practices

- Commit `go.mod` and `go.sum`
- Ignore binaries
- Validate routes
- Never store secrets in repos

