import { NextPage } from "next";
import React, { useEffect } from "react";
import { LoginForm } from "../components/auth/login_form";
import { defaultLayout } from "../components/layouts/default";
import { redirectIfAuthenticated } from "../lib/services/redirect_service";

const Page: NextPage<any> = ({token}) => {
  console.log({token});
  // Similar to componentDidMount and componentDidUpdate:
  useEffect(() => {
    // Update the document title using the browser API
    redirectIfAuthenticated(token);
  });

  return <>
    <LoginForm/>
  </>;
};

// Page.getInitialProps = async (ctx) => {
//   await redirectIfAuthenticated(ctx);
//   return {};
// };

export default defaultLayout(Page);
