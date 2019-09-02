import React from "react";
import { APP_ROUTES } from "../../lib/routes";
import { App } from "../../lib/services/api/photos";

export function AppList({ apps }: { apps: App[] }) {
  return <ul>
    {apps.map(app => <li>
      <a href={APP_ROUTES.app.index.create({ appId: app.ID, appSlug: app.Name })}>{app.Name}</a>
    </li>)}
  </ul>;
}