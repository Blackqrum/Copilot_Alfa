import React from "react";
import "../styles/HistoryPanel.css";

export default function HistoryPanel({ isOpen, onClose }) {
    return (
        <div className={`history-panel ${isOpen ? "open" : ""}`}>
            <div className="history-header">
                <span>История</span>
                <button className="close-btn" onClick={onClose}>×</button>
            </div>

            <div className="history-content">
                <h3>Сегодня</h3>
                <div className="history-item">оптимизация задач…</div>

                <h3>Вчера</h3>
                <div className="history-item">...</div>
            </div>
        </div>
    );
}
