import { SignUp } from "../../../components/sign_up_form";
import { post } from "../../rest_client";
import { catchAxiosError } from "../error";

export async function signUp(inputs: SignUp | any) {
  try {
    const foo = await post("/sign-up", inputs);
    console.log(foo);
  } catch (e) {
    return catchAxiosError(e);
  }
  return "hello sunshine my best friend";
}
