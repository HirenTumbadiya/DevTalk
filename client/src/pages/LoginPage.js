import { useState } from 'react';
import RegisterForm from '../components/RegisterForm';
import LoginForm from '../components/LoginForm';
import logo from '../assets/devtalk.jpg';

const LoginPage = ({setIsLoggedIn}) => {
  const [isRegistering, setIsRegistering] = useState(false);

  const toggleForm = () => {
    setIsRegistering(!isRegistering);
  };

  return (
    <>
      <div className="flex w-full">
        <div className="bg-black w-6/12 text-white h-screen flex flex-col">
          <div className="flex-1 flex flex-col justify-center items-center">
            <h1 className="text-6xl font-bold">DEVTALKS</h1>
            <p>connect chat & develop</p>
          </div>
          <div className="flex-1 overflow-hidden">
            <img
              src={logo}
              className="object-cover rounded-2xl shadow-teal-300 shadow-2xl h-full"
              alt="Logo"
            />
          </div>
        </div>

        <div className="w-6/12 flex justify-center items-center h-screen">
          <div className="w-3/5 h-fit p-5 shadow-2xl rounded-2xl transition-opacity duration-500">
            <h1 className="text-2xl font-bold flex justify-center text-center p-5 pb-0">
              {isRegistering ? 'Create Your Account' : 'Login First to Your Account'}
            </h1>
            {isRegistering ? <RegisterForm toggleForm={toggleForm} /> : <LoginForm setIsLoggedIn={setIsLoggedIn} />}
            <p className="text-sm p-5 flex justify-center items-center">
              {isRegistering ? 'Already have an account? ' : 'Don’t have an account? '}
              <button className="text-sky-600" onClick={toggleForm}>
                {isRegistering ? ' Login' : ' Register Here'}
              </button>
            </p>
          </div>
        </div>
      </div>
    </>
  );
};

export default LoginPage;
