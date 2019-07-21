import React from "react";

import { Head } from "../components/layout/head";
import Header from "../components/layout/header";
import { Login } from "../components/login_form";

function Home() {
  return (
    <>
      <Head/>
      <Header/>
      <Login/>
    </>
  );
}


export default Home;
