import express from "express";
import authRoute from "./auth";
import userRoute from "./user";

const Routes = express.Router();

Routes.use(authRoute)
Routes.use(userRoute)

export default Routes;