import React, { useState, useRef, useEffect } from 'react';
import '../styles/Chat.css';
import HistoryPanel from "./HistoryPanel";

const Chat = ({ onBackToAuth }) => {
    const [isHistoryOpen, setIsHistoryOpen] = useState(false);
    const [messages, setMessages] = useState([]);
    const [inputMessage, setInputMessage] = useState('');
    const [isLoading, setIsLoading] = useState(false);
    const messagesEndRef = useRef(null);

    const scrollToBottom = () => {
        messagesEndRef.current?.scrollIntoView({ behavior: "smooth" });
    };

    useEffect(() => {
        scrollToBottom();
    }, [messages]);

    const handleSendMessage = async (e) => {
        e.preventDefault();
        if (!inputMessage.trim()) return;

        const userMessage = { text: inputMessage, sender: 'user' };
        setMessages(prev => [...prev, userMessage]);
        setInputMessage('');
        setIsLoading(true);

        setTimeout(() => {
            const botMessage = {
                text: `Это ответ на: "${inputMessage}". Я ваш AI-ассистент!`,
                sender: 'bot'
            };
            setMessages(prev => [...prev, botMessage]);
            setIsLoading(false);
        }, 1000);
    };

    return (
        <div className="chat-container">

            {/* Панель истории */}
            <HistoryPanel 
                isOpen={isHistoryOpen} 
                onClose={() => setIsHistoryOpen(false)} 
            />

            <div className="chat-header">
                <button className="history-btn" onClick={() => setIsHistoryOpen(true)}>
                    ☰ История
                </button>
                <h2>Альфа-Бизнес Ассистент</h2>
            </div>

            <div className="messages-container">
                {messages.map((message, index) => (
                    <div key={index} className={`message ${message.sender}`}>
                        {message.text}
                    </div>
                ))}

                {isLoading && (
                    <div className="message bot loading">
                        <div className="typing-indicator">
                            <span></span><span></span><span></span>
                        </div>
                    </div>
                )}

                <div ref={messagesEndRef} />
            </div>

            <form onSubmit={handleSendMessage} className="message-input-form">
                <input
                    type="text"
                    value={inputMessage}
                    onChange={(e) => setInputMessage(e.target.value)}
                    placeholder="Введите ваше сообщение..."
                    className="message-input"
                    disabled={isLoading}
                />
                <button type="submit" disabled={isLoading} className="send-button">
                    📨
                </button>
            </form>
        </div>
    );
};

export default Chat;
