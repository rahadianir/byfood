"use client";

import { useState } from "react";

interface AddBookFormProps {
  onSuccess: () => void; // Called when the book is successfully added
}

export default function AddBookForm({ onSuccess }: AddBookFormProps) {
  const [title, setTitle] = useState("");
  const [author, setAuthor] = useState("");
  const [year, setYear] = useState("");
  const [error, setError] = useState("");

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");

    // Basic validation
    if (!title.trim() || !author.trim() || !year) {
      setError("All fields are required.");
      return;
    }

    if (Number(year) < 0 || Number(year) > new Date().getFullYear()) {
      setError("Please enter a valid publication year.");
      return;
    }

    try {
      const res = await fetch("http://localhost:8080/books", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ title, author, publish_year: Number(year) }),
      });

      if (!res.ok) {
        throw new Error(`Error: ${res.statusText}`);
      }

      // Reset form
      setTitle("");
      setAuthor("");
      setYear("");

      onSuccess(); // refresh list or close modal
    } catch (err: any) {
      setError("Failed to add book. Please try again.");
      console.error(err);
    }
  };

  return (
    <form onSubmit={handleSubmit} style={formStyle}>
      {error && <p style={errorStyle}>{error}</p>}

      <div style={fieldStyle}>
        <label style={labelStyle}>Title:</label>
        <input
          type="text"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          style={inputStyle}
          required
        />
      </div>

      <div style={fieldStyle}>
        <label style={labelStyle}>Author:</label>
        <input
          type="text"
          value={author}
          onChange={(e) => setAuthor(e.target.value)}
          style={inputStyle}
          required
        />
      </div>

      <div style={fieldStyle}>
        <label style={labelStyle}>Year:</label>
        <input
          type="number"
          value={year}
          onChange={(e) => setYear(e.target.value)}
          style={inputStyle}
          required
        />
      </div>

      <button type="submit" style={buttonStyle}>
        Save
      </button>
    </form>
  );
}

// Styles
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
