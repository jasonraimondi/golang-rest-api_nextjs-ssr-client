import Link from "next/link";
import React from "react";
import { APP_ROUTES } from "../../../lib/routes";

export const Header = () => {
  return <header>
    <nav>
      <ul className="flex justify-around">
        <li>
          <Link href={APP_ROUTES.home.create()}>
            <a>Home</a>
          </Link>
        </li>
        <li>
          <Link href={APP_ROUTES.auth.login.create()}>
            <a>Login</a>
          </Link>
        </li>
        <li>
          <Link href={APP_ROUTES.signUp.create()}>
            <a>Sign Up</a>
          </Link>
        </li>
      </ul>
    </nav>
  </header>;
};
