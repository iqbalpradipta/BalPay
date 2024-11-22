import express from "express";
import cors from "cors";
import Routes from "./Routes/route";
import 'dotenv/config'

const app = express();
const port = 8000;

app.use(cors());
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.use("/api/v1/", Routes);

app.listen(port, () => {
    console.log(`Server running at port ${port}`);
});
