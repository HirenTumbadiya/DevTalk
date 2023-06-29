import React, { useEffect, useState } from 'react';
import axios from 'axios';

const Notification = () => {
  const [friendRequests, setFriendRequests] = useState([]);

  useEffect(() => {
    const fetchFriendRequests = async () => {
      try {
        const response = await axios.get('http://localhost:8000/friend-requests');
        setFriendRequests(response.data);
      } catch (error) {
        console.error('Error fetching friend requests:', error);
      }
    };

    fetchFriendRequests();
  }, []);

  const handleAcceptRequest = async (requestId) => {
    try {
      await axios.post(`http://localhost:8000/friend-requests/${requestId}/accept`);
      // Optional: Update the local state or perform any necessary actions
      // after the friend request is accepted
    } catch (error) {
      console.error('Error accepting friend request:', error);
    }
  };

  return (
    <div>
      <h2 className="text-xl font-bold mb-4">Notifications</h2>
      {friendRequests.length > 0 ? (
        <ul>
          {friendRequests.map((request) => (
            <li key={request.id} className="p-2 mb-2 bg-gray-200 rounded-lg">
              <strong>{request.senderName}</strong> sent you a friend request.
              <button
                className="ml-2 bg-blue-500 hover:bg-blue-700 text-white py-2 px-4 rounded"
                onClick={() => handleAcceptRequest(request.id)}
              >
                Accept
              </button>
            </li>
          ))}
        </ul>
      ) : (
        <p className="text-gray-500">No friend requests</p>
      )}
    </div>
  );
};

export default Notification;
