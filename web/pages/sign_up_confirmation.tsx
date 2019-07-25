import React from "react";
import { defaultLayout } from "../elements/layouts/default";
import { AuthService } from "../lib/auth/auth_service";

function Page({user_id, token}) {
  AuthService.redirectIfAuthenticated();

  return <div>Here: {user_id}, {token}</div>;
}

Page.getInitialProps = async ({res, query}) => {
  const {user_id, token} = query;
  if (!user_id || !token) {
    AuthService.redirectToLogin(res);
  }
  return {user_id, token};
};

export default defaultLayout(Page);
