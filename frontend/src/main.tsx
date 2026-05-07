import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { BrowserRouter as Router, Routes, Route } from "react-router";
import { CookiesProvider } from "react-cookie";

import "./index.css";

import Landing from "./pages/landing.tsx";
import Auth from "./pages/auth.tsx";
import Book from "./pages/book.tsx";
import Notification from "./modules/notification.tsx";
import Profile from "./pages/profile.tsx";
import Catalog from "./pages/catalog.tsx";
import NotFound from "./pages/notfound.tsx";
import Library from "./pages/library.tsx";
import Trade from "./pages/trade.tsx";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <CookiesProvider>
      <Router>
        <Routes>
          <Route path="/" element={<Landing />} />
          <Route path="/auth" element={<Auth />} />
          <Route path="/book" element={<Book />} />
          <Route path="/test" element={<Notification />} />
          <Route path="/profile" element={<Profile />} />
          <Route path="/library" element={<Library />} />
          <Route path="/book/:uuid" element={<Book />} />
          <Route path="/book/catalog" element={<Catalog />} />
          <Route path="/landing" element={<Landing />} />
          <Route path="/profile/trade" element={<Trade />} />
          <Route path="/*" element={<NotFound />} />
        </Routes>
      </Router>
    </CookiesProvider>
  </StrictMode>,
);
