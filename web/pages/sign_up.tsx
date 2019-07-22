import Head from "next/head";
import React, { useState } from "react";
import { AuthService } from "../lib/auth/auth_service";
import { defaultLayout } from "../elements/layouts/default";

import { SignUpForm } from "../components/sign_up_form";


function SignUpPage() {
  AuthService.redirectIfAuthenticated();

  const [message, setMessage] = useState("");
  const [submitted, setSubmitted] = useState(false);

  return (
    <>
      <Head>
        <title>SIGN UP TO WIN</title>
      </Head>
      {submitted ? message : <SignUpForm setMessage={setMessage} setSubmitted={setSubmitted}/>}
    </>
  );
}

export default defaultLayout(SignUpPage);
