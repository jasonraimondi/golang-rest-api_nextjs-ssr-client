import NextHead from "next/head";
import { createGlobalStyle } from "styled-components";
import "../../../styles/style.css";

const GlobalStyles = createGlobalStyle`
  html, body, #__next {
    width: 100%;
    height: 100%;
  }
`;

export const Head = () => {
  return <>
    <NextHead>
      <title>This page has a title 🤔</title>
      <meta charSet='utf-8'/>
      <meta key="viewport" name='viewport' content='initial-scale=1.0, width=device-width'/>
    </NextHead>
    <GlobalStyles/>
  </>;
};
