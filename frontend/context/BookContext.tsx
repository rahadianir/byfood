"use client";

import React, { createContext, useContext, useState, ReactNode } from "react";

export interface Book {
  id: number;
  title: string;
  author: string;
  publish_year: number;
}

interface BookContextType {
  books: Book[];
  fetchBooks: () => Promise<void>;
  addBook: (book: Omit<Book, "id">) => Promise<void>;
  updateBook: (id: number, updatedBook: Omit<Book, "id">) => Promise<void>;
  deleteBook: (id: number) => Promise<void>;
}

interface PaginationType {

}

const BookContext = createContext<BookContextType | undefined>(undefined);

export const BookProvider = ({ children }: { children: ReactNode }) => {
  const [books, setBooks] = useState<Book[]>([]);
  

  const fetchBooks = async () => {
    try {
        const baseURL = new URL("http://localhost:8080/books")

        const res = await fetch(baseURL);
        
        if (!res.ok) throw new Error("Failed to fetch books");
        
        const body = await res.json();
        
        setBooks(Array.isArray(body.data) ? body.data : []);
    } catch (error) {
        console.error("Error fetching books:", error);
        setBooks([]);
    }
  };

  const addBook = async (book: Omit<Book, "id">) => {
    try {
      const res = await fetch("http://localhost:8080/books", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(book),
      });
      if (!res.ok) throw new Error("Failed to add book");
      const newBook = await res.json();
      setBooks((prev) => [...prev, newBook]);
    } catch (error) {
      console.error("Error adding book:", error);
    }
  };

  const updateBook = async (id: number, updatedBook: Omit<Book, "id">) => {
    try {
      const res = await fetch(`http://localhost:8080/books/${id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(updatedBook),
      });
      if (!res.ok) throw new Error("Failed to update book");
      const updated = await res.json();
      setBooks((prev) =>
        prev.map((book) => (book.id === id ? updated : book))
      );
    } catch (error) {
      console.error("Error updating book:", error);
    }
  };

  const deleteBook = async (id: number) => {
    try {
      const res = await fetch(`http://localhost:8080/books/${id}`, {
        method: "DELETE",
      });
      if (!res.ok) throw new Error("Failed to delete book");
      setBooks((prev) => prev.filter((book) => book.id !== id));
    } catch (error) {
      console.error("Error deleting book:", error);
    }
  };

  return (
    <BookContext.Provider
      value={{ books, fetchBooks, addBook, updateBook, deleteBook }}
    >
      {children}
    </BookContext.Provider>
  );
};

export const useBooks = (): BookContextType => {
  const context = useContext(BookContext);
  if (!context) throw new Error("useBooks must be used within a BookProvider");
  return context;
};
