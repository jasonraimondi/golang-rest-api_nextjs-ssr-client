import React, { useState } from "react";

import { Head } from "../components/layout/head";
import Header from "../components/layout/header";
import { SignUpForm } from "../components/sign_up_form";


function Home() {
  const [message, setMessage] = useState("");
  const [submitted, setSubmitted] = useState(false);

  return (
    <>
      <Head/>
      <Header/>
      <h1>This page has a title🤔</h1>
      {submitted ? message : (
        <SignUpForm setMessage={setMessage} setSubmitted={setSubmitted}/>
      )}
    </>
  );
}


export default Home;
