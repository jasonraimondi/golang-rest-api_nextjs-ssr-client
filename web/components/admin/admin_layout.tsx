import { NextPage } from "next";
import React from "react";

import { AuthProps, privateRoute } from "../auth/private_route";
import { AdminHead } from "./parts/admin_head";
import { AdminHeader } from "./parts/admin_header";

type Props = AuthProps

export function adminLayout(Page: any) {
  const DefaultLayout: NextPage<any> = (props: Props) => {
    return <>
      <AdminHead/>
      <AdminHeader/>
      <Page auth={props.auth} {...props}/>
    </>;
  };

  DefaultLayout.getInitialProps = async (ctx) => {
    return {
      ...(Page.getInitialProps ? await Page.getInitialProps(ctx) : {}),
    };
  };

  return privateRoute(DefaultLayout);
}
