# About

This is a basic websocket server that I did following a tutorial to get my bearings writing websocket servers in Go.
I was originally planning on having this be the actual server for my eventual app, but I will likely completely rewrite
it to make sure I understand what is going on well enough to troubleshoot issues.

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

## Resources

My backend knowledge and Go knowledge are both fairly limited. So my goal is to use tutorials to help me figure this out and then mostly rewrite the project when I have a somewhat better idea what I'm doing (hopefully pretty much immediately after follow the tutorials). I will post helpful resources I find here:

[Current tutorial](https://www.youtube.com/watch?v=W9SuX9c40s8)
