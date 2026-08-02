// All the calls to the Go backend live here, one plain function per
// endpoint. Nothing fancy: fetch, check if it worked, return JSON.

import type { User, Room, RankingEntry, Meme } from "./types";

const API_URL = import.meta.env.VITE_API_URL || "http://localhost:4040/api/v1";

async function handleResponse<T>(response: Response): Promise<T> {
  const data = await response.json().catch(() => null);
  if (!response.ok) {
    const message = data?.error || "Algo salio mal, intenta de nuevo";
    throw new Error(message);
  }
  return data as T;
}

export async function getUserByName(name: string): Promise<User> {
  const response = await fetch(`${API_URL}/users/${encodeURIComponent(name)}`);
  return handleResponse<User>(response);
}

export async function createUser(name: string): Promise<User> {
  const response = await fetch(`${API_URL}/users`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ name }),
  });
  return handleResponse<User>(response);
}

export async function createRoom(name: string): Promise<Room> {
  const response = await fetch(`${API_URL}/rooms`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ name }),
  });
  return handleResponse<Room>(response);
}

export async function getRoom(roomId: string | number): Promise<Room> {
  const response = await fetch(`${API_URL}/rooms/${roomId}`);
  return handleResponse<Room>(response);
}

export async function getRanking(roomId: string | number): Promise<RankingEntry[]> {
  const response = await fetch(`${API_URL}/rooms/${roomId}/ranking`);
  return handleResponse<RankingEntry[]>(response);
}

export async function uploadMeme(
  roomId: string | number,
  userId: number,
  text: string,
  imageFile: File,
): Promise<Meme> {
  const formData = new FormData();
  formData.append("user_id", String(userId));
  formData.append("text", text);
  formData.append("image", imageFile);

  const response = await fetch(`${API_URL}/memes/${roomId}`, {
    method: "POST",
    body: formData,
  });
  return handleResponse<Meme>(response);
}

export async function voteMeme(memeId: number, userId: number, value = 1) {
  const response = await fetch(`${API_URL}/vote`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ meme_id: memeId, user_id: userId, value }),
  });
  return handleResponse(response);
}

// The backend serves uploaded images from /static, but image_url comes
// back as an absolute path like "/static/memes/room_1/xxx.png", so we
// just need to prefix it with the server's origin.
export function imageUrl(path: string): string {
  const origin = API_URL.replace(/\/api\/v1\/?$/, "");
  return `${origin}${path}`;
}
