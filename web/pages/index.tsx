import React from "react";
import { defaultLayout } from "../elements/layouts/default";

function Page() {
  return <div className="w-full h-full flex items-center justify-center">
    <p className="text-red-500">Home</p>
  </div>;
}

Page.getInitialProps = async () => {
  return {};
};

export default defaultLayout(Page);
