import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { Login, Logout, Register } from "../../actions/user/register";
import { ApiResponse } from "../../utils/api";

interface User {
  userId?: string;
  username: string;
  email: string;
}

interface UserState {
  user: User;
  status: "idle" | "loading" | "succeeded" | "failed";
  error: string | null;
}

const initialState: UserState = {
  user: {
    username: "",
    email: "",
    userId: "",
  },
  status: "idle",
  error: null,
};

export const userSlice = createSlice({
  name: "user",
  initialState,
  reducers: {},
  extraReducers(builder) {
    builder
      .addCase(Login.pending, (state) => {
        state.status = "loading";
      })
      .addCase(
        Login.fulfilled,
        (state, action: PayloadAction<ApiResponse<unknown>>) => {
          state.status = "succeeded";
          state.user = action.payload.data as User;
        }
      )
      .addCase(Login.rejected, (state, action) => {
        state.status = "failed";
        state.error = action.error.message || "An error occurred";
      })
      .addCase(Register.pending, (state) => {
        state.status = "loading";
      })
      .addCase(
        Register.fulfilled,
        (state, action: PayloadAction<ApiResponse<unknown>>) => {
          state.status = "succeeded";
          state.user = action.payload.data as User;
        }
      )
      .addCase(Register.rejected, (state, action) => {
        state.status = "failed";
        state.error = action.error.message || "An error occurred";
      })
      .addCase(Logout.pending, (state) => {
        state.status = "loading";
      })
      .addCase(Logout.fulfilled, (state) => {
        state.status = "succeeded";
        state.user = {
          username: "",
          email: "",
          userId: "",
        };
      })
      .addCase(Logout.rejected, (state, action) => {
        state.status = "failed";
        state.error = action.error.message || "An error occurred";
      });
  },
});
export default userSlice.reducer;
