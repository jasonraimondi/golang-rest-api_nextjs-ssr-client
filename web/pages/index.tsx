import Cookie from "js-cookie";
import React from "react";
import Header from "../components/header";

function Home() {
  const handleStoreToken = () => {
    Cookie.set("isAuthenticated", true);
  };

  const handleRemoveToken = () => {
    Cookie.remove("isAuthenticated");
  };

  return <>
    <Header/>
    <p>Hello Index</p>
    <button onClick={handleStoreToken}>Login</button>
    <button onClick={handleRemoveToken}>Logout</button>
  </>;
}

export default Home;
