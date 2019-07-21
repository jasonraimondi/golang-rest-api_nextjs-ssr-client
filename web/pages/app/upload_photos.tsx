import React from "react";
import { privateRoute } from "../../components/auth/private_route";
import { defaultLayout } from "../../components/layouts/default";

function UploadPhotosPage() {
  return <>
    <p>Hello Upload Photos</p>
  </>;
}

export default defaultLayout(privateRoute(UploadPhotosPage));
