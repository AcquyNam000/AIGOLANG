import React, { useState } from "react";
import ChatBox from "../components/ChatBox";
import ChatInput from "../components/ChatInput";
import DualChatInput from "../components/DualChatInput";
import VocabularyList from "../components/VocabularyList";
import { processDialog, manualDialog } from "../services/api";

function ChatPage() {
  const [mode, setMode] = useState("process"); // "process" hoặc "manual"
  const [dialog, setDialog] = useState(null);
  const [words, setWords] = useState([]);

  const handleSubmit = async (content) => {
    const response =
      mode === "process" ? await processDialog(content) : await manualDialog(content);
    setDialog(response.dialog);
    setWords(response.words);
  };

  return (
    <div className="chat-page">
      <div className="mode-toggle">
        <button onClick={() => setMode("process")} className={mode === "process" ? "active" : ""}>
          Nhập đoạn hội thoại
        </button>
        <button onClick={() => setMode("manual")} className={mode === "manual" ? "active" : ""}>
          Chat giữa hai người
        </button>
      </div>
      <ChatBox dialog={dialog} />
      {mode === "process" ? (
        <ChatInput onSubmit={handleSubmit} />
      ) : (
        <DualChatInput onSubmit={handleSubmit} />
      )}
      <VocabularyList words={words} />
    </div>
  );
}

export default ChatPage;
