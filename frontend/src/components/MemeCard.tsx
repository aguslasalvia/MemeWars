import { useEffect, useState } from "react";
import { createPortal } from "react-dom";
import { imageUrl } from "../api.ts";
import type { Meme } from "../types";

interface MemeCardProps {
  meme: Meme;
  votes: number;
  hasVoted: boolean;
  onVote: (memeId: number) => void;
}

function MemeCard({ meme, votes, hasVoted, onVote }: MemeCardProps) {
  const [isExpanded, setIsExpanded] = useState(false);

  // Let people close the enlarged view with Escape, not just by clicking out.
  useEffect(() => {
    if (!isExpanded) return;

    function handleKeyDown(event: KeyboardEvent) {
      if (event.key === "Escape") setIsExpanded(false);
    }

    window.addEventListener("keydown", handleKeyDown);
    return () => window.removeEventListener("keydown", handleKeyDown);
  }, [isExpanded]);

  return (
    <div className="card meme-card">
      <img
        className="meme-image"
        src={imageUrl(meme.image_url)}
        alt={meme.text}
        onClick={() => setIsExpanded(true)}
      />
      {meme.text && <p className="meme-text">{meme.text}</p>}

      <button
        className={`vote-button ${hasVoted ? "card-pressed" : "card"}`}
        onClick={() => onVote(meme.ID)}
        disabled={hasVoted}
      >
        <span className="vote-count">{votes}</span>
        <span className="vote-label">{hasVoted ? "votado" : "votar"}</span>
      </button>

      {isExpanded &&
        // Rendered straight onto <body> instead of here: the meme card is
        // rotated with a CSS transform, and a transform on an ancestor
        // makes position:fixed children stick to THAT box instead of the
        // full screen. A portal skips the rotated ancestor entirely.
        createPortal(
          <div className="modal-backdrop" onClick={() => setIsExpanded(false)}>
            <div className="modal-card" onClick={(event) => event.stopPropagation()}>
              <button className="modal-close" onClick={() => setIsExpanded(false)} aria-label="Cerrar" />
              <img className="modal-image" src={imageUrl(meme.image_url)} alt={meme.text} />
              {meme.text && <p className="modal-text">{meme.text}</p>}
            </div>
          </div>,
          document.body,
        )}
    </div>
  );
}

export default MemeCard;
