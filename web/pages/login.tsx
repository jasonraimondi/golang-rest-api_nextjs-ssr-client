import {Formik, FormikProps} from "formik";
import React from "react";
import {TextInput} from "../elements/forms/text";
import {defaultLayout} from "../elements/layouts/default";
import {AuthService} from "../lib/auth/auth_service";
import {SubmitButton} from "../elements/forms/button";

function LoginPage() {
  AuthService.redirectIfAuthenticated();
  return <>
    <LoginForm/>
  </>;
}

const emailRegex = /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,4}$/i;

export type LoginInputs = {
  email: string
  password: string
}

function LoginForm() {
  const initialValues = {email: "", password: ""};

  const validate = (values: LoginInputs) => {
    let errors: Partial<LoginInputs> = {};

    if (!values.email) {
      errors.email = 'Required';
    } else if (!emailRegex.test(values.email)) {
      errors.email = 'Invalid email address';
    }

    return errors;
  };

  const onSubmit = async (values, {setSubmitting, setError }) => {
    setError("hi error");
    await AuthService.login(values);
    setSubmitting(false);
  };

  return <Formik
    initialValues={initialValues}
    validate={validate}
    onSubmit={onSubmit}
  >
    {({
      values,
      error,
      errors,
      touched,
      handleChange,
      handleBlur,
      handleSubmit,
      isSubmitting,
    }: FormikProps<LoginInputs>) => <form className="container mx-auto max-w-sm" onSubmit={handleSubmit}>
      {error ? error : null}
      <TextInput type="email"
                 label="Email"
                 name="email"
                 touched={touched.email}
                 value={values.email}
                 error={errors.email}
                 handleBlur={handleBlur}
                 handleChange={handleChange}
                 submitting={isSubmitting}
                 required
      />
      <TextInput type="password"
                 label="Password"
                 name="password"
                 touched={touched.password}
                 value={values.password}
                 error={errors.password}
                 handleBlur={handleBlur}
                 handleChange={handleChange}
                 submitting={isSubmitting}
                 required
      />
      <SubmitButton label="Submit" type="submit" disabled={isSubmitting}/>
    </form>}
  </Formik>;
}


export default defaultLayout(LoginPage);
