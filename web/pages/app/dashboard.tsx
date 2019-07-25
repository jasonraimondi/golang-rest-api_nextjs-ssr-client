import React from "react";
import { defaultLayout } from "../../elements/layouts/default";
import { AuthProps, privateRoute } from "../../lib/auth/private_route";

type Props = {}

function Page({ auth }: Props & AuthProps) {
  return <>
    <pre>{JSON.stringify(auth.user)}</pre>
  </>;
}

export default privateRoute(defaultLayout(Page));
