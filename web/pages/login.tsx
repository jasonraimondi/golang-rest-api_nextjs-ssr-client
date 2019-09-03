import { NextPageContext } from "next";
import React from "react";
import { LoginForm } from "../components/auth/login_form";
import { defaultLayout } from "../components/layouts/default";
import { redirectIfAuthenticated } from "../lib/services/redirect_service";

function Page() {
  return <>
    <LoginForm/>
  </>;
}

Page.getInitialProps = async (ctx: NextPageContext) => {
  await redirectIfAuthenticated(ctx);
  return {};
};

export default defaultLayout(Page);
