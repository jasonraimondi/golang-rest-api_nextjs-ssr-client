import { NextPageContext } from "next";
import React from "react";
import { AuthToken } from "../../lib/services/auth_token";
import { Head } from "./parts/head";
import Header from "./parts/header";

type Props = { auth?: AuthToken }

export function defaultLayout(Page: any) {
  const DefaultLayout = (props: Props) => {
    return <>
      <Head/>
      <Header auth={props.auth}/>
      <Page auth={props.auth} {...props}/>
    </>;
  };

  DefaultLayout.getInitialProps = async (ctx: NextPageContext) => {
    return {
      ...(Page.getInitialProps ? await Page.getInitialProps(ctx) : {}),
    };
  };

  return DefaultLayout;
}
