import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import toast from "react-hot-toast";
import { getUserByName, createUser } from "../api.ts";
import type { User } from "../types";

interface EntryPageProps {
  user: User | null;
  onLogin: (user: User) => void;
}

function EntryPage({ user, onLogin }: EntryPageProps) {
  const [name, setName] = useState("");
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  // If we already know who you are, skip straight to the lobby.
  useEffect(() => {
    if (user) {
      navigate("/lobby");
    }
  }, [user, navigate]);

  async function handleSubmit(event: React.FormEvent) {
    event.preventDefault();
    const trimmedName = name.trim();
    if (!trimmedName) {
      toast.error("Escribi un nombre para entrar");
      return;
    }

    setLoading(true);

    try {
      // If the name already exists, reuse that user. Otherwise, create it.
      let foundUser: User;
      try {
        foundUser = await getUserByName(trimmedName);
      } catch {
        foundUser = await createUser(trimmedName);
      }
      onLogin(foundUser);
      navigate("/lobby");
    } catch (err) {
      toast.error(err instanceof Error ? err.message : "Algo salio mal");
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className="page-center">
      <div className="card entry-card">
        <h1 className="wordmark">
          MEME<span className="wordmark-pop">WARS</span>
        </h1>
        <p className="subtitle">Subi memes. Ganate los votos.</p>

        <form onSubmit={handleSubmit} className="entry-form">
          <input
            className="field"
            type="text"
            placeholder="Tu nombre de guerra"
            value={name}
            onChange={(event) => setName(event.target.value)}
            maxLength={40}
          />
          <button className="btn" type="submit" disabled={loading}>
            {loading ? "Entrando..." : "Entrar al campo"}
          </button>
        </form>
      </div>
    </div>
  );
}

export default EntryPage;
