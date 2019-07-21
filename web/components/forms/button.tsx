import React from "react";

export interface SubmitFields {
  label: string;
}

export function SubmitButton({ label }: SubmitFields) {
  return <button className="p-2 mt-4 bg-blue-500 text-white rounded-full" type="submit">{label}</button>;
}
