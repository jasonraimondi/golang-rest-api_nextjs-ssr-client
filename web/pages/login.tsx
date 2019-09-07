import { NextPage } from "next";
import React from "react";

import { LoginForm } from "../components/auth/login_form";
import { defaultLayout } from "../components/layouts/default";
import { redirectIfAuthenticated } from "../lib/services/redirect_service";

const Page: NextPage<any> = () => {
  return <>
    <LoginForm/>
  </>;
};

Page.getInitialProps = async (ctx) => {
  await redirectIfAuthenticated(ctx);
  return {
    message: "by including this message, next will stop complaining about this page's get initial props returning empty"
  };
};

export default defaultLayout(Page);
