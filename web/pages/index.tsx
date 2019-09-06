import { NextPage } from "next";
import React from "react";
import { AppList } from "../components/apps/app_list";

import { defaultLayout } from "../components/layouts/default";
import { ApiResponse } from "../lib/services/api/api_response";
import { listApps } from "../lib/services/api/apps";
import { App } from "../lib/services/api/photos";

type Props = {
  apps: ApiResponse<App[]>
};

const Page: NextPage<Props> = ({ apps }: Props) => {
  const [list, error] = apps;
  return <AppList apps={list} error={error}/>;
};

Page.getInitialProps = async () => {
  const apps = await listApps(1, 250);
  return { apps };
};

export default defaultLayout(Page);
