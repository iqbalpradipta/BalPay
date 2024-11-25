import { Router } from "express";
import authRoute from "./auth";
import userRoute from "./user";
import gameRoute from "./game";
import gameProductRoute from "./gameProduct";

const Routes = Router();

Routes.use(authRoute);
Routes.use(userRoute);
Routes.use(gameRoute);
Routes.use(gameProductRoute);

export default Routes;
