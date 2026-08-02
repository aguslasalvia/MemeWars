import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { createRoom, getRoom } from "../api.ts";
import type { User } from "../types";

interface LobbyPageProps {
  user: User | null;
  onLogout: () => void;
}

function LobbyPage({ user, onLogout }: LobbyPageProps) {
  const [roomName, setRoomName] = useState("");
  const [roomCode, setRoomCode] = useState("");
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  useEffect(() => {
    if (!user) {
      navigate("/");
    }
  }, [user, navigate]);

  async function handleCreateRoom(event: React.FormEvent) {
    event.preventDefault();
    if (!roomName.trim()) {
      setError("Ponele un nombre a la sala");
      return;
    }

    setLoading(true);
    setError("");
    try {
      const room = await createRoom(roomName.trim());
      navigate(`/room/${room.ID}`);
    } catch (err) {
      setError(err instanceof Error ? err.message : "Algo salio mal");
    } finally {
      setLoading(false);
    }
  }

  async function handleJoinRoom(event: React.FormEvent) {
    event.preventDefault();
    if (!roomCode.trim()) {
      setError("Escribi el codigo de la sala");
      return;
    }

    setLoading(true);
    setError("");
    try {
      const room = await getRoom(roomCode.trim());
      navigate(`/room/${room.ID}`);
    } catch {
      setError("No encontramos esa sala");
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className="page-center">
      <div className="lobby-header">
        <p className="subtitle">Hola, {user?.name}</p>
        <button className="btn" onClick={onLogout}>
          Salir
        </button>
      </div>

      <div className="lobby-grid">
        <form onSubmit={handleCreateRoom} className="card lobby-card">
          <h2>Crear sala</h2>
          <p className="subtitle">Armate tu propia arena de memes</p>
          <input
            className="field"
            type="text"
            placeholder="Nombre de la sala"
            value={roomName}
            onChange={(event) => setRoomName(event.target.value)}
          />
          <button className="btn" type="submit" disabled={loading}>
            Crear sala
          </button>
        </form>

        <form onSubmit={handleJoinRoom} className="card lobby-card">
          <h2>Unirse a sala</h2>
          <p className="subtitle">Metete a pelear con un codigo</p>
          <input
            className="field"
            type="text"
            placeholder="Codigo de sala"
            value={roomCode}
            onChange={(event) => setRoomCode(event.target.value)}
          />
          <button className="btn" type="submit" disabled={loading}>
            Unirse
          </button>
        </form>
      </div>

      {error && <p className="error-text">{error}</p>}
    </div>
  );
}

export default LobbyPage;
