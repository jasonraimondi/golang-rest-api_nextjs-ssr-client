import { Formik, FormikProps } from "formik";
import { NextPageContext } from "next";
import React, { CSSProperties } from "react";
import { SubmitButton } from "../components/forms/button";
import { TextInput } from "../components/forms/text";
import { APP_ROUTES } from "../lib/routes";
import { signUp } from "../lib/services/api/sign_up";
import { redirectIfAuthenticated, redirectToLogin } from "../lib/services/redirect_service";
import { emailRegex } from "./login";

export type SignUpInputs = {
  email: string
  password: string
  first: string
  last: string
}

function Page() {
  const initialValues: SignUpInputs = { email: "", password: "", first: "", last: "" };

  const validate = (values: SignUpInputs) => {
    let errors: Partial<SignUpInputs> = {};

    if (!values.email) {
      errors.email = "Required";
    } else if (!emailRegex.test(values.email)) {
      errors.email = "Invalid email address";
    }

    return errors;
  };

  const onSubmit = async (values: any, { setSubmitting, setStatus }: any) => {
    const res: any = await signUp(values);
    if (res.error) {
      setStatus(res.error);
    } else {
      setStatus("Thank you for signing up!");
      await redirectToLogin();
    }
    setSubmitting(false);
  };

  const gridStyle: CSSProperties = {
    display: "grid",
    gridGap: 10,
    gridTemplateColumns: "1fr 1fr",
  };

  return <>
    <a href={APP_ROUTES.auth.login.create({})}>Back to Login</a>
    <div className="flex flex-col justify-center h-full">
      <Formik initialValues={initialValues}
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
          <div style={gridStyle}>
            <TextInput type="text"
                       label="First"
                       name="first"
                       touched={touched.first}
                       value={values.first}
                       error={errors.first}
                       handleBlur={handleBlur}
                       handleChange={handleChange}
                       submitting={isSubmitting}
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
            />
          </div>
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
          />
          <SubmitButton label="Submit" type="submit" disabled={isSubmitting}/>
        </form>}
      </Formik>
    </div>
  </>;
}

Page.getInitialProps = async (ctx: NextPageContext) => {
  await redirectIfAuthenticated(ctx);
  return {};
};

export default Page;
