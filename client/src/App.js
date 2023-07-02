import React, { useState, useEffect } from 'react';
import LoginPage from './pages/LoginPage';
import ChatRoomPage from './pages/ChatRoomPage';

const App = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  useEffect(() => {
    const storedIsLoggedIn = localStorage.getItem('isLoggedIn');
    if (storedIsLoggedIn === 'true') {
      setIsLoggedIn(true);
    }
  }, []);

  const handleLogin = () => {
    setIsLoggedIn(true);
    localStorage.setItem('isLoggedIn', 'true');
  };

  const handleLogout = () => {
    setIsLoggedIn(false);
    localStorage.removeItem('isLoggedIn');
    localStorage.removeItem('id');
    localStorage.removeItem('token');
  };

  useEffect(() => {
    // Establish a WebSocket connection
    const webSocket = new WebSocket('ws://localhost:3000/ws');

    // Listen for WebSocket connection open event
    webSocket.onopen = () => {
      console.log('WebSocket connection established');
    };

    // Listen for WebSocket message event
    webSocket.onmessage = (event) => {
      const message = JSON.parse(event.data);
      // Handle the incoming WebSocket message here
      console.log('Received WebSocket message:', message);
    };

    // Clean up the WebSocket connection on component unmount
    return () => {
      webSocket.close();
    };
  }, []);

  return (
    <div>
      {isLoggedIn ? (
        <ChatRoomPage handleLogout={handleLogout} />
      ) : (
        <LoginPage setIsLoggedIn={handleLogin} />
      )}
    </div>
  );
};

export default App;
