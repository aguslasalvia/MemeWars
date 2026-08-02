// Shapes of the data the Go backend sends back. Only the fields the
// frontend actually uses are listed here.

export interface User {
  ID: number;
  name: string;
}

export interface Meme {
  ID: number;
  room_id: number;
  user_id: number;
  image_url: string;
  text: string;
}

export interface Room {
  ID: number;
  name: string;
  active: boolean;
  memes?: Meme[];
}

export interface RankingEntry {
  meme_id: number;
  image_url: string;
  text: string;
  total_votes: number;
}
