import React from "react";
import { privateRoute } from "../../components/auth/private_route";
import { defaultLayout } from "../../components/layouts/default";

function DashboardPage({ authService }) {
  console.log('dashboard', authService);
  return <>
    <p>Hello Dashboard</p>
  </>;
}

export default defaultLayout(privateRoute(DashboardPage));
