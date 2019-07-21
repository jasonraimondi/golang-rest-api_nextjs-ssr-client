import React, { Component } from "react";
import { AuthService } from "../auth/auth_service";
import { Head } from "./parts/head";
import Header from "./parts/header";

export function defaultLayout(WrappedComponent) {
  return class extends Component<{ auth?: AuthService }> {
    render() {
      return <>
        <Head/>
        <Header auth={this.props.auth}/>
        <WrappedComponent {...this.props}/>
      </>;
    }
  };
}
