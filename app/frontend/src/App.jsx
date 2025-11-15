import React, { useState } from 'react';
import Login from './components/Login.jsx';
import Register from './components/Register.jsx';
import Chat from './components/Chat.jsx';
import './styles/Auth.css';
import './styles/Chat.css';

function App() {
    const [currentView, setCurrentView] = useState('auth');
    const [isLogin, setIsLogin] = useState(true);

    const toggleForm = () => {
        setIsLogin(!isLogin);
    };

    const handleAuthSuccess = () => {
        setCurrentView('chat');
    };

    const handleBackToAuth = () => {
        setCurrentView('auth');
    };

    if (currentView === 'chat') {
        return <Chat onBackToAuth={handleBackToAuth} />;
    }

    return (
        <div className="App">
            {isLogin ? (
                <Login
                    onToggleForm={toggleForm}
                    onLoginSuccess={handleAuthSuccess}
                />
            ) : (
                <Register
                    onToggleForm={toggleForm}
                    onRegisterSuccess={handleAuthSuccess}
                />
            )}
        </div>
    );
}

export default App;