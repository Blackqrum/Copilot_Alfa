import React, { useState } from 'react';
import '../styles/Auth.css';

const Login = ({ onToggleForm }) => {
  const [formData, setFormData] = useState({
    login: '',
    password: ''
  });

  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value
    });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log('Login data:', formData);
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
              <label className="input-label">Логин</label>
              <input
                type="text"
                name="login"
                value={formData.login}
                onChange={handleChange}
                required
                className="auth-input"
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