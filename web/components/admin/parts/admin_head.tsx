import NextHead from "next/head";
import "../../../styles/style.css";

export const AdminHead = ({children}: any) => {
  console.log({children});
  return <>
    <NextHead>
      <title>Admin Pages</title>
      <meta charSet='utf-8'/>
      <meta key="viewport" name='viewport' content='initial-scale=1.0, width=device-width'/>
    </NextHead>
  </>;
};
