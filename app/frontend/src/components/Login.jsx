import React, { useState } from 'react';
import '../styles/Auth.css';

const Login = ({ onToggleForm, onLoginSuccess }) => {
  const [formData, setFormData] = useState({
    email: '',
    password: ''
  });

  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    
    try {
      const response = await fetch('/api/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData)
      });

      const data = await response.json();
      
      if (response.ok) {
        alert('Вход успешен!');
        onLoginSuccess();
      } else {
        alert(data.error || 'Ошибка входа');
      }
    } catch (error) {
      alert('Ошибка подключения к серверу');
    }
  };

  return (
    <div className="auth-container">
      <div className="auth-header">
        <h1 className="auth-main-title">Альфа-Бизнес Ассистент</h1>
      </div>
      
      <div className="auth-content">
        <div className="auth-card">
          <h2 className="auth-title">Войдите в аккаунт</h2>
          <form onSubmit={handleSubmit} className="auth-form">
            <div className="input-group">
              <label className="input-label">Email</label>
              <input
                type="email"
                name="email"
                value={formData.email}
                onChange={handleChange}
                required
                className="auth-input"
                placeholder="example@mail.com"
              />
            </div>
            
            <div className="input-group">
              <label className="input-label">Пароль</label>
              <input
                type="password"
                name="password"
                value={formData.password}
                onChange={handleChange}
                required
                className="auth-input"
              />
            </div>
            
            <button type="submit" className="auth-button">
              Войти
            </button>
          </form>
          
          <p className="auth-switch">
            Нет аккаунта?{' '}
            <span className="switch-link" onClick={onToggleForm}>
              Зарегистрироваться
            </span>
          </p>
        </div>
      </div>
    </div>
  );
};

export default Login;