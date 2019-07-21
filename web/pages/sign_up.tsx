import Head from "next/head";
import React, { useState } from "react";
import { defaultLayout } from "../components/layouts/default";

import { SignUpForm } from "../components/sign_up_form";


function SignUpPage() {
  const [message, setMessage] = useState("");
  const [submitted, setSubmitted] = useState(false);

  return (
    <>
      <Head>
        <title>SIGNUP TO WIN</title>
      </Head>
      <h1>This page has a title🤔</h1>
      {submitted ? message : <SignUpForm setMessage={setMessage} setSubmitted={setSubmitted}/>}
    </>
  );
}

export default defaultLayout(SignUpPage);
