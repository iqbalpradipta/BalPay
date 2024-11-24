import { Router } from "express";
import authRoute from "./auth";
import userRoute from "./user";
import gameRoute from "./game";

const Routes = Router();

Routes.use(authRoute);
Routes.use(userRoute);
Routes.use(gameRoute);

export default Routes;
