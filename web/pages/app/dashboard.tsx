import React from "react";
import { defaultLayout } from "../../elements/layouts/default";
import { AuthService } from "../../lib/auth/auth_service";
import { privateRoute } from "../../lib/auth/private_route";

type AuthProps = {
  auth: AuthService
}

type Props = {}

function Page({auth}: Props & AuthProps) {
  console.log(auth);
  return <>
    <p>{JSON.stringify(auth.user)}</p>
  </>;
}

export default privateRoute(defaultLayout(Page));
