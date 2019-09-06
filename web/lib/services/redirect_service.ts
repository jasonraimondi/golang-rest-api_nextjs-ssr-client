import { ServerResponse } from "http";
// import { NextPageContext } from "next";
import Router from "next/router";
import { APP_ROUTES } from "../routes";
import { AuthToken } from "./auth_token";

export const redirectToLogin = async (server?: ServerResponse) => {
  if (server) console.log("toLogin", server.getHeaders());
  await redirectTo(`${APP_ROUTES.auth.login.create()}#/redirected`, server);
};

export const redirectIfAuthenticated = async (token: string) => {
  try {
    const auth = AuthToken.fromToken(token);
    if (auth.isValid) {
      await redirectTo(`${APP_ROUTES.admin.dashboard.create()}#/redirected`);
    }
  } catch (e) {
  }
};

const redirectTo = async (url: string, server?: ServerResponse) => {
  if (server) {
    // @see https://github.com/zeit/next.js/wiki/Redirecting-in-%60getInitialProps%60
    // server rendered pages do not have access to "next/router", thus they need to redirect
    server.writeHead(302, {
      Location: url,
    });
    server.end();
  } else {
    await Router.push(url);
  }
};
