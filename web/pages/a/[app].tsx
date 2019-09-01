import { NextPageContext } from "next";
import React from "react";

function Page() {
  return <p>hello page</p>
}

Page.getInitialProps = async ({ query }: NextPageContext) => {
  const { app } = query;
  if (!app) {
    // ERROR RETURN 404
    return {}
  }
  console.log(app);
  // await listPhotosForApp
  return {};
};

export default Page;