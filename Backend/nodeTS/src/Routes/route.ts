import express from "express";
import authRoute from "./auth";

const Routes = express.Router();

Routes.use(authRoute)

export default Routes;