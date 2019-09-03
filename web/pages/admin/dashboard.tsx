import React from "react";
import { defaultLayout } from "../../components/layouts/default";
import { AuthProps, privateRoute } from "../../components/auth/private_route";

type Props = {}

function Page({ token }: Props & AuthProps) {
  return <>
    <pre>{JSON.stringify(token)}</pre>
  </>;
}

export default privateRoute(defaultLayout(Page));
