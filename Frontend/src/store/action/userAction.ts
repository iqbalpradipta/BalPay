import { createAsyncThunk } from "@reduxjs/toolkit";
import { UserState } from "../../interfaces/UserInterface";
import axios from "axios";
import Cookies from 'js-cookie'

export const loginUser = createAsyncThunk<
  UserState,
  { name: string; email: string; password: string; phoneNumber: string },
  { rejectValue: string }
>("user/loginUser", async (userData) => {
  try {
    const response = await axios.post(
      "http://localhost:8000/api/v1/user/login",
      userData
    );
    Cookies.set('token', (response.data as any).token as string, { expires: 1 })
    return response.data as UserState;
  } catch (error) {
    throw error;
  }
});

export const createUser = createAsyncThunk<
  UserState,
  { name: string; email: string; password: string; phoneNumber: string },
  { rejectValue: string }
>("user/createUser", async (userData) => {
  try {
    const response = await axios.post(
      "http://localhost:8000/api/v1/user/register",
      userData
    );
    return response.data as UserState;
  } catch (error) {
    throw error;
  }
});
