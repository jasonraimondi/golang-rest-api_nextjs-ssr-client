import { NextPage } from "next";
import React from "react";
import { LoginForm } from "../components/auth/login_form";
import { defaultLayout } from "../components/layouts/default";

const Page: NextPage<any> = () => {
  return <>
    <LoginForm/>
  </>;
};

// Page.getInitialProps = async (ctx) => {
//   await redirectIfAuthenticated(ctx);
//   return {};
// };

export default defaultLayout(Page);
