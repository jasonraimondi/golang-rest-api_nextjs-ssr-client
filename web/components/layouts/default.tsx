import React, { Component } from "react";
import { Head } from "../parts/head";
import Header from "../parts/header";

export function defaultLayout(WrappedComponent: any) {
  return class extends Component {
    render() {
      return <>
        <Head/>
        <Header authService={undefined}/>
        <WrappedComponent/>
      </>;
    }
  };
}
