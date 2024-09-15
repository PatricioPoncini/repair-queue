# `🏁 🏎️ repair-queue`

## `🛠️ Stack`

[![My Stack](https://skillicons.dev/icons?i=go,mysql)](https://skillicons.dev)

## `📜 Commands`

### `🚀 Run Server`

```bash
make run
```

### `Setting Up Environment Configuration for Docker Compose`
1. Create a `.env` file in the root directory of your project with the following content:
```bash
DB_USER=root
DB_PASSWORD=my_secret_password
DB_HOST=127.0.0.1
DB_NAME=repair_queue
PUBLIC_HOST=http://127.0.0.1
PORT=3000
JWT_SECRET=my_super_secret_key
JWT_EXPIRATION_TIME=60 # seconds
```
2. Start your Docker containers by `docker compose up -d` in your terminal
3. Set up the project locally by running `make run`

### `🔬 Run Tests`

```bash
make test
```

### `✨ Run Lint`

```bash
make lint
# or
make lint-fix
```

### `📦 Run Migrations`

`Up Migrations ⬆️`

```bash
make migrate-up
```

`Down Migrations ⬇️`

```bash
make migrate-down
```
