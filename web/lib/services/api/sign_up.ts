import { SignUpInputs } from "../../../pages/sign_up";
import { get, post } from "../../rest_client";
import { catchAxiosError } from "../error";

export async function signUp(inputs: SignUpInputs | any) {
  try {
    return await post("/sign_up", inputs);
  } catch (e) {
    return catchAxiosError(e);
  }
}

export async function signUpConfirmation(inputs: { u: string, t: string }) {
  try {
    return await get("/sign_up_confirmation", inputs);
  } catch (e) {
    return catchAxiosError(e);
  }
}
