import Cookie from "js-cookie";
import { LoginForm } from "../../components/login_form";
import { SignUp } from "../../components/sign_up_form";
import { COOKIES } from "../cookie";
import { appRestClient } from "../rest_client";

export async function login(inputs: LoginForm) {
  const res = await appRestClient.post<{ token: string }>("/login", inputs);
  if (res.data.token) {
    Cookie.set(COOKIES.authToken, res.data.token);
  }
}

export async function signUp(inputs: SignUp) {
  const foo = await appRestClient.post("/sign-up", inputs).catch(catchAxiosError);
  console.log(foo);
  return "hello sunshine my best friend";
}

function catchAxiosError(error: any) {
  if (error.response) {
    // The request was made and the server responded with a status code
    // that falls out of the range of 2xx
    console.log(error.response.data.message);
    console.log(error.response.status);
    console.log(error.response.headers);
  } else if (error.request) {
    // The request was made but no response was received
    // `error.request` is an instance of XMLHttpRequest in the browser and an instance of
    // http.ClientRequest in node.js
    console.log(error.request);
  } else {
    // Something happened in setting up the request that triggered an Error
    console.log("Error", error.message);
  }
  console.log(error.config);
}
