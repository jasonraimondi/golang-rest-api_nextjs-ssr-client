import Cookies from "js-cookie";
import Router from "next/router";
import React, { Component } from "react";
import { GlobalHeader } from "./head";

export function PrivateRoute(WrappedComponent: any) {
  console.log(Cookies.get("isAuthenticated"));


  return class extends Component<{ isAuthenticated: boolean }> {
    state = {
      isAuthenticated: false,
    };

    render() {
      const { isAuthenticated } = this.state;

      if (!isAuthenticated) {
        return Router.push("/about");
      }
      return <>
        <GlobalHeader/>
        <button onClick={() => this.setState({ isAuthenticated: false })}>Logout</button>
        <p>IS AUTHENTICATED {isAuthenticated ? "YES" : "NO"}</p>
        <WrappedComponent {...this.props} />
      </>;
    }
  };
}