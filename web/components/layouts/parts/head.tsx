import NextHead from "next/head";

export const Head = () => {
  return <>
    <NextHead>
      <title>This page has a title 🤔</title>
      <meta charSet='utf-8'/>
      <meta key="viewport" name='viewport' content='initial-scale=1.0, width=device-width'/>
    </NextHead>
  </>;
};
