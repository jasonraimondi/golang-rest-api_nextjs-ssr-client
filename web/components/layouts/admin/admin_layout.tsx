import { NextPage } from "next";
import React from "react";

import { AuthToken } from "../../../lib/services/auth_token";
import { privateRoute } from "../../auth/private_route";
import { AdminHead } from "./parts/admin_head";
import { AdminHeader } from "./parts/admin_header";

type Props = { auth?: AuthToken }

export function adminLayout(Page: any) {
  const DefaultLayout: NextPage<any> = (props: Props) => {
    return <>
      <AdminHead/>
      <AdminHeader auth={props.auth}/>
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
