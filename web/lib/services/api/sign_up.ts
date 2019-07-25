import { SignUpInputs } from "../../../pages/sign_up";
import { post } from "../../rest_client";
import { catchAxiosError } from "../error";

export async function signUp(inputs: SignUpInputs | any) {
  try {
    const foo = await post("/sign-up", inputs);
    console.log(foo);
  } catch (e) {
    return catchAxiosError(e);
  }
  return "hello sunshine my best friend";
}
