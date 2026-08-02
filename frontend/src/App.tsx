import { useState } from "react";
import { Routes, Route } from "react-router-dom";
import EntryPage from "./pages/EntryPage.tsx";
import LobbyPage from "./pages/LobbyPage.tsx";
import RoomPage from "./pages/RoomPage.tsx";
import RankingPage from "./pages/RankingPage.tsx";
import type { User } from "./types";

// The user is just "a name". We keep it in localStorage so refreshing
// the page (or coming back later) doesn't ask again.
function loadSavedUser(): User | null {
  const saved = localStorage.getItem("memewars_user");
  return saved ? JSON.parse(saved) : null;
}

function App() {
  const [user, setUser] = useState<User | null>(loadSavedUser);

  function handleLogin(newUser: User) {
    localStorage.setItem("memewars_user", JSON.stringify(newUser));
    setUser(newUser);
  }

  function handleLogout() {
    localStorage.removeItem("memewars_user");
    setUser(null);
  }

  return (
    <Routes>
      <Route path="/" element={<EntryPage user={user} onLogin={handleLogin} />} />
      <Route path="/lobby" element={<LobbyPage user={user} onLogout={handleLogout} />} />
      <Route path="/room/:roomId" element={<RoomPage user={user} />} />
      <Route path="/room/:roomId/ranking" element={<RankingPage />} />
    </Routes>
  );
}

export default App;
