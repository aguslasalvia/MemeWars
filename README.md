# MemeWars

App para crear salas donde los usuarios suben memes y votan por sus favoritos, con un ranking por sala. Backend en Go (Gin + GORM/SQLite), frontend en React servido por el mismo binario.

## Estructura

```
cmd/server        entrypoint (main.go)
internal/config    helpers de configuración (env vars)
internal/db        conexión y migraciones (SQLite vía GORM)
internal/models     entidades: User, Room, Meme, Vote
internal/handlers   controllers Gin
internal/services   lógica de negocio
internal/routes     definición de rutas
frontend/           SPA en React + Vite (ver frontend/README.md)
static/             assets servidos por el server
  memes/              imágenes subidas por los usuarios (gitignored)
  dist/               build del frontend (gitignored, generado por `bun run build`)
```

## Requisitos

- Go 1.26+
- Bun (o npm/node) para el frontend

## Correr en desarrollo

Backend (puerto 4040 por defecto):

```bash
go run ./cmd/server
```

Frontend con hot reload (apunta a `http://localhost:4040/api/v1` por defecto):

```bash
cd frontend
bun install
bun run dev
```

Variables de entorno (`.env` en la raíz, opcional):

- `PORT` — puerto del server (default `4040`)

## Build para producción

El server sirve el frontend compilado, así que hay que buildear antes de levantar el binario:

```bash
cd frontend && bun install && bun run build && cd ..
go build -o memewars ./cmd/server
./memewars
```

`bun run build` deja el output en `static/dist`, que Gin sirve en `/`, `/assets` y `/favicon.svg`, con fallback SPA para rutas del cliente (ej. `/room/1`).

## API (`/api/v1`)

| Método | Ruta                        | Descripción                          |
|--------|-----------------------------|---------------------------------------|
| POST   | `/rooms`                    | Crea una sala (`{ "name": string }`) |
| GET    | `/rooms/:id`                 | Obtiene una sala                     |
| GET    | `/rooms/:id/ranking`         | Ranking de memes de la sala por votos |
| POST   | `/users`                     | Crea un usuario (`{ "name": string }`) |
| GET    | `/users/:name`               | Busca un usuario por nombre          |
| POST   | `/memes/:room_id`            | Sube un meme (multipart: `user_id`, `text`, `image`) |
| GET    | `/memes/:room_id`            | Lista los memes de una sala          |
| GET    | `/memes/:room_id/:meme_id`   | Obtiene un meme puntual              |
| POST   | `/vote`                      | Vota un meme (`{ "meme_id", "user_id", "value" }`) |

Las imágenes subidas se guardan en `static/memes/room_<id>/` y quedan servidas en `/static/memes/room_<id>/<archivo>`.
