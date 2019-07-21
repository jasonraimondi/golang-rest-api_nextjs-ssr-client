import { SignUp } from "../../components/sign_up_form";
import { appRestClient } from "../rest_client";
import { catchAxiosError } from "./error";

export async function signUp(inputs: SignUp) {
  const foo = await appRestClient.post("/sign-up", inputs).catch(catchAxiosError);
  console.log(foo);
  return "hello sunshine my best friend";
}
