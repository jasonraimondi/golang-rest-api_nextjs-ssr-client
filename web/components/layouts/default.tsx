import { NextPageContext } from "next";
import React, { Component } from "react";
import { AuthProps } from "../../lib/auth/private_route";
import { AuthToken } from "../../lib/services/auth_token";
import { Head } from "./parts/head";
import Header from "./parts/header";

type Props = { auth?: AuthToken }

export function defaultLayout(Page: any) {
  return class extends Component<Props> {
    static async getInitialProps(ctx: NextPageContext & AuthProps) {
      let result = {
        ...(Page.getInitialProps ? await Page.getInitialProps(ctx) : {}),
      };
      console.log("default layout", result);
      return result;
    }

    render() {
      return <>
        <Head/>
        <Header auth={this.props.auth}/>
        <Page {...this.props}/>
      </>;
    }
  };
}
