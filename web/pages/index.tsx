import React from "react";
import { defaultLayout } from "../components/layouts/default";
// import { listPhotos } from "../lib/services/api/photos";

function Page() {
  return <div className="w-full h-full flex items-center justify-center">
    <p className="text-red-500">Home</p>
  </div>;
}

Page.getInitialProps = async () => {
  // const photos = await listPhotos();
  return {};
};

export default defaultLayout(Page);
