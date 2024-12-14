import { createSlice } from "@reduxjs/toolkit";
import { UserState } from "../../interfaces/UserInterface";
import { createUser, loginUser } from "../action/userAction";

const initialState: UserState = {
  id: null,
  name: null,
  email: null,
  password: null,
  phoneNumber: null,
  role: null,
  status: "idle",
};

const userSlice = createSlice({
  name: "user",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      // Login
      .addCase(loginUser.pending, (state) => {
        state.status = "loading";
      })
      .addCase(loginUser.fulfilled, (state, action) => {
        state.status = "success";
        state.email = action.payload.email;
        state.password = action.payload.password;
      })
      .addCase(loginUser.rejected, (state, action) => {
        state.status = "error";
        console.error(action.payload);
      })

      // Register
      .addCase(createUser.pending, (state) => {
        state.status = "loading";
      })
      .addCase(createUser.fulfilled, (state, action) => {
        state.status = "success";
        state.id = action.payload.id;
        state.name = action.payload.name;
        state.email = action.payload.email;
        state.password = action.payload.password;
        state.phoneNumber = action.payload.phoneNumber;
        state.role = action.payload.role;
      })
      .addCase(createUser.rejected, (state, action) => {
        state.status = "error";
        console.error(action.payload);
      });
  },
});

export default userSlice.reducer;
