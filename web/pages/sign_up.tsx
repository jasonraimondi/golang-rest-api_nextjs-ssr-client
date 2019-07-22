import Head from "next/head";
import React from "react";
import { SignUpForm } from "../components/sign_up_form";
import { defaultLayout } from "../elements/layouts/default";
import { AuthService } from "../lib/auth/auth_service";

function SignUpPage() {
  AuthService.redirectIfAuthenticated();

  return (
    <>
      <Head>
        <title>SIGN UP TO WIN</title>
      </Head>
      <SignUpForm/>
    </>
  );
}

export default defaultLayout(SignUpPage);
