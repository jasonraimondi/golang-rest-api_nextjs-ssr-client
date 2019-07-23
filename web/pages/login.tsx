import {Formik, FormikProps} from "formik";
import React from "react";
import {TextInput} from "../elements/forms/text";
import {defaultLayout} from "../elements/layouts/default";
import {AuthService} from "../lib/auth/auth_service";
import {SubmitButton} from "../elements/forms/button";

function LoginPage() {
  AuthService.redirectIfAuthenticated();
  return <>
    <SoFo/>
  </>;
}

export default defaultLayout(LoginPage);

function SoFo() {
  const initialValues = {email: "", password: ""};
  const validate = values => {
    let errors: any = {};
    if (!values.email) {
      errors.email = "Required";
    }
    return errors;
  };
  const onSubmit = (values, {setSubmitting}) => {
    setTimeout(() => {
      const foo = JSON.stringify(values, null, 2);
      console.log(foo);
      setSubmitting(false);
    }, 400);
  };
  return <Formik
    initialValues={initialValues}
    validate={validate}
    onSubmit={onSubmit}
  >
    {LoginForm}
  </Formik>;
}

export type LoginInputs = {
  email: string
  password: string
}

function LoginForm({
  values,
  errors,
  touched,
  handleChange,
  handleBlur,
  handleSubmit,
  isSubmitting,
}: FormikProps<LoginInputs>) {
  return <form onSubmit={handleSubmit}>
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
    <SubmitButton label="Submit" type="submit" disabled={isSubmitting} />
  </form>;
}