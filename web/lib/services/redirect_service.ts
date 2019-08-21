import { ServerResponse } from "http";
import { NextPageContext } from "next";
import Router from "next/router";
import { APP_ROUTES } from "../routes";
import { AuthToken } from "./auth_token";

export const redirectToLogin = (server?: ServerResponse) => {
  redirectTo(`${APP_ROUTES.auth.login}?redirected=true`, server);
};

export const redirectIfAuthenticated = async (ctx: NextPageContext) => {
  try {
    const auth = AuthToken.fromNext(ctx);
    if (auth.isValid) {
      redirectTo(`${APP_ROUTES.admin.dashboard}?redirected=true`, ctx.res);
    }
  } catch (e) {
  }
};

const redirectTo = (url: string, server?: ServerResponse) => {
  if (server) {
    // @see https://github.com/zeit/next.js/wiki/Redirecting-in-%60getInitialProps%60
    // server rendered pages do not have access to "next/router", thus they need to redirect
    server.writeHead(302, {
      Location: url,
    });
    server.end();
  } else {
    Router.push(url);
  }
};
