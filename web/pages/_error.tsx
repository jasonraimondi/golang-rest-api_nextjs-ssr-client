import { NextPage } from "next";

import { defaultLayout } from "@/components/layouts/default_layout";

const Error: NextPage<{ statusCode?: number }> = ({ statusCode }) => {
  return <p>
    {statusCode
      ? `An error ${statusCode} occurred on server`
      : "An error occurred on client"}
  </p>;
};

Error.getInitialProps = async ({ res, err }) => {
  let statusCode: number | undefined;
  if (res) {
    statusCode = res.statusCode;
  } else {
    statusCode = err ? err.statusCode : undefined;
  }
  return { statusCode };
};

export default defaultLayout(Error);