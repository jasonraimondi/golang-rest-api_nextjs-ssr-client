import Cookie from "js-cookie";
import React from "react";
import { PrivateRoute } from "../../components/private_route";

function UploadPhotos() {
  const handleStoreToken = () => {
    Cookie.set("isAuthenticated", true);
  };
  const handleRemoveToken = () => {
    Cookie.remove("isAuthenticated");
  };
  return <>
    <p>Hello Upload Photos</p>
    <button onClick={handleStoreToken}>Login</button>
    <button onClick={handleRemoveToken}>Logout</button>
  </>;
}

export default PrivateRoute(UploadPhotos);
