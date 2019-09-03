import React from "react";
import { defaultLayout } from "../../components/layouts/default";
import { AuthProps, privateRoute } from "../../components/auth/private_route";

type Props = {}

function Page({ token }: Props & AuthProps) {
  return <p>Hi Dashboard: {token}</p>;
}

export default privateRoute(defaultLayout(Page));
