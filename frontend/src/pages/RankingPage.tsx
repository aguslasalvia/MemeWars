import { useEffect, useState } from "react";
import { Link, useParams } from "react-router-dom";
import { getRanking, imageUrl } from "../api.ts";
import type { RankingEntry } from "../types";

const MEDAL_CLASSES = ["medal-gold", "medal-silver", "medal-bronze"];

function RankingPage() {
  const { roomId } = useParams();
  const [ranking, setRanking] = useState<RankingEntry[]>([]);
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    if (!roomId) return;
    getRanking(roomId)
      .then(setRanking)
      .catch((err) => setError(err instanceof Error ? err.message : "No se pudo cargar el ranking"))
      .finally(() => setLoading(false));
  }, [roomId]);

  if (loading) {
    return <div className="page-center">Cargando ranking...</div>;
  }

  const podium = ranking.slice(0, 3);
  const rest = ranking.slice(3);

  return (
    <div className="ranking-page">
      <header className="room-header">
        <h1>Ranking</h1>
        <div className="header-actions">
          <Link className="btn" to="/lobby">
            Inicio
          </Link>
          <Link className="btn" to={`/room/${roomId}`}>
            Volver a la sala
          </Link>
        </div>
      </header>

      {error && <p className="error-text">{error}</p>}
      {!loading && ranking.length === 0 && <p className="subtitle">Todavia no hay memes en esta sala.</p>}

      <div className="podium">
        {podium.map((entry, index) => (
          <div key={entry.meme_id} className={`card podium-spot ${MEDAL_CLASSES[index]}`}>
            <span className="podium-rank">#{index + 1}</span>
            <img className="podium-image" src={imageUrl(entry.image_url)} alt={entry.text} />
            <span className="podium-votes">{entry.total_votes} votos</span>
          </div>
        ))}
      </div>

      {rest.length > 0 && (
        <ul className="scoreboard">
          {rest.map((entry, index) => (
            <li key={entry.meme_id} className="card scoreboard-row">
              <span className="scoreboard-rank">#{index + 4}</span>
              <img className="scoreboard-image" src={imageUrl(entry.image_url)} alt={entry.text} />
              <span className="scoreboard-text">{entry.text}</span>
              <span className="scoreboard-votes">{entry.total_votes}</span>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}

export default RankingPage;
