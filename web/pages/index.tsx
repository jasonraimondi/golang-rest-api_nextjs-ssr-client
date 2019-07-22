import React from "react";
import { defaultLayout } from "../elements/layouts/default";

function Index() {
  return <>
    <p className="text-red-500">Home</p>
  </>;
}

export default defaultLayout(Index);
