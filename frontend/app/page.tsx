import Button from "@/components/Button";
import Modal from "@/components/Modal";

export default function Home() {
  const books = [{
    id: 1,
    title: "One Piece",
    author: "Eiichiro Oda",
    year: 1997,
  },{
    id: 2,
    title: "Naruto",
    author: "Masashi Kishimoto",
    year: 1997,
  },{
    id: 3,
    title: "Bleach",
    author: "Tite Kubo",
    year: 1997,
  }]
  return (
    <div>
      <Modal>asdf</Modal>
      <div style={{ padding: "20px" }}>
      {/* ===== Header ===== */}
      <div style={{ display: "flex", justifyContent: "space-between", alignItems: "center" }}>
        <h1>📚 Book Dashboard</h1>
        <Button 
          style={{ padding: "8px 16px", background: "#4CAF50", color: "white", border: "none", borderRadius: "4px" }}
          label="➕ Add New Book"
        />
      </div>

      {/* ===== Book List ===== */}
      <div style={{ marginTop: "20px" }}>
        {books.length === 0 ? (
          <p>No books found.</p>
        ) : (
          <table style={{ width: "100%", borderCollapse: "collapse" }}>
            <thead>
              <tr style={{ borderBottom: "2px solid #ccc" }}>
                <th style={{ textAlign: "left", padding: "8px" }}>Title</th>
                <th style={{ textAlign: "left", padding: "8px" }}>Author</th>
                <th style={{ textAlign: "left", padding: "8px" }}>Year</th>
                <th style={{ textAlign: "right", padding: "8px" }}>Actions</th>
              </tr>
            </thead>
            <tbody>
              {books.map((book) => (
                <tr key={book.id} style={{ borderBottom: "1px solid #eee" }}>
                  <td style={{ padding: "8px" }}>{book.title}</td>
                  <td style={{ padding: "8px" }}>{book.author}</td>
                  <td style={{ padding: "8px" }}>{book.year}</td>
                  <td style={{ padding: "8px", textAlign:"right" }}>
                    <Button
                      style={{ marginRight: "6px" }}
                      label="🔍 View"
                    />
                    <Button
                      style={{ marginRight: "6px" }}
                      label="✏️ Edit"
                    />
                    <Button 
                      style={{color: "red"}} 
                      label="🗑️ Delete"
                    />
                    
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        )}
      </div>
    </div>
    </div>
  );
}
