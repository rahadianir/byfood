"use client";

import { useEffect, useState } from "react";
import { useParams, useRouter } from "next/navigation";
import "./BookDetail.css";

interface Book {
  id: number;
  title: string;
  author: string;
  publish_year: number;
  created_at?: string;
  updated_at?: string;
}

export default function BookDetailPage() {
  const { id } = useParams();
  const router = useRouter();
  const [book, setBook] = useState<Book | null>(null);

  useEffect(() => {
    if (!id) return;
    const fetchBook = async () => {
      try {
        const res = await fetch(`http://localhost:8080/books/${id}`);
        if (res.status === 404) {
          alert("Book not found");
          router.push("/");
          return;
        }
        const body = await res.json();
        setBook(body.data);
      } catch (err) {
        console.error("Failed to fetch book:", err);
      }
    };
    fetchBook();
  }, [id, router]);

  if (!book) return <div className="loading">Loading...</div>;

  return (
    <div className="container">
      <h1 className="page-title">{book.title}</h1>
      <div className="card">
        <p><strong>Author:</strong> {book.author}</p>
        <p><strong>Year:</strong> {book.publish_year}</p>
        {book.created_at && <p><strong>Created At:</strong> {new Date(book.created_at).toLocaleString()}</p>}
        {book.updated_at && <p><strong>Updated At:</strong> {new Date(book.updated_at).toLocaleString()}</p>}
      </div>
      <button onClick={() => router.push("/")} className="btn">
        Back to Dashboard
      </button>
    </div>
  );
}

