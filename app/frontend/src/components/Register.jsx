import React, { useState } from 'react';
import '../styles/Auth.css';

const Register = ({ onToggleForm, onRegisterSuccess }) => {
  const [formData, setFormData] = useState({
    name: '',
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
      const response = await fetch('/api/register', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData)
      });

      const data = await response.json();
      
      if (response.ok) {
        alert('Регистрация успешна!');
        onRegisterSuccess();
      } else {
        alert(data.error || 'Ошибка регистрации');
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
          <h2 className="auth-title">Создайте аккаунт</h2>
          <form onSubmit={handleSubmit} className="auth-form">
            <div className="input-group">
              <label className="input-label">Имя</label>
              <input
                type="text"
                name="name"
                value={formData.name}
                onChange={handleChange}
                required
                className="auth-input"
              />
            </div>
            
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
              Зарегистрироваться
            </button>
          </form>
          
          <p className="auth-switch">
            Уже есть аккаунт?{' '}
            <span className="switch-link" onClick={onToggleForm}>
              Войти
            </span>
          </p>
        </div>
      </div>
    </div>
  );
};

export default Register;