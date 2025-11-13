import React, { useState } from 'react';
import Login from './components/Login.jsx';
import Register from './components/Register.jsx';
import './styles/Auth.css';

function App() {
  const [isLogin, setIsLogin] = useState(true);

  const toggleForm = () => {
    setIsLogin(!isLogin);
  };

  return (
    <div className="App">
      {isLogin ? (
        <Login onToggleForm={toggleForm} />
      ) : (
        <Register onToggleForm={toggleForm} />
      )}
    </div>
  );
}

export default App;