import { Formik, FormikProps } from "formik";
import React from "react";
import { SubmitButton } from "../elements/forms/button";
import { TextInput } from "../elements/forms/text";
import { defaultLayout } from "../elements/layouts/default";
import { AuthService } from "../lib/auth/auth_service";
import { signUp } from "../lib/services/api/sign_up";
import { emailRegex } from "./login";

export type SignUpInputs = {
  email: string
  password: string
  first: string
  last: string
}

function Page() {
  AuthService.redirectIfAuthenticated();

  const initialValues = { email: "", password: "", first: "", last: "" };

  const validate = (values: SignUpInputs) => {
    let errors: Partial<SignUpInputs> = {};

    if (!values.email) {
      errors.email = "Required";
    } else if (!emailRegex.test(values.email)) {
      errors.email = "Invalid email address";
    }

    return errors;
  };

  const onSubmit = async (values, { setSubmitting, setStatus }) => {
    setStatus("hi error");
    await signUp(values);
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
    }: FormikProps<SignUpInputs>) => <form className="container mx-auto max-w-sm" onSubmit={handleSubmit}>
      {status ? status : null}
      <TextInput type="text"
                 label="First"
                 name="first"
                 touched={touched.first}
                 value={values.first}
                 error={errors.first}
                 handleBlur={handleBlur}
                 handleChange={handleChange}
                 submitting={isSubmitting}
                 required
      />
      <TextInput type="text"
                 label="Last"
                 name="last"
                 touched={touched.last}
                 value={values.last}
                 error={errors.last}
                 handleBlur={handleBlur}
                 handleChange={handleChange}
                 submitting={isSubmitting}
                 required
      />
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

export default defaultLayout(Page);
