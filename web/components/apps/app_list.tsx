import React from "react";
import { APP_ROUTES } from "../../lib/routes";
import { App } from "../../lib/services/api/photos";

type Props = {
  apps: App[],
  error?: string,
};

export function AppList({ apps, error }: Props) {
  if (error) {
    return <p>{error}</p>;
  }

  return <ul>
    {apps.map(app => <li key={app.ID}>
      <a href={APP_ROUTES.app.index.create({ appId: app.ID, appSlug: app.Name })}>{app.Name}</a>
    </li>)}
  </ul>;
}