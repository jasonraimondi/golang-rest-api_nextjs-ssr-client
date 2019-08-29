import { NextPageContext } from "next";
import React, { Component } from "react";
import { AuthToken } from "../../lib/services/auth_token";
import { Head } from "./parts/head";
import Header from "./parts/header";

type Props = { auth?: AuthToken }

export function defaultLayout(Wrappe: any) {
  return class extends Component<Props> {
    static async getInitialProps(ctx: NextPageContext) {
      if (Wrappe.getInitialProps) {
        const wrappedProps = await Wrappe.getInitialProps(ctx);
        return { ...wrappedProps }
      }
      return {};
    }

    render() {
      return <>
        <Head/>
        <Header auth={this.props.auth}/>
        <Wrappe {...this.props}/>
      </>;
    }
  };
}
