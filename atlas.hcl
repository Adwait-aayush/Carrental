env "local" {
    src = "file://.database/schema.sql"
    url = "postgres://carrental:carrental@postgres:5432/carrental?search_path=public&sslmode=disable"
    dev = "docker://postgres/15/dev?search_path=public"
    migration {
        dir = "file://.database/migrations"
    }
} 