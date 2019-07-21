import Cookies from "js-cookie";
import Router from "next/router";
import React, { Component } from "react";
import { COOKIES } from "../../lib/cookie";
import { GlobalHeader } from "../layout/head";
import { AuthService } from "./auth_service";

export function PrivateRoute(WrappedComponent: any) {
  const authService = new AuthService(Cookies.get(COOKIES.authToken));

  return class extends Component {
    state = {
      authService,
    };

    constructor(props) {
      super(props);
      this.logout = this.logout.bind(this);
    }

    get isAuthenticated(): boolean {
      console.log(this.state.authService);
      return this.state.authService.isAuthenticated;
    }

    componentDidMount(): void {
      if (!this.isAuthenticated) {
        Router.push("/");
      }
    }

    componentDidUpdate(): void {
      if (!this.isAuthenticated) {
        Router.push("/");
      }
    }

    logout() {
      Cookies.remove(COOKIES.authToken);
    }

    render() {

      return <>
        <GlobalHeader/>
        <WrappedComponent {...this.props} />
      </>;
    }
  };
}