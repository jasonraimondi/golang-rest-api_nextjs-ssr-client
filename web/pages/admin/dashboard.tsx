import React from "react";
import { privateRoute } from "../../components/auth/private_route";
import { defaultLayout } from "../../components/layouts/default";

const Page = () => {
  return <p>Hi Dashboard</p>;
};

export default privateRoute(defaultLayout(Page));
