import { useEffect, useState } from "react";
import { Link, useNavigate, useParams } from "react-router-dom";
import { getRoom, getRanking, uploadMeme, voteMeme } from "../api.ts";
import MemeCard from "../components/MemeCard.tsx";
import type { Room, User } from "../types";

interface RoomPageProps {
  user: User | null;
}

function RoomPage({ user }: RoomPageProps) {
  const { roomId } = useParams();
  const navigate = useNavigate();

  const [room, setRoom] = useState<Room | null>(null);
  const [votesByMeme, setVotesByMeme] = useState<Record<number, number>>({});
  const [votedMemeIds, setVotedMemeIds] = useState<Set<number>>(new Set());
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(true);

  const [showUploadForm, setShowUploadForm] = useState(false);
  const [uploadText, setUploadText] = useState("");
  const [uploadFile, setUploadFile] = useState<File | null>(null);
  const [uploading, setUploading] = useState(false);

  useEffect(() => {
    if (!user) {
      navigate("/");
    }
  }, [user, navigate]);

  // Just fetches the data, doesn't touch state. Used both on the first
  // load and whenever we need to refresh after voting/uploading.
  async function fetchRoomData(id: string) {
    const [roomData, ranking] = await Promise.all([getRoom(id), getRanking(id)]);
    const voteMap: Record<number, number> = {};
    ranking.forEach((entry) => {
      voteMap[entry.meme_id] = entry.total_votes;
    });
    return { roomData, voteMap };
  }

  useEffect(() => {
    if (!roomId) return;
    let ignore = false; // avoids updating state if roomId changes mid-fetch

    fetchRoomData(roomId)
      .then(({ roomData, voteMap }) => {
        if (ignore) return;
        setRoom(roomData);
        setVotesByMeme(voteMap);
      })
      .catch((err) => {
        if (!ignore) setError(err instanceof Error ? err.message : "No se pudo cargar la sala");
      })
      .finally(() => {
        if (!ignore) setLoading(false);
      });

    return () => {
      ignore = true;
    };
  }, [roomId]);

  async function reloadRoom() {
    if (!roomId) return;
    try {
      const { roomData, voteMap } = await fetchRoomData(roomId);
      setRoom(roomData);
      setVotesByMeme(voteMap);
    } catch (err) {
      setError(err instanceof Error ? err.message : "No se pudo cargar la sala");
    }
  }

  async function handleVote(memeId: number) {
    if (!user) return;
    setError("");
    try {
      await voteMeme(memeId, user.ID);
      setVotedMemeIds((prev) => new Set(prev).add(memeId));
      await reloadRoom();
    } catch (err) {
      setError(err instanceof Error ? err.message : "No se pudo votar");
    }
  }

  async function handleUpload(event: React.FormEvent) {
    event.preventDefault();
    if (!user || !roomId) return;
    if (!uploadFile) {
      setError("Elegi una imagen para subir");
      return;
    }

    setUploading(true);
    setError("");
    try {
      await uploadMeme(roomId, user.ID, uploadText.trim(), uploadFile);
      setUploadText("");
      setUploadFile(null);
      setShowUploadForm(false);
      await reloadRoom();
    } catch (err) {
      setError(err instanceof Error ? err.message : "No se pudo subir el meme");
    } finally {
      setUploading(false);
    }
  }

  if (loading) {
    return <div className="page-center">Cargando sala...</div>;
  }

  if (!room) {
    return <div className="page-center">{error || "No encontramos esta sala"}</div>;
  }

  return (
    <div className="room-page">
      <header className="room-header">
        <div>
          <h1>{room.name}</h1>
          <p className="subtitle">Codigo de sala: {room.ID}</p>
        </div>
        <div className="header-actions">
          <Link className="btn" to="/lobby">
            Inicio
          </Link>
          <Link className="btn" to={`/room/${room.ID}/ranking`}>
            Ver ranking
          </Link>
        </div>
      </header>

      {error && <p className="error-text">{error}</p>}

      <div className="meme-grid">
        {room.memes?.map((meme) => (
          <MemeCard
            key={meme.ID}
            meme={meme}
            votes={votesByMeme[meme.ID] || 0}
            hasVoted={votedMemeIds.has(meme.ID)}
            onVote={handleVote}
          />
        ))}
      </div>

      {!room.memes?.length && <p className="subtitle">Todavia no hay memes, se el primero.</p>}

      <button
        className="card fab"
        onClick={() => setShowUploadForm((show) => !show)}
        aria-label="Subir meme"
      />

      {showUploadForm && (
        <form onSubmit={handleUpload} className="card upload-panel">
          <h2>Subir meme</h2>
          <input
            className="field"
            type="text"
            placeholder="Un texto para tu meme (opcional)"
            value={uploadText}
            onChange={(event) => setUploadText(event.target.value)}
          />
          <input
            className="field"
            type="file"
            accept="image/*"
            onChange={(event) => setUploadFile(event.target.files?.[0] || null)}
          />
          <button className="btn" type="submit" disabled={uploading}>
            {uploading ? "Subiendo..." : "Subir al campo de batalla"}
          </button>
        </form>
      )}
    </div>
  );
}

export default RoomPage;
