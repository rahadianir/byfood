"use client";

import React from "react";

interface ModalProps {
  onClose?: () => void;
  title?: string;
  children: React.ReactNode;
}

export default function Modal({ onClose, title, children }: ModalProps) {

  return (
    <div style={overlayStyle}>
      <div style={modalStyle}>
        {/* Modal Header */}
        <div style={{ display: "flex", justifyContent: "space-between", alignItems: "center" }}>
          <h2>{title}</h2>
          <button onClick={onClose} style={closeButtonStyle}>✖</button>
        </div>

        {/* Modal Body */}
        <div style={{ marginTop: "10px" }}>
          {children}
        </div>
      </div>
    </div>
  );
}

const overlayStyle: React.CSSProperties = {
  position: "fixed",
  top: 0,
  left: 0,
  width: "100%",
  height: "100%",
  background: "rgba(0,0,0,0.5)",
  display: "flex",
  alignItems: "center",
  justifyContent: "center",
  zIndex: 1000,
};

const modalStyle: React.CSSProperties = {
  background: "white",
  padding: "20px",
  borderRadius: "8px",
  width: "400px",
  boxShadow: "0 5px 15px rgba(0,0,0,0.3)",
};

const closeButtonStyle: React.CSSProperties = {
  background: "transparent",
  border: "none",
  fontSize: "18px",
  cursor: "pointer",
};
