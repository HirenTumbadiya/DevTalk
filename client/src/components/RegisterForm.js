import React from 'react';
import { useFormik } from 'formik';
import * as Yup from 'yup';

const RegisterForm = () => {
  const formik = useFormik({
    initialValues: {
      name: '',
      email: '',
      password: '',
    },
    validationSchema: Yup.object({
      name: Yup.string().required('Required'),
      email: Yup.string().email('Invalid email address').required('Required'),
      password: Yup.string().required('Required'),
    }),
    onSubmit: (values) => {
      // Handle form submission (e.g., send registration request to backend)
      console.log(values);
    },
  });

  return (
    <form className="flex flex-col gap-2" onSubmit={formik.handleSubmit}>
    <div className='flex flex-col py-5'>
      <label htmlFor="name">Name</label>
      <input
      className='py-2 border px-2 rounded-2xl'
        type="text"
        id="name"
        placeholder="Username"
        {...formik.getFieldProps('name')}
      />
      {formik.touched.name && formik.errors.name ? (
        <div className='text-red-600 absolute mt-16'>{formik.errors.name}</div>
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
