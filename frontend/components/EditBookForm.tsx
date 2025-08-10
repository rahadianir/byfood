"use client";

import { useState } from "react";
import { useBooks } from "../context/BookContext";

interface EditBookFormProps {
  book: { id: number; title: string; author: string; publish_year: number };
  onSuccess: () => void;
}

export default function EditBookForm({ book, onSuccess }: EditBookFormProps) {
  const { updateBook } = useBooks();
  const [formData, setFormData] = useState({
    title: book.title,
    author: book.author,
    year: book.publish_year,
  });
  const [error, setError] = useState("");

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");

    if (!formData.title.trim() || !formData.author.trim() || !formData.year) {
      setError("All fields are required.");
      return;
    }

    if (Number(formData.year) < 0 || Number(formData.year) > new Date().getFullYear()) {
      setError("Please enter a valid publication year.");
      return;
    }

    try {
      await updateBook(book.id, {
        title: formData.title,
        author: formData.author,
        publish_year: formData.year,
      });
      onSuccess();
    } catch (err) {
      console.error(err);
      setError("Failed to update book. Please try again.");
    }
  };

  return (
    <form onSubmit={handleSubmit} style={formStyle}>
      {error && <p style={errorStyle}>{error}</p>}

      <div style={fieldStyle}>
        <label style={labelStyle}>Title:</label>
        <input
          type="text"
          value={formData.title}
          onChange={(e) => setFormData({ ...formData, title: e.target.value })}
          style={inputStyle}
        />
      </div>

      <div style={fieldStyle}>
        <label style={labelStyle}>Author:</label>
        <input
          type="text"
          value={formData.author}
          onChange={(e) => setFormData({ ...formData, author: e.target.value })}
          style={inputStyle}
        />
      </div>

      <div style={fieldStyle}>
        <label style={labelStyle}>Year:</label>
        <input
          type="text"
          value={formData.year}
          onChange={(e) =>
            setFormData({ ...formData, year: parseInt(e.target.value, 10) })
          }
          style={inputStyle}
        />
      </div>

      <button type="submit" style={buttonStyle}>
        Save Changes
      </button>
    </form>
  );
}

// Styles (same as AddBookForm for consistency)
const formStyle: React.CSSProperties = {
  display: "flex",
  flexDirection: "column",
};

const fieldStyle: React.CSSProperties = {
  marginBottom: "12px",
};

const labelStyle: React.CSSProperties = {
  display: "block",
  marginBottom: "4px",
  fontWeight: "bold",
  color: "#ddd",
};

const inputStyle: React.CSSProperties = {
  width: "100%",
  padding: "8px",
  borderRadius: "6px",
  border: "1px solid #555",
  backgroundColor: "#2a2a2a",
  color: "#f5f5f5",
};

const errorStyle: React.CSSProperties = {
  color: "#ff6b6b",
  marginBottom: "10px",
};

const buttonStyle: React.CSSProperties = {
  padding: "10px",
  backgroundColor: "#4CAF50",
  color: "#fff",
  border: "none",
  borderRadius: "6px",
  cursor: "pointer",
  fontWeight: "bold",
};
