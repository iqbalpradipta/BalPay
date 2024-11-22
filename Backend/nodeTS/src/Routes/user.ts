import { Router } from "express";
import JWTAuth from "../middlewares/JWTAuth";
import UserController from "../Controllers/user";

const userRoute = Router()

userRoute.get('/users', JWTAuth.Authentication(['admin']), UserController.GetUsers)
userRoute.get('/users/:id', UserController.GetUsersById)
userRoute.put('/users/:id', JWTAuth.Authentication(['admin', 'member']), UserController.UpdateUsers)
userRoute.delete('/users/:id', JWTAuth.Authentication(['admin', 'member']), UserController.DeleteUsers)

export default userRoute