import React from "react";

import { GlobalHeader } from "../components/layout/head";
import Header from "../components/layout/header";
import { Login } from "../components/login_form";

function Home() {
  return (
    <>
      <GlobalHeader/>
      <Header/>
      <Login/>
    </>
  );
}


export default Home;
