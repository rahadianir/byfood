"use client";

export default function Button({callFunction, label, style}:{callFunction?: () => void, label?: string, style?: React.CSSProperties}) {
    return (
        <button style={style} onClick={callFunction}>{label}</button>
    );
}