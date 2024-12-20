import { createSlice } from "@reduxjs/toolkit";
import { GameData, GameState } from "../../interfaces/GameInterface";
import { gameFetch } from "../action/gameAction";

const initialState: GameData = {
    data: null,
    status: 'idle'
}

const gameSlice = createSlice({
    name: "game",
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder
            .addCase(gameFetch.pending, (state) => {
                state.status = 'loading'
            })
            .addCase(gameFetch.fulfilled, (state, action) => {
                state.status = "success"
                state.data = action.payload.data as GameState[]
            })
            .addCase(gameFetch.rejected, (state) => {
                state.status = "error"
            })
    }
})

export default gameSlice.reducer