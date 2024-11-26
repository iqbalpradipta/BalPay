import express from "express";
import cors from "cors";
import Routes from "./Routes/route";
import "dotenv/config";
import { refreshToken } from "./middlewares/dropBoxGenerateToken";
import nodemon from "nodemon";

const app = express();
const port = 8000;

app.use(cors());
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.use("/api/v1/", Routes);

setInterval(async () => {
    await refreshToken()
    console.log("Refresh Token ....");
    nodemon.emit("restart", ["."]);
    process.exit(0);
  }, 4 * 60 * 60 * 1000);

app.listen(port, () => {
  console.log(`Server running at port ${port}`);
});
