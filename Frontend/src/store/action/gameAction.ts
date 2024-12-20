import { createAsyncThunk } from "@reduxjs/toolkit";
import { GameData } from "../../interfaces/GameInterface";
import axios from 'axios'

export const gameFetch = createAsyncThunk<
  GameData,
  void,
  { rejectValue: string }
>("game/getGames", async () => {
  try {
    const response = await axios.get<GameData>(
      "http://localhost:8000/api/v1/games"
    );
    return response.data;
  } catch (error) {
    console.error(error);
    throw error;
  }
});
