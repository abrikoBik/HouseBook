import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { BrowserRouter as Router, Routes, Route } from "react-router"
import { CookiesProvider } from "react-cookie";

import "./index.css";
<<<<<<< HEAD
import Auth from "./pages/auth.tsx";
import Book from "./pages/book.tsx";
=======

import Auth from "./pages/auth.tsx";
>>>>>>> 23825df14795742c332de28d6a6d4044dbcd4920

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <CookiesProvider>
        <Router>
          <Routes>
            <Route path="/auth" element={<Auth/>} />
            <Route path="/book" element={<Book/>} />
          </Routes>
        </Router>
      </CookiesProvider>
  </StrictMode>,
);
