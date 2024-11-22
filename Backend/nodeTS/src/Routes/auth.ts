import express from 'express'
import AuthController from '../Controllers/auth'
import JWTAuth from '../middlewares/JWTAuth'

const authRoute = express.Router()

authRoute.post('/user/register', AuthController.Register)
authRoute.post('/user/register/admin', AuthController.RegisterAdmin)
authRoute.post('/user/login', AuthController.Login)

export default authRoute
