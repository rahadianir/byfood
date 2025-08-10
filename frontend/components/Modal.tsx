"use client";

import React from "react";

interface ModalProps {
  isOpen: boolean;
  onClose?: () => void;
  title?: string;
  children: React.ReactNode;
}

export default function Modal({ isOpen, onClose, title, children }: ModalProps) {
  if (!isOpen) return null;

  return (
    <div style={overlayStyle}>
      <div style={modalStyle}>
        {/* Modal Header */}
        <div style={headerStyle}>
          <h2 style={titleStyle}>{title}</h2>
          <button onClick={onClose} style={closeButtonStyle}>✖</button>
        </div>

        {/* Modal Body */}
        <div style={bodyStyle}>
          {children}
        </div>
      </div>
    </div>
  );
}

// Dark overlay background
const overlayStyle: React.CSSProperties = {
  position: "fixed",
  top: 0,
  left: 0,
  width: "100%",
  height: "100%",
  background: "rgba(0, 0, 0, 0.75)",
  display: "flex",
  alignItems: "center",
  justifyContent: "center",
  zIndex: 1000,
};

// Dark modal card
const modalStyle: React.CSSProperties = {
  background: "#1e1e1e",
  padding: "20px",
  borderRadius: "10px",
  width: "400px",
  boxShadow: "0 5px 25px rgba(0, 0, 0, 0.6)",
  color: "#f5f5f5",
};

// Header styles
const headerStyle: React.CSSProperties = {
  display: "flex",
  justifyContent: "space-between",
  alignItems: "center",
  borderBottom: "1px solid #333",
  paddingBottom: "8px",
};

const titleStyle: React.CSSProperties = {
  margin: 0,
  fontSize: "1.25rem",
  fontWeight: "bold",
  color: "#ffffff",
};

// Close button in dark mode
const closeButtonStyle: React.CSSProperties = {
  background: "transparent",
  border: "none",
  fontSize: "18px",
  cursor: "pointer",
  color: "#bbb",
  transition: "color 0.2s",
};

// Body section
const bodyStyle: React.CSSProperties = {
  marginTop: "12px",
};
