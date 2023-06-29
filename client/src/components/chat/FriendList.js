import React, { useState, useEffect } from 'react';
import axios from 'axios';

const FriendList = () => {
  const [searchTerm, setSearchTerm] = useState('');
  const [searchResults, setSearchResults] = useState(null);
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
      setSearchResults([]); // Set searchResults to an empty array when no users are found or there is an error
    }
  };
  

  const handleSendRequest = async (receiver) => {
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
      <h1 className='flex justify-center text-center items-center text-3xl font-semibold text-white'>Friend List</h1>
      
      <div className="p-5">
        <input
          className='w-full py-3 px-2 rounded-xl bg-[#16171B] text-white focus:outline-none'
          type="text"
          placeholder="Search for users"
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
        />
      </div>
      
      {searchResults !== null && searchResults.length > 0 ? (
        <ul>
          {searchResults.map((user) => (
            <li key={user.id} className="flex items-center py-2">
              <img
                src={user.image} // Replace with the actual image source for the user
                alt={user.name}
                className="w-10 h-10 rounded-full mr-4"
              />
              <div>
                <p className="font-bold">{user.name}</p>
              </div>
              <button
                onClick={() => handleSendRequest(user.id)}
                className="ml-auto bg-blue-500 hover:bg-blue-700 text-white py-2 px-4 rounded"
              >
                <i className="fas fa-user-plus"></i> {/* Add the appropriate icon for sending a request */}
              </button>
            </li>
          ))}
        </ul>
      ) : (
        searchResults === null ? (
          <p>Loading...</p>
        ) : (
          <p>No users found</p>
        )
      )}
      
      {notification && (
        <div>
          <p>{notification}</p>
        </div>
      )}
    </div>
  );
};

export default FriendList;
