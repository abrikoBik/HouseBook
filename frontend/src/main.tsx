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
import NotFound from "./pages/notfound.tsx";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <CookiesProvider>
      <Router>
        <Routes>
          <Route path="/" element={<Landing />} />
          <Route path="/auth" element={<Auth />} />
          <Route path="/book" element={<Book />} />
          <Route path="/test" element={<Notification />}/>
          <Route path="/profile" element={<Profile />} />
          <Route path="/book/:uuid" element={<Book />} />
          <Route path="/*" element={<NotFound />} />
        </Routes>
      </Router>
    </CookiesProvider>
  </StrictMode>,
);
