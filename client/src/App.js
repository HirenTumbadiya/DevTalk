import React, { useState } from 'react';
import LoginPage from './pages/LoginPage';
import ChatRoomPage from './pages/ChatRoomPage';

const App = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(true);

  return (
    <div>
      {isLoggedIn ? <ChatRoomPage /> : <LoginPage setIsLoggedIn={setIsLoggedIn} />}
    </div>
  );
};

export default App;
