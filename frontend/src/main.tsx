import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { BrowserRouter as Router, Routes, Route } from "react-router";
import { CookiesProvider } from "react-cookie";

import "./index.css";

import Auth from "./pages/auth.tsx";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <CookiesProvider>
      <Router>
        <Routes>
          <Route path="/auth" element={<Auth />} />
        </Routes>
      </Router>
    </CookiesProvider>
  </StrictMode>,
);
