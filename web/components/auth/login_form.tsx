import { Formik, FormikActions, FormikProps } from "formik";
import React from "react";
import { login } from "../../lib/api/login";
import { SubmitButton } from "../forms/button";
import { TextInput } from "../forms/text";

export type LoginInputs = {
  email: string
  password: string
}

export const emailRegex = /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,4}$/i;

export function LoginForm() {
  const initialValues = { email: "", password: "" };

  const validate = (values: LoginInputs) => {
    let errors: Partial<LoginInputs> = {};

    if (!values.email) {
      errors.email = "Required";
    } else if (!emailRegex.test(values.email)) {
      errors.email = "Invalid email address";
    }

    return errors;
  };

  const onSubmit = async (values: LoginInputs, { setSubmitting, setStatus }: FormikActions<LoginInputs>) => {
    const errorMessage: any = await login((values));
    if (errorMessage) setStatus(errorMessage);
    setSubmitting(false);
  };

  return <Formik
    initialValues={initialValues}
    validate={validate}
    onSubmit={onSubmit}
  >
    {({
      values,
      status,
      errors,
      touched,
      handleChange,
      handleBlur,
      handleSubmit,
      isSubmitting,
    }: FormikProps<LoginInputs>) => <form className="container mx-auto max-w-sm" onSubmit={handleSubmit}>
      <p>{status ? status : null}</p>
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