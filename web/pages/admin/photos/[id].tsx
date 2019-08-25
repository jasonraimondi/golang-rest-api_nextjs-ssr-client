import { NextPageContext } from "next";
import React from "react";
import { defaultLayout } from "../../../elements/layouts/default";
import { AuthProps, privateRoute } from "../../../lib/auth/private_route";
import { getPhoto } from "../../../lib/services/api/photos";

type Props = {} & AuthProps

function Page({ }: Props) {
  return <>
    Hi ya slugg
  </>;
}

Page.getInitialProps = async ({ query }: NextPageContext) => {
  const id: any = query["id"];
  const res = await getPhoto(id);
  console.log(res);
  return {}
};

export default privateRoute(defaultLayout(Page));
