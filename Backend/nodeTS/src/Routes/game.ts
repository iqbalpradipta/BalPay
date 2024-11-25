import express from "express";
import GameControllers from "../Controllers/game";
import JWTAuth from "../middlewares/JWTAuth";
import Upload from "../middlewares/UploadFile";

const gameRoute = express.Router()

gameRoute.get('/games', GameControllers.GetGames)
gameRoute.get('/games/:id', GameControllers.GetGameById)
gameRoute.post('/games', JWTAuth.Authentication(['admin']), Upload.single('icon'), GameControllers.CreateGame)
gameRoute.put('/games/:id', JWTAuth.Authentication(['admin']), Upload.single('icon'), GameControllers.UpdateGame)
gameRoute.delete('/games/:id',JWTAuth.Authentication(['admin']), GameControllers.DeleteGame)

export default gameRoute