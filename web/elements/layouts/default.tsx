import React, { Component } from "react";
import { AuthToken } from "../../lib/services/auth_token";
import { Head } from "./parts/head";
import Header from "./parts/header";

type Props = { auth?: AuthToken }

export function defaultLayout(C: any) {
  return class extends Component<Props> {
    static async getInitialProps(ctx: any) {
      const props: Props = {};
      if (ctx.auth) props.auth = ctx.auth;
      if (C.getInitialProps) return await C.getInitialProps(props);
      return props;
    }

    render() {
      return <>
        <Head/>
        <Header auth={this.props.auth}/>
        <C {...this.props}/>
      </>;
    }
  };
}
