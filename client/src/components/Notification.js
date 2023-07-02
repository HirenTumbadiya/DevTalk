import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { FaCheckCircle, FaTimesCircle, FaChevronUp, FaChevronDown } from 'react-icons/fa';

const Notification = () => {
  const [friendRequests, setFriendRequests] = useState([]);
  const [isExpanded, setIsExpanded] = useState(true);

  useEffect(() => {
    const socket = new WebSocket('ws://localhost:3000/ws');

    socket.onopen = () => {
      console.log('WebSocket connection established');
    };

    socket.onmessage = (event) => {
      const friendRequest = JSON.parse(event.data);
      const currentUserId = localStorage.getItem('id');
      if (friendRequest.recipientId === currentUserId) {
        setFriendRequests((prevFriendRequests) => (
          prevFriendRequests ? [...prevFriendRequests, friendRequest] : [friendRequest]
        ));
      }
    };

    return () => {
      socket.close();
    };
  }, []);

  useEffect(() => {
    const userID = localStorage.getItem('id');
    const fetchFriendRequests = async () => {
      try {
        const response = await axios.get(`http://localhost:8000/friend-requests?userId=${userID}`);
        console.log(response);
        setFriendRequests(response.data);
      } catch (error) {
        console.error('Error fetching friend requests:', error);
      }
    };

    fetchFriendRequests();
  }, []);

  const handleAcceptRequest = async (requestId) => {
    try {
      await axios.post(`http://localhost:8000/friend-requests/accept`, { requestId });
      // Optional: Update the local state or perform any necessary actions
      // after the friend request is accepted
    } catch (error) {
      console.error('Error accepting friend request:', error);
    }
  };

  const handleRejectRequest = async (requestId) => {
    try {
      await axios.post(`http://localhost:8000/friend-requests/reject`, { requestId });
      // Optional: Update the local state or perform any necessary actions
      // after the friend request is rejected
    } catch (error) {
      console.error('Error rejecting friend request:', error);
    }
  };

  const toggleExpand = () => {
    setIsExpanded((prevExpanded) => !prevExpanded);
  };

  return (
    <div>
      <h2 className="flex justify-center text-center items-center text-3xl font-semibold text-white p-5">Notifications</h2>
      <div className="flex items-center justify-between px-5 cursor-pointer" onClick={toggleExpand}>
        <p className='text-white'>friends requests</p>
        {isExpanded ? (
          <FaChevronUp size={20} color={"#fafafa"} />
        ) : (
          <FaChevronDown size={20} color={"#fafafa"} />
        )}
      </div>
      {isExpanded && (
        <>
          {friendRequests && friendRequests.length > 0 ? (
            <ul className="px-5 mt-2">
              {friendRequests.map((request) => (
                <li key={request.id} className="flex items-center p-3 hover:bg-gray-700 mx-3 rounded-xl text-white">
                  <div className="flex items-center flex-grow">
                    <strong>{request.username}</strong>
                  </div>
                  <div className="flex items-center">
                    <button
                      className="text-green-500 hover:text-green-700 focus:outline-none"
                      onClick={() => handleAcceptRequest(request.id)}
                    >
                      <FaCheckCircle size={24} />
                    </button>
                    <button
                      className="text-red-500 hover:text-red-700 focus:outline-none ml-2"
                      onClick={() => handleRejectRequest(request.id)}
                    >
                      <FaTimesCircle size={24} />
                    </button>
                  </div>
                </li>
              ))}
            </ul>
          ) : (
            <p className="text-gray-500 p-5">No notifications</p>
          )}
        </>
      )}
    </div>
  );
};

export default Notification;
