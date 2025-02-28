# notesplace

There are notes, and this is their place.

Create notebooks and share them with a link. You can protect them using a password.

- Backend in `server/`. Using Go, Gin, PostgreSQL.
- Web interface in `web/`.

## Setup

Thanks to there being a docker compose file you just need to run it `docker-compose up`.

Environment variables:

- `DB_USER`, `DB_PASSWORD`, `DB_HOST`, `DB_PORT`, `DB_NAME` - database related variables. Those do not need to be modified assuming default config.
- `PRIVATE_KEY` - a private key to use with bcrypt
