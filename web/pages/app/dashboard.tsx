import React from "react";
import { privateRoute } from "../../components/auth/private_route";
import { defaultLayout } from "../../components/layouts/default";

function DashboardPage() {
  return <>
    <p>Hello Dashboard</p>
  </>;
}

export default privateRoute(defaultLayout(DashboardPage));
