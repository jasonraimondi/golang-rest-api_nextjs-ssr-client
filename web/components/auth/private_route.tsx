import Cookie from "js-cookie";
import Router from "next/router";
import React, { Component } from "react";
import { COOKIES } from "../../lib/cookie";
import { AuthService } from "./auth_service";

export function privateRoute(WrappedComponent: any) {
  const authService = new AuthService(Cookie.get(COOKIES.authToken));

  return class extends Component {
    componentDidMount(): void {
      if (authService.isExpired) Router.push("/");
    }

    componentDidUpdate(): void {
      if (authService.isExpired) Router.push("/");
    }

    render() {
      return <WrappedComponent auth={authService} {...this.props} />;
    }
  };
}
