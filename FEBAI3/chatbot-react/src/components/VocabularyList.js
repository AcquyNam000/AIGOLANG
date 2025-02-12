import React from "react";

function VocabularyList({ words = [] }) {
  if (!words || words.length === 0) {
    return <p>Chưa có từ vựng nào được trích xuất.</p>;
  }

  return (
    <div className="vocabulary-list">
      <h3>Từ vựng tiếng Anh</h3>
      <ul>
        {words.map((word) => (
          <li key={word.id}>
            <strong>{word.vi}</strong> → {word.en}
          </li>
        ))}
      </ul>
    </div>
  );
}

export default VocabularyList;
