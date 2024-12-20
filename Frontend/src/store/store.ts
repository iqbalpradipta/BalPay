import { configureStore } from "@reduxjs/toolkit";
import userReducer from "./reducer/userReducer";
import gameReducer from "./reducer/gameReducer";

const store = configureStore({
  reducer: {
    users: userReducer,
    games: gameReducer,
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;

export default store;
