import Cookie from "js-cookie";
import React from "react";
import Header from "../components/layout/header";

function Home() {
  const handleStoreToken = () => {
    Cookie.set("isAuthenticated", true);
  };

  return <>
    <Header/>
    <p>Hello Index</p>
    <button onClick={handleStoreToken}>Login</button>
  </>;
}

export default Home;
