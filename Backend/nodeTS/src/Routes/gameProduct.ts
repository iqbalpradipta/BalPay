import express from "express";
import GameProductController from "../Controllers/gameProduct";
import JWTAuth from "../middlewares/JWTAuth";
import Upload from "../middlewares/UploadFile";

const gameProductRoute = express.Router()

gameProductRoute.get('/gameProduct', GameProductController.GetGameProduct)
gameProductRoute.get('/gameProduct/:id', GameProductController.GetGameProductById)
gameProductRoute.post('/gameProduct', JWTAuth.Authentication(['admin']), Upload.array('image', 10), GameProductController.CreateGameProduct)
gameProductRoute.put('/gameProduct/:id', JWTAuth.Authentication(['admin']), Upload.array('image', 10), GameProductController.UpdateGameProduct)
gameProductRoute.delete('/gameProduct/:id',JWTAuth.Authentication(['admin']), GameProductController.DeleteGameProduct)

export default gameProductRoute