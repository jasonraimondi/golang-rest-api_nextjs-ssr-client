import React from "react";
import { PrivateRoute } from "../../components/auth/private_route";
import Header from "../../components/layout/header";

function UploadPhotos() {
  return <>
    <Header/>
    <p>Hello Upload Photos</p>
  </>;
}

export default PrivateRoute(UploadPhotos);
