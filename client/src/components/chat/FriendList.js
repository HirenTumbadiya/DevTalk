import React, { useState, useEffect } from 'react';
import axios from 'axios';

const FriendList = () => {
  const [searchTerm, setSearchTerm] = useState('');
  const [searchResults, setSearchResults] = useState([]);
  const [notification, setNotification] = useState(null);
  const currentUserId = "64982623a6b4dcd03b6884c8"; // Placeholder value, replace with your actual current user ID
  const friendRequests = []; // Placeholder value, replace with your actual friend requests data

  const handleSearch = async () => {
    try {
      const response = await axios.post(`http://localhost:8000/users/search?username=${searchTerm}`);
      console.log(response);
      setSearchResults(response.data);
    } catch (error) {
      console.error('Error searching for users:', error);
    }
  };

  const handleSendRequest = async (sender, receiver) => {
    try {
      const response = await axios.post('http://localhost:8000/friend-requests', {
        sender: currentUserId,
        receiver: receiver,
    });
    console.log(response);
    setNotification(response.data.message);
  } catch (error) {
    console.error('Error sending friend request:', error);
  }
};

  const handleAcceptRequest = async (requestId) => {
    try {
      const response = await axios.post(`/friend-request/accept/${requestId}`);
      setNotification(response.data.message);
    } catch (error) {
      console.error('Error accepting friend request:', error);
    }
  };

  useEffect(() => {
    // Perform the search when the search term changes
    handleSearch();
  }, [searchTerm]);

  return (
    <div>
      <h1>Friend List</h1>
      
      <div>
        <input
          type="text"
          placeholder="Search for users"
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
        />
      </div>
      
      <ul>
        {searchResults.map((user) => (
          <li key={user.id}>
            {user.name} - {user.email}
            <button onClick={() => handleSendRequest(user.id)}>Send Request</button>
          </li>
        ))}
      </ul>
      
      {notification && (
        <div>
          <p>{notification}</p>
        </div>
      )}
      
      <h2>Friend Requests</h2>
      
      <ul>
        {friendRequests.map((request) => (
          <li key={request.id}>
            {request.senderName} sent you a friend request.
            <button onClick={() => handleAcceptRequest(request.id)}>Accept</button>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default FriendList;
