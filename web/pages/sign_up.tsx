import React, { useState } from "react";

import { GlobalHeader } from "../components/head";
import Header from "../components/header";
import { UserForm } from "../components/user_form";


function Home() {
  const [message, setMessage] = useState("");
  const [submitted, setSubmitted] = useState(false);

  return (
    <>
      <GlobalHeader/>
      <Header/>
      <h1>This page has a title🤔</h1>
      {submitted ? message : (
        <UserForm setMessage={setMessage} setSubmitted={setSubmitted}/>
      )}
    </>
  );
}


export default Home;
