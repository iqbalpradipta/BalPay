import { Router } from "express";
import JWTAuth from "../middlewares/JWTAuth";
import UserController from "../Controllers/user";
import Upload from "../middlewares/UploadFile";

const userRoute = Router()

userRoute.get('/users', JWTAuth.Authentication(['admin']), UserController.GetUsers)
userRoute.get('/users/:id', UserController.GetUsersById)
userRoute.put('/users', JWTAuth.Authentication(['admin', 'member']), Upload.single('photoProfile'), UserController.UpdateUsers)
userRoute.delete('/users', JWTAuth.Authentication(['admin', 'member']), UserController.DeleteUsers)

export default userRoute