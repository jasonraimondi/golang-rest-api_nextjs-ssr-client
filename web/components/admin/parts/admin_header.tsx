import Link from "next/link";
import React from "react";

import { APP_ROUTES } from "@/lib/routes";

export const AdminHeader = () => {
  return <header>
    <nav>
      <ul className="flex justify-around">
        <li>PRIVATE ROUTE</li>
        <PrivateRoutes/>
      </ul>
    </nav>
  </header>;
};

function PrivateRoutes() {
  return <>
    <li>
      <Link href={APP_ROUTES.admin.dashboard.create()}>
        <a>Dashboard</a>
      </Link>
    </li>
    <li>
      <Link href={APP_ROUTES.admin.photos.index.create()}>
        <a>Photo</a>
      </Link>
    </li>
    <li>
      <Link href={APP_ROUTES.admin.photos.upload.create()}>
        <a>Upload</a>
      </Link>
    </li>
    <li>
      <Link href={APP_ROUTES.auth.logout.create()}>
        <a>Logout</a>
      </Link>
    </li>
  </>;
}
