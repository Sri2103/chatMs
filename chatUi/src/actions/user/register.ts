import { createAsyncThunk } from "@reduxjs/toolkit";
import Api from "../../utils/api";

export const userApi = new Api("http://localhost:7080/user");

export type RegisterPayload = {
  username: string;
  email: string;
  password: string;
};
export type LoginPayload = {
  email: string;
  password: string;
};
export const Register = createAsyncThunk(
  "users/Register",
  async (data: RegisterPayload) => {
    const resp = await userApi.post("/register", JSON.stringify(data));
    return resp;
  }
);

export const Login = createAsyncThunk(
  "users/Login",
  async (data: LoginPayload) => {
    const resp = await userApi.post("/login", JSON.stringify(data));
    return resp;
  }
);

export const Logout = createAsyncThunk("users/Logout", async () => {
  const resp = await userApi.post("/logout");
  return resp;
});
