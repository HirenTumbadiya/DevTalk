import React from 'react';
import { useFormik } from 'formik';
import * as Yup from 'yup';

const RegisterForm = ({ toggleForm }) => {
  const formik = useFormik({
    initialValues: {
      username: '',
      email: '',
      password: '',
    },
    validationSchema: Yup.object({
      username: Yup.string().required('Required'),
      email: Yup.string().email('Invalid email address').required('Required'),
      password: Yup.string().required('Required'),
    }),
    onSubmit: async (values) => {
      try {
        // Send registration request to backend
        const response = await fetch('http://localhost:8000/register', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(values),
        });
    
        if (response.ok) {
          try {
            const data = await response.json();
            console.log(data);
            setTimeout(() => {
              toggleForm(); // Toggle to the login form
            }, 1000);
          } catch (error) {
            console.log('Registration response error:', error);
            // Handle error when response body is empty or not valid JSON
          }
        } else {
          const errorData = await response.json();
          console.log('Registration failed:', errorData);
          // Handle registration error (e.g., display error message)
        }
      } catch (error) {
        console.log('Registration error:', error);
        // Handle registration error (e.g., display error message)
      }
    }
    
  });

  return (
    <form className="flex flex-col gap-2" onSubmit={formik.handleSubmit}>
    <div className='flex flex-col py-5'>
      <label htmlFor="username">Name</label>
      <input
      className='py-2 border px-2 rounded-2xl'
        type="text"
        id="username"
        placeholder="Username"
        {...formik.getFieldProps('username')}
      />
      {formik.touched.username && formik.errors.username ? (
        <div className='text-red-600 absolute mt-16'>{formik.errors.username}</div>
      ) : null}
        </div>
        <div className='flex flex-col pb-5'>
      <label htmlFor="email">Email</label>
      <input
      className='py-2 border px-2 rounded-2xl'
        type="email"
        id="email"
        placeholder="Email"
        {...formik.getFieldProps('email')}
      />
      {formik.touched.email && formik.errors.email ? (
        <div className='text-red-600 absolute mt-16'>{formik.errors.email}</div>
      ) : null}
        </div>
        <div className='flex flex-col pb-5'>
      <label htmlFor="password">Password</label>
      <input
      className='py-2 border px-2 rounded-2xl'
        type="password"
        id="password"
        placeholder="Password"
        {...formik.getFieldProps('password')}
      />
      {formik.touched.password && formik.errors.password ? (
        <div className='text-red-600 absolute mt-16'>{formik.errors.password}</div>
      ) : null}
</div>
      <button className=' bg-sky-600 text-white rounded-2xl py-2' type="submit">Register</button>
    </form>
  );
};

export default RegisterForm;
