# MemeWars — Frontend

SPA en React + TypeScript (Vite) para MemeWars: crear/entrar a salas, subir memes, votar y ver el ranking.

## Stack

- React 19 + React Router
- Vite 8 (build)
- Bun como package manager

## Estructura

```
src/
  api.ts              cliente HTTP hacia el backend (VITE_API_URL, default http://localhost:4040/api/v1)
  App.tsx              rutas de la app
  pages/                EntryPage, LobbyPage, RoomPage, RankingPage
  components/           MemeCard, etc.
  types.ts              tipos compartidos
```

## Desarrollo

```bash
bun install
bun run dev
```

Por default pega contra `http://localhost:4040/api/v1`. Para apuntar a otro backend, crear un `.env` con:

```
VITE_API_URL=http://localhost:4040/api/v1
```

## Build

```bash
bun run build
```

El output va a `../static/dist` (configurado en `vite.config.ts`), que es la carpeta que sirve el backend Go — ver [`../README.md`](../README.md) para cómo levantar todo junto.

## Otros comandos

```bash
bun run lint      # eslint
bun run preview   # sirve el build de forma local, sin el backend Go
```
