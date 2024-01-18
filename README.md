# Local Development

I haven't fully automated local database development yet.

To start the database locally use:

```sh
make postgresInitDev
```

(might need to `docker pull postgres:16-alpine` first)

To create the table, you can use psql:

```sh
# psql in docker container shortcut
make postgres
```

Then

```
\c stranger-danger
```

And copy the contents of `db/migrations/createUserTable.sql`

I will automate process this soon.
