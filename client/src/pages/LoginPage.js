import { useState } from 'react';
import RegisterForm from '../components/RegisterForm';
import LoginForm from '../components/LoginForm';

const LoginPage = () => {
  const [isRegistering, setIsRegistering] = useState(false);

  const toggleForm = () => {
    setIsRegistering(!isRegistering);
  };

  return (
    <>   
    <div className='flex w-full'>
    <div className='w-6/12'>

    </div>
    <div className='w-6/12 flex justify-center items-center h-screen'>
        <div className=' w-3/5 h-fit p-5  shadow-2xl rounded-2xl'>
      <h1 className='text-2xl font-bold flex justify-center text-center p-5 pb-0'>{isRegistering ? 'Create Your Account' : 'Login First to Your Account'}</h1>
      {isRegistering ? <RegisterForm /> : <LoginForm />}
      <p className='text-sm p-5 flex justify-center items-center'>
        {isRegistering
          ? 'Already have an account? '
          : 'Don’t have an account? '}
        <button className='text-sky-600' onClick={toggleForm}>
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
