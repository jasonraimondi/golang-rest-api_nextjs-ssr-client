import React, { Component } from "react";
import { AuthService } from "../../lib/auth/auth_service";
import { Head } from "./parts/head";
import Header from "./parts/header";

export function defaultLayout(WrappedComponent: any) {
  return class extends Component<{ auth?: AuthService }> {
    static async getInitialProps(props: any) {
      if (WrappedComponent.getInitialProps) {
        return await WrappedComponent.getInitialProps(props);
      }
      return {};
    }

    render() {
      return <>
        <Head/>
        <Header auth={this.props.auth}/>
        <WrappedComponent {...this.props}/>
      </>;
    }
  };
}
