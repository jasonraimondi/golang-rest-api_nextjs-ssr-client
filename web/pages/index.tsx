import React from "react";
import { AppList } from "../components/apps/app_list";

import { defaultLayout } from "../components/layouts/default";
import { listApps } from "../lib/services/api/apps";
import { App } from "../lib/services/api/photos";

function Page({ apps }: { apps: App[] }) {
  return <AppList apps={apps ? apps : []}/>;
}

Page.getInitialProps = async () => {
  const apps = await listApps(1, 250);
  return { apps };
};

export default defaultLayout(Page);
