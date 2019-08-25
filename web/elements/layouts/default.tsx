import { NextPageContext } from "next";
import React, { Component } from "react";
import { AuthToken } from "../../lib/services/auth_token";
import { Head } from "./parts/head";
import Header from "./parts/header";

type Props = { auth?: AuthToken }

export function defaultLayout(C: any) {
  return class extends Component<Props> {
    static async getInitialProps(ctx: NextPageContext) {
      if (C.getInitialProps) {
        const wrappedProps = await C.getInitialProps(ctx);
        return { ...wrappedProps }
      }
      return ctx;
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
