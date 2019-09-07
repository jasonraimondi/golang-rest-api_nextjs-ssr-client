import Link from "next/link";
import React from "react";
import { App } from "../../lib/entity/app";
import { APP_ROUTES } from "../../lib/routes";

type Props = {
  apps: App[],
  error?: string,
};

export function AppList({ apps, error }: Props) {
  if (error) {
    return <p>{error}</p>;
  }

  if (apps.length === 0) {
    return <p>No apps.</p>
  }

  return <ul>
    {apps.map(app => <li key={app.ID}>
      <Link href={APP_ROUTES.app.index.create({ appId: app.ID, appSlug: app.Name })}>
        <a>{app.Name}</a>
      </Link>
    </li>)}
  </ul>;
}