RUN APPS

Breakdown Step by Step

1. git clone this
2. copy .env.example to .env (cp .env.example .env)
3. go mod tidy
4. go run main.go


Migrate & Seeder (DB)
command :
- migrate
  go run index.go migrate
- seed :
  go run index.go seed
