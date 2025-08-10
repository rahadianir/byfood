"use client";

import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import { useBooks } from "@/context/BookContext";
import Modal from "../components/Modal";
import AddBookForm from "@/components/AddBookForm";
import EditBookForm from "@/components/EditBookForm";

export default function DashboardPage() {
  const router = useRouter();
  const { books, fetchBooks, deleteBook } = useBooks();

  // Modal states
  const [isAddModalOpen, setIsAddModalOpen] = useState(false);
  const [isEditModalOpen, setIsEditModalOpen] = useState(false);
  const [isDeleteModalOpen, setIsDeleteModalOpen] = useState(false);

  // Selected book for edit/delete
  const [selectedBook, setSelectedBook] = useState<any>(null);

  useEffect(() => {
    fetchBooks();
  }, [fetchBooks]);

  const handleDeleteConfirm = async () => {
    if (selectedBook) {
      await deleteBook(selectedBook.id);
      fetchBooks();
    }
    setIsDeleteModalOpen(false);
    setSelectedBook(null);
  };

  return (
    <div style={{ padding: "20px" }}>
      {/* Header */}
      <div
        style={{
          display: "flex",
          justifyContent: "space-between",
          marginBottom: "20px",
        }}
      >
        <h1>📚 Book Dashboard</h1>
        <button
          onClick={() => setIsAddModalOpen(true)}
          style={{
            padding: "8px 16px",
            background: "#4CAF50",
            color: "white",
            border: "none",
            borderRadius: "4px",
          }}
        >
          ➕ Add Book
        </button>
      </div>

      {/* Book List */}
      <table width="100%" border={1} cellPadding={8} style={{ borderCollapse: "collapse" }}>
        <thead>
          <tr style={{ background: "black" }}>
            <th>Title</th>
            <th>Author</th>
            <th>Year</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {books.map((book) => (
            <tr key={book.id}>
              <td>{book.title}</td>
              <td>{book.author}</td>
              <td align="center">{book.publish_year}</td>
              <td align="right">
                <button
                  onClick={() => router.push(`/books/${book.id}`)}
                  style={{ marginRight: "5px" }}
                >
                  View
                </button>
                <button
                  onClick={() => {
                    setSelectedBook(book);
                    setIsEditModalOpen(true);
                  }}
                  style={{ marginRight: "5px" }}
                >
                  Edit
                </button>
                <button
                  onClick={() => {
                    setSelectedBook(book);
                    setIsDeleteModalOpen(true);
                  }}
                  style={{ background: "red", color: "white" }}
                >
                  Delete
                </button>
              </td>
            </tr>
          ))}

          {books.length === 0 && (
            <tr>
              <td colSpan={4} style={{ textAlign: "center" }}>
                No books found
              </td>
            </tr>
          )}
        </tbody>
      </table>

      

      {/* Add Modal */}
      <Modal
        isOpen={isAddModalOpen}
        onClose={() => setIsAddModalOpen(false)}
        title="Add New Book"
      >
        <AddBookForm
          onSuccess={() => {
            fetchBooks();
            setIsAddModalOpen(false);
          }}
        />
      </Modal>

      {/* Edit Modal */}
      <Modal
        isOpen={isEditModalOpen}
        onClose={() => setIsEditModalOpen(false)}
        title="Edit Book"
      >
        {selectedBook && (
          <EditBookForm
            book={selectedBook}
            onSuccess={() => {
              fetchBooks();
              setIsEditModalOpen(false);
              setSelectedBook(null);
            }}
          />
        )}
      </Modal>

      {/* Delete Confirmation Modal */}
      <Modal
        isOpen={isDeleteModalOpen}
        onClose={() => setIsDeleteModalOpen(false)}
        title="Delete Book"
      >
        <p>Are you sure you want to delete &quot;{selectedBook?.title}&quot;?</p>
        <div style={{ marginTop: "20px", textAlign: "right" }}>
          <button onClick={() => setIsDeleteModalOpen(false)} style={{ marginRight: "10px" }}>
            Cancel
          </button>
          <button onClick={handleDeleteConfirm} style={{ background: "red", color: "white" }}>
            Delete
          </button>
        </div>
      </Modal>
    </div>
  );
}
