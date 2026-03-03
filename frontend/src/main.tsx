import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { BrowserRouter as Router, Routes, Route } from "react-router";
import { CookiesProvider } from "react-cookie";

import "./index.css";

import Auth from "./pages/auth.tsx";
import Book from "./pages/book.tsx";
import NotFound from "./pages/notfound.tsx";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <CookiesProvider>
      <Router>
        <Routes>
          <Route path="/auth" element={<Auth />} />
          <Route path="/book/:uuid" element={<Book />} />
          <Route path="/*" element={<NotFound />} />
        </Routes>
      </Router>
    </CookiesProvider>
  </StrictMode>,
);
