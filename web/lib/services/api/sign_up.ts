import { SignUpInputs } from "../../../pages/sign_up";
import { get, post } from "../../rest_client";

export async function signUp(inputs: SignUpInputs) {
  return await post("/sign_up", new URLSearchParams(inputs));
}

export async function signUpConfirmation(inputs: { u: string, t: string }) {
  return await get("/sign_up_confirmation", inputs);
}
