import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.jsx'

console.log('🚀 React app starting...')

try {
  const root = ReactDOM.createRoot(document.getElementById('root'));
  root.render(
    <React.StrictMode>
      <App />
    </React.StrictMode>
  );
  console.log('✅ App rendered successfully');
} catch (error) {
  console.error('❌ Error rendering app:', error);
}