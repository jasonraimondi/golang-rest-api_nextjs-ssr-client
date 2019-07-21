import Cookies from "js-cookie";
import Router from "next/router";
import React, { Component } from "react";
import { COOKIES } from "../../lib/cookie";
import { Head } from "../parts/head";
import { AuthService } from "./auth_service";

export function privateRoute(WrappedComponent: any) {
  return class extends Component {
    state = {
      authService: new AuthService(Cookies.get(COOKIES.authToken)),
    };

    componentDidMount(): void {
      if (this.state.authService.isExpired) Router.push("/");
    }

    componentDidUpdate(): void {
      if (this.state.authService.isExpired) Router.push("/");
    }

    render() {
      return <>
        <Head/>
        <LogOut auth={this.state.authService}/>
        <WrappedComponent auth={this.state.authService} {...this.props} />
      </>;
    }
  };
}

interface Props {
  auth: AuthService
}

function LogOut({ auth }: Props) {
  return <>
    <button onClick={auth.logout}>LOGOUT YEAH</button>
  </>;
}