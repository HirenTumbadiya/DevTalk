import React from 'react';
import { useFormik } from 'formik';
import * as Yup from 'yup';

const LoginForm = () => {
  const formik = useFormik({
    initialValues: {
      email: '',
      password: '',
    },
    validationSchema: Yup.object({
      email: Yup.string().email('Invalid email address').required('Required'),
      password: Yup.string().required('Required'),
    }),
    onSubmit: async (values, { setSubmitting, resetForm }) => {
      try {
        const response = await fetch('http://localhost:8000/login', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(values),
        });

        if (response.ok) {
          const responseData = await response.text();
          if (responseData) {
            const data = JSON.parse(responseData);
            console.log('Login successful:', data);
          } else {
            console.log('Login response is empty');
          }
        } else {
          // Login failed
          const error = await response.text();
          console.log('Login failed:', error);
        }
      } catch (error) {
        console.error('Error occurred during login:', error);
      }

      // Reset the form after login attempt
      resetForm();

      // Set submitting state to false
      setSubmitting(false);
    },
  });

  return (
    <form className="flex flex-col gap-2" onSubmit={formik.handleSubmit}>
    <div className='flex flex-col py-5'>
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
      <button className=' bg-sky-600 text-white rounded-2xl py-2' type="submit">Login</button>
    </form>
  );
};

export default LoginForm;
