import React, { useState } from "react";
import ChatPage from "./pages/ChatPage";
import "./styles/styles.css";

function App() {
  return (
    <div className="app-container">
      <h1>Chatbot Trích Xuất Từ Vựng</h1>
      <ChatPage />
    </div>
  );
}

export default App;
