import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { BrowserRouter } from "react-router-dom";
import { Toaster } from "react-hot-toast";
import "./index.css";
import App from "./App.tsx";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <BrowserRouter>
      <App />
      <Toaster
        position="top-center"
        toastOptions={{
          duration: 3500,
          style: {
            background: "var(--paper)",
            color: "var(--ink)",
            border: "var(--border)",
            borderRadius: "var(--radius)",
            boxShadow: "6px 6px 0 var(--hot)",
            fontFamily: "var(--font-body)",
            fontWeight: 600,
          },
          success: {
            iconTheme: { primary: "var(--gold)", secondary: "var(--bg)" },
            style: { boxShadow: "6px 6px 0 var(--sky)" },
          },
          error: {
            iconTheme: { primary: "var(--hot)", secondary: "var(--ink)" },
          },
        }}
      />
    </BrowserRouter>
  </StrictMode>,
);
